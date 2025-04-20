package main

import (
	"deck"
	"fmt"
	"math"
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
}

func dealCards(cards []deck.Card, players []*Player) {

	dealerPos := len(players) - 1
	var playerIndex int
	for i := 0; i < len(players)*2; i++ {
		playerIndex = i % len(players)
		players[playerIndex].Hands = append(players[playerIndex].Hands, Hand{FaceUp: true, Card: cards[i]})
		if playerIndex == dealerPos {
			if i == len(players)*2-1 {
				players[playerIndex].Hands[1].FaceUp = false
			}
		}
	}

}

func createPlayers(numberOfPlayers int) []*Player {
	players := make([]*Player, numberOfPlayers+1)
	for i := 0; i < numberOfPlayers; i++ {
		fmt.Println("Player ", i+1, " name")
		var name string
		n, err := fmt.Scanf("%s\n", &name)
		if err != nil || n != 1 {
			// handle invalid input
			fmt.Println(n, err)
			return []*Player{}
		}
		players[i] = &Player{
			Name:   name,
			Points: 500,
		}
	}
	players[len(players)-1] = &Player{
		Name:   "Dealer",
		Points: math.MaxInt64,
	}
	return players
}

func showCards(players []*Player) {
	for _, player := range players {
		fmt.Println("\n", player.Name, "has:")
		for _, hand := range player.Hands {
			if hand.FaceUp {
				fmt.Print(hand.Card.String(), "   ")
			} else {
				fmt.Println("Hidden Card")
			}
		}
	}
}

func calculatePoints(hands []Hand) int {
	//This function will calculate the points of a player based on the cards they have
	//if it is a face value, it will return 10 card, if it is an ace, it will return 11 or 1 depending on the other cards
	//if it is a number card, it will return the number of the card
	points := 0
	for _, hand := range hands {
		switch hand.Card.Face {
		case deck.Ace:
			if points+11 > 21 {
				points += 1
			} else {
				points += 11
			}
		case deck.Jack, deck.Queen, deck.King:
			points += 10
		default:
			points += int(hand.Card.Face)
		}
	}
	return points
}

func hit(player *Player, cards []deck.Card) {

}

func game(players []*Player, cards []deck.Card) {
	//This function will start the game and calculate points
	//We will create a loop to ask for the players to hit or stay
	//If they hit, we will deal them another card and calculate their points again
	//If they stay, we will move on to the next player
	//If the dealer has 17 or more points, they will not hit anymore
	//If the dealer has less than 17 points, they will hit until they have 17 or more points
	//At the end of the game, we will show the winner and their points
	round := 1
	for _, player := range players {
		fmt.Println(player.Name+"'s", "turn")
		fmt.Println("You score is", calculatePoints(player.Hands))
		fmt.Println("Do you want to hit or stand")

	}
}

func main() {
	//We will first create how many players to put in a blackjack game
	fmt.Println("Enter the number of players:")
	var numberOfPlayers int
	fmt.Scanf("%d\n", &numberOfPlayers)
	fmt.Println("Creating Lobby...")
	players := createPlayers(numberOfPlayers)
	fmt.Println("Creating Deck...")
	cards := deck.New(deck.Shuffle)
	fmt.Println("Dealing Cards...")
	dealCards(cards, players)
	showCards(players)

	fmt.Print("\n")
	game(players, cards[len(players)*2:])
	//Now we will create a function to start the game and calculate points
	// for _, player := range players[:len(players)-1] {
	// 	fmt.Println(player.Name, "has", calculatePoints(player.Hands), "points")
	// }

}
