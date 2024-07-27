package mappers

func Map[E any, D any](dbItems []E, mapper func(E) (D, error)) ([]D, error) {
	domainItems := make([]D, len(dbItems))

	for idx, dbItem := range dbItems {
		item, err := mapper(dbItem)
		if err != nil {
			return []D{}, err
		}
		domainItems[idx] = item
	}

	return domainItems, nil
}
