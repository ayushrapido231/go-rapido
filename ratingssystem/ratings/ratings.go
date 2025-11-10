package ratings

import (
	"errors"
	"fmt"
	"time"
)

type UserRating struct {
	UID      int
	Rating   float64
	Comments Comment
}

type Comment struct {
	Comment string
	Date    time.Time
}

type Rating struct {
	ID        int
	AvgRating float64
	Ratings   []UserRating // slice of UserRating
}

// Implement String() for UserRating : Value receiver - no star
func (u UserRating) String() string {
	return fmt.Sprintf(("\n User ID: %d\n Rating: %.1f\n %s"), u.UID, u.Rating, u.Comments)
}

// Implement String() for Comment
func (c Comment) String() string {
	return fmt.Sprintf("Comment: %q, Date: %s", c.Comment, c.Date.Format("2006-01-02 15:04:05"))
}

// Implement String() for Rating
func (r Rating) String() string {
	return fmt.Sprintf("Rating{id: %d, avgRating: %.2f, ratings: %v}", r.ID, r.AvgRating, r.Ratings)
}

// Add adds a new rating to the Rating struct
// Return an error if the rating is negative
func (r *Rating) Add(uid int, rating float64, comment string) error {
	if rating < 0 {
		return errors.New("rating cannot be negative")
	}

	newRating := UserRating{
		UID:    uid,
		Rating: rating,
		Comments: Comment{
			Comment: comment,
			Date:    time.Now(),
		},
	}
	r.Ratings = append(r.Ratings, newRating)
	r.calculateAvgRating()
	return nil
}

// calculateAvgRating recalculates the average rating
func (r *Rating) calculateAvgRating() {
	if len(r.Ratings) == 0 {
		r.AvgRating = 0
		return
	}

	var sum float64 = 0.0

	// Loop through each rating in the ratings slice
	for _, userRating := range r.Ratings {
		// Add each rating value to the sum
		sum = sum + userRating.Rating
	}

	numberOfRatings := float64(len(r.Ratings))
	r.AvgRating = sum / numberOfRatings
}
