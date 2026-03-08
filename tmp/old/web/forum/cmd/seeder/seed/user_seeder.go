package seed

import (
	"context"
	"fmt"
	"log"

	"forum.id/internal/domain"
)

func UserSeeder(total int, userRepo domain.UserRepository) {
	var successCount int

	for i := 0; i < total; i++ {
		user := domain.User{
			Name:     RandomName(),
			Username: RandomUsername(10),
			Email:    RandomEmail(10),
			Password: RandomPassword(10),
		}

		_, err := userRepo.Save(context.Background(), &user)
		if err != nil {
			log.Printf("Failed to save user %s: %v", user.Username, err)
			continue // Skip to the next iteration on error
		}

		successCount++ // Increment the success counter
		log.Printf("Successfully inserted user: %s", user.Username)
	}

	fmt.Printf("Total users inserted: %d\n", successCount)
}

func UserCleaner(userRepo domain.UserRepository) {
	ctx := context.Background()

	// Fetch all users
	users, err := userRepo.FetchAll(ctx)
	if err != nil {
		log.Printf("Failed to fetch users: %v", err)
		return
	}

	var successCount int

	// Loop through each user and delete them
	for _, user := range users {
		if err := userRepo.Delete(ctx, user); err != nil {
			log.Printf("Failed to delete user %s: %v", user.Username, err)
			continue // Skip to the next user if there's an error
		}
		successCount++
		log.Printf("Successfully deleted user: %s", user.Username)
	}

	fmt.Printf("Total users deleted: %d\n", successCount)
}
