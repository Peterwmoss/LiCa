package mappers

import "fmt"

func Map[E any, D any](dbItems []E, mapper func(E) (D, error)) ([]D, error) {
	domainItems := make([]D, len(dbItems))

	for idx, dbItem := range dbItems {
		item, err := mapper(dbItem)
		if err != nil {
			return []D{}, fmt.Errorf("mappers.Map: Failed to map item: %v:. Error: %w", dbItem, err)
		}
		domainItems[idx] = item
	}

	return domainItems, nil
}
