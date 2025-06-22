package main

import (
	"package/utils"
	"strconv"
)

func main() {
	random := utils.RandomInt(15)
	utils.PrintMessage("Random number: ", strconv.Itoa(random))
}
