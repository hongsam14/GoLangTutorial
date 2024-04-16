package myLittleAgent

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"example/smEventXml"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type MAgent struct {
	//pipe fd
	stdout io.ReadCloser
	rl     *bufio.Reader
	//process
	tail *exec.Cmd
}

// myLittleAgent is based on tail process
/*
https://stackoverflow.com/questions/24095661/os-exec-sudo-command-in-go/24095983#24095983
*/
func (aObj *MAgent) Init() {
	var err error
	//init cmd
	//aObj.tail = exec.Command("/bin/sh", "-c", "sudo tail -f /var/log/syslog")
	aObj.tail = exec.Command("tail", "-f", "/var/log/syslog")
	//fd redirection
	aObj.tail.Stderr = os.Stderr
	aObj.stdout, err = aObj.tail.StdoutPipe()
	if err != nil {
		log.Fatalf("pipe error :%v\n", err)
	}
	aObj.rl = bufio.NewReader(aObj.stdout)
}

func (aObj *MAgent) Run() {
	fmt.Println("process start")
	err := aObj.tail.Start() //if use run, it will wait until cmd end
	if err != nil {
		log.Fatalf("cmd start error :%v\n", err)
	}
	fmt.Println("decode start")
	decode_log(aObj)
}

// decode log
func decode_log(aObj *MAgent) {
	var (
		//isPrefix bool = true
		bytes []byte
		err   error
		data  *smEventXml.SysmonEventXml
	)

	data = new(smEventXml.SysmonEventXml)
	for {
		fmt.Println("-----------------------------------")
		bytes, _, err = aObj.rl.ReadLine() //read line
		//origin data
		fmt.Println(string(bytes))
		fmt.Println(">>>>>>>>>>>>>>>")
		if err != nil {
			log.Fatalf("readline error :%v\n", err)
			break
		}
		//decode xml
		err = xml.Unmarshal(bytes, data)
		if err != nil {
			log.Fatalf("xml decode error :%v\n", err)
			break
		}
		//fmt.Println(data)
		b, err := json.Marshal(data)
		if err != nil {
			log.Fatalf("json encode error :%v\n", err)
			break
		}
		fmt.Println(string(b))
	}
}
