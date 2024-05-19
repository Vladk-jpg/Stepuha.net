package service

func (serv *GoodService) CheckBelonging(userId int, goodId int) error {
	return serv.repo.CheckBelonging(userId, goodId)
}
