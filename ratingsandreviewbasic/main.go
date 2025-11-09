package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	var productID string = "Rapido-Golang"
	var productName string = "Go session"
	var productDescription string = "Session by Adarsh"

	var reviewer string = "Ayush"
	var rating int = 5
	var comment string = "Excellent Session! "

	fmt.Printf("Product ID: %s\n", productID)
	fmt.Printf("Product Name: %s\n", productName)
	fmt.Printf("Description: %s\n\n", productDescription)

	fmt.Printf("Reviewer: %s\n", reviewer)
	fmt.Printf("Rating: %d stars\n", rating)
	fmt.Printf("Comment: %s\n", comment)

	var ratings []int
	fmt.Println("\n Collecting Ratings ")

	ratings = append(ratings, 5)
	fmt.Println("Ratings so far:", ratings)

	ratings = append(ratings, 4)
	fmt.Println("Ratings so far:", ratings)

	ratings = append(ratings, 3, 2, 1)
	fmt.Println("All ratings:", ratings)

	reviewers := []string{"Ayush"}
	reviewers = append(reviewers, "Adarsh", "Rapido")
	fmt.Println("\nReviewers:", reviewers)

	for i, r := range ratings {
		fmt.Printf("Rating %d: %d stars\n", i+1, r)
	}

	// Maps :

	map1 := make(map[string]int)
	map1["apple"] = 5
	map1["banana"] = 3
	fmt.Println(map1)

	v, exists := map1["orange"]
	if exists {
		fmt.Println(v)
	} else {
		fmt.Println("orange not found")
	}

	map1["orange"] = 7
	map1["pineapple"] = 10

	for key, value := range map1 {
		fmt.Println(key, value)
	}

	// Structs :
	type Review struct {
		Reviewer string
		Rating   int
		Comment  string
	}

	var review Review = Review{
		Reviewer: "Ayush",
		Rating:   5,
		Comment:  "Excellent Session On Go!",
	}

	fmt.Printf("\nReviewer: %s\n", review.Reviewer)
	fmt.Printf("Rating: %d stars\n", review.Rating)
	fmt.Printf("Comment: %s\n", review.Comment)

	//  if-else :
	for i, rating := range ratings {
		if rating >= 4 {
			fmt.Printf("Rating %d (%d stars): Recommended\n", i+1, rating)
		} else {
			fmt.Printf("Rating %d (%d stars): Not Recommended\n", i+1, rating)
		}
	}

	//   switch case :
	for i, rating := range ratings {
		var category string
		switch rating {
		case 5:
			category = "Excellent"
		case 4:
			category = "Good"
		case 3:
			category = "Average"
		case 2:
			category = "Poor"
		default:
			category = "Very Poor"
		}
		fmt.Printf("Rating %d (%d stars): %s\n", i+1, rating, category)
	}

	// Code where i am taking user input and Printing the stars
	// on the basis of the provided rating in the expected color.

	type Rating struct {
		ID       string
		Comment  string
		Stars    int
		Username string
	}

	var userRating Rating
	fmt.Println("Enter the product id:")
	fmt.Scanln(&userRating.ID)

	fmt.Println("Enter the rating (1-5):")
	fmt.Scanln(&userRating.Stars)

	fmt.Println("Enter the comment:")
	fmt.Scanln(&userRating.Comment)

	fmt.Println("Enter the username:")
	fmt.Scanln(&userRating.Username)

	if userRating.Stars > 3 {
		color.Green("Thanks for the Amazing review!")
	} else {
		color.Red("Sorry for Bad Experience , We Will look into the feedback")
	}

	// Switch case to display stars
	fmt.Print("Stars: ")
	switch userRating.Stars {
	case 1:
		fmt.Println("*")
	case 2:
		fmt.Println("**")
	case 3:
		fmt.Println("***")
	case 4:
		fmt.Println("****")
	case 5:
		fmt.Println("*****")
	default:
		fmt.Println("Invalid rating: Please enter a rating between 1 and 5")
	}

	fmt.Printf("\nThe rating for product id %s:\n  Rating: %d stars\n  Comment: %s\n  Username: %s\n", userRating.ID, userRating.Stars, userRating.Comment, userRating.Username)

}
