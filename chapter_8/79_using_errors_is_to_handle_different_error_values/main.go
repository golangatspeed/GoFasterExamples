package main

import (
	"errors"
	"fmt"
)

var ErrScoreLessThanMin = errors.New("score is less than minimum")
var ErrScoreOverMax = errors.New("score is greater than maximum")

func checkRating(score int) (int, error) {
	min := 0
	max := 5

	if score < min {
		return 0, ErrScoreLessThanMin
	}
	if score > max {
		return 0, ErrScoreOverMax
	}

	return score, nil
}

func main() {
	rating, err := checkRating(-1)
	if err != nil {
		switch {
		case errors.Is(err, ErrScoreLessThanMin):
			fmt.Println("rating score too low")
		case errors.Is(err, ErrScoreOverMax):
			fmt.Println("rating score too high")
		default:
			fmt.Printf("unexpected error: %s\n", err)
		}
		return
	}
	fmt.Printf("Rating : %d\n", rating)
}