package repositories

import (
	"github.com/go-xorm/xorm"
	"go-trial/datamodels"
)

type TreeRepository struct {
	engine *xorm.Engine
}

func NewTree(engine *xorm.Engine) *TreeRepository {
	return &TreeRepository{
		engine: engine,
	}
}

func (my *TreeRepository) GetAll() []datamodels.Tree {

	var result []datamodels.Tree
	my.engine.Find(&result)
	return result
}
