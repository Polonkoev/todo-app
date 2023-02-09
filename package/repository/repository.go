package sevice

type Authorization interface{

}
type TodoList interface{

}
type TodoItem interface{

}
type Repository struct{
	Authorization
	TodoList
TodoItem
}

func NewRepository(repos *repository.Repository) *Service{
	return &Repository{}
}