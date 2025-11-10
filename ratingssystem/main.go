package main

import (
	"fmt"

	"ratingssystem/ratings"
)

func main() {
	rating := ratings.Rating{
		ID:        1,
		AvgRating: 0,
		Ratings:   []ratings.UserRating{},
	}

	// Add ratings using the Add method
	if err := rating.Add(1, 5, "Great product"); err != nil {
		fmt.Printf("Error adding rating: %v\n", err)
	}
	if err := rating.Add(2, -4, "Good product"); err != nil {
		fmt.Printf("Error adding rating: %v\n", err)
	}

	fmt.Println(rating)
}
