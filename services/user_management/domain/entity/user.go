package entity

import "go-micro-service/pkg/base_entity"

type User struct {
	base_entity.Model
	Email    string
	Name     string
	Password string
}
