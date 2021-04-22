package model

type User struct {
	ID                     string
	CountryCode            string
	CurrencyCode           string
	Phone                  string
	FirstName              string `json:"firstName"`
	LastName               string `json:"lastName"`
	Email                  string `json:"email"`
	Role                   string `json:"role"`
	IsPhoneNumberConfirmed bool
	IsEmailConfirmed       bool
}

const (
	defaultCurrencyCode = "EUR"
	defaultCountryCode  = "DEFAULT"
)

func (u *User) RecognizeCountry() {
	if len(u.Phone) < 4 {
		u.CountryCode = defaultCountryCode
		u.CurrencyCode = defaultCurrencyCode
		return
	}
	switch u.Phone[0:4] {
	case "+234":
		u.CountryCode = "NGA"
		//u.CurrencyCode = "NGN"
	case "+254":
		u.CountryCode = "KEN"
		//u.CurrencyCode = "KES"
	case "+233":
		u.CountryCode = "GHA"
		//u.CurrencyCode = "GHS"
	default:
		u.CountryCode = defaultCountryCode
	}
	u.CurrencyCode = defaultCurrencyCode
}
