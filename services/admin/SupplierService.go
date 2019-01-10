package admin

type SupplierService interface{}

type supplierService struct {
}

func NewSupplier() SupplierService {
	return &supplierService{}
}
