package request

import "github.com/google/uuid"

type Basket struct {
	Product_id uuid.UUID `json:"product_id"`
}
