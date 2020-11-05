package server

import (
	"bytes"
	"io"
	"strings"
//	"strconv"
//	"fmt"
	"log"
	"net"
	"sync"
	//“sync”
)

//HandlerFunc for
type HandlerFunc func(conn net.Conn)

//Server for
type Server struct {
	addr string

	mu sync.RWMutex

	handlers map[string]HandlerFunc
}


//NewServer for
func NewServer(addr string) *Server {
	return &Server{addr: addr, handlers: make(map[string]HandlerFunc)}

}

//Register for
func (s *Server) Register(path string, handler HandlerFunc) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers[path] = handler

}


//Start for
func (s *Server) Start() error {
	// TODO: start server on host & port
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Print(err)
		return err
	}
	defer func() {
		if cerr := listener.Close(); cerr != nil {

			if err == nil {
				err = cerr
				return
			}
			log.Print(cerr)
		}
	}()
	// TODO: server code

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		
		go s.handle(conn)
		// if err != nil {
		// 	log.Print(err)
		// 	continue
		// }

	}

	

//	return nil

}

func (s *Server) handle(conn net.Conn) {
	var err error
	//mu := s.mu
	defer func() {
		if cerr := conn.Close(); cerr != nil {
			if err == nil {
				err = cerr
				log.Print(err)
			}
			log.Print(err)
		}
	}()

	buf := make([]byte, 4096)

	n, err := conn.Read(buf)
	if err == io.EOF {
		log.Printf("%s", buf[:n])
		log.Print(err)
	}
	if err != nil {
		log.Print(err)
	}
	log.Printf("%s", buf[:n])

	data := buf[:n]
	requestLineDeLim := []byte{'\r', '\n'}
	requestLineEnd := bytes.Index(data, requestLineDeLim)
	if requestLineEnd == -1 {

	}

	requestLine := string(data[:requestLineEnd])
	parts := strings.Split(requestLine, " ")
	if len(parts) != 3 {

	}

	method, path, version := parts[0], parts[1], parts[2]

	if method != "GET" {

	}

	if version != "HTTP/1.1" {

	}

	if path == "/" {
		s.mu.RLock()
		handler := s.handlers["/"]
		s.mu.RUnlock()
		handler(conn)
	// 	//log.Print(npm)
	// 	body := "Ok!"
	// //	body, err := ioutil.ReadFile("static/index.html")
	// 	if err != nil {
	// 		return fmt.Errorf("can't read index.html: %w", err)
	// 	}
	// 	_, err = conn.Write([]byte(
	// 		"HTTP/1.1 200 OK\r\n" +
	// 			"Conect-Length: " + strconv.Itoa(len(body)) + "\r\n" +
	// 			"Content-Type: text/html	\r\n" +
	// 			"Connection: close\r\n" +
	// 			"\r\n" +
	// 			string(body),
	// 	))
	// 	if err != nil {
	// 		return err
	// 	}
	}
	
	if path == "/about" {
		s.mu.RLock()
		handler := s.handlers["/about"]
		s.mu.RUnlock()
		handler(conn)
	}

	
}