package DataAccess

type Coin struct {
	//gorm.model
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`

	//PasswordHash string `json:"passwordhash"`
}

type Wallet struct {
	Id     int `json:"id"`
	UserId int `json:"userid"`
	CoinId int `json:"coinid"`
	//CoinAmount double `json:"coinamount"`
}
