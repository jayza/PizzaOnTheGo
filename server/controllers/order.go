package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jayza/pizzaonthego/helpers"
	"github.com/jayza/pizzaonthego/models"
	repository "github.com/jayza/pizzaonthego/repositories"
)

// GetOneOrderHandler ...
// swagger:route GET /api/v1/orders/:id Orders findOrder
//
// Finds order by order id parameter.
//
// This will return the entire order, complete with
// line items and shipping information.
//
//
// Responses:
//   default: JSONResponse
//   200: JSONResponse
func GetOneOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderID, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		helpers.RespondWithError(w, r, 400, err)
		return
	}

	order, err := repository.OneOrder(orderID)
	helpers.RespondWithJSON(w, r, order, err)
}

// CreateOrderHandler ...
// swagger:route POST /api/v1/orders Orders addOrder
//
// Creates a new order
//
// This will return the entire order, complete with
// line items and shipping information after creation.
//
//
//
//     Responses:
//       default: JSONResponse
func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		helpers.RespondWithError(w, r, 400, err)
		return
	}

	o := models.Order{}
	json.Unmarshal(body, &o)

	order, err := repository.CreateOrder(o)

	helpers.RespondWithJSON(w, r, order, err)
}

// DownloadOrderReceipt ...
// swagger:route GET /api/v1/orders/{id}/receipt Orders findOrder
//
// Downloads the order receipt PDF
//
// This will return the entire order, complete with
// line items and shipping information in the form of a PDF.
//
// Produces:
// application/pdf
//
//
//     Responses:
//       default: JSONResponse
func DownloadOrderReceipt(w http.ResponseWriter, r *http.Request) {
	orderID := mux.Vars(r)["id"]
	filename := "Receipt-Ordernumber-" + orderID + ".pdf"
	url := "http://localhost:8080/public/receipts/" + filename

	timeout := time.Duration(5) * time.Second
	transport := &http.Transport{
		ResponseHeaderTimeout: timeout,
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, timeout)
		},
		DisableKeepAlives: true,
	}
	client := &http.Client{
		Transport: transport,
	}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", fmt.Sprintf("%d", resp.ContentLength))

	io.Copy(w, resp.Body)
}
