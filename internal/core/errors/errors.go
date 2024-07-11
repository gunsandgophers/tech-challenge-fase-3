package errors

import "errors"

var (
	ErrInvalidCategory = errors.New("invalid category. The category must be sandwich, sidedishes, drinks or desserts")
	ErrProductNotFound = errors.New("product not found")
	ValidCategories    = map[string]bool{
		"sandwich":   true,
		"sidedishes": true,
		"drinks":     true,
		"desserts":   true,
	}
)
