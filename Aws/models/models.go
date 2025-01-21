package models
import (
	"time"
)
type RegionState struct {
    RegionName string `json:"region_name"`
    State      string `json:"state"`
}
type Provider struct {
	ProviderID   uint   `gorm:"primaryKey"`
	ProviderName string `gorm:"unique"`
}

type Service struct {
	ServiceID   uint   `gorm:"primaryKey"`
	ServiceName string
	ProviderID  uint `gorm:"not null;constraint:OnDelete:CASCADE;"` // Foreign key with cascade delete
}

type Region struct {
	RegionID   uint   `gorm:"primaryKey"`
	RegionCode string `gorm:"unique"`
	ServiceID  uint   `gorm:"not null;constraint:OnDelete:CASCADE;"` // Foreign key with cascade delete
}

type SKU struct {
	ID              uint   `gorm:"primaryKey"`
	SKUCode         string `gorm:"unique"`
	ProductFamily   string
	VCPU            int
	OperatingSystem string
	InstanceType    string
	Storage         string
	Network         string
	InstanceSKU     string
	Processor       string
	UsageType       string
	RegionID        uint `gorm:"not null;constraint:OnDelete:CASCADE;"` // Foreign key with cascade delete
}

type Price struct {
	PriceID       uint      `gorm:"primaryKey;autoIncrement"`
	SKU_ID        uint      `gorm:"not null;constraint:OnDelete:CASCADE;"` // Foreign key with cascade delete
	RateCode      string    `gorm:"type:varchar(255)"`
	EffectiveDate string    `gorm:"type:varchar(255)"`
	Unit          string    `gorm:"type:varchar(50)"`
	PricePerUnit  string    `gorm:"type:varchar(50)"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	ModifiedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DisableFlag   bool      `gorm:"default:false"`
}

type Term struct {
	OfferTermID         int    `gorm:"primaryKey;autoIncrement"`
	SKU_ID              uint   `gorm:"not null;constraint:OnDelete:CASCADE;"` // Foreign key with cascade delete
	PriceID             uint   `gorm:"not null;constraint:OnDelete:CASCADE;"` // Foreign key with cascade delete
	LeaseContractLength string `gorm:"size:255"`
	PurchaseOption      string `gorm:"size:255"`
	OfferingClass       string `gorm:"size:255"`
}

type SavingPlan struct {
	ID                  uint   `gorm:"primaryKey"`
	DiscountedSku       string
	Sku                 string
	LeaseContractLength string
	DiscountedRate      string
	RegionID            uint `gorm:"not null;constraint:OnDelete:CASCADE;"` // Foreign key with cascade delete
}


// ! Iska usage kya hai pata nhi filhal
type JSON map[string]interface{}
type PricingData struct {
	Products map[string]Product                           `json:"products"`
	Terms    map[string]map[string]map[string]TermDetails `json:"terms"`
}
type Product struct {
	SKU           string            `json:"sku"`
	ProductFamily string            `json:"productFamily"`
	Attributes    map[string]string `json:"attributes"`
}
type PricePerUnit struct {
	USD string `json:"USD"`
}
type PriceDimension struct {
	RateCode     string            `json:"rateCode"`
	Description  string            `json:"description"`
	BeginRange   string            `json:"beginRange"`
	EndRange     string            `json:"endRange"`
	Unit         string            `json:"unit"`
	PricePerUnit map[string]string `json:"pricePerUnit"`
	AppliesTo    []string          `json:"appliesTo"`
}
type TermAttributes struct {
	LeaseContractLength string `json:"LeaseContractLength"`
	PurchaseOption      string `json:"PurchaseOption"`
	OfferingClass       string `json:"OfferingClass"`
}
type TermDetails struct {
	OfferTermCode   string                    `json:"offerTermCode"`
	Sku             string                    `json:"sku"`
	EffectiveDate   string                    `json:"effectiveDate"`
	PriceDimensions map[string]PriceDimension `json:"priceDimensions"`
	TermAttributes  TermAttributes            `json:"termAttributes"` // This should be a struct, not a map
}

type SavingTermDetails struct {
	Sku                 string `json:"sku"`
	Description         string `json:"description"`
	EffectiveDate       string `json:"effectiveDate"`
	LeaseContractLength struct {
		Duration int    `json:"duration"`
		Unit     string `json:"unit"`
	} `json:"leaseContractLength"`
	Rates []RateDetails `json:"rates"`
}
type RateDetails struct {
	DiscountedSku         string `json:"discountedSku"`
	DiscountedUsageType   string `json:"discountedUsageType"`
	DiscountedOperation   string `json:"discountedOperation"`
	DiscountedServiceCode string `json:"discountedServiceCode"`
	RateCode              string `json:"rateCode"`
	Unit                  string `json:"unit"`
	DiscountedRate        struct {
		Price    string `json:"price"`
		Currency string `json:"currency"`
	} `json:"discountedRate"`
	DiscountedRegionCode   string `json:"discountedRegionCode"`
	DiscountedInstanceType string `json:"discountedInstanceType"`
}
type SavingData struct {
	TermsPlan struct {
		SavingsPlan []SavingTermDetails `json:"savingsPlan"`
	} `json:"terms"`
}
