package main

import (
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

func main() {
	server_ip := ""
	port := "8080"
	server_address := server_ip + ":" + port

	/* Listen */
	listner, err := net.Listen("tcp", server_address)
	if err != nil {
		fmt.Println("[+] Unable to listen")
		panic(err)
	}

	/* Accept */
	server_conn, err := listner.Accept()
	if err != nil {
		fmt.Println("[+] Unable to accept")
		panic(err)
	}
	client_ip := server_conn.RemoteAddr().String()

	fmt.Println("[+] Client IP is ", client_ip)

	decoder := gob.NewDecoder(server_conn)

	fs := &FileStruct{}
	decoder.Decode(fs)

	file, err := os.Create("RecievedFile.txt")
	if err != nil {
		fmt.Println("[+] unable to create file")
	}

	file.Write(fs.FileContent)
	file.Write(([]byte)("Wrote by Server"))
	file.Close()

}
