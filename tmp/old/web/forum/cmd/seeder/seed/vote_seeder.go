package seed

import (
	"context"
	"fmt"
	"log"

	"forum.id/internal/domain"
)

func VoteThreadSeeder(total int, threadRepo domain.VoteRepository) {
	var successCount int

	for i := 0; i < total*20; i++ {
		voted := domain.VoteThread{
			VoteType: string(RandomVoteType()),
			UserID:   RandomInt64(int64(total)),
			ThreadID: RandomInt64(int64(total)),
		}

		_, err := threadRepo.SetVoteThread(context.Background(), &voted)
		if err != nil {
			log.Printf("Failed to save vote thread %v: %v", voted.ID, err)
			continue // Skip to the next iteration on error
		}

		successCount++ // Increment the success counter
		log.Printf("Successfully inserted thread: %v", voted.ID)
	}

	fmt.Printf("Total threads inserted: %d\n", successCount)
}

func VoteCommentSeeder(total int, threadRepo domain.VoteRepository) {
	var successCount int

	for i := 0; i < total*20; i++ {
		voted := domain.VoteComment{
			VoteType:  string(RandomVoteType()),
			UserID:    RandomInt64(int64(total)),
			CommentID: RandomInt64(int64(total)),
		}

		_, err := threadRepo.SetVoteComment(context.Background(), &voted)
		if err != nil {
			log.Printf("Failed to save vote thread %v: %v", voted.ID, err)
			continue // Skip to the next iteration on error
		}

		successCount++ // Increment the success counter
		log.Printf("Successfully inserted thread: %v", voted.ID)
	}

	fmt.Printf("Total threads inserted: %d\n", successCount)
}
