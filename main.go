package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/bits"

	"github.com/valyala/fasthttp"
)

type transaction struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Value    int    `json:"value"`
	Time     int    `json:"time"`
	Category string `json:"category"` // optional
}

type addRequest transaction

type transactionsRequest struct {
	// all fields are optional
	sender         string `json:"sender"`
	receiver       string `json:"receiver"`
	timeRangeStart int    `json:"time_range_start"`
	timeRangeEnd   int    `json:"time_range_end"`
	category       string `json:"category"`
}

// UserID is a user id provided by client
type UserID int

type dbMap map[UserID][]transaction

func addHandlerFunc(ctx *fasthttp.RequestCtx, db *dbMap) {
	body := ctx.PostBody()
	addReq := addRequest{}

	err := json.Unmarshal(body, &addReq)
	if err != nil {
		log.Println("Error occured while json parsing: ", err.Error())
		return
	}

	trans := transaction(addReq)
	(*db)[0] = append((*db)[0], trans)

	log.Println("transaction: ", trans, " added to db")

	ctx.SetBody([]byte{})
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func transactionsHandlerFunc(ctx *fasthttp.RequestCtx, db *dbMap) {
	queryArgs := ctx.URI().QueryArgs()

	transReq := transactionsRequest{}

	transReq.category = string(queryArgs.Peek("category"))
	transReq.receiver = string(queryArgs.Peek("receiver"))
	transReq.sender = string(queryArgs.Peek("sender"))

	parsedTimeStart, err := queryArgs.GetUint("time_range_start")

	if err != nil {
		transReq.timeRangeStart = 0
	} else {
		transReq.timeRangeStart = parsedTimeStart
	}

	parsedTimeEnd, err := queryArgs.GetUint("time_range_end")

	if err != nil {
		transReq.timeRangeEnd = 1<<(bits.UintSize-1) - 1
	} else {
		transReq.timeRangeEnd = parsedTimeEnd
	}

	log.Println("Request received: ", transReq)
	transactions := make([]transaction, 0)

	for _, trans := range (*db)[0] {
		if trans.Category != transReq.category && transReq.category != "" {
			log.Println("category check failed")
			continue
		}
		if trans.Receiver != transReq.receiver && transReq.receiver != "" {
			log.Println("receiver check failed")
			continue
		}
		if trans.Sender != transReq.sender && transReq.sender != "" {
			log.Println("sender check failed")
			continue
		}
		if trans.Time < transReq.timeRangeStart || trans.Time >= transReq.timeRangeEnd {
			log.Println("time check failed")
			continue
		}
		transactions = append(transactions, trans)
	}

	log.Printf("Send next transactions: ", transactions)
	jsonBytes, err := json.Marshal(transactions)

	log.Printf("JsonBytes sent: ", jsonBytes)
	if err != nil {
		log.Println("can't marshal transactions into json, fix structure: ", err.Error())
		ctx.SetBody([]byte{})
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetBody(jsonBytes)
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
}

func main() {
	db := make(dbMap)

	db[0] = append(db[0], transaction{"Bank of America", "VTB", 3, 10000, "Unknown"})

	serve := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/add":
			addHandlerFunc(ctx, &db)
		case "/transactions":
			transactionsHandlerFunc(ctx, &db)
		default:
			log.Printf("Uknown path %v \n", ctx.Path())
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}

	port := "8080"
	log.Println("Started at: localhost:", port)

	err := fasthttp.ListenAndServe(fmt.Sprintf(":%s", port), serve)

	if err != nil {
		log.Println("Something went wrong: ", err.Error())
	}
}
