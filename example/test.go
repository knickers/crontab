// A simple example crontab implementation
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

	err = cron.AddString("* * * * * ls -al ~/", `{"name":"John","age":"27"}`)
	if err != nil {
		fmt.Println("Cron.AddString", err.Error())
	}

	for i, e := range cron.Jobs {
		fmt.Printf("%d) %s\n", i, e)
		fmt.Println("    " + e.Comment)
	}

	fmt.Println()
	cron.RemoveJob(cron.Jobs[len(cron.Jobs)-1])

	for i, e := range cron.Jobs {
		fmt.Printf("%d) %s\n", i, e)
		fmt.Println("    " + e.Comment)
	}

	err = cron.Save("test.cron")
	if err != nil {
		fmt.Println("Cron.Save", err.Error())
	}

	fmt.Println("Bye")
}
