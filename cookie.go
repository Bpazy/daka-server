package main

type dakaCookie struct {
	Cookie string `json:"cookie"`
}

func (d dakaCookie) getUserId() (userId string) {
	return d.Cookie
}
