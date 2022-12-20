package types

type Model interface {
	FindById(id any)
}

type ModelSearcher interface {
	Search(term any) (any, error)
}
