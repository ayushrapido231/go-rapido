package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type Rating struct {
	ID       string
	Reviewer string
	Ratings  int
	Feedback string
}

// Note For Revision :
//  Function : Has no Receiver (e.g. : IsValidRating below)
//  Method : Has a Receiver (e.g.: TakeUserInput, ShowRating -> both have r *Rating)

// Taking User Input here and storing it in the Rating struct
func (r *Rating) TakeUserInput() {
	fmt.Println("Enter the product Id:")
	fmt.Scanln(&r.ID)

	fmt.Println("Enter your Username:")
	fmt.Scanln(&r.Reviewer)

	fmt.Println("Enter your rating (1-5):")
	fmt.Scanln(&r.Ratings)

	fmt.Println("Enter your comment:")
	fmt.Scanln(&r.Feedback)
}

func (r *Rating) ShowRating() {
	stars := strings.Repeat("🌟", r.Ratings)
	fmt.Printf("\nReview by %s:\n", r.Reviewer)
	fmt.Printf("Rating: %d stars %s\n", r.Ratings, stars)
	fmt.Printf("Comment: %s\n", r.Feedback)

	if r.Ratings >= 4 {
		color.Green("Thank you for your positive feedback!")
	} else if r.Ratings == 3 {
		color.Yellow("Thank you for your feedback!")
	} else {
		color.Red("We appreciate your feedback and will get back to you soon!")
	}
}

// Checking if the rating is valid or not [It should be between 1 and 5]
func IsValidRating(rating int) bool {
	return rating >= 1 && rating <= 5
}

func main() {
	var review *Rating = &Rating{} // review := &Rating{} is also valid

	review.TakeUserInput()

	if !IsValidRating(review.Ratings) {
		color.Red("Invalid rating. Please enter a rating between 1 and 5.")
		return
	}

	review.ShowRating()
}
