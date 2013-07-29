package crontab

import (
	"bufio"
	"errors"
	"fmt"
	//"io"
	//"log"
	"os"
	"strings"
)

var x = "# This file was automatically generated by crontab.go\n# Do not manually make changes to this file, they will be overwriten\n#\n# For more information visit http://www.github.com/knickers/crontab\n# Also see the Linux manual pages of crontab(5) and cron(8)\n\n"

// This Job struct holds all the cron job time values as specefied in the Linux
// manual pages of crontab(5) and cron(8)
type Job struct {
	Min, Hour, Dom, Month, Dow []string
	Cmd, Comment string
	index *int
}

// Concatinates the Job's time values into a cron job formatted string
func (j *Job) String() string {
	return fmt.Sprint(
		strings.Join(j.Min, ","), " ",
		strings.Join(j.Hour, ","), " ",
		strings.Join(j.Dom, ","), " ",
		strings.Join(j.Month, ","), " ",
		strings.Join(j.Dow, ","), " ",
		j.Cmd,
	)
}

// Return Job's index in the cron array, the index is nil if Job has not yet
// been inserted into a cron array.
func (j *Job) Index() int {
	return *j.index
}

// Hold a collection of cron jobs
type Cron struct {
	Jobs []*Job
}

// Create a new collection of cron jobs
func New() *Cron {
	cron := new(Cron)
	return cron
}

// Parses a cron job formatted string into a Job struct object
func Parse(cmd string) (job *Job, err error) {
	args := strings.Split(cmd, " ")

	if len(args) > 5 {
		job = new(Job)
		job.Min = strings.Split(args[0], ",")
		job.Hour = strings.Split(args[1], ",")
		job.Dom = strings.Split(args[2], ",")
		job.Month = strings.Split(args[3], ",")
		job.Dow = strings.Split(args[4], ",")
		job.Cmd = strings.Join(args[5:], " ")
	} else {
		err = errors.New("Not enough arguments to create a cron job: " + cmd)
	}

	return
}

// Insert a new cron job into the list of jobs as specefied by the object Job.
func (c *Cron) AddJob(job *Job) error {
	job.index = new(int)
	*job.index = len(c.Jobs)
	c.Jobs = append(c.Jobs, job)
	return nil
}

// Insert a new cron job into the list of jobs as specefied by the string cmd.
// Also sets the Job comment.
func (c *Cron) AddString(cmd, comment string) error {
	job, err := Parse(cmd)
	if err != nil {
		return err
	}
	job.Comment = comment
	return c.AddJob(job)
}

func (c *Cron) RemoveJob(job *Job) error {
	fmt.Println("Removing", job)
	return nil
}

// Read in the cron jobs from file. Each cron job must be on it's own line, an
// optional comment can directly follow the cron job also on it's own line.
func (c *Cron) Load(file string) error {
	fd, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		if job := strings.Trim(scanner.Text(), " "); len(job) > 0 {
			if job[0] == '#' && len(c.Jobs) > 0 {
				c.Jobs[len(c.Jobs)-1].Comment += job[1:]
			} else if job[0] != '#' {
				if err = c.AddString(job, ""); err != nil {
					return err
				}
			}
		}
	}
	if err = scanner.Err(); err != nil {
		return err
	}

	return nil
}

// Write out the cron jobs to file. Each cron job will be on it's own line, the
// optional comment will directly follow the cron job also on it's own line.
func (c *Cron) Save(file string) error {
	fmt.Println("Saving to", file)
	read, err := os.Open(file)
	if err != nil {
		return err
	}
	defer read.Close()
	w, err := os.OpenFile(file, os.O_WRONLY, os.FileMode(0644))
	if err != nil {
		return err
	}
	defer w.Close()

	_, err = w.WriteString(x)
	if err != nil {
		return err
	}

	for i := range c.Jobs {
		_, err = w.WriteString(c.Jobs[i].String() + "\n")
		if err != nil {
			return err
		}
		_, err = w.WriteString("#" + c.Jobs[i].Comment + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
