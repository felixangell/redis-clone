package main

import "github.com/bat-labs/krake/internal"

func main() {
	k := internal.NewKrakeServer()
	k.ListenAndServe("localhost:9093")
}
