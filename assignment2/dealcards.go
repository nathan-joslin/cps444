//dealcards displays cards by specification or at random. Two args minimum.
// Arg 1: v (vertical) / h (horizontal)
// Arg 2+: use {[2-10],J,Q,K,A} to return a desired card or "r" for random
// Example: Args[v 3 r] Will display the cards vertically, a "3" and a random card.
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const two string = `/------------\` +
	"\n" + `|    ____    |` +
	"\n" + `|        |   |` +
	"\n" + `|        |   |` +
	"\n" + `|     ---    |` +
	"\n" + `|    |       |` +
	"\n" + `|    |____   |` +
	"\n" + `\------------/` + "\n"

const three string = `/------------\` +
	"\n" + `|    ____    |` +
	"\n" + `|        \   |` +
	"\n" + `|        /   |` +
	"\n" + `|     ---    |` +
	"\n" + `|        \   |` +
	"\n" + `|    ____/   |` +
	"\n" + `\------------/` + "\n"

const four string = `/------------\` +
	"\n" + `|   |    |   |` +
	"\n" + `|   |    |   |` +
	"\n" + `|   |    |   |` +
	"\n" + `|   \----|-  |` +
	"\n" + `|        |   |` +
	"\n" + `|        |   |` +
	"\n" + `\------------/` + "\n"

const five string = `/------------\` +
	"\n" + `|    ____    |` +
	"\n" + `|   |        |` +
	"\n" + `|   |        |` +
	"\n" + `|   \----\   |` +
	"\n" + `|        |   |` +
	"\n" + `|    ____/   |` +
	"\n" + `\------------/` + "\n"

const six string = `/------------\` +
	"\n" + `|    ____    |` +
	"\n" + `|   |        |` +
	"\n" + `|   |        |` +
	"\n" + `|   |----\   |` +
	"\n" + `|   |    |   |` +
	"\n" + `|   |____|   |` +
	"\n" + `\------------/` + "\n"

const seven string = `/------------\` +
	"\n" + `|    ____    |` +
	"\n" + `|        |   |` +
	"\n" + `|        |   |` +
	"\n" + `|        |   |` +
	"\n" + `|        |   |` +
	"\n" + `|            |` +
	"\n" + `\------------/` + "\n"

const eight string = `/------------\` +
	"\n" + `|    ____    |` +
	"\n" + `|   |    |   |` +
	"\n" + `|   |    |   |` +
	"\n" + `|   |----|   |` +
	"\n" + `|   |    |   |` +
	"\n" + `|   |____|   |` +
	"\n" + `\------------/` + "\n"

const nine string = `/------------\` +
	"\n" + `|    ____    |` +
	"\n" + `|   |    |   |` +
	"\n" + `|   |    |   |` +
	"\n" + `|   |____|   |` +
	"\n" + `|        |   |` +
	"\n" + `|        |   |` +
	"\n" + `\------------/` + "\n"

const ten string = `/------------\` +
	"\n" + `|      ___   |` +
	"\n" + `|  |  |   |  |` +
	"\n" + `|  |  |   |  |` +
	"\n" + `|  |  |   |  |` +
	"\n" + `|  |  |___|  |` +
	"\n" + `|            |` +
	"\n" + `\------------/` + "\n"

const jack string = `/------------\` +
	"\n" + `|   _____    |` +
	"\n" + `|      |     |` +
	"\n" + `|      |     |` +
	"\n" + `|      |     |` +
	"\n" + `|  \___/     |` +
	"\n" + `|            |` +
	"\n" + `\------------/` + "\n"

const queen string = `/------------\` +
	"\n" + `|    ____    |` +
	"\n" + `|   |    |   |` +
	"\n" + `|   |    |   |` +
	"\n" + `|   |    |   |` +
	"\n" + `|   |____\   |` +
	"\n" + `|         \  |` +
	"\n" + `\------------/` + "\n"

const king string = `/------------\` +
	"\n" + `|   |  /     |` +
	"\n" + `|   | /      |` +
	"\n" + `|   |/       |` +
	"\n" + `|   |\       |` +
	"\n" + `|   | \      |` +
	"\n" + `|   |  \     |` +
	"\n" + `\------------/` + "\n"

const ace string = `/------------\` +
	"\n" + `|    ____    |` +
	"\n" + `|   |    |   |` +
	"\n" + `|   |    |   |` +
	"\n" + `|   |----|   |` +
	"\n" + `|   |    |   |` +
	"\n" + `|   |    |   |` +
	"\n" + `\------------/` + "\n"

// for testing horizontal option
const card string = `line1 ` +
	"\n" + `line2 ` +
	"\n" + `line3 ` +
	"\n" + `line4 ` +
	"\n" + `line5 ` +
	"\n" + `line6 ` +
	"\n" + `line7 ` +
	"\n" + `line8 ` + "\n"

const cardLines int = 8

func main() {
	//Check command line arguments
	validateArgs()

	//vertical by default
	horizontal := false
	if os.Args[1] == "h" {
		horizontal = true
	}

	//vertical display
	if !horizontal {
		fmt.Fprintf(os.Stdout, "Displaying vertically...\n")
		//iterate over args
		for _, arg := range os.Args[2:] {
			//fmt.Fprintf(os.Stdout, "Got arg: %s\n", arg)
			thisCard := getCard(arg)
			fmt.Fprintf(os.Stdout, thisCard)
		}
	} else { //horizontal display
		fmt.Fprintf(os.Stdout, "Displaying horizontally...\n")
		//initialize hand with the first requested card
		firstCard := getCard(os.Args[2])
		if len(os.Args[2:]) == 1 {
			fmt.Fprintf(os.Stdout, firstCard)
		} else {
			//process subsequent cards using slices
			cards := strings.Split(firstCard, "\n")
			//fmt.Fprintf(os.Stdout, "%v\n", cards)
			for _, arg := range os.Args[3:] {
				//slice the cards
				nextCard := getCard(arg)
				nextCardSlice := strings.Split(nextCard, "\n")
				//concat each line of the cards
				for i := 0; i < cardLines; i++ {
					cards[i] = strings.TrimSuffix(cards[i], "\n")
					cards[i] = cards[i] + nextCardSlice[i] + "\n"
				}
			}
			fmt.Fprintf(os.Stdout, "%v\n", strings.Join(cards, ""))
		}
	}
}

// randCard generates and returns a random card number as a string
func randCard() string {
	rand.Seed(time.Now().UnixNano())                //changes behavior of pseudo-random num gen each run based on current time
	return getCard(strconv.Itoa(rand.Intn(13) + 2)) //2 is the smallest card
}

// getCard returns the ASCII card, given its letter representation {[2-10],J,Q,K,A,r}
func getCard(arg string) string {
	var thisCard = ""
	switch arg {
	case "r":
		thisCard = randCard()
	case "2":
		thisCard = two
	case "3":
		thisCard = three
	case "4":
		thisCard = four
	case "5":
		thisCard = five
	case "6":
		thisCard = six
	case "7":
		thisCard = seven
	case "8":
		thisCard = eight
	case "9":
		thisCard = nine
	case "10":
		thisCard = ten
	case "J":
		thisCard = jack
	case "11": //from rand num gen
		thisCard = jack
	case "Q":
		thisCard = queen
	case "12": //from rand num gen
		thisCard = queen
	case "K":
		thisCard = king
	case "13": //from rand num gen
		thisCard = king
	case "A":
		thisCard = ace
	case "14": //from rand num gen
		thisCard = ace
		/*case "card": //for testing
		thisCard = card*/
	}
	return thisCard
}

// validateArgs terminates the program if invalid command line arguments are provided.
func validateArgs() {
	//check arg count
	if len(os.Args[1:]) < 2 {
		fmt.Fprintf(os.Stderr, "dealcards: Invalid number of args provided.\n")
		os.Exit(1)
	}

	if os.Args[1] != "v" && os.Args[1] != "h" {
		fmt.Fprintf(os.Stderr, "dealcards: Invalid args provided: %s.\n", os.Args[1])
		fmt.Fprintf(os.Stderr, "dealcards: Please provide a valid first argument: {v,h}.\n")
		os.Exit(1)
	}

	//check args
	for _, arg := range os.Args[2:] {
		//use regex package here?
		if arg != "2" && arg != "3" && arg != "4" && arg != "5" && arg != "6" && arg != "7" &&
			arg != "8" && arg != "9" && arg != "10" && arg != "J" && arg != "Q" && arg != "K" &&
			arg != "A" && arg != "r" {
			fmt.Fprintf(os.Stderr, "dealcards: Invalid args provided: %s.\n", arg)
			fmt.Fprintf(os.Stderr, "dealcards: Please provide valid args: {[2-10],J,Q,K,A,r}.\n")
			os.Exit(1)
		}
	}
}
