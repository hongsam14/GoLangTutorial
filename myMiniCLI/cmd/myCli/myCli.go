package main

import (
	"bufio"
	ml "example/myLittleAgent"
	"example/stdlibwithgo"
	"fmt"
	"os"
)

type myCli struct {
	//fd
	cmdFd *bufio.Reader
	//cmd exitcode
	exitCode int
	//data field
	cmdData       string
	cmdParsedData []string
}

func (cObj *myCli) Init() {
	cObj.cmdFd = bufio.NewReader(os.Stdin)
	//Default exitcode = 0
	cObj.exitCode = 0
}

func (cObj *myCli) Loop() {
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
		cObj.exitCode = simpleExec(cObj.cmdData)
	}
}

func simpleExec(s ...string) (exitcode int) {
	if s[0] == "" {
		fmt.Println("empty string")
		return 1
	}
	fmt.Println(s)
	return 0
}

func main() {
	var m *ml.MAgent

	/*
		cObj = new(myCli)
		cObj.Init()
		cObj.Loop()
	*/
	m = new(ml.MAgent)
	m.Init()
	m.Run()
}