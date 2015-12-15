package main 

import (
    "encoding/json"
    "fmt"

    "github.com/ianberinger/stockfighter/api"
)

const (
    apiKey string = ""
    account = "FTB60965847"
    venue = "VFEX"
    symbol = "EQGU"
)

func main () {
    sfInstance := api.NewInstance(apiKey, account, venue, symbol)
    prettyPrint("current orderbook:", sfInstance.Orderbook())

    stream := i.Quotes(false)
    for tick := range stream.Values {
        prettyPrint("tick:", tick)
        currentOrderBook := sfInstance.Orderbook()

        for i := range len(currentOrderBook["asks"]) {
            price := currentOrderBook["asks"][i]["price"]
            qty := currentOrderBook["asks"][i]["qty"]
            if price <= 4991 {
                order := i.NewOrder(price, qty, api.Buy, api.Limit)

                prettyPrint("created order:", order)
            }
        }       
    }
}

func prettyPrint(description, v interface{}) {
    x, _ := json.MarshalIndent(v, "", "    ")
    fmt.Printf("%s %+v\n", description, string(x))
}
