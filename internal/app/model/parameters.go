package model

import "time"

var (
	// StatusSuccess to store a status response success
	StatusSuccess = "Success"

	// MessageDataSuccess to store a success message response of data
	MessageDataSuccess = "Data Berhasil Dikirim"

	// MessageUnprocessableEntity to store a message response of unproccessable entity
	MessageUnprocessableEntity = "Entitas Tidak Dapat Diproses"
)

// NowUTC to get real current datetime but UTC format
func NowUTC() time.Time {
	return time.Now().UTC().Add(7 * time.Hour)
}
