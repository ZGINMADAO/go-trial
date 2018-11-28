package services

import (
	"go-trial/datamodels"
	"go-trial/repositories"
	"go-trial/datasource"
)

type AuthService interface {
	Get(id int) *datamodels.Admin
	//GetAll() []datamodels.Admin
	//Search(username string) []datamodels.Admin

	//Delete(id int) error
	//Update(user *datamodels.Admin, columns []string) error
	//Create(user *datamodels.Admin) error
}

type authService struct {
	repository *repositories.AdminRepository
}

func NewAuth() AuthService {
	return &authService{
		repository: repositories.New(datasource.Instance()),
	}
}

//func NewUserRepository(source map[int64]datamodels.User) UserRepository {
//	return &userMemoryRepository{source: source}
//}

func (my *authService) Get(id int) *datamodels.Admin {
	return my.repository.Get(id)
}

//func (s *AuthService)GetAll() []models.StarInfo {
//	return s.dao.GetAll()
//}
//
//func (s *AuthService)Search(country string) []models.StarInfo {
//	return s.dao.Search(country)
//}

//func (s *AuthService)Delete(id int) error {
//	return s.dao.Delete(id)
//}
//func (s *AuthService)Update(user *models.StarInfo, columns []string) error {
//	return s.dao.Update(user, columns)
//}
//func (s *AuthService)Create(user *models.StarInfo) error {
//	return s.dao.Create(user)
//}
