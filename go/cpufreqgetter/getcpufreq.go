package main

import (
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strings"
)

var (
	supportedOSArch = map[string]bool{
		"linux.amd64":   true,
		"linux.ppc64le": true,
		"linux.s390x":   true,
	}
)

func main() {
	hostName, _ := os.Hostname()
	osArch := runtime.GOOS + "." + runtime.GOARCH
	if _, ok := supportedOSArch[osArch]; !ok {
		fmt.Printf("The OSArch[%v] on the node[%v] is not supported!", osArch, hostName)
		return
	}

	cpuinfo, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		fmt.Printf("Can't open the cpu proc file on the node[%v]:%v", hostName, err)
	}

	switch osArch {
	case "linux.amd64", "linux.s390x":
		re, err := regexp.Compile(`cpu MHz.*`)
		if err != nil {
			fmt.Printf("Can't parse regular expression for the node[%v]: %v", hostName, err)
			return
		}
		cpufreq := re.FindString(string(cpuinfo))
		str := strings.Split(cpufreq, ":")
		if len(str) != 2 {
			fmt.Printf("Get an invalid cpu frequency string[%v] for the node[%v]", cpufreq, hostName)
			return
		}
		fmt.Printf(strings.TrimSpace(str[1]))
	case "linux.ppc64le":
		re, err := regexp.Compile(`clock.*`)
		if err != nil {
			fmt.Printf("Can't parse regular expression for the node[%v]: %v", hostName, err)
			return
		}
		cpufreq := re.FindString(string(cpuinfo))
		str := strings.Split(cpufreq, ":")
		if len(str) != 2 {
			fmt.Printf("Get an invalid cpu frequency string[%v] for the node[%v]", cpufreq, hostName)
			return
		}
		clock := strings.TrimSpace(str[1])
		// TODO: I cannot find the specification of /proc/cpuinfo output for Power linux. It is not clear if the MHz unit
		//   could potentially be different, such as GHz, on different node. For now, assume the unit is always MHz
		fmt.Printf(strings.TrimSuffix(clock, "MHz"))
	}

}
