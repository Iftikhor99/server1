package server

import (
	"bytes"
	"io"
	"runtime"
	"strconv"
	"strings"

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
	//var wg sync.WaitGroup
	log.Println(len(s.handlers))

	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Print(err)
		return err
	}
	if len(s.handlers) == 0 {
		listener.Close()
		log.Print("if no handlerrrrrrrrrrrrr")
		return err
	}
	defer func() {
		//	wg.Wait()
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
			conn.Close()
			continue
		}

		//		for i := 0; i < len(s.handlers); i++ {

		go s.handle(conn)
		//	wg.Done()
		log.Println("number of gorutines: ", runtime.NumGoroutine())

		//		}

		// if err != nil {
		// 	log.Print(err)
		// 	continue
		// }

	}

	//	wg.Wait()

	return nil

}

func (s *Server) handle(conn net.Conn) {
	var err error
	//var wg sync.WaitGroup
	//wg.Add(1)
	//mu := s.mu
	//	wg := sync.WaitGroup{}
	//	defer wg.Done()
	//wg.Add(1)
	defer func() {
		//	wg.Done()
		if cerr := conn.Close(); cerr != nil {
			if err == nil {
				err = cerr
				log.Print(err)
			}
			log.Print(err)
		}
	}()
	for {
		buf := make([]byte, 4096)

		n, err := conn.Read(buf)
		if n == 0 || err == io.EOF {
			log.Printf("%s", buf[:n])
			log.Print(err)
			log.Print("We are hereeeeeeeeeee")
		//	conn.Close()
			break
		}
		// if err != nil {
		// 	log.Print(err)
		// }
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
		// if path == "/" {

		// }

		if path == "/" {
			s.mu.RLock()
			//		handler := s.handlers["/"]
			//
			s.mu.RUnlock()
			//		handler(conn)

			//	wg.Done()
		}
		//log.Print(npm)
		body := "Ok!"
		//	body, err := ioutil.ReadFile("static/index.html")

		_, err = conn.Write([]byte(
			"HTTP/1.1 200 OK\r\n" +
				"Conect-Length: " + strconv.Itoa(len(body)) + "\r\n" +
				"Content-Type: text/html	\r\n" +
				"Connection: close\r\n" +
				"\r\n" +
				string(body),
		))
		if err != nil {
			log.Print(err)
		}

		if path == "/about" {
			s.mu.RLock()
			//		handler := s.handlers["/about"]
			//		handler(conn)
			s.mu.RUnlock()
			//		handler(conn)
			//	wg.Done()
		}
	}
	//	wg.Wait()
}
