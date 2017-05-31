package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	key := fmt.Sprint(rand.Int63n(time.Now().UnixNano()))
	fmt.Println(key)
}
