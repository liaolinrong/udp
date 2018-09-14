// udpServer project main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("udpserver")

	serverIp := flag.String("s", "0.0.0.0", "Local UDP server listen address")
	serverPort := flag.Int("p", 8080, "Local UDP server listen port")
	flag.Parse()

	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		//IP:   net.IPv4(0, 0, 0, 0),
		IP:   net.ParseIP(*serverIp),
		Port: *serverPort,
	})

	if err != nil {
		fmt.Println("listern ", err)
		return
	}

	defer socket.Close()

	ipString := getIpString()

	for {
		data := make([]byte, 100)
		read, remoteAddr, err := socket.ReadFromUDP(data)
		if err != nil {
			fmt.Println("read data ", err)
			continue
		}

		fmt.Println(read, remoteAddr)
		fmt.Printf("%s\n", data)
		send_data := []byte("hi client! " + ipString)
		_, err = socket.WriteToUDP(send_data, remoteAddr)
		if err != nil {
			return
			fmt.Println("send fail!", err)
		}
	}
}

func getIpString() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		log.Fatal("InterfaceAddrs: ", err.Error())
	}

	ipString := "This is "
	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipString += ipnet.IP.String() + " "
			}

		}
	}

	return ipString
}
