package models

import "time"

type Provider struct {
	ProviderID   uint      `gorm:"primaryKey;autoIncrement"`
	ProviderName string    `gorm:"size:50;not null"`
	CreatedDate  time.Time `gorm:"default:current_timestamp"`
	ModifiedDate time.Time `gorm:"default:current_timestamp"`
	DisableFlag  bool      `gorm:"default:false"`
}

func (Provider) TableName() string {
	return "providers" // Explicitly specify the table name
}

type Region struct {
	RegionID    uint      `gorm:"primaryKey;autoIncrement"`
	ProviderID  uint      `gorm:"not null"`
	RegionCode  string    `gorm:"size:20;not null"`
	CreatedDate time.Time `gorm:"default:current_timestamp"`
	ModifiedDate time.Time `gorm:"default:current_timestamp"`
	DisableFlag bool      `gorm:"default:false"`
}

func (Region) TableName() string {
	return "regions" // Explicitly specify the table name
}

type Sku struct {
    ID                  uint      `gorm:"primaryKey;column:id"` // Change to ID
    RegionID            uint      `gorm:"column:region_id"`
    Armskuname          string    `gorm:"column:armskuname"`
    Name                string    `gorm:"column:name"`
    UsageType           string    `gorm:"column:type"`
    SkuCode             *string   `gorm:"column:sku_id_api"`
    ProductName         *string   `gorm:"column:product_name"`
    ProductFamily       *string   `gorm:"column:service_family"`
    VCPU                int       `gorm:"column:v_cpus"`
    Memory              string    `gorm:"column:memory_gb"`
    CpuArchitectureType string    `gorm:"column:cpu_architecture_type"`
    Network             string   `gorm:"column:max_network_interfaces"`
    CreatedAt           time.Time `gorm:"column:created_at"`
    UpdatedAt           time.Time `gorm:"column:modified_at"` 
    DisableFlag         bool      `gorm:"column:disable_flag"`
}

func (Sku) TableName() string {
    return "skus"
}

// Term represents the terms table
type Term struct {
    OfferTermID         uint       `gorm:"primaryKey"`
    OfferTermCode       *string    `gorm:"size:255"`
    PriceID             uint       `gorm:"not null"`
    SkuID               int        `gorm:"not null"`
    PurchaseOption      *string    `gorm:"size:100"`
    LeaseContractLength *string    `gorm:"size:50"`
    DiscountedSku       *string    `gorm:"size:255"`
    DiscountedRate      *float64   `gorm:"type:decimal(10,2)"`
    OfferingClass       *string    `gorm:"size:50"`
    CreatedDate         time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
    ModifiedDate        time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
    DisableFlag         bool       `gorm:"default:false"`
}

// TableName specifies the table name for Term
func (Term) TableName() string {
	return "terms"
}

type Price struct {
	PriceID       int       `gorm:"primaryKey;autoIncrement"`    // Primary Key, Auto-incremented
	SkuID         int       `gorm:"not null"`                    // Foreign key referencing sku table
	RetailPrice   float64   `gorm:"type:numeric(15,6)"` // Retail price (numeric field with precision)
	Unit          string    `gorm:"size:255;not null"`           // Unit of measurement
	EffectiveDate time.Time `gorm:"not null"`                    // Effective date for the price
	CreatedAt     time.Time `gorm:"default:current_timestamp"`   // Creation timestamp
	ModifiedAt    time.Time `gorm:"default:current_timestamp"`   // Last modification timestamp
	DisableFlag   bool      `gorm:"default:false"`               // Disable flag (defaults to false)
}

// TableName specifies the table name for Price
func (Price) TableName() string {
	return "prices"
}