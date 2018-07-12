package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/Eiryyy/agrec/record"
	"github.com/robfig/cron"
)

type config struct {
	Programs []program
}

type program struct {
	Title  string
	Cron   string
	Minute int
}

func main() {
	var c config
	if _, err := toml.DecodeFile("programs.toml", &c); err != nil {
		fmt.Println("toml:", err)
		return
	}
	location, _ := time.LoadLocation("Asia/Tokyo")
	cron := cron.NewWithLocation(location)
	for _, p := range c.Programs {
		title := p.Title
		minute := p.Minute
		job := func() {
			record.Do(title, minute)
		}
		cron.AddFunc(p.Cron, job)
		fmt.Println("Recording " + p.Title + " is scheduled")
	}
	cron.Start()
	runtime.Goexit()
}
