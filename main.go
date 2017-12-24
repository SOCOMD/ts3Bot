package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/zanven42/ts3Query"
)

func main() {

	helpFlag := flag.Bool("help", false, "If Defined will print the help menu")
	flag.Parse()
	tsIP := os.Getenv("TSBOT_TSIP")
	tsPort := os.Getenv("TSBOT_TSPORT")
	tsUser := os.Getenv("TSBOT_TSUSER")
	tsPass := os.Getenv("TSBOT_TSPASS")
	tsLogFile := os.Getenv("TSBOT_LOGFILE")
	tsDelay := os.Getenv("TSBOT_COMMANDDELAY_MILLISECONDS")

	if *helpFlag == true {
		//print all help things and leave
		fmt.Printf(`Environment variable Settings:
			Teamspeak:
			TSBOT_TSIP=` + tsIP + `
			TSBOT_TSPORT=` + tsPort + `
			TSBOT_TSUSER=` + tsUser + `
			TSBOT_TSPASS=` + tsPass + `
			TSBOT_COMMANDDELAY_MILLISECONDS= ` + tsDelay + `
			Misc:
			TSBOT_LOGFILE=` + tsLogFile + `
			`)
		return
	}

	// if a log file is specified make the system log to the file and to stdout
	if tsLogFile != "" {
		f, err := os.OpenFile(tsLogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file: %s", err)
		}
		defer f.Close()
		log.SetOutput(io.MultiWriter(os.Stdout))
	}
	if tsDelay == "" {
		tsDelay = "20"
	}
	fmt.Printf("Connecting to server:" + tsIP + ":" + tsPort + " \n")
	// establish a connection to the teamspeak server
	conn, err := net.Dial("tcp", tsIP+":"+tsPort)
	if err != nil {
		log.Fatalf("Failed dialing ts server: %s\n", err)
	}
	defer conn.Close()
	// get the initial message from dialing and read it into the void
	for {
		buff := make([]byte, 2048)
		n, err := conn.Read(buff)
		if n != 0 || err != nil {

			break
		}
	}

	delay, err := strconv.Atoi(tsDelay)
	if err != nil {
		log.Fatalf("Can't convert delay to int")
	}
	// Create the main Query Object

	query := ts3Query.New(conn, time.Millisecond*time.Duration(delay))
	if err := query.Login(tsUser, tsPass); err != nil {
		log.Fatalf("Failed to log in to teamspeak server: %s", err)
	}

	if err := query.Use("1"); err != nil {
		log.Fatalf("Failed to use the main virtual Server: %s", err)
	}

	/*
		res, err := query.ServerGroupList()
		if err != nil {
			log.Fatalf("Failed to check help: %s", err)
		}
		log.Println(res)
	*/
	clients := query.ClientDBList()
	for _, v := range clients {
		fmt.Println(v)
	}

	memberclients, err := query.ServerGroupClientList("12")
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}
	for _, client := range clients {
		for _, cldbid := range memberclients {
			if cldbid == client.DBID {
				fmt.Printf("%#v\n", client)
			}

		}
	}
	fmt.Printf("Client IDS in member group\n%#v\n", memberclients)

}
