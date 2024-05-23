package service

func (serv *GoodService) CheckBelonging(userId int, goodId int) error {
	return serv.repo.CheckBelonging(userId, goodId)
}

func (serv *SupplementaryService) CheckIfModer(userId int) (bool, error) {
	return serv.repo.CheckIfModer(userId)
}
