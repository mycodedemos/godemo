package main

import (
	"github.com/mycodedemos/godemo/utils"
	"fmt"
	"time"
)

func main() {
	begin := time.Now().Unix()
	for i := 0; i < 100000; i++ {
		utils.GetRandStr(1000)
		utils.GetRandStr1(1000)
	}

	end := time.Now().Unix()
	fmt.Println(begin, end)

	fmt.Println("time: ", (end - begin))
}
