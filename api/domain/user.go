package domain

type Users []User

type (
	User struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		Email          string `json:"email"`
		HashedPassword string `json:"-"`
	}

	UserForm struct {
		User
		Password string `json:"password"`
	}
)

func ConvertUserFormToUser(userForm UserForm) (user User) {
	user.ID = userForm.ID
	user.Name = userForm.Name
	user.Email = userForm.Email
	user.HashedPassword = userForm.HashedPassword
	return
}
