package main

import (
	"log"

	"gorm.io/gorm"
)

func insertProvince(db *gorm.DB, province Province) int {
	dbProvince := DBProvince{
		Name: province.Name,
	}

	err := db.Create(&dbProvince).Error
	if err != nil {
		log.Fatal(err)
	}
	return dbProvince.ID
}

func insertCity(db *gorm.DB, city City, provinceId int) int {
	dbCity := DBCity{
		ProvinceID: provinceId,
		Name:       city.Name,
	}

	err := db.Create(&dbCity).Error
	if err != nil {
		log.Fatal(err)
	}
	return dbCity.ID
}

func insertSubdistrict(db *gorm.DB, subdistrict Subdistrict, provinceId int, cityId int) int {
	dbSubdistrict := DBSubdistrict{
		ProvinceID: provinceId,
		CityID:     cityId,
		Name:       subdistrict.Name,
	}

	err := db.Create(&dbSubdistrict).Error
	if err != nil {
		log.Fatal(err)
	}
	return dbSubdistrict.ID
}

func insertVillage(db *gorm.DB, village Village, provinceId int, cityId int, subdistrictId int) int {
	dbVillage := DBVillage{
		ProvinceID:    provinceId,
		CityID:        cityId,
		SubdistrictID: subdistrictId,
		Name:          village.Name,
	}

	err := db.Create(&dbVillage).Error
	if err != nil {
		log.Fatal(err)
	}
	return dbVillage.ID
}
