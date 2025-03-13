package interactive

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func getCoefficient(reader *bufio.Reader, coefficientName string) float64 {
	for {
		fmt.Printf("%s = ", coefficientName)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		coefficient, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Printf("Error. Expected a valid real number, got %s instead\n", input)
			continue
		}

		if coefficient == 0 && coefficientName == "a" {
			fmt.Println("Error: a cannot be equal to 0")
			continue
		}

		return coefficient
	}
}



func RunInteractiveMode() (a, b, c float64) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter coefficients of a quadratic equation:")
	a = getCoefficient(reader, "a")
	b = getCoefficient(reader, "b")
	c = getCoefficient(reader, "c")

	return
}