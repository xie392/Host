package main

import (
	"fmt"
	"github.com/xie392/restful-api/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
