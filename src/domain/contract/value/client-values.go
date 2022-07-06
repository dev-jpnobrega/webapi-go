package domain

import "github.com/google/uuid"

// A SearchArgs represent filtering by fields
type SearchArgs struct {
	ID   uuid.UUID `param:"id" query:"id" form:"id"`
	Name string    `param:"name" query:"name" form:"name"`
}

// A CreateArgs represent fields to create client
type CreateArgs struct {
	Name      string `json:"name" form:"name" query:"name" validate:"required"`
	Email     string `json:"email" form:"email" query:"email" validate:"required"`
	Age       int    `json:"age" form:"age" query:"age" validate:""`
	CompanyID int    `json:"companyId" form:"companyId" query:"companyId" validate:"required"`
}

// A PutArgs represent fields to edit client
type PutArgs struct {
	ID   uuid.UUID `json:"id" param:"id"`
	Name string    `json:"name" form:"name" query:"name" validate:""`
}
