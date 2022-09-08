package events

import (
	"github.com/SerkanKutlu/orderService/model"
	"time"
)

type OrderUpdated struct {
	Id         string        `bson:"_id" json:"_id" validate:"isdefault"`              //calculate auto
	CustomerId string        `bson:"customerId" json:"customerId" validate:"required"` //from user
	Quantity   int           `bson:"quantity" json:"quantity" validate:"isdefault"`    //calculate auto
	Total      float32       `bson:"total" json:"total" validate:"isdefault"`          //calculate auto
	Status     string        `bson:"status" json:"status" validate:"required"`         //from user
	Address    model.Address `bson:"address" json:"address" validate:"isdefault"`      //get from customer
	ProductIds []string      `bson:"productIds" json:"productIds" validate:"required"` //from user
	UpdatedAt  time.Time     `bson:"updatedAt" json:"updatedAt" validate:"isdefault"`  //calculate auto
}

func (oc *OrderUpdated) FillUpdated(order *model.Order) {

	oc.Id = order.Id
	oc.CustomerId = order.CustomerId
	oc.Quantity = order.Quantity
	oc.Total = order.Total
	oc.Status = order.Status
	oc.Address = order.Address
	oc.ProductIds = order.ProductIds
	oc.UpdatedAt = order.UpdatedAt
}
