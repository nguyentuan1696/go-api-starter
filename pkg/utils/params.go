package utils

import (
	"github.com/labstack/echo/v4"
)

const (
	DefaultPageNumber = 1
	DefaultPageSize   = 10
)

type QueryParams struct {
	PageNumber int
	PageSize   int
	Search     string
	Filters    map[string]string
	OrderBy    string
}

func NewQueryParams(c echo.Context) *QueryParams {
	filters := make(map[string]string)

	// Lấy province_code từ query parameter nếu có
	if provinceCode := c.QueryParam("province_code"); provinceCode != "" {
		filters["province_code"] = provinceCode
	}

	if districtCode := c.QueryParam("district_code"); districtCode != "" {
		filters["district_code"] = districtCode
	}

	return &QueryParams{
		PageNumber: ToNumberWithDefault(c.QueryParam("page_number"), DefaultPageNumber),
		PageSize:   ToNumberWithDefault(c.QueryParam("page_size"), DefaultPageSize),
		Search:     c.QueryParam("search"),
		Filters:    filters,
		OrderBy:    c.QueryParam("order_by"),
	}
}
