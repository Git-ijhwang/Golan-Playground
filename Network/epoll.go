package main

import (
	"fmt"
	"net"
	"os"
)
import "syscall"

const (
	EPOLLET        = 1 << 31
	MaxEpollEvents = 32
)

type epoll struct {
	fd int
}

func initListenFd(ip string, port int) (int, error) {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.O_NONBLOCK|syscall.SOCK_STREAM, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer syscall.Close(fd)

	if err = syscall.SetNonblock(fd, true); err != nil {
		fmt.Println("setnonblock1: ", err)
		os.Exit(1)
	}

	addr := syscall.SockaddrInet4{Port: port}
	copy(addr.Addr[:], net.ParseIP(ip).To4())

	syscall.Bind(fd, &addr)
	syscall.Listen(fd, 10)

	return fd, err
}

func initEpoll() (epoll, error) {
	epfd, err := syscall.EP
}

func main() {
	var listenFd int
	var err error

	ip := "127.0.0.1"

	listenFd, err = initListenFd(ip, 8080)
	if err != nil {
		panic(err)
	}
	defer syscall.Close(listenFd)

	var ep epoll
	ep, err = initEpoll()
	if err != nil {
		panic(err)
	}
}
