// udpClient project main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("udpClient")
	serverIp := flag.String("s", "127.0.0.1", "UDP server address")
	serverPort := flag.Int("p", 8080, "UDP server port")
	flag.Parse()
	socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		//IP:   net.IPv4(127, 0, 0, 1),
		IP:   net.ParseIP(*serverIp),
		Port: *serverPort,
	})

	defer socket.Close()

	ipString := getIpString()
	senddata := []byte("hi server! " + ipString)
	_, err = socket.Write(senddata)
	if err != nil {
		fmt.Println("send fail !", err)
		return
	}

	data := make([]byte, 100)
	read, remoteAddr, err := socket.ReadFromUDP(data)
	if err != nil {
		fmt.Println("read fail !", err)
		return
	}
	fmt.Println(read, remoteAddr)
	fmt.Printf("%s\n", data)
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
