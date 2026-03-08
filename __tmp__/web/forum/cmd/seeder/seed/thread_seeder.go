package seed

import (
	"context"
	"fmt"
	"log"

	"forum.id/internal/domain"
)

func ThreadSeeder(total int, threadRepo domain.ThreadRepository) {
	var successCount int

	for i := 0; i < total*2; i++ {
		thread := domain.Thread{
			Title:    RandomString(10),
			Body:     RandomString(255),
			Category: RandomString(10),
			UserID:   RandomInt64(int64(total)),
		}

		_, err := threadRepo.Save(context.Background(), &thread)
		if err != nil {
			log.Printf("Failed to save thread %v: %v", thread.ID, err)
			continue // Skip to the next iteration on error
		}

		successCount++ // Increment the success counter
		log.Printf("Successfully inserted thread: %v", thread.ID)
	}

	fmt.Printf("Total threads inserted: %d\n", successCount)
}

func ThreadCleaner(threadRepo domain.ThreadRepository) {
	ctx := context.Background()

	// Fetch all threads
	threads, err := threadRepo.FetchAll(ctx)
	if err != nil {
		log.Printf("Failed to fetch threads: %v", err)
		return
	}

	var successCount int

	// Loop through each thread and delete them
	for _, thread := range threads {
		if err := threadRepo.Delete(ctx, thread); err != nil {
			log.Printf("Failed to delete thread %v: %v", thread.ID, err)
			continue // Skip to the next thread if there's an error
		}
		successCount++
		log.Printf("Successfully deleted thread: %v", thread.ID)
	}

	fmt.Printf("Total threads deleted: %d\n", successCount)
}
