package repositories

import (
	"github.com/go-xorm/xorm"
	"go-trial/datamodels"
)

type AdminRepository struct {
	engine *xorm.Engine
}

func New(engine *xorm.Engine) *AdminRepository {
	return &AdminRepository{
		engine: engine,
	}
}

func (my *AdminRepository) Get(id int) *datamodels.Admin {
	data := &datamodels.Admin{Id: id}
	ok, err := my.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}
