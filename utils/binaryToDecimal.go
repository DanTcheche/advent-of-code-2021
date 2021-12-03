package utils

import (
	"strconv"
)

func BinaryToDecimal (binary string) int64 {
	output, err := strconv.ParseInt(binary, 2, 64)  
	if err != nil {  
		panic("Invalid input")
	} 
	return output
}