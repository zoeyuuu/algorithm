package huawei

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func HJ2() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	if scanner.Scan() {
		str1 := scanner.Text()
		scanner.Scan()
		str2 := scanner.Text()
		str1 = strings.ToLower(str1)
		str2 = strings.ToLower(str2)
		for i := 0; i < len(str1); i++ {
			if str1[i] == str2[0] {
				sum++
			}
		}
	}
	fmt.Println(sum)
}
