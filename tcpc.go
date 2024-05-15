package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arg := os.Args
	if len(arg)==1{
		fmt.Println("Please provide host:port name")
		return
	}
	connect:=arg[1]
	c, err:=net.Dial("tcp", connect)
	if err !=nil{
		fmt.Println(err)
        return
	}
	for{
		reader:=bufio.NewReader(os.Stdin)
		fmt.Print(">>")
		text,_:=reader.ReadString('\n')
		fmt.Fprintf(c,text+"\n")

		message,_:=bufio.NewReader(c).ReadString('\n')
		fmt.Print("->:", message)

		if strings.TrimSpace(string(text))=="STOP"{
			fmt.Println("TCP client exit ...")
			return
		}
	}

}