package main

import (
	"bufio"
	"example/stdlibwithgo"
	"fmt"
	"os"
)

type cliObj struct {
	//fd
	cmdFd *bufio.Reader
	//cmd exitcode
	exitCode int
	//data field
	cmdData string
}

func (cObj *cliObj) Init() {
	cObj.cmdFd = bufio.NewReader(os.Stdin)
	//Default exitcode = 0
	cObj.exitCode = 0
}

func (cObj *cliObj) Loop() {
	var err error

	for cObj.exitCode == 0 {
		var tmp string

		fmt.Print(">>")
		//read cmd
		tmp, err = cObj.cmdFd.ReadString('\n') //include '\n'
		if err != nil {
			cObj.exitCode = 1
		}
		//store data
		cObj.cmdData, _, _ = stdlibwithgo.Cut(tmp, "\n") //use Cut func to clean '\n'
		//logic
		cObj.exitCode = sampleFunc(cObj.cmdData)
	}
}

func sampleFunc(s string) (exitcode int) {
	if s == "" {
		fmt.Println("empty string")
		return 1
	}
	fmt.Println(s)
	return 0
}

func main() {
	var (
		cObj *cliObj
	)

	//allocate
	cObj = new(cliObj)
	cObj.Init()
	//run loop
	cObj.Loop()
}
