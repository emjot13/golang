package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var welcomeString = `Hello!
In this game you will be trying to guess some random number in the fewest tries you can.
After each guess you will be shown a number of icons depending on the difficulty level.
Among them is hidden one of 2 icons: up :"◓" and down: "◒".
Up icon means your guess was too high, and down icon means your guess was too low.
However you have to be quick, because the icons will disappear!
If you want to stop playing type in the terminal "finish"
Good luck!`

var firstGameOfSession = true

var path string

type score struct {
	name   string
	result map[int]string
}

type allScores struct {
	name    string
	results []map[int]string
}

var playersAndScores []allScores
var low, high int

func generateNumber() int {
	for {
		var lowString, highString string
		fmt.Println("Enter low-end of the range from which the random number will be generated")
		_, err1 := fmt.Scan(&lowString)
		low, err1 = strconv.Atoi(lowString)
		fmt.Println("Enter high-end of the range")
		_, err := fmt.Scan(&highString)
		high, err = strconv.Atoi(highString)
		if err1 != nil || err != nil {
			fmt.Print("\nError, enter valid integers\n\n")
		} else if low > high {
			fmt.Printf("\nError, low-end of range cannot be bigger than high-end\n\n")
		} else {
			rand.Seed(time.Now().UnixNano())
			numToGuess := low + rand.Intn(high-low+1) // ensures generating number from the specified range
			return numToGuess
		}
	}
}

func createResultString() []string {
	var stringResults []string
	for _, value := range playersAndScores {
		tmp := "" // here is sorting slice of maps by key, its probably not the best way but it is what it is
		type results struct {
			tries int
			time  string
		}
		var playerResults []results
		for _, val := range value.results {
			for k, val1 := range val {
				playerResults = append(playerResults, results{k, val1})
			}
		}
		sort.Slice(playerResults, func(i, j int) bool { return playerResults[i].tries < playerResults[j].tries })
		for i, v := range playerResults {
			tmp += fmt.Sprintf("%v, %s", v.tries, v.time)
			if i != len(playerResults)-1 {
				tmp += " – "
			}
		}
		stringResults = append(stringResults, tmp)
	}
	return stringResults
}

func printScores() {
	fmt.Println()
	fmt.Println("Scores: ")
	for i, value := range playersAndScores {
		fmt.Printf("Player: %s, results (guesses, time): ", value.name)
		fmt.Printf("%s", createResultString()[i])
		fmt.Println()
	}
	fmt.Println()
}

func addScore(player score) {
	if len(playersAndScores) == 0 {
		playersAndScores = append(playersAndScores, allScores{player.name, []map[int]string{player.result}}) // first player
	} else {
		for i, players := range playersAndScores { // looking for the player in the current session and changing his
			if players.name == player.name { // results to the appened slice
				players.results = append(players.results, player.result)
				playersAndScores[i] = allScores{player.name, players.results}
				break
			} else if i == len(playersAndScores)-1 { // if we didn't find player in current session we create a new one
				playersAndScores = append(playersAndScores, allScores{player.name, []map[int]string{player.result}})
			}
		}
	}
}

func getName() string {
	fmt.Println("Enter your name")
	var name string
	fmt.Scan(&name)
	return name
}

func createFile() {
	fmt.Println("Type \"y\" if you want to create a file for results or type anything else if you don't want to")
	var ifNewFile string
	fmt.Scan(&ifNewFile)
	if strings.Contains("Yy", ifNewFile) {
		fmt.Println("Enter the filename with csv extension")
		var fileName string
		fmt.Scan(&fileName)
		csvFile, _ := os.Create(fileName)
		csvwriter := csv.NewWriter(csvFile)
		csvwriter.Write([]string{"Name", "Results", "Date"})
		csvwriter.Flush()
		csvFile.Close()
	}

}

func writeToFile() {
	fmt.Println("Do you want to save the results to a file? (y/n): ")
	var ifSave string
	fmt.Scan(&ifSave)
	switch ifSave {
	case "y", "Y":
		fmt.Println("Enter the absolute path of the csv file")
		var path string
		fmt.Scan(&path)
		file, _ := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		var data [][]string
		for i, val := range playersAndScores {
			date := time.Now().Format("Monday 2006-01-02 15:04")
			data = append(data, []string{val.name, createResultString()[i], date}) // appending data with every iteration
		}

		writer := csv.NewWriter(file)
		writer.WriteAll(data)
		writer.Flush()
		file.Close()
	}
}

func readFile() string {
	fmt.Println("If you want to load a file with results type \"y\" or click any button to continue")
	var ifLoadingFile string
	fmt.Scan(&ifLoadingFile)
	if strings.Contains("Yy", ifLoadingFile) {
		fmt.Println("Enter the absolute path to the file")
		var path string
		fmt.Scan(&path)
		return path
	}
	return ""
}

