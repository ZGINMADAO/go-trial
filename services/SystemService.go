package services

type SystemService interface {
	Permissions()
}

type systemService struct {
}

func NewSystem() SystemService {
	return &systemService{}
}

func (my *systemService) Permissions() {

}
