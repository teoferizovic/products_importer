package model

type ProductImage struct {
	ID  		int		      `json:"id"`
	Path  		string    	  `json:"path" binding:"required"`
	ProductId  	int			  `json:"product_id" binding:"required"`
}


