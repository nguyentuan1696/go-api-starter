package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"go-api-starter/core/logger"
	"go-api-starter/core/params"
	"go-api-starter/modules/product/entity"
)

func (r *ProductRepository) PublicGetProvinces(ctx context.Context, params params.QueryParams) (*entity.PaginatedProvinceEntity, error) {
	// Tính offset cho pagination
	offset := (params.PageNumber - 1) * params.PageSize

	// Xây dựng điều kiện WHERE cho search
	whereClause := ""
	args := []interface{}{}
	argIndex := 1

	if params.Search != "" {
		whereClause = "WHERE name ILIKE $" + fmt.Sprintf("%d", argIndex)
		args = append(args, "%"+params.Search+"%")
		argIndex++
	}

	// Query để đếm tổng số records
	countQuery := "SELECT COUNT(*) FROM provinces " + whereClause
	var totalItems int
	err := r.DB.GetContext(ctx, &totalItems, countQuery, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("ProductRepository:GetProvinces - Count", err)
		return nil, err
	}

	// Query để lấy data với pagination
	dataQuery := "SELECT * FROM provinces " + whereClause + " ORDER BY code ASC OFFSET $" + fmt.Sprintf("%d", argIndex) + " ROWS FETCH NEXT $" + fmt.Sprintf("%d", argIndex+1) + " ROWS ONLY"
	args = append(args, offset, params.PageSize)

	var provinces []entity.Province
	err = r.DB.SelectContext(ctx, &provinces, dataQuery, args...)
	if err != nil {
		logger.Error("ProductRepository:GetProvinces - Select", err)
		return nil, err
	}

	return &entity.PaginatedProvinceEntity{
		Items:      provinces,
		TotalItems: totalItems,
		PageNumber: params.PageNumber,
		PageSize:   params.PageSize,
	}, nil
}

func (r *ProductRepository) PublicGetDistricts(ctx context.Context, params params.QueryParams) (*entity.PaginatedDistrictEntity, error) {
	// Tính offset cho pagination
	offset := (params.PageNumber - 1) * params.PageSize

	// Base query để lấy districts
	baseQuery := `FROM districts d`

	// Thêm điều kiện search nếu có
	var whereClause string
	var args []interface{}

	conditions := []string{}
	argIndex := 1

	// Thêm filter theo province_code nếu có trong params
	if provinceCode := params.Filters["province_code"]; provinceCode != "" {
		conditions = append(conditions, fmt.Sprintf("d.province_code = $%d", argIndex))
		args = append(args, provinceCode)
		argIndex++
	}

	// Thêm search theo tên district nếu có
	if params.Search != "" {
		conditions = append(conditions, fmt.Sprintf("d.name ILIKE $%d", argIndex))
		args = append(args, "%"+params.Search+"%")
		argIndex++
	}

	if len(conditions) > 0 {
		whereClause = " WHERE " + strings.Join(conditions, " AND ")
	}

	// Query để đếm tổng số records
	countQuery := `SELECT COUNT(*) ` + baseQuery + whereClause
	var totalItems int
	err := r.DB.GetContext(ctx, &totalItems, countQuery, args...)
	if err != nil {
		logger.Error("ProductRepository:GetDistricts - Count", err)
		return nil, err
	}

	// Query để lấy data với pagination
	dataQuery := `
		SELECT 
			d.*
		` + baseQuery + whereClause + `
		ORDER BY d.name ASC
		OFFSET $` + fmt.Sprintf("%d", argIndex) + ` ROWS FETCH NEXT $` + fmt.Sprintf("%d", argIndex+1) + ` ROWS ONLY`

	// Thêm params cho pagination
	args = append(args, offset, params.PageSize)

	var districts []entity.District
	err = r.DB.SelectContext(ctx, &districts, dataQuery, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("ProductRepository:GetDistricts - Select", err)
		return nil, err
	}

	return &entity.PaginatedDistrictEntity{
		Items:      districts,
		TotalItems: totalItems,
		PageNumber: params.PageNumber,
		PageSize:   params.PageSize,
	}, nil
}

func (r *ProductRepository) PublicGetWards(ctx context.Context, params params.QueryParams) (*entity.PaginatedWardEntity, error) {
	// Tính offset cho pagination
	offset := (params.PageNumber - 1) * params.PageSize

	// Base query để lấy wards
	baseQuery := `FROM wards w`

	// Thêm điều kiện search nếu có
	var whereClause string
	var args []interface{}

	conditions := []string{}
	argIndex := 1

	// Thêm filter theo district_code nếu có trong params
	if districtCode := params.Filters["district_code"]; districtCode != "" {
		conditions = append(conditions, fmt.Sprintf("w.district_code = $%d", argIndex))
		args = append(args, districtCode)
		argIndex++
	}

	// Thêm search theo tên ward nếu có
	if params.Search != "" {
		conditions = append(conditions, fmt.Sprintf("w.name ILIKE $%d", argIndex))
		args = append(args, "%"+params.Search+"%")
		argIndex++
	}

	if len(conditions) > 0 {
		whereClause = " WHERE " + strings.Join(conditions, " AND ")
	}

	// Query để đếm tổng số records
	countQuery := `SELECT COUNT(*) ` + baseQuery + whereClause
	var totalItems int
	err := r.DB.GetContext(ctx, &totalItems, countQuery, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("ProductRepository:GetWards - Count", err)
		return nil, err
	}

	// Query để lấy data với pagination
	dataQuery := `
		SELECT 
			w.*
		` + baseQuery + whereClause + `
		ORDER BY w.name ASC
		OFFSET $` + fmt.Sprintf("%d", argIndex) + ` ROWS FETCH NEXT $` + fmt.Sprintf("%d", argIndex+1) + ` ROWS ONLY`

	// Thêm params cho pagination
	args = append(args, offset, params.PageSize)

	var wards []entity.Ward
	err = r.DB.SelectContext(ctx, &wards, dataQuery, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("ProductRepository:GetWards - Select", err)
		return nil, err
	}

	return &entity.PaginatedWardEntity{
		Items:      wards,
		TotalItems: totalItems,
		PageNumber: params.PageNumber,
		PageSize:   params.PageSize,
	}, nil
}
