package utils

import (
	"fmt"
	"github.com/fatih/color"
)

// 预定义常用颜色属性
var (
	ColorHiRed     = color.FgHiRed
	ColorHiGreen   = color.FgHiGreen
	ColorHiYellow  = color.FgHiYellow
	ColorHiBlue    = color.FgHiBlue
	ColorHiMagenta = color.FgHiMagenta
	ColorHiCyan    = color.FgHiCyan
	ColorHiWhite   = color.FgHiWhite
	ColorHiBlack   = color.FgHiBlack
	ColorRed       = color.FgRed
	ColorGreen     = color.FgGreen
	ColorWhite     = color.FgWhite
)

type ColorPrinter struct {
	*color.Color
}

func NewColorPrinter(attr color.Attribute) *ColorPrinter {
	return &ColorPrinter{color.New(attr)}
}

// ColorPrint 使用指定颜色打印文本（自动换行）
func ColorPrint(attr color.Attribute, format string, a ...interface{}) {
	c := color.New(attr)
	c.Println(fmt.Sprintf(format, a...))
}

// ColorPrintf 使用指定颜色打印文本（不自动换行）
func ColorPrintf(attr color.Attribute, format string, a ...interface{}) {
	c := color.New(attr)
	c.Printf(format, a...)
}

// ColorSprint 返回带颜色的字符串
func ColorSprint(attr color.Attribute, text string) string {
	return color.New(attr).Sprint(text)
}

// 预定义常用颜色的快捷函数
func PrintRed(text string) {
	color.New(color.FgRed).Println(text)
}

func PrintGreen(text string) {
	color.New(color.FgGreen).Println(text)
}

func PrintYellow(text string) {
	color.New(color.FgYellow).Println(text)
}

func PrintHiRed(text string) {
	color.New(color.FgHiRed).Println(text)
}