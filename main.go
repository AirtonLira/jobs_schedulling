package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func hello(name string) {
	message := fmt.Sprintf("Olá, %v", name)
	fmt.Println(message)
}

func runCronJobs() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(5).Seconds().Do(func() {
		hello("Airton")
	})

	s.StartBlocking()
}

func main() {
	runCronJobs()
}
