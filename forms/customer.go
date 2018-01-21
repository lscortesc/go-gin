package forms

// RegisterForm Type
type RegisterForm struct {
	Firstname string `form:"firstname" json:"firstname" binding:"required"`
	Lastname  string `form:"lastname" json:"lastname" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required,email"`
	Password  string `form:"password" json:"password" binding:"required"`
	CountryID string `form:"country_id" json:"country_id" binding:"required"`
}
