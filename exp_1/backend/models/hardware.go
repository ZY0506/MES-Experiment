package models

import "time"

type Model struct {
	ID       uint64
	CreateAt time.Time `json:"-"`
	UpdateAt time.Time `json:"-"`
}

type Product struct {
	Model
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    int     `json:"quantity"`
}
