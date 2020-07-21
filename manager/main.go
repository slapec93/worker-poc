package main

import (
	"fmt"
	"os"
	"time"

	workers "github.com/digitalocean/go-workers2"
)

func myJob(message *workers.Msg) error {
	fmt.Printf("PID: %d\n", os.Getpid())
	fmt.Println(message.Args().MustArray())
	time.Sleep(5 * time.Second)
	fmt.Println("sup dude")
	return nil
}

func main() {
	fmt.Printf("PID: %d\n", os.Getpid())
	manager, err := workers.NewManager(workers.Options{
		// location of redis instance
		ServerAddr: "localhost:6379",
		// instance of the database
		Database: 0,
		// number of connections to keep open with redis
		PoolSize: 30,
		// unique process id for this instance of workers (for proper recovery of inprogress jobs on crash)
		ProcessID: "1",
	})

	if err != nil {
		fmt.Println(err)
	}

	manager.AddWorker("default", 10, myJob)

	manager.Run()
}
