package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var arg = os.Args[1]
	cmd := exec.Command(arg)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	s := string(out)
	fmt.Println(s)
}
