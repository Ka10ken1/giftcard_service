package model

type OrderStatus string

const (
	OrderPending   OrderStatus = "PENDING"
	OrderCompleted OrderStatus = "COMPLETED"
	OrderCancelled OrderStatus = "CANCELLED"
)

func (s OrderStatus) isValid() bool {
	switch s {
	 case OrderPending, OrderCancelled, OrderCompleted:
	    return true
	}
	return false
}
