package render

import "fmt"

func Print(lines []string) {
	for _, l := range lines {
		fmt.Println(l)
	}
}
