package models

// Validate => Checks for validation errors
// func (u *User) Validate(action string) map[string]string {
// 	var errorMessages = make(map[string]string)
// 	var err error

// 	switch strings.ToLower(action) {
// 	case "login":
// 		if u.Password == "" {
// 			err = errors.New("Password required")
// 			errorMessages["Password_Required"] = err.Error()
// 		}

// 		if u.Email == "" {
// 			err = errors.New("Email required")
// 			errorMessages["Email_Required"] = err.Error()
// 		}

// 	default:
// 		if u.Email == "" {
// 			err = errors.New("Email required")
// 			errorMessages["Email_Required"] = err.Error()
// 		}

// 		if u.Username == "" {
// 			err = errors.New("Username required")
// 			errorMessages["Username_Required"] = err.Error()
// 		}

// 		if u.Password == "" {
// 			err = errors.New("Password required")
// 			errorMessages["Password_Required"] = err.Error()
// 		}

// 		if len(u.Password) < 8 {
// 			err = errors.New("Invalid password length")
// 			errorMessages["Invalid_Password_Length"] = err.Error()
// 		}
// 	}

// 	return errorMessages
// }
