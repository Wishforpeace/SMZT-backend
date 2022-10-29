package user

// loginRequest Login 请求
type loginRequest struct {
	StudentID string `json:"student_id" binding:"required"`
	Password  string `json:"password" binding:"required"`
} // @name loginRequest

// loginResponse Login 请求响应
type loginResponse struct {
	Token string `json:"token"`
} // @name loginResponse
