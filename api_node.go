package shared

import (
	"time"
)

// APINodeList represents list of nodes
type APINodeList []ApiNode

// APINode contains single node
type APINode struct {
	Name         string    `json:"name"`
	Slug         string    `json:"slug"`
	Calories     float64   `json:"calories"`
	Fat          float64   `json:"fat"`
	Carbohydrate float64   `json:"carbohydrate"`
	Protein      float64   `json:"protein"`
	Barcode      string    `json:"barcode,omitempty"`
	UserID       string    `json:"-"`
	Created      time.Time `json:"created"`
}