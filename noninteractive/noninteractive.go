package noninteractive

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func RunNonInteractiveMode(filePath string) (a, b, c float64) {
	data, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Printf("Error: file %s does not exist\n", filePath)
		return math.NaN(), math.NaN(), math.NaN()
	}

	if !strings.HasSuffix(string(data), "\\n") {
		fmt.Println("Error: invalid file format")
		return math.NaN(), math.NaN(), math.NaN()
	}

	data = data[:len(data)-2]

	coefficients := strings.Split(string(data), "\\s")
	if len(coefficients) != 3 {
		fmt.Println("Error: invalid number of coefficients")
		return math.NaN(), math.NaN(), math.NaN()
	}

	a, err = strconv.ParseFloat(coefficients[0], 64)
	if err != nil {
		fmt.Println("Error: invalid file format")
		return math.NaN(), math.NaN(), math.NaN()
	}

	b, err = strconv.ParseFloat(coefficients[1], 64)
	if err != nil {
		fmt.Println("Error: invalid file format")
		return math.NaN(), math.NaN(), math.NaN()
	}

	c, err = strconv.ParseFloat(coefficients[2], 64)
	if err != nil {
		fmt.Println("Error: invalid file format")
		return math.NaN(), math.NaN(), math.NaN()
	}

	if a == 0 {
		fmt.Println("Error: a cannot be equal to 0")
		return math.NaN(), math.NaN(), math.NaN()
	}
	return
}
