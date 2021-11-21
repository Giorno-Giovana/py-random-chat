package dto

type TinderRegisterUserRequest struct {
	Email       string `json:"email"`
	Username    string `json:"username"`
	Mode        string `json:"mode"`
	Occupation  string `json:"occupation"`
	PlaceOfWork string `json:"place_of_work"`
	PhotoURL    string `json:"photo_url"`
}

type TinderRegisterUserResponse struct {
	UserID string `json:"user_id"`
}
