crontab
=======

Manage a custom crontab file in the Go programming language

Created as a simple way for Go programs to access the Linux Cron Daemon.

The simplest way to implement this package is to have a cron file in the local program directory that is sym-linked into /etc/cron.d/. The Linux cron daemon will automatically check it every minute for changes.

example/test.go
---------------
`
package main

import (
	"fmt"
	"github.com/knickers/crontab"
)

func main() {
	fmt.Println("Hi")

	cron := crontab.New()

	err := cron.Load("test.cron")
	if err != nil {
		fmt.Println("Cron.Load", err.Error())
	}

	err = cron.AddString("* * * * * ls -al ~/")
	if err != nil {
		fmt.Println("Cron.AddString", err.Error())
	}

	for i, e := range cron.Jobs {
		fmt.Printf("%d) %s\n", i, e)
	}

	err = cron.Save("test.cron")
	if err != nil {
		fmt.Println("Cron.Save", err.Error())
	}

	fmt.Println("Bye")
}
`
