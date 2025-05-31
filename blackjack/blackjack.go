package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	card1Points := ParseCard(card1)
	card2Points := ParseCard(card2)
	playerSum := card1Points + card2Points
	dealerCardPoints := ParseCard(dealerCard)
	if card1 == "ace" && card2 == "ace" {
		return "P" // split
	} else if playerSum >= 17 && playerSum <= 20 {
		return "S" // stand
	} else if playerSum >= 12 && playerSum <= 16 {
		if dealerCardPoints >= 7 {
			return "H" // hit
		} else {
			return "S" // stand
		}
	} else if playerSum <= 11 {
		return "H" // hit
	} else if playerSum == 21 {
		isDealerHighCard := dealerCardPoints == 10 || dealerCardPoints == 11
		if !isDealerHighCard {
			return "W" // win
		} else {
			return "S" // stand
		}
	}
	return "" // if no case matches
}
