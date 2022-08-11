package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

type Data struct {
	Email string `json:"email"`
}

func (h *Handler) SignUp(d *json.RawMessage) {
	// надо понять, что придет в сообщении, и исходя из этого сделать структуру
	// десериализовать данные и отправлять их в сообщение
	var data Data
	err := json.Unmarshal(*d, &data)
	if err != nil {
		logrus.Errorf("error send sign up email: %s", err.Error())
	}

	err = h.mailer.Send(data.Email, "Sign Up to Plantpay", "sign-up.html", nil)
	if err != nil {
		logrus.Errorf("error send sign up email: %s", err.Error())
	}
	logrus.Infof("email to %s was sent", data.Email)
}
