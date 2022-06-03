package testuser

type TestUser struct {
	ID               int    `json:"id"`
	Nickname         string `json:"nickname"`
	Password         string `json:"password"`
	LoginAccessToken string `json:"login_access_token"`
	Status           string `json:"status"`
	SiteID           string `json:"site_id"`
	Type             string `json:"type"`
	Owner            string `json:"owner"`
	SellerExperience string `json:"seller_experience"`
	Credit           struct {
		Level    string `json:"level"`
		Consumed int    `json:"consumed"`
		Rank     string `json:"rank"`
	} `json:"credit"`
	Email                 string      `json:"email"`
	MercadoEnvios         string      `json:"mercado_envios"`
	BuyEqualsPayAsSeller  bool        `json:"buy_equals_pay_as_seller"`
	PurchaseID            interface{} `json:"purchase_id"`
	OneTimePassword       string      `json:"one_time_password"`
	AvailableAccountMoney float64     `json:"available_account_money"`
	IsVendor              bool        `json:"is_vendor"`
}

type TestUserError struct {
	Message  string        `json:"message"`
	Resource string        `json:"resource"`
	Error    string        `json:"error"`
	Status   int           `json:"status"`
	Cause    []interface{} `json:"cause"`
}

type TestUsersList struct {
	ID, Site, User, Pass, Estado, FechaEstado, Esquema string
}