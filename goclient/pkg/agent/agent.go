package agent

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

//Agent information
type Agent struct {
	OpSys     string
	ShellType string
	ShellFlag string
	IFace     string
	ServerIP  []byte
	MyIP      []byte
}

//Setup - sets up NTP configurations based on OS, sends out first beacon, add firewall rule every 5?
func Setup(newAgent Agent) {
	var commandList []string
	if newAgent.OpSys == "Windows" {
		commandList = []string{
			"net start w32time",
			"w32tm /config /syncfromflags:manual /manualpeerlist:" + string(newAgent.ServerIP),
			"w32tm /config /update",
			"w32tm /resync"} //TODO add firewall rule?
	} else {
		commandList = []string{
			"echo working",
			"echo yes!"} //TODO add actual command
	}

	for _, s := range commandList {
		output, err := exec.Command(newAgent.ShellType, newAgent.ShellFlag, s).Output()

		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Couldn't execute command")
		}

		fmt.Println(string(output))
	}

}

//DetectOS - detects which OS agent is running on
func DetectOS() (string, string, string) {
	sys := "Unknown"
	shell := "temp"
	flag := "temp"
	if runtime.GOOS == "windows" {
		sys = "Windows"
		shell = "cmd"
		flag = "/c"
	} else if runtime.GOOS == "linux" {
		sys = "Linux"
		shell = "/bin/sh"
		flag = "-c"
	} else if runtime.GOOS == "darwin" {
		sys = "macOS"
		shell = "/bin/sh"
		flag = "-c"
	} else {
		fmt.Println("operating system not detected")
		os.Exit(1)
	}

	return sys, shell, flag
}

//GetNetAdapter - gets network interface of agent
func GetNetAdapter(newAgent Agent) string { //TODO there has got to be a better way of doing this
	var iface string
	if runtime.GOOS == "windows" {
		output, err := exec.Command(newAgent.ShellType, newAgent.ShellFlag, "getmac /fo csv /v | findstr \"Ethernet\"").Output() //getting ethernet description for pcap
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Couldn't execute command")
		}
		startIndex := strings.Index(string(output), "_{")
		finalIndex := strings.Index(string(output), "}")

		temp := string(output)[startIndex+2 : finalIndex]
		iface := "\\Device\\NPF_{" + temp + "}"

		return iface
	} else {
		potentials := [4]string{"eth0", "en0", "ens33"}

		devices, err := net.Interfaces()

		if err != nil {
			fmt.Println("error gathering nics")
		}

		iface = "eth0" //default
		for _, device := range devices {
			for i := 0; i < len(potentials); i++ {
				if strings.Contains(strings.ToLower(device.Name), strings.ToLower(potentials[i])) {
					iface = device.Name
					goto End
				}
			}
		}
	}
End:
	return iface
}

//GetServerIP - gets IP address of NTP server
func GetServerIP() []byte {
	input := os.Args[1]
	addr := net.ParseIP(input)

	if addr == nil {
		fmt.Println("Invalid server IP address")
		os.Exit(1)
	}
	return addr
}

//GetMyIP - gets local IP
func GetMyIP() []byte {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP
			}
		}
	}
	return nil
} //function code taken from github.com/emmuanuel/DiscordGo
