// package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"math"
// 	"net/http"
// 	"strconv"

// 	"math/rand"

// 	"github.com/gliderlabs/ssh"
// )

// type Tunnel struct {
// 	w      io.Writer
// 	doench chan struct{}
// }

// var tunnels = map[int]chan Tunnel{}

// func main() {
// 	ssh.Handle(func(s ssh.Session) {
// 		id := rand.Intn(math.MaxInt)
// 		tunnels[id] = make(chan Tunnel)
// 		fmt.Println("tunnel ID", id)

// 		tunnel := <-tunnels[id]

// 		_, err := io.Copy(tunnel.w, s)

// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		close(tunnel.doench)
// 		s.Write([]byte("We are Done!!!!"))

// 	})

// 	log.Fatal(ssh.ListenAndServe(":2222", nil))
// 	fmt.Println("welcome")
// }

// func handleRequest(w http.ResponseWriter, r *http.Request) {
// 	idstr := r.URL.Query().Get("id")
// 	id, _ := strconv.Atoi(idstr)
// 	tunnel, ok := tunnels[id]
// 	if !ok {
// 		w.Write([]byte("tunnel not found"))
// 		return
// 	}
// 	doench := make(chan struct{})
// 	tunnel <- Tunnel{
// 		w:      w,
// 		doench: doench,
// 	}
// 	<-doench
// }

package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"

	"math/rand"

	"github.com/gliderlabs/ssh"
)

type Tunnel struct {
	w      io.Writer
	doench chan struct{}
}

var tunnels = map[int]chan Tunnel{}

func main() {
	// Set up HTTP server for handling requests
	http.HandleFunc("/tunnel", handleRequest)
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil)) // HTTP server on port 8080
	}()

	// Set up SSH server
	ssh.Handle(func(s ssh.Session) {
		id := rand.Intn(math.MaxInt)
		tunnels[id] = make(chan Tunnel)
		fmt.Println("Tunnel ID:", id)

		tunnel := <-tunnels[id]

		_, err := io.Copy(tunnel.w, s)
		if err != nil {
			log.Printf("Error copying data: %v\n", err)
		}
		close(tunnel.doench)
		s.Write([]byte("We are Done!!!!"))
	})

	err := ssh.ListenAndServe(":2222", nil) // SSH server on port 2222
	if err != nil {
		log.Fatalf("Failed to start SSH server: %v", err)
	}
	fmt.Println("Welcome")
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid tunnel ID", http.StatusBadRequest)
		return
	}

	tunnel, ok := tunnels[id]
	if !ok {
		w.Write([]byte("Tunnel not found"))
		return
	}
	doench := make(chan struct{})
	tunnel <- Tunnel{
		w:      w,
		doench: doench,
	}
	<-doench
}
