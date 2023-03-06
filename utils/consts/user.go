package consts

// enums
const (
	E_USER_Mentor    string = "Mentor"
	E_USER_Placement string = "Placement"
	E_USER_People    string = "People"
	E_USER_Admission string = "Admission"
	E_USER_Academic  string = "Academic"
)

const (
	E_USER_User  string = "User"
	E_USER_Admin string = "Admin"
)

const (
	E_USER_Active    string = "Active"
	E_USERr_NotActive string = "Not-Active"
	E_USER_Deleted   string = "Deleted"
)

// Validation
const (
	USER_EmptyCredentialError string = "email and password must be filled"
)

// Bind Error
const (
	USER_ErrorBindUserData string = "error bind user data"
)

// Response Error
const (
	// Login
	USER_UserNotFound string = "user not found"
	USER_WrongPassword string = "wrong password"

	// Register
	USER_EmailAlreadyUsed string = "email is already used"

	// Select
	USER_FailedSelect string = "failed select user data"

	// Update
	USER_FailedUpdate string = "failed update user data"

	// Delete
	USER_FailedDelete string = "failed delete user data"
)

// Response Success
const (
	// Login
	USER_LoginSuccess string = "login succeed"

	// Register
	USER_RegisterSuccess string = "succesfully insert user data"

	// Select
	USER_SuccessReadUserData string = "succesfully read user data"

	// Update
	USER_SuccessUpdateUserData string = "succesfully update user data"

	// Delete
	USER_SuccessDelete string = "succesfully delete user"
)