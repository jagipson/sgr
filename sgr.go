// sgr utilizes features in fmt to add ANSI SGR codes to strings.
package sgr

import (
	"fmt"
	"math"
)

type SGR string

var Config struct {
	Depth int
}

func init() {
	Config.Depth = 24
}

const (
	csi             = string(rune(27)) + string('[')
	end             = "m"
	Reset       SGR = "0"
	Bold        SGR = "1"
	Faint       SGR = "2"
	Italic      SGR = "3"
	Underline   SGR = "4"
	BlinkSlow   SGR = "5"
	BlinkFast   SGR = "6"
	Inverse     SGR = "7"
	Conceal     SGR = "8" // not widely supported
	Strike      SGR = "9" // not widely supported
	Font0       SGR = "10"
	Font1       SGR = "11"
	Font2       SGR = "12"
	Font3       SGR = "13"
	Font4       SGR = "14"
	Font5       SGR = "15"
	Font6       SGR = "16"
	Font7       SGR = "17"
	Font8       SGR = "18"
	Font9       SGR = "19"
	Fractur     SGR = "20" // Blackletter
	BOUD        SGR = "21" // Bold off or Underline Double - not widely supported
	Normal      SGR = "22" // normal intensity (Bold off, Faint off)
	IOFO        SGR = "23" // Italic off, Fractur off
	UO          SGR = "24" // underline off
	Steady      SGR = "25" // (no blink)
	Obverse     SGR = "27" // (no inverse)
	Reveal      SGR = "28" // (no conceal)
	Unstrike    SGR = "29"
	ForeBlack   SGR = "30"
	ForeRed     SGR = "31"
	ForeGreen   SGR = "32"
	ForeYellow  SGR = "33"
	ForeBlue    SGR = "34"
	ForeMagenta SGR = "35"
	ForeCyan    SGR = "36"
	ForeWhite   SGR = "37"
	ForeExtd    SGR = "38"
	ForeDefault SGR = "39"
	BackBlack   SGR = "40"
	BackRed     SGR = "41"
	BackGreen   SGR = "42"
	BackYellow  SGR = "43"
	BackBlue    SGR = "44"
	BackMagenta SGR = "45"
	BackCyan    SGR = "46"
	BackWhite   SGR = "47"
	BackExtd    SGR = "48"
	BackDefault SGR = "49"
	Framed      SGR = "51"
	Circled     SGR = "52"
	Overline    SGR = "53"
	NFNC        SGR = "54" // Not framed, not circled
	NoOverline  SGR = "55"
)

type Style []SGR

func (s Style) String() string {
	codes := ""
	for i := range s {
		if i > 0 {
			codes += ";"
		}
		codes += string(s[i])
	}
	return string(csi) + codes + string(end)
}

type Text struct {
	Value string
	Style Style
}

func (t Text) String() string {
	return t.Value
}

func (t Text) Format(f fmt.State, c rune) {
	pre := ""
	post := ""
	newRune := c
	if c == 'a' {
		pre = t.Style.String()
		post = string(csi) + string(Reset) + string(end)
		newRune = 's'
	}
	format := "%"
	if width, ok := f.Width(); ok {
		format += fmt.Sprintf("%d", width)
	}
	if precision, ok := f.Precision(); ok {
		format += fmt.Sprintf(".%d", precision)
	}
	format += string(newRune)

	f.Write([]byte(fmt.Sprintf(format, pre+t.Value+post)))
}

// representation of a 24 bit color
type Color24 struct {
	R byte
	G byte
	B byte
}

// reduce nearest color in new color depth
func (c Color24) To256SGR(background bool) SGR {
	rgd := math.Abs(float64(c.R - c.G))
	rbd := math.Abs(float64(c.R - c.B))
	gbd := math.Abs(float64(c.G - c.B))

	var fgbg SGR
	if background {
		fgbg = "48;5;"
	} else {
		fgbg = "38;5;"
	}

	// Use the sum of the differences to decide whether C is a color or a grey
	if rgd+rbd+gbd > 8 {
		r := int((float32(c.R)/255*5)+0.5) * 36
		g := int((float32(c.G)/255*5)+0.5) * 6
		b := int((float32(c.B) / 255 * 5) + 0.5)
		return SGR(fmt.Sprintf("%s%d", fgbg, 16+r+g+b))
	} else {
		// process as a grey.
		// average the rgb values and match nearest:
		result := greyOrder256[int((float32(c.R)+float32(c.G)+float32(c.B))/3/255*30+0.5)]
		return SGR(fmt.Sprintf("%s%d", fgbg, result))
	}
}

// reduce nearest color in new color depth
func (c Color24) To16SGR(background bool) SGR {
	var fgbg int
	if background {
		fgbg = 40
	} else {
		fgbg = 30
	}

	r := float64(c.R) / 127
	g := float64(c.G) / 127
	b := float64(c.B) / 127

	var s SGR
	var result int

	switch {
	case r > 1 && g > 1 && b > 1:
		result = 7
	case r > 0 && g > 1 && b > 1:
		result = 6
	case r > 1 && g > 0 && b > 1:
		result = 5
	case r > 0 && g > 0 && b > 1:
		result = 4
	case r > 1 && g > 1 && b > 0:
		result = 3
	case r > 0 && g > 1 && b > 0:
		result = 2
	case r > 1 && g > 0 && b > 0:
		result = 1
	case r > 0 && g > 0 && b > 0:
		result = 0
	}

	s += SGR(fmt.Sprintf("%d", result+fgbg))
	avg := (r + g + b) / 3.0
	if avg > 1.1 && !background {
		s += ";1"
	}
	return s
}

func (c Color24) ToRealSGR(background bool) SGR {
	var result SGR
	var fgbg SGR
	if background {
		fgbg = "48"
	} else {
		fgbg = "38"
	}
	result = SGR(fmt.Sprintf("%s;2;%d;%d;%d", fgbg, c.R, c.G, c.B))
	return result
}

var greyOrder256 = [...]int{
	16,
	232,
	233,
	234,
	235,
	236,
	237,
	238,
	239,
	240,
	59,
	241,
	242,
	243,
	244,
	102,
	245,
	246,
	247,
	248,
	139,
	145,
	249,
	250,
	251,
	252,
	188,
	253,
	254,
	255,
	231,
}

func (c Color24) String() string {
	return fmt.Sprintf("#%0.2x%0.2x%0.2x", c.R, c.G, c.B)
}
