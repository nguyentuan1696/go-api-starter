package entity

import "go-api-starter/core/entity"

type Province struct {
	Code                   string  `db:"code"`
	Name                   string  `db:"name"`
	NameEn                 *string `db:"name_en"`
	FullName               string  `db:"full_name"`
	FullNameEn             *string `db:"full_name_en"`
	CodeName               *string `db:"code_name"`
	AdministrativeUnitID   *int    `db:"administrative_unit_id"`
	AdministrativeRegionID *int    `db:"administrative_region_id"`
}

type District struct {
	Code                 string  `db:"code"`
	Name                 string  `db:"name"`
	NameEn               *string `db:"name_en"`
	FullName             *string `db:"full_name"`
	FullNameEn           *string `db:"full_name_en"`
	CodeName             *string `db:"code_name"`
	ProvinceCode         *string `db:"province_code"`
	AdministrativeUnitID *int    `db:"administrative_unit_id"`
}

type Ward struct {
	Code                 string  `db:"code"`
	Name                 string  `db:"name"`
	NameEn               *string `db:"name_en"`
	FullName             *string `db:"full_name"`
	FullNameEn           *string `db:"full_name_en"`
	CodeName             *string `db:"code_name"`
	DistrictCode         *string `db:"district_code"`
	AdministrativeUnitID *int    `db:"administrative_unit_id"`
}

type PaginatedProvinceEntity = entity.Pagination[Province]
type PaginatedDistrictEntity = entity.Pagination[District]
type PaginatedWardEntity = entity.Pagination[Ward]
