package model

import "time"

type GeoProvince struct {
	ID              float64   `json:"id"`
	Code            string    `json:"code"`
	Name            string    `json:"name"`
	Flag            int       `json:"flag"`
	CreatedDate     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_date"`
	CreatedBy       string    `json:"created_by"`
	CreatedIp       string    `json:"createdIp"`
	LastUpdatedDate time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"last_updated_date"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedIp   string    `json:"last_updated_ip"`
}

type CreateGeoProvince struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type ModifyCreateGeoProvince struct {
	ID   float64 `json:"id"`
	Code string  `json:"code"`
	Name string  `json:"name"`
}
