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

type UserModel struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Id           int
	Name         string
	Email        string
	PasswordHash string
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
	BuyCost    float32 `json:"buycost"`
	SellCost   float32 `json:"sellcost"`
}

type Wallet struct {
	Id     int    `json:"id"`
	UserId int    `json:"userid"`
	Name   string `json:"name"`
}

type WalletOperation struct {
	Id          int `json:"id"`
	WalletId    int `json:"walletid"`
	OperationId int `json:"operationid"`
}

type WalletOperationDTO struct {
	Id          int     `json:"id"`
	WalletId    int     `json:"walletid"`
	OperationId int     `json:"operationid"`
	UserId      int     `json:"userid"`
	CoinSymbol  string  `json:"coinsymbol"`
	CoinAmount  float32 `json:"coinamount"`
	BuyCost     float32 `json:"buycost"`
}
