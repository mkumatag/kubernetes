package image

import (
	"reflect"
)

type Catalog struct {
	Busybox string
	Ubuntu string
}

func New() *Catalog {
	return &Catalog{
		Busybox: getBusyBoxImage(),
		Ubuntu: getUbuntuImage(),
	}
}

func (c *Catalog) List() []string{
	v := reflect.ValueOf(c).Elem()
	values := make([]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface().(string)
	}

	return values
}