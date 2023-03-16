package dtos

type ProductDTO struct {
	Name  string `json:"name" form:"name" validate:"required,max=50"`
	Stock int16  `json:"stock" form:"stock"`
	Price int    `json:"price" form:"price" validate:"required"`
}
