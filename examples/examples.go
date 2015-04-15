package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jagipson/sgr"
)

func main() {
	//ok := sgr.Text{"OKAY", sgr.Style{sgr.ForeGreen}}
	//warn := sgr.Text{"Danger!", sgr.Style{sgr.ForeYellow, sgr.BackRed}}
	//crit := sgr.Text{"CRITICAL", sgr.Style{sgr.ForeBlack, sgr.BackRed, sgr.Bold}}

	//fmt.Printf("%s\n%s\n%s\n", ok, warn, crit)
	//fmt.Printf("%a\n%a\n%a\n", ok, warn, crit)

	//pink := sgr.Color24{0xff, 0x53, 0xaa}

	//pinkStyle1 := sgr.Style{pink.To256SGR(false)}
	//t := sgr.Text{"and the brain", pinkStyle1}
	//fmt.Printf("%a\n", t)

	//pinkStyle1 = sgr.Style{pink.ToRealSGR(false), sgr.BlinkSlow}
	//t = sgr.Text{"and the brain", pinkStyle1}
	//fmt.Printf("%a\n", t)

	//pinkStyle1 = sgr.Style{pink.ToRealSGR(false), sgr.Fractur}
	//t = sgr.Text{"and the brain", pinkStyle1}
	//fmt.Printf("%a\n", t)

	rand.Seed(time.Now().UTC().UnixNano())
	// colors test
	for x := 0; x < 20; x++ {
		color := sgr.Color24{byte(rand.Intn(256)),
			byte(rand.Intn(256)),
			byte(rand.Intn(256)),
		}
		hex := color.String()

		loStyle := sgr.Style{color.To16SGR(false)}
		mdStyle := sgr.Style{color.To256SGR(false)}
		hiStyle := sgr.Style{color.ToRealSGR(false)}

		lo := sgr.Text{"LOW COLOR", loStyle}
		med := sgr.Text{"MED COLOR", mdStyle}
		hi := sgr.Text{"HIGH COLOR", hiStyle}

		fmt.Printf("%s\t%a\t%a\t%a\n", hex, lo, med, hi)

	}
}
