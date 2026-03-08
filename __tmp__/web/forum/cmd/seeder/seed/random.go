package seed

import (
	"fmt"

	"forum.id/internal/domain"
	"golang.org/x/exp/rand"
)

func RandomName() string {
	firstNames := []string{"John", "Jane", "Alex", "Emily", "Chris", "Michael", "Sarah", "David", "Laura", "James"}
	lastNames := []string{"Doe", "Smith", "Johnson", "Brown", "Williams", "Jones", "Garcia", "Miller", "Davis", "Martinez"}

	firstName := firstNames[rand.Intn(len(firstNames))]
	lastName := lastNames[rand.Intn(len(lastNames))]

	return firstName + " " + lastName
}

func RandomUsername(length int) string {
	letterPart := RandomString(length)
	numberPart := rand.Intn(9999-1000) + 1000 // Ensures the number is always 4 digits

	return fmt.Sprintf("%s%d", letterPart, numberPart)
}

func RandomEmail(length int) string {
	domains := []string{"example.com", "test.com", "mail.com"}
	username := RandomString(length)
	domain := domains[rand.Intn(len(domains))]
	return username + "@" + domain
}

func RandomVoteType() domain.VoteType {
	voteTypes := []domain.VoteType{domain.UpVote, domain.DownVote, domain.NeutralVote}
	return voteTypes[rand.Intn(len(voteTypes))]
}

func RandomString(length int) string {
	if length <= 0 {
		length = 10
	}

	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]rune, length)

	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func RandomPassword(length int) string {
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	special := "!@#$%^&*"

	all := lower + upper + digits + special
	result := make([]rune, length)

	result[0] = rune(lower[rand.Intn(len(lower))])
	result[1] = rune(upper[rand.Intn(len(upper))])
	result[2] = rune(digits[rand.Intn(len(digits))])
	result[3] = rune(special[rand.Intn(len(special))])

	for i := 4; i < length; i++ {
		result[i] = rune(all[rand.Intn(len(all))])
	}

	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return string(result)
}

func RandomInt64(max int64) int64 {
	if 0 > max {
		panic("Invalid range: min must be less than or equal to max")
	}
	return rand.Int63n(max-0+1) + 0
}
