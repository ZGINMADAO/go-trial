package repositories

import (
	"github.com/go-xorm/xorm"
	"go-trial/datamodels"
	"log"
)

type AdminRepository struct {
	engine *xorm.Engine
}

func NewAdmin(engine *xorm.Engine) *AdminRepository {
	return &AdminRepository{
		engine: engine,
	}
}

func (my *AdminRepository) Get(id int) *datamodels.Admin {
	data := &datamodels.Admin{Id: id}
	_, err := my.engine.Get(data)
	if err != nil {
		log.Fatal("Get Error")
	}
	return data
}

func (my *AdminRepository) GetAll() []datamodels.Admin {
	dataList := make([]datamodels.Admin, 0)
	err := my.engine.Desc("id").Find(&dataList)
	if err != nil {
		log.Fatal("Admin GetAll Error")
	}
	return dataList
}

func (my *AdminRepository) Search(username, password string) bool {

	var admin datamodels.Admin
	admin.Username = username
	admin.Password = password
	ok, _ := my.engine.Get(&admin)
	return ok
}
