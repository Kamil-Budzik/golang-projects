package models

type Character struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Role   string `json:"role"`
	Bounty int    `json:"bounty"`
}
