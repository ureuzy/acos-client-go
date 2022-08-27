package auth

type Request struct {
	Credentials `json:"credentials"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	*Body `json:"authresponse"`
}

type Body struct {
	Signature   string `json:"signature"`
	Description string `json:"description"`
}
