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
