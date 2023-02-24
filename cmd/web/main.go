package main

import (
	"fmt"
	"time"

	"github.com/iamhi/cloudy-memory-go/src/core/bootservice"
	"github.com/iamhi/cloudy-memory-go/src/http/routes"
)

func main() {
	fmt.Println("Testing")

	bootservice.BootUp()

	time.Sleep(2 * time.Second)

	routes.Setup()
}
