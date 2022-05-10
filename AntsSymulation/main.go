package main

import (
	"fmt"
	"math/rand"
	// "os"
	// "os/exec"
	"strings"
	"time"
)

type antPosition struct {
	x int
	y int
	withLeaf bool
}

type leafPosition struct {
	x int
	y int
	onGround bool
}

var allAntsPositions []antPosition
var allLeafsPositions []leafPosition


func getPossibleMoves(index int, matrix [][]string)[][]int{
	matrixLength := len(matrix)
	matrixWidth := len(matrix[0])
	x, y := allAntsPositions[index].x, allAntsPositions[index].y
	var possibleMoves [][]int
	can_left := allAntsPositions[index].x - 1 >= 0
	can_right := allAntsPositions[index].x + 1 < matrixLength
	can_up := allAntsPositions[index].y - 1 >= 0
	can_down := allAntsPositions[index].y + 1 < matrixWidth
	left, right, up, down, leftup, leftdown, rightup, rightdown := []int{x - 1, y}, []int{x + 1, y}, []int{x, y - 1}, []int{x, y + 1}, []int{x - 1, y - 1}, []int{x - 1, y + 1}, []int{x + 1, y - 1}, []int{x + 1, y + 1}

	switch {
	case can_left && can_right && can_up && can_down:
		possibleMoves = append(possibleMoves, left, right, up, down, leftup, leftdown, rightup, rightdown)
	case can_left && can_right && can_up:
		possibleMoves = append(possibleMoves, left, right, up, leftup, rightup)
	case can_left && can_right && can_down:
		possibleMoves = append(possibleMoves, left, right, down, leftdown, rightdown)
	case can_up && can_down && can_left:
		possibleMoves = append(possibleMoves, up, down, left, leftup, leftdown)	
	case can_up && can_down && can_right:
		possibleMoves = append(possibleMoves, up, down, right, rightup, rightdown)
	case can_up && can_right:
		possibleMoves = append(possibleMoves, up, right, rightup)
	case can_up && can_left:
		possibleMoves = append(possibleMoves, up, left, leftup)
	case can_down && can_right:
		possibleMoves = append(possibleMoves, down, right, rightdown)
	case can_down && can_left:
		possibleMoves = append(possibleMoves, down, left, leftdown)
	}
	return possibleMoves
}




// func action(matrix [][]string, x, y int){

// }

func moveAnts(matrix [][]string){
	// fmt.Println(allAntsPositions, len(allAntsPositions))
	// fmt.Println(allLeafsPositions)
	for i, ant := range allAntsPositions{
		availableMoves := getPossibleMoves(i, matrix)
		randomIndex := rand.Intn(len(availableMoves))
		x, y := availableMoves[randomIndex][0], availableMoves[randomIndex][1]
		fmt.Println(x, y, matrix[y][x])
		var foundLeaf bool
		for _, leaf := range allLeafsPositions{
			if leaf.onGround{
				if matrix[y][x] == "🍂"{
					fmt.Println("here")
					foundLeaf = true
					if !ant.withLeaf{
					ant.withLeaf = true
					leaf.onGround = false
					matrix[ant.y][ant.x] = "🪲"
					matrix[y][x] = "🔲"
					break
				}
				if ant.withLeaf{
					var indexes []int
					for i := 0; i < len(availableMoves); i++{
						indexes = append(indexes, i)
					}
					rand.Shuffle(len(indexes), func(i, j int) { indexes[i], indexes[j] = indexes[j], indexes[i]})
					for i := range indexes{
						x1, y1 := availableMoves[indexes[i]][0], availableMoves[indexes[i]][1]
						fmt.Println("x mrówki", ant.x, "y mrówki", ant.y, "x liścia", leaf.x, "y liścia", leaf.y, "nowy ruch", x1, y1, matrix[y1][x1])
						if matrix[y1][x1] == "🔲"{
							matrix[ant.y][ant.x] = "🐜"
							// fmt.Println("x mrówki", ant.x, "y mrówki", ant.y, "x liścia", leaf.x, "y liścia", leaf.y, "nowy ruch", x1, y1, matrix[y1][x1])
							matrix[y1][x1] = "🍂"
							allLeafsPositions = append(allLeafsPositions, leafPosition{x1, y1, true})
							ant.withLeaf = false
							break
						}
					}


				}

				}
			}
		}

		if !foundLeaf {
			matrix[ant.y][ant.x] = "🔲"
			if ant.withLeaf {
				matrix[y][x] = "🪲"
			}else{
			matrix[y][x] = "🐜"}
			ant.x, ant.y = x, y
		}
		allAntsPositions[i] = antPosition{ant.x, ant.y, ant.withLeaf}
		fmt.Println(allAntsPositions[i])

	}

}




func createMatrix(matrixLength, matrixWidth, antsNumber, leafsNumber int)[][]string{
	var smaller, bigger int
	if matrixLength < matrixWidth {
		smaller = matrixLength
		bigger = matrixWidth
	}else{
		smaller = matrixWidth
		bigger = matrixLength
	}

	matrix := make([][]string, matrixWidth)
	indexes := make([][]int, matrixLength*matrixWidth)
	i, j := 0, 0
	for y := range indexes {
		if j < bigger{
			indexes[y] = []int{i, j}
			j++
		}else{
			i++
			j = 0
			if i < smaller{
				indexes[y] = []int{i, j}
				j++
			}
		}
}       

	for i := range matrix {
		matrix[i] = make([]string, matrixLength)
	}

	for _, list := range matrix {
		for y := range list {
			list[y] = "🔲"
		}
	}

	shuffled := make([][]int, len(indexes))
	perm := rand.Perm(len(indexes))
	for i, randIndex := range perm {
	shuffled[i] = indexes[randIndex]
	}
	
	for i := 0; i < antsNumber; i++ {
		x, y := shuffled[i][1], shuffled[i][0]
		allAntsPositions = append(allAntsPositions, antPosition{y, x, false})
		matrix[x][y] = "🐜"
	}

	for i := antsNumber; i < antsNumber + leafsNumber; i++ {
		x, y := shuffled[i][0], shuffled[i][1]
		allLeafsPositions = append(allLeafsPositions, leafPosition{x, y, true})
		matrix[y][x] = "🍂"
	}

	return matrix
}

func printMatrix(matrix [][]string){
	j := 0
	fmt.Printf("   ")
	for i := 0; i<len(matrix); i++{
		fmt.Printf("%d  ", i)
	}
	fmt.Printf("\n")
	for _, list := range matrix {
		fmt.Println(j, strings.Join(list[:], " "))
		j++
	}
}

func main(){
	rand.Seed(time.Now().UnixNano())
	matrix := createMatrix(10, 10, 1, 10)
	// printMatrix(matrix)
	for i := 0; i < 50; i++ {
		moveAnts(matrix)
		fmt.Printf("\n\n\n")
		printMatrix(matrix)
		// time.Sleep(1 * time.Second)
	// 	if i != 35{
	// 	cmd := exec.Command("clear")
	// 	cmd.Stdout = os.Stdout
	// 	cmd.Run()
	// }	

}
}

