package services

import (
	"cco_backend/config"
	"cco_backend/models"
	"cco_backend/utils"
	"fmt"
	"log"
	"time"
)

func ImportPricesData() error {
	// Prices API URL (Initial URL to start fetching)
	priceApiUrl := "https://prices.azure.com/api/retail/prices?api-version=2023-01-01-preview&$filter=serviceName%20eq%20%27Virtual%20Machines%27"

	// Loop to handle pagination
	for {
		// Fetch price data
		priceData, err := utils.FetchData(priceApiUrl)
		if err != nil {
			return fmt.Errorf("error fetching price data: %w", err)
		}

		// Parse price items from the response
		priceItems, ok := priceData["Items"].([]interface{})
		if !ok {
			return fmt.Errorf("invalid format for price items")
		}

		// Iterate over each price item
		for _, priceItemInterface := range priceItems {
			priceItem := priceItemInterface.(map[string]interface{})

			// Extract required fields from the API response
			skuID, _ := priceItem["skuId"].(string)
			retailPrice, _ := priceItem["retailPrice"].(float64)
			unitOfMeasure, _ := priceItem["unitOfMeasure"].(string)
			effectiveStartDate, _ := priceItem["effectiveStartDate"].(string)

			// Find the corresponding SKU in the database
			sku := models.Sku{}
			if err := config.DB.Where("sku_id_api = ?", skuID).First(&sku).Error; err != nil {
				log.Printf("SKU not found for skuId: %s, skipping...", skuID)
				continue
			}

			// Parse the effective start date
			effectiveDate, err := time.Parse(time.RFC3339, effectiveStartDate)
			if err != nil {
				log.Printf("Invalid effective start date for skuId: %s, skipping...", skuID)
				continue
			}

			// Create a new Price entry
			price := models.Price{
				SkuID:        int(sku.ID),        // Foreign key referencing SKU table
				RetailPrice:  retailPrice,         // Retail price
				Unit:         unitOfMeasure,       // Unit of measurement
				EffectiveDate: effectiveDate,      // Effective date for the price
				CreatedAt:     time.Now(),         // Current timestamp for created date
				ModifiedAt:    time.Now(),         // Current timestamp for modified date
				DisableFlag:   false,              // Disable flag (defaults to false)
			}

			// Insert the Price into the database
			result := config.DB.Create(&price)
			if result.Error != nil {
				log.Printf("Error inserting price for skuId: %s, error: %v", skuID, result.Error)
			} else {
				log.Printf("Price inserted successfully for skuId: %s with retail price: %.6f", skuID, retailPrice)
			}
		}

		// Check for the next page using the NextPageLink field
		nextPageLink, exists := priceData["NextPageLink"].(string)
		if !exists || nextPageLink == "" {
			log.Println("All pages fetched successfully.")
			break
		}

		// Update the API URL to the next page URL for the next iteration
		priceApiUrl = nextPageLink
	}

	log.Println("Prices data import completed successfully.")
	return nil
}