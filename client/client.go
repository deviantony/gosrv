package main // import "github.com/deviantony/goclient"

import "net"
import "fmt"
import "bufio"
import "os"
import "bytes"

const (
	server = "localhost:7777"
)

func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp", server)

	go handler(conn)

	fmt.Print("Start talking with the server: ")
	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text)
	}
}

func handler(conn net.Conn) {
	for {
		// listen for reply
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("ERROR: something wrong occured:", err.Error())
			return
		}
		n := bytes.Index(buf, []byte{0})
		fmt.Print("echo: " + string(buf[:n]))
	}
}
