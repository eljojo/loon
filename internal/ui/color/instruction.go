package color

import "fmt"

type InstructionType int

const (
	Reset InstructionType = iota
	Bold
	Faint
	Italic
	Underline
)

const (
	Black InstructionType = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

const (
	BrightBlack InstructionType = iota + 90
	BrightRed
	BrightGreen
	BrightYellow
	BrightBlue
	BrightMagenta
	BrightCyan
	BrightWhite
)

var InstructionCodes = map[string]string{
	"reset":          Reset.String(),
	"bold":           Bold.String(),
	"faint":          Faint.String(),
	"italic":         Italic.String(),
	"underline":      Underline.String(),
	"black":          Black.String(),
	"red":            Red.String(),
	"green":          Green.String(),
	"yellow":         Yellow.String(),
	"blue":           Blue.String(),
	"magenta":        Magenta.String(),
	"cyan":           Cyan.String(),
	"white":          White.String(),
	"bright_black":   BrightBlack.String(),
	"bright_red":     BrightRed.String(),
	"bright_green":   BrightGreen.String(),
	"bright_yellow":  BrightYellow.String(),
	"bright_blue":    BrightBlue.String(),
	"bright_magenta": BrightMagenta.String(),
	"bright_cyan":    BrightCyan.String(),
	"bright_white":   BrightWhite.String(),
}

func (i InstructionType) String() string {
	return fmt.Sprintf("\x1b[%dm", i)
}
