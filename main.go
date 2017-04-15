package main

import (
	"os/exec"
	"os"
	"fmt"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "<Command or program to run>")
	}
	args := ""
	for _, k := range os.Args[2:] {
		args += k
	}
	pre := time.Now()
	cmd := exec.Command(os.Args[1], args)
	cmd.Run()
	fmt.Println("Time used:", time.Now().Sub(pre).String())
}
