package model


type Product struct {
	Name  string    `json:"name" binding:"required"`
	Status  int		`json:"status" binding:"required"`
	CategoryId int	`json:"category_id" json:"omitempty"`
	Price float64	`json:"price" json:"omitempty"`
}

