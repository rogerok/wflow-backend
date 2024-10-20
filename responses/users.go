package responses

type TokensModel struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}
