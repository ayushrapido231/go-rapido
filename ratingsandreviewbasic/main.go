package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

type ProductRating struct {
	UserID  string
	Rating  float64
	Comment string
	Date    time.Time
}

type Product struct {
	ID      int
	Ratings []ProductRating
}

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

	// Long-running process example with graceful exit

	r := Product{ // r is a Product struct
		ID:      1,
		Ratings: []ProductRating{}, // slice of ProductRating
	}

	scanner := bufio.NewScanner(os.Stdin) // to read standard input

	for { // infinite loop
		uid := readUserInput(scanner, "Enter user ID (or type 'exit' to quit): ")
		if isExit(uid) {
			color.Green("✔ Exiting program.")
			fmt.Printf("%+v\n", r)
			break
		}

		if addNewRating(scanner, &r, uid) {
			displayRatings(r)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

// Check if user wants to exit
func isExit(input string) bool {
	return input == "" || strings.ToLower(input) == "exit"
}

// Read user input
func readUserInput(scanner *bufio.Scanner, prompt string) string {
	fmt.Print(prompt)
	if !scanner.Scan() {
		return ""
	}
	return strings.TrimSpace(scanner.Text())
}

// Read and validate rating (1-5)
func readRating(scanner *bufio.Scanner) float64 {
	ratingStr := readUserInput(scanner, "Enter rating (1 to 5): ")
	if ratingStr == "" {
		return 0
	}

	rating, err := strconv.ParseFloat(ratingStr, 64)
	if err != nil {
		color.Red("✗ Invalid rating. Please enter a number.")
		return 0
	}

	if rating < 1 || rating > 5 {
		color.Red("▲ Rating must be between 1 and 5")
		return 0
	}

	return rating
}

// Add a new rating - returns true if successful
func addNewRating(scanner *bufio.Scanner, product *Product, userID string) bool {
	rating := readRating(scanner)
	if rating == 0 {
		return false // Invalid rating
	}

	comment := readUserInput(scanner, "Enter comment: ")
	if comment == "" {
		return false // No comment
	}

	addRating(product, userID, rating, comment)
	return true // Success
}

// Helper function to calculate average rating
func calculateAverage(ratings []ProductRating) float64 {
	if len(ratings) == 0 {
		return 0
	}
	var sum float64
	for _, r := range ratings {
		sum += r.Rating
	}
	return sum / float64(len(ratings))
}

// Helper function to add a rating to a product
func addRating(product *Product, userID string, rating float64, comment string) {
	newRating := ProductRating{
		UserID:  userID,
		Rating:  rating,
		Comment: comment,
		Date:    time.Now(),
	}
	product.Ratings = append(product.Ratings, newRating)
	color.Green("✔ Rating added successfully!")
}

// Helper function to display all ratings
func displayRatings(product Product) {
	avgRating := calculateAverage(product.Ratings)
	fmt.Printf("Product ID: %d, Avg Rating: %.1f, Ratings: [\n", product.ID, avgRating)
	for i, rt := range product.Ratings {
		if i > 0 {
			fmt.Println()
		}
		fmt.Printf("  User ID: %s, Rating: %.1f, Comment: \"%s\", Date: %s",
			rt.UserID, rt.Rating, rt.Comment, rt.Date.Format("2006-01-02 15:04:05"))
	}
	fmt.Println("\n]")
}
