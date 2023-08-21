package models

type AccountRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Response struct {
	Success bool   `json:"success" binding:"required"`
	Reason  string `json:"reason" binding:"required"`
}

func ResponseType() {
	type Response struct {
		Success bool   `json:"success" example:"true"`
		Reason  string `json:"reason" example:""`
	}
	type BadRequestResponse struct {
		Success bool   `json:"success" example:"false"`
		Reason  string `json:"reason" example:"failed reason"`
	}
}
