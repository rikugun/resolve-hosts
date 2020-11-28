package main

import (
	"bufio"
	"fmt"
	"log"
	"github.com/robfig/cron"
	"net"
	"os"
	"path/filepath"
)

func AutoResolveHost() {
	//每10分钟自动解析
	spec := "10/* * * * *"
	c := cron.New()
	c.AddFunc(spec, resolveHost)
	c.Start()
	log.Println("AutoResolveHost")
	select {}
}



func cwd() string {
	path, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Dir(path)
}

func readServerNames() *bufio.Scanner {
	fmt.Printf("resolve file %s", filepath.Join(cwd(), "servers.txt \n"))
	if servernames, err := os.Open(filepath.Join(cwd(), "servers.txt")); err == nil {
		return bufio.NewScanner(servernames)
	}
	return nil
}

func resolveHost()  {
	namesScanner := readServerNames()
	hostsFile, _ := os.Create(filepath.Join(cwd(), "static","hosts"))

	for namesScanner.Scan() {
		hostname :=namesScanner.Text()
		if hosts, err := net.LookupHost(hostname);err ==nil{
			for i := 0; i < len(hosts); i++ {
				hostsFile.WriteString(fmt.Sprintf("%s %s \n",hosts[i],hostname))
				fmt.Printf("%s %s \n",hosts[i],hostname)
			}
		}
	}
}
