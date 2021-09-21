/*
 * @Author: your name
 * @Date: 2021-08-03 12:40:56
 * @LastEditTime: 2021-09-09 09:12:11
 * @LastEditors: your name
 * @Description: In User Settings Edit
 * @FilePath: /digital_twin_fedavg_main_sort/server.go
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

var urlName = "localhost:%d"

type Server struct {
	node *Node
	url	string
}

func nodeIdToPort(nodeId int) int{
	return nodeId + 9080
}

func NewServer(nodeId int) *Server{
	server :=  &Server{
		NewNode(nodeId),
		fmt.Sprintf(urlName, nodeIdToPort(nodeId)),
	}
	return server
}

func (s *Server) Start(){
	s.node.Start()
	ln, err := net.Listen("tcp", s.url)
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	fmt.Printf("server start at %s\n",s.url)
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn){
	req, err := ioutil.ReadAll(conn)
	if err != nil {
		panic(err)
	}
	s.node.msgQueue <- req
}
