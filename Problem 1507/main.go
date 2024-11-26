package main

import (
	"fmt"
)

func reformatDate(date string) string {

	dict := map[string]string{
		"Jan": "01",
		"Feb": "02",
		"Mar": "03",
		"Apr": "04",
		"May": "05",
		"Jun": "06",
		"Jul": "07",
		"Aug": "08",
		"Sep": "09",
		"Oct": "10",
		"Nov": "11",
		"Dec": "12",
	}

	day := date[0:2]
	if day[1] < '0' || day[1] > '9' {
		day = "0" + day[0:1]
	}

	return date[len(date)-4:] + "-" + dict[date[len(date)-8:len(date)-5]] + "-" + day
}

func main() {
	fmt.Println(reformatDate("20th Oct 2052"))
}
