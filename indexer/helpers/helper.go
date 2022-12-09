package helpers

func Filter[T any](items *[]T, cb func(*T) bool) []T {

	var filteredItems []T
	for i := 0; i < len(*items); i++ {
		if cb(&(*items)[i]) {
			filteredItems = append(filteredItems, (*items)[i])
		}
	}
	return filteredItems
}
