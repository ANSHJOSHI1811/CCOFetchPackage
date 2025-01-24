package main

import (
	"log"
	"ccofetchpackage/aws"   // Corrected to lowercase
	"ccofetchpackage/azure" // Corrected to lowercase
)

func main() {
	// Run AWS data-fetching code
	log.Println("Starting AWS data fetch...")
	err := aws.Run() // Use lowercase `aws` for the package name
	if err != nil {
		log.Fatalf("Error running AWS data fetch: %v", err)
	}
	log.Println("AWS data fetch completed.")

	// Run Azure data-fetching code
	log.Println("Starting Azure data fetch...")
	err = azure.Run() // Use lowercase `azure` for the package name
	if err != nil {
		log.Fatalf("Error running Azure data fetch: %v", err)
	}
	log.Println("Azure data fetch completed.")
}
