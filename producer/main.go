package main

import (
	"fmt"

	workers "github.com/digitalocean/go-workers2"
	"github.com/eiannone/keyboard"
)

func main() {
	producer, err := workers.NewProducer(workers.Options{
		ServerAddr: "localhost:6379",
		ProcessID:  "2",
	})
	if err != nil {
		panic(err)
	}
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	scheduledIndex := 0

	go workers.StartAPIServer(8080)

	fmt.Println("Press ESC to quit")
	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeySpace {
			fmt.Println("hello")
			producer.Enqueue("default", "myJob", []int{scheduledIndex})
		}
		if key == keyboard.KeyEsc || key == keyboard.KeyCtrlC {
			break
		}
		scheduledIndex++
	}
}
