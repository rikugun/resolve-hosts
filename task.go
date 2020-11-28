package main

import (
	"bufio"
	"fmt"
	"github.com/robfig/cron"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func AutoResolveHost() {
	resolveHost()
	//每10分钟自动解析
	spec := "10/* * * * *"
	c := cron.New()
	c.AddFunc(spec, resolveHost)
	c.Start()
	select {}
}

func readServerNames() *bufio.Scanner {
	if servernames, err := os.Open("./servers.txt"); err != nil {
		log.Fatalf("无效的文件 %v", err)
	} else {
		return bufio.NewScanner(servernames)
	}
	return nil
}

func resolveHost() {
	log.Println("AutoResolveHost")
	namesScanner := readServerNames()
	hostsFile, _ := os.Create(filepath.Join("static", "hosts"))

	for namesScanner.Scan() {
		hostname :=  strings.TrimSpace(namesScanner.Text())
		if len(hostname) == 0{
			continue
		}
		if strings.Index(hostname, "#") == 0 {
			hostsFile.WriteString(fmt.Sprintf("%s \n", hostname))
		}else{
			if hosts, err := net.LookupHost(hostname); err == nil {
				for i := 0; i < len(hosts); i++ {
					hostsFile.WriteString(fmt.Sprintf("%s %s \n", hosts[i], hostname))
					fmt.Printf("%s %s \n", hosts[i], hostname)
				}
			}
		}
	}
}
