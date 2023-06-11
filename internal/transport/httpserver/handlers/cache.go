package handlers

import (
	"strings"
	"test/internal/models"
)

func (h *Handler) CacheWarming(data []models.Data) {
	for _, link := range data {
		index := strings.LastIndex(link.Active_link, "/")
		substring := link.Active_link[index+1:]
		if Contains(substring) {
			h.service.Cache.Add(link.Active_link, link.History_link)
		}
	}
}

func Contains(word string) bool {
	smartfon := []string{"smartfon", "apple", "iphone", "samsung", "xiaomi", "huawai", "oppo"}
	flag := false

	for _, str := range smartfon {
		if strings.Contains(str, word) {
			flag = true
		}
	}

	return flag
}
