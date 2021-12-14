package domain

import "github.com/google/uuid"

// SearchArgs - filter search
type SearchArgs struct {
	Name string `param:"name" query:"name" form:"name"`
}

// CreateArgs - create search
type CreateArgs struct {
	Name      string `json:"name" form:"name" query:"name" validate:"required"`
	Email     string `json:"email" form:"email" query:"email" validate:"required"`
	Age       int    `json:"age" form:"age" query:"age" validate:""`
	CompanyID int    `json:"companyId" form:"companyId" query:"companyId" validate:"required"`
}

// PutArgs - TODO
type PutArgs struct {
	ID   uuid.UUID `json:"id" param:"id"`
	Name string    `json:"name" form:"name" query:"name" validate:""`
}

func (p *PutArgs) get() {

}
