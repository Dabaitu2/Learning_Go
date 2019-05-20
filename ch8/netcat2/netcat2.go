package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// 只要输入没有终止，这里就会把服务端响应打印出来
	go mustCopy(os.Stdout, conn)
	// 把输入写入连接, 除非遇到command+c/D 这个函数才终止
	mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
