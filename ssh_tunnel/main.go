package main

import (
	"fmt"
	"io"
	"log"
	"math"

	"math/rand"

	"github.com/gliderlabs/ssh"
)

type Tunnel struct{
	w io.Writer
	doench chan struct{}
}

var tunnels = map[int]chan Tunnel{}

func main (){
	ssh.Handle(func(s ssh.Session) {
		id := rand.Intn(math.MaxInt)
		tunnels[id] = make(chan Tunnel)
		fmt.Println("tunnel ID", id)

		tunnel  := <- tunnels[id]

		_, err := 	io.Copy(tunnel.w, s)

		if err != nil {
			log.Fatal(err)
		}
		close(tunnel.doench)
		s.Write([]byte("We are Done!!!!"))

	})  

	log.Fatal(ssh.ListenAndServe(":2222", nil))
	fmt.Println("welcome")
}
