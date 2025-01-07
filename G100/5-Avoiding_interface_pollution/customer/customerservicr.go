package customer

type Customer struct {
	Id int
}

type customerStorer interface {
	StoreCustomer(Customer) error
}

type CustomerService struct {
	Storer customerStorer
}

func (cs CustomerService) CreateNewCustomer(id int) error {
	customer := Customer{Id: id}
	return cs.Storer.StoreCustomer(customer)
}

//type CustomerService struct {
//	store mysql.Stone // 依赖于具体实现
//}
//
//func (cs CustomerService) CreateNewCustomer(id string) error {
//	customer := Customer{id: id}
//	return cs.store.StroeCustomer(customer)
//}
