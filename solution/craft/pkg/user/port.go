package user

type Lister interface {
	List() ([]User, error)
}
