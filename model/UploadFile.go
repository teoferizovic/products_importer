package model


type UploadFile struct {
	ID  int		          `json:"id"`
	Name  string    	  `json:"name" binding:"required"`
}

