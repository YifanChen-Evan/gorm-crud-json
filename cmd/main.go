package main

import (
	"fmt"

	"github.com/YifanChen-Evan/gorm-read-json/cli"
)

func main() {
	cli := cli.CLI{} // create an instance of CLI
	cli.Run()

	fmt.Println("--------- End ---------")
}
