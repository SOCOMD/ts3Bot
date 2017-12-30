package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	pb "github.com/SOCOMD/ts3Bot"
	"github.com/zanven42/ts3Query"
	"google.golang.org/grpc"
)

var (
	env struct {
		TSIP           string
		TSPort         string
		TSUsername     string
		TSPassword     string
		TSCommandDelay string
		TSLogFile      string

		GRPCPort string
	}
)

func interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// establish a connection to the teamspeak server
	conn, err := net.Dial("tcp", env.TSIP+":"+env.TSPort)
	if err != nil {
		err = fmt.Errorf("Failed dialing ts server: %s", err)
		return
	}
	defer conn.Close()

	delay, err := strconv.Atoi(env.TSCommandDelay)
	if err != nil {
		err = fmt.Errorf("Can't convert delay to int")
		return
	}
	q, err := QueryConnect(conn, time.Millisecond*time.Duration(delay), env.TSUsername, env.TSPassword)
	if err != nil {
		err = fmt.Errorf("Failed to connect to teamspeak Server: %s", err)
		return
	}

	server, ok := info.Server.(*server)
	if ok == false {
		err = fmt.Errorf("Fucked the casting")
		return
	}
	server.Query = &q
	resp, err = handler(ctx, req)
	server.Query.Logout()
	server.Query.Quit()
	return
}

type server struct {
	Query *ts3Query.Ts3Query
}

// Returns a list of all users.
func (s *server) GetUsers(context.Context, *pb.Nil) (users *pb.UserList, err error) {
	users = &pb.UserList{}
	clients := s.Query.ClientDBList()
	for _, client := range clients {
		users.Users = append(users.Users, &pb.User{Name: client.Name, Dbid: client.DBID, Uuid: client.UUID, Created: client.Created, Lastconnected: client.LastConnected})
	}
	fmt.Println("GetUsers found", len(users.Users), "users")
	return
}

// Not Implemented
func (s *server) GetUser(ctx context.Context, in *pb.User) (user *pb.User, err error) {
	user = &pb.User{}
	clients := s.Query.ClientDBList()
	for _, client := range clients {
		if client.DBID == in.Dbid {
			user.Dbid = client.DBID
			user.Name = client.Name
			user.Uuid = client.UUID
			user.Created = client.Created
			user.Lastconnected = client.LastConnected
			fmt.Printf("Found user: \n%#v\n", user)
			return
		}
	}
	err = fmt.Errorf("no user was found")
	return
}

//
func (s *server) GetServerGroups(context.Context, *pb.Nil) (result *pb.ServerGroupList, err error) {
	result = &pb.ServerGroupList{}
	groups, err := s.Query.ServerGroupList()
	if err != nil {
		return
	}
	for _, sg := range groups {
		result.Groups = append(result.Groups, &pb.ServerGroup{Name: sg.Name, Sgid: sg.SGID})
	}
	return
}

// TODO
func (s *server) GetUsersInGroup(ctx context.Context, in *pb.ServerGroup) (users *pb.UserList, err error) {
	users = &pb.UserList{}
	ids, err := s.Query.ServerGroupClientList(in.Sgid)
	if err != nil {
		return
	}
	for _, v := range ids {
		users.Users = append(users.Users, &pb.User{Dbid: v})
	}
	return
}

func (s *server) AddUserToGroup(ctx context.Context, in *pb.UserAndGroup) (n *pb.Nil, err error) {
	err = s.Query.ServerGroupAddClient(in.User.Dbid, in.Group.Sgid)
	return
}

func (s *server) DelUserFromGroup(ctx context.Context, in *pb.UserAndGroup) (n *pb.Nil, err error) {
	err = s.Query.ServerGroupDelClient(in.User.Dbid, in.Group.Sgid)
	return nil, nil
}

func main() {
	err := passENV()
	if err != nil {
		fmt.Println(err)
		return
	}
	// if a log file is specified make the system log to the file and to stdout
	if env.TSLogFile != "" {
		f, err := os.OpenFile(env.TSLogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file: %s", err)
		}
		defer f.Close()
		log.SetOutput(io.MultiWriter(os.Stdout))
	}
	if env.TSCommandDelay == "" {
		env.TSCommandDelay = "20"
	}

	opt := grpc.UnaryInterceptor(interceptor)
	grpcServer := grpc.NewServer(opt)
	pb.RegisterTs3BotServer(grpcServer, &server{})
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", env.GRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen on GRPC PORT: %v", err)
	}
	log.Printf("GRPC Listening on Port: %s\n", env.GRPCPort)
	grpcServer.Serve(lis)
}

func login() (query ts3Query.Ts3Query, connection *net.Conn, err error) {
	if env.TSCommandDelay == "" {
		env.TSCommandDelay = "20"
	}
	// establish a connection to the teamspeak server
	conn, err := net.Dial("tcp", env.TSIP+":"+env.TSPort)
	if err != nil {
		err = fmt.Errorf("Failed dialing ts server: %s", err)
		return
	}
	connection = &conn

	delay, err := strconv.Atoi(env.TSCommandDelay)
	if err != nil {
		err = fmt.Errorf("Can't convert delay to int")
		return
	}
	query, err = QueryConnect(conn, time.Millisecond*time.Duration(delay), env.TSUsername, env.TSPassword)
	err = fmt.Errorf("Testing the error yo")
	return
}

// QueryConnect takes the tcp connection and connects to the ts3 query. it first removes the connection response followed by logging into the ts server
func QueryConnect(rw io.ReadWriter, commandDelay time.Duration, tsUser string, tsPass string) (query ts3Query.Ts3Query, err error) {
	// get the initial message from dialing and read it into the void
	for {
		buff := make([]byte, 2048)
		n, err := rw.Read(buff)
		if n != 0 || err != nil {

			break
		}
	}

	// Create the main Query Object

	query = ts3Query.New(rw, commandDelay)

	if err = query.Login(tsUser, tsPass); err != nil {
		log.Fatalf("Failed to log in to teamspeak server: %s", err)
	}

	if err := query.Use("1"); err != nil {
		log.Fatalf("Failed to use the main virtual Server: %s", err)
	}
	return
}

func passENV() (err error) {

	helpFlag := flag.Bool("help", false, "If Defined will print the help menu")
	flag.Parse()
	env.TSIP = os.Getenv("TSBOT_TSIP")
	env.TSPort = os.Getenv("TSBOT_TSPORT")
	env.TSUsername = os.Getenv("TSBOT_TSUSER")
	env.TSPassword = os.Getenv("TSBOT_TSPASS")
	env.TSLogFile = os.Getenv("TSBOT_LOGFILE")
	env.TSCommandDelay = os.Getenv("TSBOT_COMMANDDELAY_MILLISECONDS")
	env.GRPCPort = os.Getenv("TSBOT_GRPC_PORT")

	if *helpFlag == true {
		//print all help things and leave
		err = fmt.Errorf(`Environment variable Settings:
Teamspeak:
  TSBOT_TSIP=` + env.TSIP + `
  TSBOT_TSPORT=` + env.TSPort + `
  TSBOT_TSUSER=` + env.TSUsername + `
  TSBOT_TSPASS=` + env.TSPassword + `
  TSBOT_COMMANDDELAY_MILLISECONDS= ` + env.TSCommandDelay + `
Misc:
  TSBOT_LOGFILE=` + env.TSLogFile + `
  TSBOT_GRPC_PORT=` + env.GRPCPort + `
`)

	}
	return
}
