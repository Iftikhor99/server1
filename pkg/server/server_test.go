package server

import (
	"strconv"
	"log"
//	"os"
//	"fmt"
//	"reflect"
	"testing"
	"net"
	
)

// func BenchmarkSerever(b *testing.B) {
	
// 	host := "0.0.0.0"
// 	port := "9999"

// 	srv := NewServer(net.JoinHostPort(host, port))

// 	srv.Register("/", func(conn net.Conn) {
// 		body := "Welcome to our web-site"

// 		_, err := conn.Write([]byte(
// 			"HTTP/1.1 200 OK\r\n" +
// 				"Content-Length: " + strconv.Itoa(len(body)) + "“\r\n" +
// 				"Content-Type: text/html\r\n" +
// 				"Connection: close\r\n" +
// 				"\r\n" +
// 				body,
// 		))
// 		if err != nil {
// 			log.Print(err)
// 		}
// 	})
// 	srv.Register("/about", func(conn net.Conn) {
// 		body := "About Golang Academy"

// 		_, err := conn.Write([]byte(
// 			"HTTP/1.1 200 OK\r\n" +
// 				"Content-Length: “ + strconv.Itoa(len(body)) + “\r\n" +
// 				"Content-Type: text/html\r\n" +
// 				"Connection: close\r\n" +
// 				"\r\n" +
// 				body,
// 		))
// 		if err != nil {
// 			log.Print(err)
// 		}
// 	})
	
// 	for i := 0; i < b.N; i++ {

// 	errN := srv.Start()

// 	conn, err := net.Dial("tcp", net.JoinHostPort(host, port)) 
//     if err != nil { 
//         log.Println(err) 
        
//     } 
//     defer conn.Close() 

// 	if errN != nil {
// 		log.Print(errN)
// 	}
// }
// }