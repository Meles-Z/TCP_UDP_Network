package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments)==1{
		fmt.Println("Please provide host:port string")
		return 
	}

	connect:=arguments[1]
	s, err:=net.ResolveUDPAddr("udp4", connect)
	c, err:=net.DialUDP("udp4", nil, s)
	if err !=nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("The udp server is %s\n", c.RemoteAddr().String())
	defer c.Close()

	for{
		reader:=bufio.NewReader(os.Stdin)
		fmt.Print(">>")

		text, _:=reader.ReadString('\n')
		data:=[]byte(text+"\n")
		_, err:=c.Write(data)
		if strings.TrimSpace(string(data))=="STOP"{
			fmt.Println("Exiting client UDP")
			return
		}
		if err !=nil{
			fmt.Println(err)
			return
		}

		buffer:=make([]byte, 1024)
		n, _,err:=c.ReadFromUDP(buffer)
		if err !=nil{
			fmt.Println(err)
			return

		}
		fmt.Printf("Reply: %s\n", string(buffer[0:n]))

	}
}