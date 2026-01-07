package changes

import "gcm/internal/model"

func CategorizeByStatus(items []model.GitChange) map[string][]model.GitChange {
	res := make(map[string][]model.GitChange)
	for _, it := range items {
		key := it.StatusKey()
		res[key] = append(res[key], it)
	}
	return res
}
