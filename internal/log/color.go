package log

import "fmt"

type Color struct{}

func (c Color) Red(s string) string {
	return fmt.Sprintf("\033[31m%s\033[0m", s)
}

func (c Color) Green(s string) string {
	return fmt.Sprintf("\033[32m%s\033[0m", s)
}

func (c Color) Yellow(s string) string {
	return fmt.Sprintf("\033[33m%s\033[0m", s)
}

func (c Color) Blue(s string) string {
	return fmt.Sprintf("\033[34m%s\033[0m", s)
}

func (c Color) Magenta(s string) string {
	return fmt.Sprintf("\033[35m%s\033[0m", s)
}

func (c Color) Cyan(s string) string {
	return fmt.Sprintf("\033[36m%s\033[0m", s)
}
