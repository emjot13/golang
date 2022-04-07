package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var welcomeString = `Hello!
In this game you will be trying to guess some random number in the fewest tries you can. 
If you want to stop playing type in the terminal "finish"
Good luck!`

var firstGameOfSession = true

var path string


type score struct {
	name string
	result int
}

type allScores struct {
	name string
	results []int
}

var playersAndScores []allScores

func generateNumber()(int){
	for {
		var lowString string
		var low int
		var high int
		var highString string
		fmt.Println("Enter low-end of the range from within the random number will be generated")
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
			numToGuess := low + rand.Intn(high - low + 1) // ensures generating number from the specified range
			return numToGuess
        }
    }
}



func printScores() {
	fmt.Println("Scores: ")
	for _, value := range playersAndScores {
		fmt.Printf("Player: %s, results (number of tries): ", value.name)
		sort.Ints(value.results)
		for i, num := range value.results{
			if i != len(value.results) - 1{			// if not last element we print with comma
				fmt.Printf("%d, ", num)
			}else{
			fmt.Printf("%d.", num)}					// if last, we print with dot
		}
		fmt.Println()
	}
}


func addScore(player score){
	if len(playersAndScores) == 0 {
		playersAndScores = append(playersAndScores, allScores{player.name, []int{player.result}})  // first player
	}else{
	for i, players := range playersAndScores {													   // looking for the player in the current session and changing his
		if players.name == player.name {														   // results to the appened slice
			players.results = append(players.results, player.result)
			playersAndScores[i] = allScores{player.name, players.results}
			break
		} else if i == len(playersAndScores) - 1 {												  // if we didn't find player in current session we create a new one
			playersAndScores = append(playersAndScores, allScores{player.name, []int{player.result}})
		}
	}
}
}

func getName(triesNumber int)string{
	fmt.Printf("Congratulations, you guessed the number in %d tries\n", triesNumber)
	fmt.Println("Enter your name to save your result")
	var name string
	fmt.Scan(&name)
	return name
}


func createFile(){
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
		csvFile.Close()}
	


}



func writeToFile(){
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
			for _, val := range playersAndScores{
				results := strings.Join(strings.Fields(strings.Trim(fmt.Sprint(val.results), "[]")), ",")			// converting slice of ints to []string and then to string
				date := time.Now().Format("Monday 2006-01-02 15:04")
				data = append(data, []string{val.name, results, date})												// appending data with every iteration
			}
		
			writer := csv.NewWriter(file)
			writer.WriteAll(data)
			writer.Flush()
			file.Close()
	}
}

func readFile()string {
	fmt.Println("If you want to load a file with results type \"y\" or click any button to continue")
	var ifLoadingFile string
	fmt.Scan(&ifLoadingFile)
	if strings.Contains("Yy", ifLoadingFile){
		fmt.Println("Enter the absolute path to the file")
		var path string
		fmt.Scan(&path)
		return path
	}
	return ""
}




func checkIfNewRecord(path, name string, result int){
	file, _ := os.Open(path)
	csvReader := csv.NewReader(file)
	data, _ := csvReader.ReadAll()
	var highScores []int 
	var date string
	for _, value := range data{
		for i, v := range value{
			if v == name{
				sessionRecord, _ := strconv.Atoi(strings.Split(value[i + 1], ",")[0])			// converting string in form of "10, 12" into slice of int and taking the first element
				highScores = append(highScores, sessionRecord)
				sort.Ints(highScores)
				if sessionRecord == highScores[0]{											// if session record is the high score we assign corresponding date to the date variable 																							
					date = value[i + 2]														// it doesn't take into account which of the same high scores was scored earliest
				}

			}
		}
	}
	if len(highScores) != 0{
	if highScores[0] > result{
		fmt.Printf("Congratulations you've beaten your previous record of %d tries from %s\n", highScores[0], date)
	}}
}

func playAgain(){
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

func main() {
	if firstGameOfSession{
		fmt.Println(welcomeString)
		createFile()
		path = readFile()
		}
	triesNumber := 1
	numToGuess := generateNumber()
	for {
		fmt.Println("Enter your guess")
		var guess string
		fmt.Scan(&guess)
		if guess == "finish"{ 		 				// first we check if string input is keyword "finish"
			printScores()
			fmt.Println("Goodbye!") 
			os.Exit(0)
		}
		guessInt, err := strconv.Atoi(guess)		// if not we can try to convert the input to integer
		if err == nil {
		switch {
		case guessInt == numToGuess:
			name := getName(triesNumber)
			addScore(score{name, triesNumber})
			if path != "" {
			checkIfNewRecord(path, name, triesNumber)}
			playAgain()
		case guessInt < numToGuess:
			fmt.Println("Your guess was too low. Try again.")
			triesNumber += 1
		case guessInt > numToGuess:
			fmt.Println("Your guess was too high. Try again.")
			triesNumber += 1
		}
		}else{
			fmt.Println("Enter valid integer as your guess")
		}	
	}
}


