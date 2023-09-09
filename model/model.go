package model

import "time"

type Response struct {
	Response string `json:"response"`
}

type PillTime struct {
	Time time.Time `json:"time"`
}
