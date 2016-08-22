j2s [![Build Status](https://drone.io/github.com/btfak/j2s/status.png)](https://drone.io/github.com/btfak/j2s/latest)
======

>j2s is a cmd tool generate golang struct from JSON

## example

```
cat example.json | ./j2s > example_output.go

package main

type Financial struct {
	Basic  string  `json:"basic"`
	Items  []Items `json:"items"`
	Normal float64 `json:"normal"`
	Unit   Unit    `json:"unit"`
}

type UnitTwo struct {
	Price         string `json:"price"`
	PriceQuantity string `json:"price_quantity"`
	Quantity      string `json:"quantity"`
}

type Items struct {
	EntrustRate   float64   `json:"entrust_rate"`
	GoodsTypeName string    `json:"goods_type_name"`
	ItemID        string    `json:"item_id"`
	ItemType      string    `json:"item_type"`
	ItemTypeID    float64   `json:"item_type_id"`
	ManifestName  string    `json:"manifest_name"`
	PaymentRate   float64   `json:"payment_rate"`
	Price         float64   `json:"price"`
	Quantity      float64   `json:"quantity"`
	SerialNumber  string    `json:"serial_number"`
	UnitTwo       []UnitTwo `json:"unit_two"`
}
type Unit struct {
	Price         string `json:"price"`
	PriceQuantity string `json:"price_quantity"`
	Quantity      string `json:"quantity"`
}
```

Installation
------------

```
$ go get github.com/lubia/j2s
```

Related Work
------------

-  github.com/tmc/json-to-struct
