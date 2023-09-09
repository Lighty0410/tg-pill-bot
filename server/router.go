package server

import (
	"net/http"
	"tg-pill-bot/helper"
	"tg-pill-bot/model"
	"time"
)

func (H *HttpHandler[С]) GetPillTime(w http.ResponseWriter, r *http.Request) error {
	t, err := helper.StructToJson(model.PillTime{Time: time.Now()})
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(t)
	if err != nil {
		return err
	}

	return nil
}
func (H *HttpHandler[С]) RemindUser()        {}
func (H *HttpHandler[С]) UpdatePreferences() {}
func (H *HttpHandler[С]) UpdatePillTime()    {}
