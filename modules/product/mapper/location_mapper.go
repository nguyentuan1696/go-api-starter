package mapper

import (
	"go-api-starter/modules/product/dto"
	"go-api-starter/modules/product/entity"
)

func ToProvinceDTO(entity *entity.Province) *dto.Province {
	return &dto.Province{
		Code:                   entity.Code,
		Name:                   entity.Name,
		NameEn:                 entity.NameEn,
		FullName:               entity.FullName,
		FullNameEn:             entity.FullNameEn,
		CodeName:               entity.CodeName,
		AdministrativeUnitID:   entity.AdministrativeUnitID,
		AdministrativeRegionID: entity.AdministrativeRegionID,
	}
}

func ToWardDTO(entity *entity.Ward) *dto.Ward {
	return &dto.Ward{
		Code:                 entity.Code,
		Name:                 entity.Name,
		NameEn:               entity.NameEn,
		FullName:             entity.FullName,
		FullNameEn:           entity.FullNameEn,
		CodeName:             entity.CodeName,
		DistrictCode:         entity.DistrictCode,
		AdministrativeUnitID: entity.AdministrativeUnitID,
	}
}

func ToDistrictDTO(entity *entity.District) *dto.District {
	return &dto.District{
		Code:                 entity.Code,
		Name:                 entity.Name,
		NameEn:               entity.NameEn,
		FullName:             entity.FullName,
		FullNameEn:           entity.FullNameEn,
		CodeName:             entity.CodeName,
		ProvinceCode:         entity.ProvinceCode,
		AdministrativeUnitID: entity.AdministrativeUnitID,
	}
}

func ToProvincePaginationDTO(entity *entity.PaginatedProvinceEntity) *dto.PaginatedProvinceDTO {

	provinceResponses := make([]dto.Province, len(entity.Items))
	for i, tag := range entity.Items {
		provinceResponses[i] = *ToProvinceDTO(&tag)
	}

	// Tính total pages
	totalPages := 0
	if entity.PageSize > 0 {
		totalPages = (entity.TotalItems + entity.PageSize - 1) / entity.PageSize
	}

	return &dto.PaginatedProvinceDTO{
		Items:      provinceResponses,
		TotalItems: entity.TotalItems,
		TotalPages: totalPages,
		PageNumber: entity.PageNumber,
		PageSize:   entity.PageSize,
	}
}

func ToWardPaginationDTO(entity *entity.PaginatedWardEntity) *dto.PaginatedWardDTO {

	wardResponses := make([]dto.Ward, len(entity.Items))
	for i, tag := range entity.Items {
		wardResponses[i] = *ToWardDTO(&tag)
	}

	// Tính total pages
	totalPages := 0
	if entity.PageSize > 0 {
		totalPages = (entity.TotalItems + entity.PageSize - 1) / entity.PageSize
	}

	return &dto.PaginatedWardDTO{
		Items:      wardResponses,
		TotalItems: entity.TotalItems,
		TotalPages: totalPages,
		PageNumber: entity.PageNumber,
		PageSize:   entity.PageSize,
	}
}

func ToDistrictPaginationDTO(entity *entity.PaginatedDistrictEntity) *dto.PaginatedDistrictDTO {

	wardResponses := make([]dto.District, len(entity.Items))
	for i, tag := range entity.Items {
		wardResponses[i] = *ToDistrictDTO(&tag)
	}

	// Tính total pages
	totalPages := 0
	if entity.PageSize > 0 {
		totalPages = (entity.TotalItems + entity.PageSize - 1) / entity.PageSize
	}

	return &dto.PaginatedDistrictDTO{
		Items:      wardResponses,
		TotalItems: entity.TotalItems,
		TotalPages: totalPages,
		PageNumber: entity.PageNumber,
		PageSize:   entity.PageSize,
	}
}
