package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min)+min
}

func main(){
	argument:=os.Args
	if len(argument)==1{
		fmt.Println("Please provide port number")
		return
	}
	PORT:=":"+argument[1]
	s, err:=net.ResolveUDPAddr("udp4", PORT)
	if err !=nil{
		fmt.Println(err)
		return
	}
	connection, err:=net.ListenUDP("udp4", s)
	if err !=nil{
		fmt.Println(err)
		return
	}
	defer connection.Close()
	buffer:=make([]byte, 1024)
	rand.Seed(time.Now().Unix())
	for{
		n, add, err:=connection.ReadFromUDP(buffer)
		fmt.Println("-> ", string(buffer[0:n-1]))

		if strings.TrimSpace(string(buffer[0:n]))=="STOP"{
			fmt.Println("UDP server is exited!")
			return
		}
		data:=[]byte(strconv.Itoa(random(1, 1001)))
		fmt.Printf("data: %s\n", string(data))
		_, err=connection.WriteToUDP(data, add)
		if err !=nil{
			fmt.Println(err)
			return
		}
	}
}