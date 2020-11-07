package conf

type Config struct {
	Mode        string `json:"mode"`
	Application Application
	Database    Database
	Gorm        Gorm
}

type Application struct {
	Domain        string `json:"domain" yml:"Application.Domain"`
	Host          string `json:"host" yml:"Application.Host"`
	IsHttps       bool   `json:"isHttps" yml:"Application.IsHttps"`
	Name          string `json:"name" yml:"Application.Name"`
	Port          string `json:"port" yml:"Application.Port"`
	ReadTimeout   int    `json:"readTimeout" yml:"Application.ReadTimeout"`
	WriterTimeout int64  `json:"writeTimeout" yml:"Application.WriteTimeout"`
	Location      string `json:"melancholy" yml:"Application.Location"`
}

type Database struct {
	DbType   string `json:"dbType" yml:"Database:DbType"`
	Host     string `json:"host" yml:"Database:Host"`
	Port     string `json:"port" yml:"Database:Port"`
	Name     string `json:"name" yml:"Database:Name"`
	Password string `json:"password" yml:"Database:Password"`
	Username string `json:"username" yml:"Database:Username"`
}

type Gorm struct {
	LogMode     int64 `json:"logMode"`
	MaxIdleConn int64 `json:"maxIdleConn"`
	MaxOpenConn int64 `json:"maxOpenConn"`
}
