package domain

// SearchArgs - filter search
type SearchArgs struct {
	Name  string `json:"name" form:"name" query:"name" validate:""`
	Email string `json:"email" form:"email" query:"email" validate:""`
}

// CreateArgs - create search
type CreateArgs struct {
	Name      string `json:"name" form:"name" query:"name" validate:"required"`
	Email     string `json:"email" form:"email" query:"email" validate:"required"`
	Age       int    `json:"age" form:"age" query:"age" validate:""`
	CompanyID int    `json:"companyId" form:"companyId" query:"companyId" validate:"required"`
}
