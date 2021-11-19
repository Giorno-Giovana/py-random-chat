package convert

import (
	"gitlab.com/procsy/attorney/api/internal/model/core"
	"gitlab.com/procsy/attorney/api/internal/model/dto"
)

func User2Core(user *dto.User) core.User {
	return core.User{
		ID:       core.UserID(user.ID),
		Email:    user.Email,
		Username: user.Username,
	}
}

func User2DTO(user *core.User) dto.User {
	return dto.User{
		ID:       string(user.ID),
		Email:    user.Email,
		Username: user.Username,
	}
}
