package booking

import (
	"encoding/json"
	"net/http"
	"os"

	"booking-service/commons"

	"github.com/gocarina/gocsv"
)

type BookingDetail struct {
	ID            string         `json:"id"`
	CustomerName  string         `json:"customer_name"`
	CustomerEmail string         `json:"customer_email"`
	PaymentType   int            `json:"payment_type"`
	RoomTypeID    int            `json:"room_type_id"`
	Price         string         `json:"price"`
	CheckinDate   string         `json:"checkin_date"`
	CheckoutDate  string         `json:"checkout_date"`
	BookingStatus string         `json:"booking_status"`
	Adult         int            `json:"adult"`
	Children      int            `json:"children"`
	IsSmoke       bool           `json:"is_smoke"`
	IsPet         bool           `json:"is_pet"`
	RoomDetail    []commons.Room `json:"room_detail"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	bookingLogFile, err := os.OpenFile("../../data/booking_log.csv", os.O_APPEND|os.O_WRONLY, os.ModePerm)

	if err != nil {
		bookingLogFile, err = os.OpenFile("../../data/booking_log.csv", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			panic(err)
		}
		initFile := []*commons.Room{}
		gocsv.MarshalFile(&initFile, bookingLogFile)
	}
	defer bookingLogFile.Close()

	res := &commons.Response{}

	decoder := json.NewDecoder(r.Body)
	bookingDetail := BookingDetail{}

	err = decoder.Decode(bookingDetail)
	if err != nil {
		res.Fail(http.StatusInternalServerError, 500, err.Error())
	}

	err = gocsv.MarshalWithoutHeaders(&bookingDetail.RoomDetail, bookingLogFile)
	if err != nil {
		panic(err)
	}
	bookingDetail.BookingStatus = "Success"

	w.WriteHeader(http.StatusOK)
	res.Success("OK", bookingDetail)
	json.NewEncoder(w).Encode(res)
}
