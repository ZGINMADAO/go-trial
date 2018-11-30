package services

import (
	"go-trial/datamodels"
	"go-trial/repositories"
	"go-trial/datasource"
)

type AuthService interface {
	AdminGet(id int) *datamodels.Admin
	AdminGetAll() []datamodels.Admin
	AdminSearch(username, password string) bool

	//Delete(id int) error
	//Update(user *datamodels.Admin, columns []string) error
	//Create(user *datamodels.Admin) error

	TreeGetAll() []datamodels.Tree
}

type authService struct {
	adminRepository *repositories.AdminRepository
	treeRepository  *repositories.TreeRepository
}

func NewAuth() AuthService {
	return &authService{
		adminRepository: repositories.NewAdmin(datasource.Instance()),
		treeRepository:  repositories.NewTree(datasource.Instance()),
	}
}

//func NewUserRepository(source map[int64]datamodels.User) UserRepository {
//	return &userMemoryRepository{source: source}
//}

func (my *authService) AdminGet(id int) *datamodels.Admin {
	return my.adminRepository.Get(id)
}

func (my *authService) AdminGetAll() []datamodels.Admin {
	return my.adminRepository.GetAll()
}

func (my *authService) AdminSearch(username, password string) bool {
	return my.adminRepository.Search(username, password)
}

func (my *authService) TreeGetAll() ([]datamodels.Tree) {
	return my.treeRepository.GetAll()
}

//func (s *AuthService)Delete(id int) error {
//	return s.dao.Delete(id)
//}
//func (s *AuthService)Update(user *models.StarInfo, columns []string) error {
//	return s.dao.Update(user, columns)
//}
//func (s *AuthService)Create(user *models.StarInfo) error {
//	return s.dao.Create(user)
//}
