package types

type GetJsonPlaceholder struct {
	Id       int                       `json:"id"`
	Name     string                    `json:"name"`
	Username string                    `json:"username"`
	Address  GetJsonPlaceholderAddress `json:"address"`
	Phone    string                    `json:"phone"`
	Website  string                    `json:"website"`
	Company  GetJsonPlaceholderCompany `json:"company"`
}

type GetJsonPlaceholderAddress struct {
	Street  string            `json:"street"`
	Suite   string            `json:"suite"`
	City    string            `json:"city"`
	Zipcode string            `json:"zipcode"`
	Geo     map[string]string `json:"geo"`
}

type GetJsonPlaceholderCompany struct {
	Name             string `json:"name"`
	CatchPhrase      string `json:"catchPhrase"`
	BusinessStrategy string `json:"bs"`
}
