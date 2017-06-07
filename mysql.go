package main

import (
	//"time"
	"github.com/mycodedemos/godemo/utils"
	"fmt"
)

func main() {
	utils.Execute("insert user set name = ?", "get")

	s := new(utils.Search)

	fmt.Print(s)

}
