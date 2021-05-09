package consts

const (
	// Application
	UserID = "user_id"
	User   = "user"
	ApiV1  = "/api/v1"

	// utils
	// email regular
	EmailPattern    = `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	UserNamePattern = `^[a-zA-Z0-9]{4,12}$`
	PasswordPattern = `^[a-zA-Z0-9]{6,12}$`

	// verify filename is valid
	//FileNamePattern = `^\s\\/:\*\?\"<>\|[^\s\\/:\\?\"<>\|\.]$`
	FileNamePattern = `[\\u4e00-\\u9fa5]{0,8}[a-zA-Z0-9]{0,8}$`
)

const (
	OssBucketPattern        = `[0-9a-z-]{3,63}`
	OssBucketGeneratePrefix = `melancholy-userid-`
	Https                   = "https://"
	RegionID                = "cn-shanghai" // 用于视频点播
)
