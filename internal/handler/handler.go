package handler

import (
	"encoding/json"
	"github.com/evgeniy-krivenko/plantpay-email-sender/pkg/mailer"
	"github.com/evgeniy-krivenko/plantpay-email-sender/pkg/qpye"
)

func testHandler(d *json.RawMessage) {

}

type Handler struct {
	mailer *mailer.Mailer
}

func NewHandler(mailer *mailer.Mailer) *Handler {
	return &Handler{mailer: mailer}
}

func (h Handler) InitRoutes() *qpye.Router {
	routes := []qpye.Route{
		{Pattern: "test_order", Cb: h.SignUp},
	}
	r := qpye.NewRouter(&routes)
	return r
}
