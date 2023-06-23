package auth

type AuthDto struct {
	Input    string `form:"input" json:"input" xml:"input" binding:"required,gte=6,lte=255"`
	Password string `form:"password" json:"password" xml:"password" binding:"required,gte=6,lte=16"`
}

type AuthResponseDto struct {
	Email  string `form:"email" json:"email" xml:"email"`
	Phone  string `form:"phone" json:"phone" xml:"phone"`
	Avatar string `form:"avatar" json:"avatar" xml:"avatar"`
	Token  string `json:"token"`
	// RefreshToken string `json:"refresh_token"`
}
