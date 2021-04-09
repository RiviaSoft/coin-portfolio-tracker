package DataAccess

// type Coin struct {
// 	//gorm.model
// 	Id   int    `json:"id"`
// 	Name string `json:"name"`
// }

// type Wallet struct {
// 	Id         int     `json:"id"`
// 	UserId     int     `json:"userid"`
// 	CoinSymbol string  `json:"coinid"`
// 	CoinAmount float64 `json:"coinamount"`
// }

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`

	//PasswordHash string `json:"passwordhash"`
}

type RecentBuyOperation struct {
	Id         int     `json:"id"`
	UserId     int     `json:"userid"`
	CoinSymbol string  `json:"coinsymbol"`
	CoinAmount float32 `json:"coinamount"`
	BuyCost    float32 `json:"buycost"`
}

//Inheritance yapılacak
type ArchivedOperation struct {
	Id         int     `json:"id"`
	UserId     int     `json:"userid"`
	CoinSymbol string  `json:"coinsymbol"`
	CoinAmount float32 `json:"coinamount"`
	SellCost   float32 `json:"sellcost"`
}
