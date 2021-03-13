package main

import (
	"context"
	"fmt"
	"github.com/brandenc40/fmcsa-qc-mobile/qcmobile"
	"log"
	"net/http"
)

func main() {
	cfg := qcmobile.Config{
		Key:        "YOUR_KEY",
		HTTPClient: &http.Client{},
	}
	client := qcmobile.NewClient(cfg)
	ctx := context.Background()

	carrier, err := client.GetCarrier(ctx, 53467)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(carrier.Carrier.SafetyRatingDate.Parse())

	auth, err := client.GetAuthority(ctx, 53467)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(auth[0].CarrierAuthority)

	opClass, err := client.GetOperationClassification(ctx, 53467)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(opClass[0])

	basics, err := client.GetBasics(ctx, 53467)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(basics[0])

	oos, err := client.GetOOS(ctx, 885213)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(oos[0].Oos)

	cargo, err := client.GetCargoCarried(ctx, 885213)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cargo[0])

	carByDock, err := client.GetCarriersByDocket(ctx, 1515)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(carByDock[0].Carrier)

	dockets, err := client.GetDocketNumbers(ctx, 885213)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dockets[0])

	searchRes, err := client.SearchCarriersByName(ctx, "werner", 1, 5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(searchRes[0].Carrier)
}
