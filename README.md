crontab
=======

Manage a custom crontab file in the Go programming language

Created as a simple way for Go programs to access the Linux Cron Daemon.

The simplest way to implement this package is to have a cron file in the local program directory that is sym-linked into /etc/cron.d/. The Linux cron daemon will automatically check it every minute for changes.
