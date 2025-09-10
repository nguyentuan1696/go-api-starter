package params

import (
	"go-api-starter/core/constants"
	"go-api-starter/core/utils"

	"github.com/labstack/echo/v4"
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
		PageNumber: utils.ToNumberWithDefault(c.QueryParam("page_number"), constants.DefaultPageNumber),
		PageSize:   utils.ToNumberWithDefault(c.QueryParam("page_size"), constants.DefaultPageSize),
		Search:     c.QueryParam("search"),
		Filters:    filters,
		OrderBy:    c.QueryParam("order_by"),
	}
}
