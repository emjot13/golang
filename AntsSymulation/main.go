package main

import (
	"fmt"
	"math/rand"
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


func moveAnts(matrix [][]string){
	// length := len(matrix)
	// width := len(matrix[0])
	// fmt.Println(allLeafsPositions)
	fmt.Println(allAntsPositions)

	for ind, ant := range allAntsPositions{
		moves := getPossibleMoves(ind, matrix)
		randInd := rand.Intn(len(moves))
		x, y := moves[randInd][0], moves[randInd][1]
		fmt.Println("x mrufki:", ant.x, ", y mrufki:", ant.y, ", nowy ruch:", x, y)
		if matrix[y][x] == 
	}
	// 	foundLeaf := false
	// 	for _, leaf := range allLeafsPositions{
	// 		if leaf.x == x && leaf.y == y {
	// 			foundLeaf = true
	// 			if ant.withLeaf {
	// 				ant.withLeaf = false
	// 				matrix[ant.y][ant.x] = "🐜"
	// 				break
	// 			}else{
	// 			matrix[ant.y][ant.x] = "🔲"
	// 			matrix[x][y] = "🪲"
	// 			ant.x, ant.y = x, y
	// 			ant.withLeaf = true}
	// 			break
	// 		}
	// 	}
	// 	if !foundLeaf{
	// 		fmt.Print(matrix[ant.x][ant.y])
	// 		matrix[ant.y][ant.x] = "🔲"
	// 		fmt.Print(matrix[ant.x][ant.y])
	// 		matrix[x][y] = "🐜"
	// 		ant.x, ant.y = x, y
	// 	}
	// 	allAntsPositions[ind] = antPosition{ant.y, ant.x, ant.withLeaf}
	// }
	// fmt.Println(allAntsPositions)
}
	// for ind, ant := range allAntsPositions{
	// 	matrix[ant.x][ant.y] = "🔲"
	// 	new_x := rand.Intn(length)
	// 	new_y := rand.Intn(width)
	// 	// ant.x, ant.y = new_x, new_y
	// 	for _, leaf := range allLeafsPositions{
	
	// 		if leaf.x == new_x && leaf.y == new_y {
	// 			if ant.withLeaf {
	// 				matrix[leaf.x][leaf.y] = "🍂"

	// 				ant.withLeaf = false
	// 			}else{
	// 			matrix[leaf.x][leaf.y] = "🪲"
	// 			ant.withLeaf = true}
	// 		}
	// 	}
	// 	if !ant.withLeaf {
	// 		matrix[ant.x][ant.y] = "🐜"
	// 	}
	// 	allAntsPositions[ind] = antsPosition{ant.x, ant.y	}









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
		x, y := shuffled[i][0], shuffled[i][1]
		allAntsPositions = append(allAntsPositions, antPosition{x, y, false})
		matrix[y][x] = "🐜"
	}

	for i := antsNumber; i < antsNumber + leafsNumber; i++ {
		x, y := shuffled[i][0], shuffled[i][1]
		allLeafsPositions = append(allLeafsPositions, leafPosition{x, y})
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
	matrix := createMatrix(10, 10, 2, 5)
	printMatrix(matrix)
	for i := 0; i < 50; i++ {
		moveAnts(matrix)
		fmt.Printf("\n\n\n")
		printMatrix(matrix)
	}	

}


