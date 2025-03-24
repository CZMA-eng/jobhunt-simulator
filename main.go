package main

import (
	"math/rand"
	"jobhunt/utils"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	utils.ShowOpening()
	mainLoop()
}