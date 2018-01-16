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

	"github.com/SOCOMD/env"

	pb "github.com/SOCOMD/ts3Bot"
	"github.com/zanven42/ts3Query"
	"google.golang.org/grpc"
)

var (
	e env.Env
)

func interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// establish a connection to the teamspeak server
	conn, err := net.Dial("tcp", e.Tsbot.TSHost+":"+e.Tsbot.TSPort)
	if err != nil {
		err = fmt.Errorf("Failed dialing ts server: %s", err)
		return
	}
	defer conn.Close()

	delay, err := strconv.Atoi(e.Tsbot.TSDelay)
	if err != nil {
		err = fmt.Errorf("Can't convert delay to int")
		return
	}
	q, err := QueryConnect(conn, time.Millisecond*time.Duration(delay), e.Tsbot.TSUser, e.Tsbot.TSPass)
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
	fmt.Println("Get User Called")
	user = &pb.User{}
	clients := s.Query.ClientDBList()
	for _, client := range clients {
		userFound := false
		if len(in.Dbid) > 0 && client.DBID == in.Dbid {
			userFound = true
		}

		if len(in.Uuid) > 0 && client.UUID == in.Uuid {
			userFound = true
		}

		if len(in.Name) > 0 && client.Name == in.Name {
			userFound = true
		}

		if userFound == false {
			continue
		}

		user.Dbid = client.DBID
		user.Name = client.Name
		user.Uuid = client.UUID
		user.Created = client.Created
		user.Lastconnected = client.LastConnected
		fmt.Printf("Found user: \n%#v\n", user)
		return
	}
	err = fmt.Errorf("no user was found")
	return
}

func (s *server) ClientList(ctx context.Context, in *pb.Nil) (users *pb.UserList, err error) {
	users = &pb.UserList{}
	clients, err := s.Query.ClientList()
	if err != nil {
		return
	}

	for _, client := range clients {
		users.Users = append(users.Users, &pb.User{Name: client.Name, Dbid: client.DBID, Uuid: client.UUID, Created: client.Created, Lastconnected: client.LastConnected})
	}

	fmt.Println("ClientList found", len(users.Users), "users")
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
	helpFlag := flag.Bool("help", false, "If Defined will print the help menu")
	devFlag := flag.Bool("dev", false, "If Defined will run dev code instead of hosting bot")
	flag.Parse()
	e = env.Get()

	// if a log file is specified make the system log to the file and to stdout
	if e.Tsbot.LogFile != "" {
		f, err := os.OpenFile(e.Tsbot.LogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file: %s", err)
		}
		defer f.Close()
		log.SetOutput(io.MultiWriter(os.Stdout))
	}

	if e.Tsbot.TSDelay == "" {
		e.Tsbot.TSDelay = "20"
	}

	if *helpFlag == true {
		//print all help things and leave
		fmt.Println(e)
		return
	}

	if *devFlag == true {
		ts3Query, _, err := login()
		if err != nil {
			panic(err.Error())
		}

		dev(ts3Query)

		return
	}

	opt := grpc.UnaryInterceptor(interceptor)
	grpcServer := grpc.NewServer(opt)
	pb.RegisterTs3BotServer(grpcServer, &server{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", e.Tsbot.GrpcHost, e.Tsbot.GrpcPort))
	if err != nil {
		log.Fatalf("Failed to listen on GRPC PORT: %v", err)
	}
	log.Printf("GRPC Listening on %s: %s\n", e.Tsbot.GrpcHost, e.Tsbot.GrpcPort)
	grpcServer.Serve(lis)
}

func login() (query ts3Query.Ts3Query, connection *net.Conn, err error) {
	// establish a connection to the teamspeak server
	conn, err := net.Dial("tcp", e.Tsbot.TSHost+":"+e.Tsbot.TSPort)
	if err != nil {
		err = fmt.Errorf("Failed dialing ts server: %s", err)
		return
	}
	connection = &conn

	delay, err := strconv.Atoi(e.Tsbot.TSDelay)
	if err != nil {
		err = fmt.Errorf("Can't convert delay to int")
		return
	}
	query, err = QueryConnect(conn, time.Millisecond*time.Duration(delay), e.Tsbot.TSUser, e.Tsbot.TSPass)
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

func dev(ts3Query ts3Query.Ts3Query) {

}

func cleanupClientTags(ts3Query ts3Query.Ts3Query) {
	dbClients := ts3Query.ClientDBList()
	if dbClients == nil {
		return
	}

	currentTime := time.Now().Unix()

	for _, client := range dbClients {

		if client.Name == "ServerQuery Guest" {
			continue
		}

		lastConnected, _ := strconv.Atoi(client.LastConnected)
		elapsed := currentTime - int64(lastConnected)
		elapsedInDays := elapsed / 86400

		if elapsedInDays <= 180 {
			continue
		}

		serverGroups, err := ts3Query.ServerGroupByClientID(client.DBID)
		if err != nil {
			continue
		}

		isRetired := false
		for _, serverGroup := range serverGroups {
			if serverGroup.SGID == "11" {
				isRetired = true
				break
			}

			if serverGroup.SGID == "8" {
				isRetired = true
				break
			}
		}

		if isRetired == true {
			continue
		}

		fmt.Printf("\n\n===== %s =====\n", client.Name)
		for _, serverGroup := range serverGroups {
			fmt.Printf("%s\n", serverGroup.Name)
		}
		fmt.Printf("Remove %s (Y/N)?\n", client.Name)

		var res string
		fmt.Scanf("%s\n", &res)
		if res != "y" && res != "Y" {
			continue
		}

		for _, serverGroup := range serverGroups {
			err = ts3Query.ServerGroupDelClient(client.DBID, serverGroup.SGID)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}
