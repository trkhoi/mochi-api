package activity

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetOne(id int) (*model.Activity, error)
	GetDefaultActivities() ([]model.Activity, error)
}
