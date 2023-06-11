package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	_ = godotenv.Load()

	db := InitGormDatabase()

	seedProvinces(db, "provinces.json")
}

func seedProvinces(db *gorm.DB, filename string) {
	provinces := make([]Province, 0)
	filePath := fmt.Sprintf("../../data/%s", filename)
	readFileToJson(filePath, &provinces)

	for _, province := range provinces {
		dbProvinceId := insertProvince(db, province)
		seedCities(db, province.ID, dbProvinceId)
	}
}

func seedCities(db *gorm.DB, provinceId int, dbProvinceId int) {
	cities := make([]City, 0)
	filePath := fmt.Sprintf("../../data/cities/%d.json", provinceId)
	readFileToJson(filePath, &cities)

	for _, city := range cities {
		dbCityId := insertCity(db, city, dbProvinceId)
		seedSubdistricts(db, provinceId, city.ID, dbProvinceId, dbCityId)
	}
}

func seedSubdistricts(db *gorm.DB, provinceId int, cityId int, dbProvinceId int, dbCityId int) {
	subdistricts := make([]Subdistrict, 0)
	filePath := fmt.Sprintf("../../data/subdistricts/%d_%d.json", provinceId, cityId)
	readFileToJson(filePath, &subdistricts)

	for _, subdistrict := range subdistricts {
		dbSubdistrictId := insertSubdistrict(db, subdistrict, dbProvinceId, dbCityId)
		seedVillages(db, provinceId, cityId, subdistrict.ID, dbProvinceId, dbCityId, dbSubdistrictId)
	}
}

func seedVillages(db *gorm.DB, provinceId int, cityId int, subdistrictId int, dbProvinceId int, dbCityId int, dbSubdistrictId int) {
	villages := make([]Village, 0)
	filePath := fmt.Sprintf("../../data/villages/%d_%d_%d.json", provinceId, cityId, subdistrictId)
	readFileToJson(filePath, &villages)

	for _, village := range villages {
		insertVillage(db, village, dbProvinceId, dbCityId, dbSubdistrictId)
	}
}
