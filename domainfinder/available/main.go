package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// mplements is in the WHOIS specification by opening a connection to port 43
// on the specified whoisServer instance with a call to net.Dial
func exists(domain string) (bool, error) {
	// The com.whois-servers.net WHOIS service supports domain names for .com and .net
	const whoisServer string = "com.whois-servers.net"
	conn, err := net.Dial("tcp", whoisServer+":43")
	if err != nil {
		return false, err
	}
	defer conn.Close()

	// we simply write the domain followed by rn
	// (the carriage return and linefeed characters)
	conn.Write([]byte(domain + "rn"))
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {

		if strings.Contains(strings.ToLower(scanner.Text()), "no match") {
			return false, nil
		}
	}

	return true, nil
}

var marks = map[bool]string{true: "Yes", false: "No"}

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		domain := s.Text()
		fmt.Print(domain, " ")
		exist, err := exists(domain)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(marks[!exist])
		// to make sure we take it easy on the WHOIS server
		time.Sleep(1 * time.Second)
	}
}
