package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct{
	err error
	output []byte
}

func main(){
	var (
		ctx context.Context
		cancelfunc context.CancelFunc
		cmd *exec.Cmd
		resChan chan *result
		res *result
	)

	resChan = make(chan *result,1000)

	ctx,cancelfunc = context.WithCancel(context.Background())

	go func(){
		var (
			err error
			output []byte
		)
		cmd = exec.CommandContext(ctx,"C:\\cygwin64\\bin\\bash.exe","-c","sleep 2;echo hello;")
		output,err = cmd.CombinedOutput()

		resChan <- &result{
			err:    err,
			output: output,
		}
	}()
	//脚本两秒才执行 这里 休眠1秒 然后取消协程，故协程没机会执行完 就被杀死
	time.Sleep(time.Second)

	cancelfunc()

	res = <- resChan

	fmt.Println(res.err,string(res.output))
}
