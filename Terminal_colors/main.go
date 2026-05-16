package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"math"
)

func rgb(i int) (int, int, int) {
    var f = 0.1
    return int(math.Sin(f*float64(i)+0)*127 + 128),
        int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
        int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}


func main() {
	reader := bufio.NewReader(os.Stdin)

	i := 0

	for {
		r, _, err := reader.ReadRune()

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "read error:", err)
			break
		}
		if r == '\n' {
			fmt.Print("\n")
			continue
		}

		r_, g, b := rgb(i)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r_, g, b, r)

		i++
	}
}