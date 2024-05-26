package main

import "github.com/smartboot/smart-historian/historian"

func main() {
	h := historian.Historian{}
	h.Insert(historian.Request{
		Database: "test",
		Table:    "test",
		Fields: []historian.Field{
			{
				Timestamp: 123456789,
				Name:      "test",
				Value:     123,
			},
			{
				Timestamp: 123456789,
				Name:      "test",
				Value:     123,
			},
		},
	})
}
