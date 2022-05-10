package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)


func main() {
	var indexes []int
	for i := 0; i < 8; i++{
		indexes = append(indexes, i)
	}
	rand.Shuffle(len(indexes), func(i, j int) { indexes[i], indexes[j] = indexes[j], indexes[i]})
	fmt.Println(indexes)
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}