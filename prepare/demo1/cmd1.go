package main

import (
	"fmt"
	"os/exec"
)

func main(){
	var (
		cmd *exec.Cmd
		err error
		output []byte
	)
	//go 执行命令行
	cmd = exec.Command("C:\\cygwin64\\bin\\bash.exe","-c","echo hello")
	if output,err = cmd.Output(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
}
