package main

import (
	"flag"
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/pkg/browser"
	"github.com/ttacon/chalk"
	"time"
)

const url = "https://aftonbladet.se"

func main() {
	var hour, minute, second int

	flag.IntVar(&hour, "hr", 0, "Add x hour(s) to timer")
	flag.IntVar(&minute, "m", 0, "Add x minute(s) to timer")
	flag.IntVar(&second, "s", 0, "Add x second(s) to timer")

	flag.Parse()

	sumSec := convertToSeconds(hour, minute, second)

	fmt.Printf("%vh%vm%vs - Total: %vs\n", hour, minute, second, sumSec)

	s := spinner.New(spinner.CharSets[69], 100*time.Millisecond)
	s.Start()
	timer := time.NewTimer(time.Duration(sumSec) * time.Second)
	<-timer.C
	browser.OpenURL(url)
	s.Stop()
	fmt.Println(chalk.Red.Color("TIMEOUT!!!"))
}

func convertToSeconds(h, m, s int) int {
	hourToSec := h * 3600
	minToSec := m * 60

	sum := hourToSec + minToSec + s
	return sum
}
