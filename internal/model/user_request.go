package model

import (
	"time"
)

var RequestStatus = struct {
	NotAvailable string
	NotFilled    string
	Canceled     string
	Pending      string
	Approved     string
}{
	"not-available",
	"not-filled",
	"canceled",
	"pending",
	"approved",
}

type UserRequest struct {
	ID        uint64 `gorm:"primary_key" json:"id"`
	TierId    uint64
	Tier      Tier      `json:"tier" gorm:"foreignkey:TierId"`
	UserId    string    `json:"userId"`
	Status    string    `json:"status"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
