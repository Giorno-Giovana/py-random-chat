package dto

// Only used in responses! Does not need validation
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type GetUserDataRequest struct {
	// empty
}

type GetUserDataResponse struct {
	User User `json:"user"`
}

type SearchUsersRequest struct {
	// query "selector"
}

type SearchUsersResponse struct {
	Users []User `json:"users"`
}
