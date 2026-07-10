package main

import (
	"fmt"
	"time"
)

// Result represent data fetched from the core database
type Result struct {
	Data string
}

// SingleFlightEngine prevents redundant queries from hitting the database simultaneously
type SingleFlightEngine struct {
	ActiveQueries map[string]bool
}

// ExecuteCollapsingQuery simulates grouping overlapping requests into a single database hit
func (s *SingleFlightEngine) ExecuteCollapsingQuery(queryKey string) Result {
	fmt.Printf("[🚀 GO ENGINE] Incoming high-concurrency request for key: %s\n", queryKey)
	
	// Abstract Boundary: If the query is already in flight, collapse duplicate requests
	if s.ActiveQueries[queryKey] {
		fmt.Printf("[✔ GO ENGINE] Collapsing duplicate request! Waiting for active database query to return...\n")
		return Result{Data: "Shared_Cached_Result_From_Single_Hit"}
	}

	// Mark key as active
	s.ActiveQueries[queryKey] = true
	
	// Simulating the ONE actual physical database read that was allowed to go through
	time.Sleep(1 * time.Millisecond) 
	
	return Result{Data: "Fresh_Data_From_DB_Scan"}
}

func main() {
	fmt.Println("[*] Activating OmniOrigin Query Collapsing Gate...")
	engine := SingleFlightEngine{ActiveQueries: make(map[string]bool)}

	// Simulating 2 users hitting the exact same data endpoint at the exact same millisecond
	engine.ExecuteCollapsingQuery("SELECT * FROM products WHERE id = 101")
	engine.ExecuteCollapsingQuery("SELECT * FROM products WHERE id = 101")
}
