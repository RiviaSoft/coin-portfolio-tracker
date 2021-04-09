package Business

import (
	. "github.com/msrexe/portfolio-tracker/Core/Result"
	. "github.com/msrexe/portfolio-tracker/DataAccess"
)

func GetAllBuyOperations(userId int) Result {
	return Result{}
}
func AddBuyOperation(userId int, operation RecentBuyOperation) Result {
	return Result{}

}
func DeleteBuyOperation(userId int, operation RecentBuyOperation) Result {
	return Result{}

}
func UpdateBuyOperation(userId float64, operation RecentBuyOperation) Result {
	return Result{Success: true, Message: "Güncelleme Başarılı !!", Data: nil}
}
