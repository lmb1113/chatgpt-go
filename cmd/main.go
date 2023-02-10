package main

import (
	"openAi/internal"
	logs "openAi/log"
)

func main() {
	logs.Init()
	for {
		internal.Handle()
	}
}
