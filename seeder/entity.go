package main

import (
	"time"

	"gorm.io/gorm"
)

type DBProvince struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (DBProvince) TableName() string {
	return "provinces"
}

type DBCity struct {
	ID         int            `json:"id"`
	ProvinceID int            `json:"province_id"`
	Name       string         `json:"name"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

func (DBCity) TableName() string {
	return "cities"
}

type DBSubdistrict struct {
	ID         int            `json:"id"`
	ProvinceID int            `json:"province_id"`
	CityID     int            `json:"city_id"`
	Name       string         `json:"name"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

func (DBSubdistrict) TableName() string {
	return "subdistricts"
}

type DBVillage struct {
	ID            int            `json:"id"`
	ProvinceID    int            `json:"province_id"`
	CityID        int            `json:"city_id"`
	SubdistrictID int            `json:"subdistrict_id"`
	Name          string         `json:"name"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}

func (DBVillage) TableName() string {
	return "villages"
}

type Province struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type City struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Subdistrict struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Village struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