func checkIfNewRecord(path, name string, result int) {
	file, _ := os.Open(path)
	csvReader := csv.NewReader(file)
	data, _ := csvReader.ReadAll()

	var personalHighScores []int
	var globalHighScores []int
	var date string
	var dateGlobal string
	var player string
	for ind, value := range data {
		if ind != 0 {
			sessionRecord, _ := strconv.Atoi(strings.Split(value[1], ",")[0])
			globalHighScores = append(globalHighScores, sessionRecord)
			sort.Ints(globalHighScores)
			if sessionRecord == globalHighScores[0] {
				dateGlobal = value[2]
				player = value[0]
			}
		}

		for _, v := range value {
			if v == name {
				PersonalSessionRecord, _ := strconv.Atoi(strings.Split(value[1], ",")[0]) // converting string in form of "10, 12" into slice of int and taking the first element
				personalHighScores = append(personalHighScores, PersonalSessionRecord)
				sort.Ints(personalHighScores)
				if PersonalSessionRecord == personalHighScores[0] { // if session record is the high score we assign corresponding date to the date variable
					date = value[2] // it doesn't take into account which of the same high scores was scored earliest
				}

			}
		}
	}
	if len(personalHighScores) != 0 {
		if personalHighScores[0] > result {
			fmt.Printf("Congratulations you've beaten your previous record of %d tries from %s\n", personalHighScores[0], date)
		}
	}

	sort.Ints(globalHighScores)
	if len(globalHighScores) != 0 {
		if globalHighScores[0] > result {
			fmt.Printf("Congratulations you've beaten the %s's global record of %d tries from %s\n", player, globalHighScores[0], dateGlobal)
		}
	}
}

func playAgain() {
	fmt.Println("Do you want to play again? (y/n):")
	for {
		var playAgain string
		fmt.Scan(&playAgain)
		switch playAgain {
		case "y", "Y":
			firstGameOfSession = false
			main()
		case "n", "N":
			printScores()
			writeToFile()
			os.Exit(0)
		default:
			fmt.Println("Incorrect input. Type your answer again")
		}
	}
}

func checkIfOptimal(gameRange, guesses int) {
	if int(math.Floor(math.Logb(float64(gameRange))))+1 < guesses {
		fmt.Println("The number can be found with less guesses. Maybe try to think about some algorithm or be quicker :D")
	}
}

func createMatrix(size int, low_high string) {
	right := "◑"
	left := "◐"
	up := "◓"
	down := "◒"
	x := rand.Intn(size)
	y := rand.Intn(size)
	matrix := make([][]string, size)
	for i := range matrix {
		matrix[i] = make([]string, size)
	}
	for _, list := range matrix {
		for i := range list {
			left_right := rand.Intn(2)
			if left_right == 0 {
				list[i] = left
			} else {
				list[i] = right
			}
		}
	}
	if low_high == "low" {
		matrix[x][y] = down
	} else {
		matrix[x][y] = up
	}

	for _, list := range matrix {
		fmt.Println(strings.Join(list[:], " "))
	}

}

func difficultyLevel() int {
	fmt.Println("Choose difficulty level")
	fmt.Println("1. Easy\t\t2. Medium\t3. Hard\t\t4. Insane")
	for {
		var answer int
		fmt.Scan(&answer)
		switch answer {
		case 1:
			return 5
		case 2:
			return 10
		case 3:
			return 15
		case 4:
			return 25
		default:
			fmt.Println("Incorrect difficulty level. Type again.")
		}
	}
}

func guessing(level int) string {

	var seconds int
	switch level {
	case 5: //easy
		seconds = 4
	case 10: // medium
		seconds = 7
	case 15: // hard
		seconds = 9
	case 25: // insane
		seconds = 14
	}
	var input string
	ch := make(chan string)
	go func() {
		fmt.Printf("Enter your guess\n")
		fmt.Scan(&input)
		ch <- input
	}()
	select {
	case <-ch:
		return input
	case <-time.After(time.Duration(seconds) * time.Second):
		fmt.Print(strings.Repeat("\n", 2000))
		return "empty"
	}
}


func main() {
	if firstGameOfSession {
		fmt.Println(welcomeString)
		createFile()
		path = readFile()
	}
	size := difficultyLevel()
	triesNumber := 1
	numToGuess := generateNumber()
	var firstGuess bool
	firstGuess = true
	var guess string
	start := time.Now()
	for {
		if !firstGuess {
			guess = guessing(size)
		}else{
			fmt.Println("Time start!")
			fmt.Println("Enter your guess")
			fmt.Scan(&guess)
		}
		if guess == "empty" { // absolutely no idea why player has to enter input twice :(
			fmt.Println("You were too slow ;P")													// but it's not a bug it's a feaature
			fmt.Println("Now you have to enter your last guess and then you can guess again")	// it doesnt do anything, though
			fmt.Scan(&guess)
		}
		firstGuess = false
		if guess == "finish" { // first we check if string input is keyword "finish"
			printScores()
			fmt.Println("Goodbye!")
			os.Exit(0)
		}
		guessInt, err := strconv.Atoi(guess) // if not we can try to convert the input to integer
		if err == nil {
			switch {
			case guessInt == numToGuess:
				stop := strings.Split(fmt.Sprintf("%.2f", time.Since(start).Minutes()), ".")
				stopString := stop[0] + " min " + stop[1] + " sec"
				fmt.Printf("Congratulations, you guessed the number in %s in %d tries\n", stopString, triesNumber)
				scoreTimeMap := make(map[int]string)
				scoreTimeMap[triesNumber] = stopString
				checkIfOptimal(1+high-low, triesNumber)
				name := getName()
				addScore(score{name, scoreTimeMap})
				if path != "" {
					checkIfNewRecord(path, name, triesNumber)
				}
				playAgain()
			case guessInt < numToGuess:
				createMatrix(size, "low")
				triesNumber += 1
			case guessInt > numToGuess:
				createMatrix(size, "high")
				triesNumber += 1
			}
		} else {
			fmt.Println("Enter valid integer as your guess")
		}
	}
}
