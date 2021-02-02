package consts

const (
	// Application
	UserID = "user_id"
	ApiV1  = "/handler/v1"

	// utils
	// email regular
	EmailPattern    = `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	UserNamePattern = `^[a-zA-Z0-9]{4,12}$`
	PasswordPattern = `^[a-zA-Z0-9]{6,12}$`

	// verify filename is valid
	FileNamePattern = `^\s\\/:\*\?\"<>\|[^\s\\/:\\?\"<>\|\.]$`
)