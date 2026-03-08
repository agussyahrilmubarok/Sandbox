package seed

import (
	"context"
	"fmt"
	"log"

	"forum.id/internal/domain"
)

func CommentSeeder(total int, commentRepo domain.CommentRepository) {
	var successCount int

	for i := 0; i < total*10; i++ {
		comment := domain.Comment{
			Content:  RandomString(100),
			UserID:   RandomInt64(int64(total)),
			ThreadID: RandomInt64(int64(total)),
		}

		_, err := commentRepo.Save(context.Background(), &comment)
		if err != nil {
			log.Printf("Failed to save comment %v: %v", comment.ID, err)
			continue // Skip to the next iteration on error
		}

		successCount++ // Increment the success counter
		log.Printf("Successfully inserted comment: %v", comment.ID)
	}

	fmt.Printf("Total comments inserted: %d\n", successCount)
}

func CommentCleaner(commentRepo domain.CommentRepository) {
	ctx := context.Background()

	// Fetch all comments
	comments, err := commentRepo.FetchAll(ctx)
	if err != nil {
		log.Printf("Failed to fetch comments: %v", err)
		return
	}

	var successCount int

	// Loop through each comment and delete them
	for _, comment := range comments {
		if err := commentRepo.Delete(ctx, comment); err != nil {
			log.Printf("Failed to delete comment %v: %v", comment.ID, err)
			continue // Skip to the next comment if there's an error
		}
		successCount++
		log.Printf("Successfully deleted comment: %v", comment.ID)
	}

	fmt.Printf("Total comments deleted: %d\n", successCount)
}
