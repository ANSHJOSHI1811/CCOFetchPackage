package services

import (
	"cco_backend/config"
	"cco_backend/models"
	"cco_backend/utils"
	"fmt"
	"log"
	"time"
)

func ImportTermsData() error {
	// Prices API base URL
	basePriceApiUrl := "https://prices.azure.com/api/retail/prices?api-version=2023-01-01-preview&$filter=serviceName%20eq%20%27Virtual%20Machines%27"

	nextPageUrl := basePriceApiUrl
	totalPagesFetched := 0 //keeps track of how many pages have been fetched from api

	for nextPageUrl != "" { //for pagination
		// Fetch pricing data for the current page
		priceData, err := utils.FetchData(nextPageUrl) //calls fetch data function to fetch data
		if err != nil {
			return fmt.Errorf("error fetching price data: %w", err)
		}

		// Parse data
		priceItems, ok := priceData["Items"].([]interface{})
		if !ok {
			return fmt.Errorf("invalid format for price items")
		}

		// Process each price item
		for _, priceItemInterface := range priceItems {
			priceItem, ok := priceItemInterface.(map[string]interface{})
			if !ok {
				log.Printf("Skipping invalid price item: %v", priceItemInterface)
				continue
			}

			// Extract required fields from the price API
			skuID, _ := priceItem["skuId"].(string)

			// Find the corresponding SKU in the database
			sku := models.Sku{}
			if err := config.DB.Where("sku_id = ?", skuID).First(&sku).Error; err != nil {
				log.Printf("SKU not found for skuId: %s, skipping...", skuID)
				continue
			}

			// Find or create the corresponding price record
			priceRecord := models.Price{}
			priceID := 0 // Initialize as 0, will be updated if price exists
			// You can adjust the condition based on what data you have available
			if err := config.DB.Where("sku_id = ?", skuID).First(&priceRecord).Error; err != nil {
				// Insert the price record if it doesn't exist
				priceRecord = models.Price{
					SkuID: int(sku.ID), // Ensure this matches your foreign key type
					// Add any other necessary fields for the priceRecord
				}
				if err := config.DB.Create(&priceRecord).Error; err != nil {
					log.Printf("Error creating price record for skuId: %s, error: %v", skuID, err)
					continue
				}
				priceID = priceRecord.PriceID
				log.Printf("Created new price record for skuId: %s", skuID)
			} else {
				priceID = priceRecord.PriceID
			}

			// Extract savingsPlan from the price API
			savingsPlans, ok := priceItem["savingsPlan"].([]interface{})
			if !ok {
				log.Printf("No savings plan available for skuId: %s, skipping...", skuID)
				continue
			}

			// Process each savings plan
			for _, planInterface := range savingsPlans {
				plan, ok := planInterface.(map[string]interface{})
				if !ok {
					log.Printf("Skipping invalid savings plan for skuId: %s", skuID)
					continue
				}

				leaseContractLength, _ := plan["term"].(string)

				// Create a new Term entry
				term := models.Term{
					PriceID:             uint(priceID), // Convert int to uint
					SkuID:               int(sku.ID),  // Convert int to uint
					OfferTermCode:       nil,            // Set to NULL
					PurchaseOption:      nil,            // Null as specified
					OfferingClass:       nil,            // Null as specified
					LeaseContractLength: &leaseContractLength, // Nullable field
					CreatedDate:         time.Now(),     // Automatically generated
					ModifiedDate:        time.Now(),     // Automatically generated
					DisableFlag:         false,          // Default value
				}

				// Insert the Term into the database
				result := config.DB.Create(&term)
				if result.Error != nil {
					log.Printf("Error inserting term for skuId: %s, error: %v", skuID, result.Error)
				} else {
					log.Printf("Term inserted successfully for skuId: %s with lease_contract_length: %s", skuID, leaseContractLength)
				}
			}
		}

		// Increment total pages fetched
		totalPagesFetched++

		// Get the next page URL
		nextPageUrl, _ = priceData["NextPageLink"].(string)
		log.Printf("Next page URL: %s", nextPageUrl)

		// Optional delay between requests to avoid rate limiting
		time.Sleep(2 * time.Second)
	}

	log.Println("Terms data import completed successfully.")
	return nil
}