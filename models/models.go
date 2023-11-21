package models

import (
    "time"
    "gorm.io/gorm"
)

// Define a constant for the maximum length of the description.
const MaxDescriptionLength = 255

// Delosfarm represents the Farm entity
type Delosfarm struct {
    gorm.Model
    DelosfarmName     string    `gorm:"unique" json:"delosfarm_name"`
    FarmLocation      string    `json:"farm_location"`
    FarmSize          float64   `json:"farm_size"`
    FarmEstablishedAt time.Time `json:"farm_established_at"`
    FarmOwner         string    `json:"farm_owner"`
    FarmContactEmail  string    `gorm:"unique" json:"farm_contact_email"`
    FarmDescription   string    `gorm:"type:text" json:"farm_description"`
    Ponds             []Pond    `gorm:"foreignKey:FarmID" json:"ponds,omitempty"`
}

// Pond represents the Pond entity
type Pond struct {
    gorm.Model
    PondName            string  `gorm:"unique" json:"pond_name"`
    FarmID              uint    `gorm:"not null" json:"farmID"`
    PondSize            float64 `json:"pond_size"`
    PondDepth           float64 `json:"pond_depth"`
    PondTemperature     float64 `json:"pond_temperature"`
    PondWaterQuality    string  `json:"pond_water_quality"`
    PondStockingDensity float64 `json:"pond_stocking_density"`
    PondInletFlowRate   float64 `json:"pond_inlet_flow_rate"`
    PondOutletFlowRate  float64 `json:"pond_outlet_flow_rate"`
    PondWaterpH         float64 `json:"pond_water_ph"`
    PondDescription     string  `gorm:"type:text" json:"pond_description"`
    Farm                Delosfarm `gorm:"foreignkey:FarmID;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"farm,omitempty"`
}

func (f *Delosfarm) BeforeSave(tx *gorm.DB) error {
    // Ensure the description field does not exceed the maximum length.
    if len(f.FarmDescription) > MaxDescriptionLength {
        f.FarmDescription = f.FarmDescription[:MaxDescriptionLength]
    }
    return nil
}

func (p *Pond) BeforeSave(tx *gorm.DB) error {
    // Ensure the description field does not exceed the maximum length.
    if len(p.PondDescription) > MaxDescriptionLength {
        p.PondDescription = p.PondDescription[:MaxDescriptionLength]
    }
    return nil
}
