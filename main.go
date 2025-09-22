// Deniss Kozlovs 241RDB330
// 1 uzdevums
package main

import "fmt"

func main() {
	var number int
	fmt.Println("Write your number:")
	_,err := fmt.Scan(&number)
	if err != nil || number <2 {
		fmt.Println("input-output error")
		return
	} else {
		var factor  = 2
		for factor <= number {
			if number % factor == 0 {
				fmt.Print(factor," ")
				number  = number / factor
			}else{
				factor++
			}
    	}
	}
}
