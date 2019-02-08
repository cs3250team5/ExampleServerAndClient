package main

import (
	"fmt"
	"net"
	"os"
)

const (
	Port = "8080"
	Host = "localhost"
)

var Messages = [...]string{
	"Help I'm Alive",
<<<<<<< Updated upstream
	"I LOVE SODA!",
=======
	"Latin is cool!",
>>>>>>> Stashed changes
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
	fmt.Fprintf(connection, s)
	connection.Close()
	fmt.Println("Connection Closed")
}
