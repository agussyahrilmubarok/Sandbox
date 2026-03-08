package main

import (
	"flag"
	"time"

	"forum.id/cmd/seeder/seed"
	"forum.id/internal/app"
	"forum.id/internal/postgres"
	"forum.id/pkg/config"
)

func main() {
	cfg := config.LoadViper()
	logger := app.NewLogger(cfg)
	database := app.OpenPostgres(cfg, logger)

	userRepo := postgres.NewUserRepository(database, logger)
	threadRepo := postgres.NewThreadRepository(database, logger)
	commentRepo := postgres.NewCommentRepository(database, logger)
	voteRepo := postgres.NewVoteRepository(database, logger)

	totalFlag := flag.Int("total", 100, "Number of data to insert (default is 100)")
	cleanFlag := flag.Bool("clean", false, "Clean all users from the database")
	flag.Parse()

	if *cleanFlag {
		seed.CommentCleaner(commentRepo)
		time.Sleep(5 * time.Second)
		seed.ThreadCleaner(threadRepo)
		time.Sleep(5 * time.Second)
		seed.UserCleaner(userRepo)
	} else {
		seed.UserSeeder(*totalFlag, userRepo)
		time.Sleep(5 * time.Second)
		seed.ThreadSeeder(*totalFlag, threadRepo)
		time.Sleep(5 * time.Second)
		seed.CommentSeeder(*totalFlag, commentRepo)
		time.Sleep(5 * time.Second)
		seed.VoteThreadSeeder(*totalFlag, voteRepo)
		time.Sleep(5 * time.Second)
		seed.VoteCommentSeeder(*totalFlag, voteRepo)
	}
}
