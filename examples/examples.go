package main

import (
	"fmt"
	"github.com/jagipson/sgr"
)

func main() {
	ok := sgr.Text{"OKAY", sgr.Style{sgr.ForeGreen}}
	warn := sgr.Text{"Danger!", sgr.Style{sgr.ForeYellow, sgr.BackRed}}
	crit := sgr.Text{"CRITICAL", sgr.Style{sgr.ForeBlack, sgr.BackRed, sgr.Bold}}

	fmt.Printf("%s\n%s\n%s\n", ok, warn, crit)
	fmt.Printf("%a\n%a\n%a\n", ok, warn, crit)

}
