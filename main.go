package main

import (
	"fmt"

	"github.com/agencyrevolution/go-microservices-example/utils"
)

func main() {
	fmt.Println(utils.NewVulcandClient("", "", 0))
}
