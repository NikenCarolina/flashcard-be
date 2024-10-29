package dto

type User struct {
	Username string `json:"username" binding:"required,gte=8,lte=15"`
	Password string `json:"password" binding:"required,omitempty,gte=8,lte=15"`
}

type Profile struct {
	Username string `json:"username" binding:"required,gte=8,lte=15"`
}
