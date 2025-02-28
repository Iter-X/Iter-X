package bo

type SignInResponse struct {
	Token        string
	RefreshToken string
	ExpiresIn    float64
}
