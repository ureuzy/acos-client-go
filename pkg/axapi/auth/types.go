package auth

type Request struct {
	Credentials `json:"credentials"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	*AuthResponse `json:"authresponse"`
}

type AuthResponse struct {
	Signature   string `json:"signature"`
	Description string `json:"description"`
}
