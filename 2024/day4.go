package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	X = 88
	M = 77
	A = 65
	S = 83
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 1)
	var width, height int
	for scanner.Scan() {
		line := scanner.Text()
		width = len(line)
		lines = append(lines, line)
		height += 1
	}

	fmt.Println(width, height)
	data := strings.Join(lines, "")
	var total int
	// fmt.Println(data)
	// fmt.Println(get_char(data, width, 5, 5))
	for y := range height {
		for x := range width {
			if get_char(data, width, x, y) == X {
				total += check_0(data, width, x, y)
				total += check_45(data, width, x, y)
				total += check_90(data, width, x, y)
				total += check_135(data, width, x, y)
				total += check_180(data, width, x, y)
				total += check_225(data, width, x, y)
				total += check_270(data, width, x, y)
				total += check_315(data, width, x, y)
			}
		}
	}

	fmt.Println("total:", total)
}

func get_char(data string, width, x, y int) byte {
	if x < 0 {
		fmt.Println("x less than zero")
		panic("!!")
	}

	if y < 0 {
		fmt.Println("y less than zero")
		panic("!!")
	}

	idx := width*y + x
	// fmt.Println(idx)
	return data[idx]
}

func check_0(data string, width, x, y int) int {
	if x+3 >= width {
		return 0
	}

	if get_char(data, width, x+1, y) == M && get_char(data, width, x+2, y) == A && get_char(data, width, x+3, y) == S {
		fmt.Printf("0 deg found at %d, %d\n", x, y)
		return 1
	}
	return 0
}

func check_45(data string, width, x, y int) int {
	if y < 3 || x+3 >= width {
		return 0
	}

	if get_char(data, width, x+1, y-1) == M && get_char(data, width, x+2, y-2) == A && get_char(data, width, x+3, y-3) == S {
		fmt.Printf("45 deg found at %d, %d\n", x, y)
		return 1
	}
	return 0
}

func check_90(data string, width, x, y int) int {
	if y < 3 {
		return 0
	}

	if get_char(data, width, x, y-1) == M && get_char(data, width, x, y-2) == A && get_char(data, width, x, y-3) == S {
		fmt.Printf("45 deg found at %d, %d\n", x, y)
		return 1
	}
	return 0
}

func check_135(data string, width, x, y int) int {
	if y < 3 || x < 3 {
		return 0
	}

	if get_char(data, width, x-1, y-1) == M && get_char(data, width, x-2, y-2) == A && get_char(data, width, x-3, y-3) == S {
		fmt.Printf("45 deg found at %d, %d\n", x, y)
		return 1
	}
	return 0
}

func check_180(data string, width, x, y int) int {
	if x < 3 {
		return 0
	}

	if get_char(data, width, x-1, y) == M && get_char(data, width, x-2, y) == A && get_char(data, width, x-3, y) == S {
		fmt.Printf("0 deg found at %d, %d\n", x, y)
		return 1
	}
	return 0
}

func check_225(data string, width, x, y int) int {
	if y+3 >= width || x < 3 {
		return 0
	}

	if get_char(data, width, x-1, y+1) == M && get_char(data, width, x-2, y+2) == A && get_char(data, width, x-3, y+3) == S {
		fmt.Printf("45 deg found at %d, %d\n", x, y)
		return 1
	}
	return 0
}

func check_270(data string, width, x, y int) int {
	if y+3 >= width {
		return 0
	}

	if get_char(data, width, x, y+1) == M && get_char(data, width, x, y+2) == A && get_char(data, width, x, y+3) == S {
		fmt.Printf("45 deg found at %d, %d\n", x, y)
		return 1
	}
	return 0
}

func check_315(data string, width, x, y int) int {
	if y+3 >= width || x+3 >= width {
		return 0
	}

	if get_char(data, width, x+1, y+1) == M && get_char(data, width, x+2, y+2) == A && get_char(data, width, x+3, y+3) == S {
		fmt.Printf("315 deg found at %d, %d\n", x, y)
		return 1
	}
	return 0
}
