// sgr utilizes features in fmt to add ANSI SGR codes to strings.
package sgr

import (
	"fmt"
)

type SGR string

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

// ForeRGB is not supported by all terms. I know it works with Xterm
func ForeRGB(r, g, b byte) SGR {
	return SGR(fmt.Sprintf("%d;2;%d;%d;%d", ForeExtd, r, g, b))
}

// TODO: write a version that can snap to 255 color and 16 color, 8 color based
// on color depth variable

// BackRGB is not supported by all terms. I know it works with Xterm
func BackRGB(r, g, b byte) SGR {
	return SGR(fmt.Sprintf("%d;2;%d;%d;%d", BackExtd, r, g, b))
}

func ForeCode(code byte) SGR {
	return SGR(fmt.Sprintf("%d;5;%d", ForeExtd, code))
}

func BackCode(code byte) SGR {
	return SGR(fmt.Sprintf("%d;5;%d", BackExtd, code))
}

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
