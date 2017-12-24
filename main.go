package main

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	configFilePath, err := getConfigFilePath()
	if !err {
		fmt.Println(".ssh config file nothing Error")
		os.Exit(1)
	}
	hosts, hasHost := readFile(configFilePath)
	if hasHost {
		fmt.Printf("%s\t%s\n", "number", "Host")
		for i, host := range hosts.Hosts {
			fmt.Printf("%d\t%v\n", i, host.Alias)
		}
		fmt.Printf("%s", "select ssh host number?: ")
		stdIn := bufio.NewScanner(os.Stdin)
		if stdIn.Scan() {
			text := stdIn.Text()
			key, err := strconv.Atoi(text)
			if err != nil {
				fmt.Println("select ssh host number must be integer")
				os.Exit(1)
			}
			cmd := exec.Command("ssh", hosts.Hosts[key].Alias)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err = cmd.Run(); err != nil {
				fmt.Println("can't ssh connection")
				panic(err)
			}
		}
	}
}

func getConfigFilePath() (string, bool) {
	var configFilePath string
	homePath, err := homedir.Dir()
	if err != nil {
		fmt.Println("nothing home dir in host")
		return configFilePath, false
	}
	configFilePath = fmt.Sprintf("%s/.ssh/config", homePath)
	if _, err := os.Stat(configFilePath); err != nil {
		fmt.Println("~/.ssh/config file nothing")
		return configFilePath, false
	}
	return configFilePath, true
}

type Host struct {
	Alias string
}

type HostList struct {
	Hosts []Host
}

func readFile(configFilePath string) (HostList, bool) {
	file, err := os.Open(configFilePath)
	if err != nil {
		fmt.Printf("%s is not exist\n", configFilePath)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReaderSize(file, 4096)
	var hosts HostList
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("can't read config file")
			panic(err)
		}
		host, hasHostObject := objectMapping((string(line)))
		if !hasHostObject {
			continue
		}
		hosts.Hosts = append(hosts.Hosts, host)
	}
	if hosts.Hosts == nil {
		return hosts, false
	}
	return hosts, true
}

func objectMapping(fileArgs string) (Host, bool) {
	var alias Host
	if strings.Contains(fileArgs, "Host ") {
		args := strings.Fields(fileArgs)
		if len(args) != 2 {
			return alias, false
		}
		alias.Alias = args[1]
		return alias, true
	}
	return alias, false
}
