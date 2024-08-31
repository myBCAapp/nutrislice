package schema

import "time"

type MenuDay struct {
	Date                time.Time `json:"date"`
	HasUnpublishedMenus bool `json:"has_unpublished_menus"`
}