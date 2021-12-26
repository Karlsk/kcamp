package main

import (
	"fmt"
	"strings"
)

func listReplace(before *[5]string) {
	for i := 0; i < 5; i++ {
		if strings.Compare((*before)[i],"stupid") == 0{
			(*before)[i] = "smart"
		}else if strings.Compare((*before)[i],"weak") == 0 {
			(*before)[i] = "strong"
		}else {
			continue
		}
	}

}

func main() {
	var before = [5]string{"I","am","stupid","and","weak"}
	fmt.Println("before=",before)
	listReplace(&before)
	fmt.Println("before=",before)

}
