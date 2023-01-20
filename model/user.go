package model

import "oos/dto"

type User struct {
	dto.UserCreate
}

type Address struct {
	dto.AddressCreate
}

// References
// https://github.com/nyaruka/phonenumbers
// https://phonenumbers.temba.io/
// https://github.com/Boostport/address
// https://chromium-i18n.appspot.com/ssl-address