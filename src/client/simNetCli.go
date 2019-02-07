package main

import (
	"fmt"
	"net"
	"os"
)

const (
	Port        = "8080"
	Host        = "localhost"
	MessageSize = 1024
)

var Messages = [...]string{
	"Help I'm Alive",
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla aliquam malesuada turpis pharetra porttitor. Ut rutrum libero odio, id blandit eros maximus ac. Etiam porttitor luctus mi eu vehicula. Integer est sem, ultricies non diam at, fringilla consectetur tortor. In hac habitasse platea dictumst. Suspendisse eros lectus, gravida eget eros a, rutrum hendrerit libero. Morbi ut enim id elit luctus euismod. Ut aliquam, nisl sit amet venenatis elementum, urna erat viverra nunc, ut maximus ipsum eros nec nunc. Etiam sapien ipsum, mattis eu magna ac, congue blandit velit. In at nibh rhoncus, euismod nisi sit amet, accumsan lectus. Donec eget turpis a velit efficitur cursus ac lobortis metus. Morbi fringilla lacus non sem dapibus ornare.",
	"Sed accumsan et ipsum ut viverra. Nunc eget volutpat libero. Nunc consequat nulla in tortor malesuada sodales. Aliquam faucibus vitae orci sed pretium. Praesent id rutrum justo, quis mattis neque. In cursus consectetur sem, vel commodo massa dignissim vel. Nam elementum mauris ut risus auctor, eu finibus enim rhoncus. Pellentesque scelerisque risus at risus placerat dictum. Sed fermentum pulvinar mollis. Donec egestas scelerisque nulla, accumsan viverra quam sodales vitae. Vestibulum vehicula auctor ex sed auctor.",
	"Morbi sagittis venenatis risus, eget tincidunt ex sollicitudin vitae. Sed ut ex non ex hendrerit laoreet quis at ante. Nunc luctus vehicula libero nec semper. Aliquam erat volutpat. Maecenas vel tristique est, at bibendum mauris. Nam molestie metus nunc, in luctus metus accumsan in. Morbi egestas, turpis ut bibendum dapibus, velit diam commodo dolor, nec sodales enim eros nec nunc. Maecenas finibus ornare feugiat. Aliquam sagittis nunc ut finibus pulvinar. Aliquam efficitur dolor in arcu molestie, quis consequat nisl sagittis. Fusce neque nisi, sagittis vel pellentesque id, porttitor non nulla. Nunc id porttitor sem. Mauris at purus vel sem tempus gravida id sit amet sapien.",
	"Duis ut nunc et velit viverra porttitor eu et risus. Cras non auctor augue, dictum vehicula justo. Suspendisse finibus tristique nisi. Sed non tellus fermentum, aliquam sem et, porta tortor. Phasellus id mauris nisi. Duis sit amet nibh vel sapien faucibus tincidunt id quis nunc. Maecenas fermentum elementum imperdiet.",
}

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Sending Message: ", i)
		send(Messages[i])
	}
}
func connect() net.Conn {
	connection, err := net.Dial("tcp", Host+":"+Port)
	if err != nil {
		fmt.Println("Error in Client main: " + err.Error())
		os.Exit(1)
	}
	fmt.Println("Connected")
	return connection
}
func send(s string) {
	connection := connect()
	defer func() {
		connection.Close()
		fmt.Println("Connection Closed")
	}()

	fmt.Fprintf(connection, s)

	buffer := make([]byte, MessageSize)
	numBytes, err := connection.Read(buffer)
	if err != nil && err.Error() != "EOF" {
		fmt.Println("Error in server:HandleConnection: " + err.Error())
		return
	}

	fmt.Print(string(buffer[:numBytes]) + "\n")

}
