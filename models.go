package main

import (
	"net/url"
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID        uuid.UUID
	Username    string
	Name        string
	DateOfBirth time.Time
	Password    string
	Avatar      []byte
}

type Product struct {
	UUID        uuid.UUID
	Name        string
	Description string
	URL         url.URL
	Image       []byte
}

// Product {
// 	ID = 1
// 	Name = "Kissie Hertog Jan"
// 	Description = ""
// 	URL = "ah.nl/fjkdfjkg"
// 	Image = "blob:data/base64;image/jpg"
// }

type Wager struct {
	UUID       uuid.UUID
	Challenger User
	Opponent   User
	Stake      Product
}
