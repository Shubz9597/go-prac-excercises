package main

import (
	"bufio"
	"deck"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Hand struct {
	FaceUp bool
	Card   deck.Card
}

type Player struct {
	Name      string
	Hands     []Hand
	Points    int
	BetAmount int
	Stand     bool
}

func dealCards(cards []deck.Card, players []*Player) {
	cardIndex := 0
	dealerPos := len(players) - 1
	for i := 0; i < 2; i++ {
		for j := 0; j < len(players); j++ {
			faceUp := true
			if j == dealerPos && i == 1 {
				faceUp = false
			}
			players[j].Hands = append(players[j].Hands, Hand{FaceUp: faceUp, Card: cards[cardIndex]})
			cardIndex++
		}
	}
}

func createPlayers(numberOfPlayers int) []*Player {
	players := make([]*Player, numberOfPlayers+1)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < numberOfPlayers; i++ {
		fmt.Printf("Player %d name: ", i+1)
		scanner.Scan()
		name := scanner.Text()
		players[i] = &Player{
			Name:   name,
			Points: 500,
			Stand:  false,
		}
	}
	players[len(players)-1] = &Player{
		Name:   "Dealer",
		Points: math.MaxInt64,
		Stand:  false,
	}
	return players
}

func showCards(player *Player) {
	fmt.Println("You have:")
	for _, hand := range player.Hands {
		if hand.FaceUp {
			fmt.Print(hand.Card.String(), "   ")
		} else {
			fmt.Println("Hidden Card")
		}
	}
	fmt.Println()
}

func calculatePoints(hands []Hand) (int, bool) {
	points := 0
	aces := 0
	for _, hand := range hands {
		switch hand.Card.Face {
		case deck.Ace:
			points += 11
			aces++
		case deck.Jack, deck.Queen, deck.King:
			points += 10
		default:
			points += int(hand.Card.Face)
		}
	}
	for points > 21 && aces > 0 {
		points -= 10
		aces--
	}
	isSoft := (points == 17 && aces > 0)
	return points, isSoft
}

func hit(player *Player, cards *[]deck.Card) {
	player.Hands = append(player.Hands, Hand{FaceUp: true, Card: (*cards)[0]})
	*cards = (*cards)[1:]
	fmt.Println("You drew a", player.Hands[len(player.Hands)-1].Card.String())
	score, _ := calculatePoints(player.Hands)
	fmt.Println("Your score is", score)
}

func gameNew(player *Player, cards *[]deck.Card, scanner *bufio.Scanner) {
	fmt.Println("Player", player.Name, "turn")
	fmt.Printf("%s, enter your bet (you have %d points): ", player.Name, player.Points)
	scanner.Scan()
	bet, _ := strconv.Atoi(scanner.Text())
	player.BetAmount = bet

	for !player.Stand {
		showCards(player)
		currentPoints, _ := calculatePoints(player.Hands)
		if currentPoints == 21 {
			fmt.Println("Blackjack! You win!")
			player.Points += player.BetAmount * 2
			return
		} else if currentPoints > 21 {
			fmt.Println("You busted! You lose!")
			return
		} else {
			for {
				fmt.Println("\nDo you want to hit or stand, 1 for hit, 2 for stand")
				scanner.Scan()
				choice := strings.TrimSpace(scanner.Text())
				if choice == "1" {
					hit(player, cards)
					break
				} else if choice == "2" {
					fmt.Println("You chose to stand")
					player.Stand = true
					break
				} else {
					fmt.Println("Invalid input. Please enter 1 or 2.")
				}
			}
		}
	}
}

func dealerGame(dealer *Player, cards *[]deck.Card) {
	fmt.Println("Dealer's turn")
	dealer.Hands[1].FaceUp = true
	showCards(dealer)
	for {
		points, isSoft := calculatePoints(dealer.Hands)
		if points < 17 || (points == 17 && isSoft) {
			fmt.Printf("Dealer hits on %d (soft: %v)\n", points, isSoft)
			hit(dealer, cards)
		} else {
			fmt.Printf("Dealer stands on %d\n", points)
			break
		}
	}
	points, _ := calculatePoints(dealer.Hands)
	if points > 21 {
		fmt.Println("Dealer busted!")
	}
}

func determineWinners(players []*Player, dealer *Player) {
	dealerPoints, _ := calculatePoints(dealer.Hands)
	for _, player := range players[:len(players)-1] {
		playerPoints, _ := calculatePoints(player.Hands)
		fmt.Printf("%s: %d vs Dealer: %d\n", player.Name, playerPoints, dealerPoints)
		switch {
		case playerPoints > 21:
			fmt.Println("Busted! You lose.")
			player.Points -= player.BetAmount
		case dealerPoints > 21 || playerPoints > dealerPoints:
			fmt.Println("You win!")
			player.Points += player.BetAmount
		case playerPoints < dealerPoints:
			fmt.Println("You lose!")
			player.Points -= player.BetAmount
		default:
			fmt.Println("Push!")
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the number of players:")
	scanner.Scan()
	numberOfPlayers, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Creating Lobby...")
	players := createPlayers(numberOfPlayers)

	fmt.Println("Creating Deck...")
	cards := deck.New(deck.Shuffle)
	cardsPointer := &cards

	fmt.Println("Dealing Cards...")
	dealCards(cards, players)
	*cardsPointer = (*cardsPointer)[len(players)*2:]
	fmt.Print("\n")

	for _, player := range players[:len(players)-1] {
		gameNew(player, cardsPointer, scanner)
	}

	dealer := players[len(players)-1]
	dealerGame(dealer, cardsPointer)
	determineWinners(players, dealer)
}
