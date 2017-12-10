package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"

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

	if *helpFlag == true {
		//print all help things and leave
		fmt.Printf(`Environment variable Settings:
 Teamspeak:
  TSBOT_TSIP=` + tsIP + `
  TSBOT_TSPORT=` + tsPort + `
  TSBOT_TSUSER=` + tsUser + `
  TSBOT_TSPASS=` + tsPass + `

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
	fmt.Printf("Connecting to server:" + tsIP + ":" + tsPort + " \n")
	// establish a connection to the teamspeak server
	conn, err := net.Dial("tcp", tsIP+":"+tsPort)
	if err != nil {
		log.Fatalf("Failed dialing ts server: %s\n", err)
	}
	// get the initial message from dialing and read it into the void
	for {
		buff := make([]byte, 2048)
		n, err := conn.Read(buff)
		if n != 0 || err != nil {

			break
		}
	}
	query := ts3Query.New(conn)

	// login to the ts server
	if err := query.Login(tsUser, tsPass); err != nil {
		log.Fatalf("Failed to log in to teamspeak server: %s", err)
	}

	if err := query.Use("1"); err != nil {
		log.Fatalf("Failed to use the main virtual Server: %s", err)
	}

	fmt.Printf("Connected\n")
	res, err := query.Help("login")
	if err != nil {
		log.Fatalf("Failed to check help: %s", err)
	}
	log.Println(res)
}
