package main

import (
	"fmt"
	"sync"
	"time"
)

//////////////////////////////////////////////////////////
// Step 1: Define the type to be implemented as a Singleton
//////////////////////////////////////////////////////////

// Singleton represents a type that will only have one instance
// throughout the lifecycle of the application.
// It can represent global configurations, database connections, or loggers.
type Singleton struct {
	CreatedAt time.Time // Used to demonstrate when the instance was created
}

//////////////////////////////////////////////////////////
// Step 2: Declare a global variable to hold the single instance
//////////////////////////////////////////////////////////

// `instance` holds the only Singleton object.
// This variable remains private within the package.
var instance *Singleton

//////////////////////////////////////////////////////////
// Step 3: Use sync.Once to ensure thread-safe initialization
//////////////////////////////////////////////////////////

// `sync.Once` guarantees that a function will only execute once,
// even when called from multiple goroutines concurrently.
var once sync.Once

//////////////////////////////////////////////////////////
// Step 4: Provide a global accessor function
//////////////////////////////////////////////////////////

// GetInstance returns the single instance of the Singleton.
// The instance is initialized only once, regardless of how many
// goroutines attempt to call this function simultaneously.
func GetInstance() *Singleton {
	once.Do(func() {
		fmt.Println("Creating new Singleton instance")
		instance = &Singleton{
			CreatedAt: time.Now(),
		}
	})
	return instance
}

//////////////////////////////////////////////////////////
// Step 5: Demonstrate the Singleton pattern in a concurrent context
//////////////////////////////////////////////////////////

func main() {
	const n = 10
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(id int) {
			defer wg.Done()
			if id%2 == 0 {
				time.Sleep(10 * time.Millisecond)
			}
			s := GetInstance()
			fmt.Printf("Goroutine %d -> instance pointer: %p, createdAt: %s\n",
				id, s, s.CreatedAt.Format(time.RFC3339Nano))
		}(i)
	}

	wg.Wait()
}

//////////////////////////////////////////////////////////
// SUMMARY OF SINGLETON DESIGN PATTERN STEPS
//////////////////////////////////////////////////////////

// Step 1: Define the structure that will act as the Singleton.
// Step 2: Declare a private global variable to hold the instance.
// Step 3: Use a synchronization mechanism (`sync.Once`) to ensure
//         the instance is created only once, even in concurrent scenarios.
// Step 4: Provide a public accessor function (`GetInstance`) to retrieve the instance.
// Step 5: Demonstrate usage in a concurrent environment to verify thread safety.

// Key Characteristics:
// - Only one instance exists throughout the program lifetime.
// - Thread-safe through `sync.Once`.
// - Provides global access to the instance.

// Appropriate use cases include:
// - Database connection management
// - Logging systems
// - Global configuration management

// Caution:
// The Singleton pattern introduces global state, which may complicate testing
// and reduce modularity if overused.

// References:
// https://refactoring.guru/design-patterns/singleton
