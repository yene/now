package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	currentTime := time.Now()
	text := fmt.Sprintf("%s", currentTime.Format(time.RFC3339))
	fmt.Println(text)

	if runtime.GOOS == "darwin" {
		err := writePasteboard(text)
		if err != nil {
			panic(err)
		}
		fmt.Println("copied to pasteboard")
	}
}
