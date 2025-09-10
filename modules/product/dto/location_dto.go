package dto

import "go-api-starter/core/dto"

type Province struct {
	Code                   string  `json:"code"`
	Name                   string  `json:"name"`
	NameEn                 *string `json:"name_en"`
	FullName               string  `json:"full_name"`
	FullNameEn             *string `json:"full_name_en"`
	CodeName               *string `json:"code_name"`
	AdministrativeUnitID   *int    `json:"administrative_unit_id"`
	AdministrativeRegionID *int    `json:"administrative_region_id"`
}

type District struct {
	Code                 string  `json:"code"`
	Name                 string  `json:"name"`
	NameEn               *string `json:"name_en"`
	FullName             *string `json:"full_name"`
	FullNameEn           *string `json:"full_name_en"`
	CodeName             *string `json:"code_name"`
	ProvinceCode         *string `json:"province_code"`
	AdministrativeUnitID *int    `json:"administrative_unit_id"`
}

type Ward struct {
	Code                 string  `json:"code"`
	Name                 string  `json:"name"`
	NameEn               *string `json:"name_en"`
	FullName             *string `json:"full_name"`
	FullNameEn           *string `json:"full_name_en"`
	CodeName             *string `json:"code_name"`
	DistrictCode         *string `json:"district_code"`
	AdministrativeUnitID *int    `json:"administrative_unit_id"`
}

type PaginatedProvinceDTO = dto.Pagination[Province]
type PaginatedDistrictDTO = dto.Pagination[District]
type PaginatedWardDTO = dto.Pagination[Ward]
