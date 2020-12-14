package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

type FileStruct struct {
	FileName    string
	FileSize    int
	FileContent []byte
}

func ReadFileContent(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	info, err := file.Stat()

	filesize := info.Size()
	buffer := make([]byte, filesize)
	reader := bufio.NewReader(file)

	reader.Read(buffer)

	return buffer, err
}

func main() {
	ip := "127.0.0.1"
	port := "8080"

	address := ip + ":" + port

	client_conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("[+] Unable to connect")
		panic(err)
	}

	server_ip := client_conn.RemoteAddr().String()

	fmt.Println("[+] Server IP is ", server_ip)
	file_name := "/Users/root1/Desktop/Development/goworkspace/src/Golan-Playground/Network/test_file"
	content, err := ReadFileContent(file_name)
	fs := &FileStruct{
		FileName:    file_name,
		FileSize:    len(content),
		FileContent: content,
	}

	encoder := gob.NewEncoder(client_conn)
	encoder.Encode(fs)

}
