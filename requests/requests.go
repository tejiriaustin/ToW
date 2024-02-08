package requests

import "time"

type CreateUserRequest struct {
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Phone         string    `json:"phone"`
	DOB           time.Time `json:"dob"`
	Country       string    `json:"country"`
	ZipCode       string    `json:"zip_code"`
	Email         string    `json:"email"`
	Profession    string    `json:"profession"`
	Income        string    `json:"income"`
	Company       string    `json:"company"`
	PersonalLinks []string  `json:"personal_links"`
}

type FreezeAccountRequest struct {
	AccountId string `json:"account_id"`
	Reason    string `json:"reason"`
}

type TradeWallyRequest struct {
	Amount           int64
	RecipientDetails string
}
