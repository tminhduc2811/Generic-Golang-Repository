package repository

type GenericRepo interface {
	Fetch(query interface{}, model interface{}) error
	Create(model interface{}) error
	Update(selector interface{}, model interface{}) error
	Delete(query interface{}) error
	FetchOne(query interface{}, model interface{}) error
	Exists(query interface{}) (bool, error)
}
