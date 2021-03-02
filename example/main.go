package main

import (
	"context"
	"fmt"
	"github.com/brandenc40/fmcsa-qc-mobile/qcmobile"
)

func main() {
	cfg := qcmobile.Config{Key: "your-key-goes-here"}
	client := qcmobile.NewClient(cfg)
	ctx := context.Background()

	carrier, err := client.GetCarrier(ctx, 53467)
	if err != nil {
		panic(err)
	}
	fmt.Println(carrier.Content.Carrier.SafetyRatingDate.Parse())

	auth, err := client.GetAuthority(ctx, 53467)
	if err != nil {
		panic(err)
	}
	fmt.Println(auth.Content[0].CarrierAuthority)

	opClass, err := client.GetOperationClassification(ctx, 53467)
	if err != nil {
		panic(err)
	}
	fmt.Println(opClass.Content[0])

	basics, err := client.GetBasics(ctx, 53467)
	if err != nil {
		panic(err)
	}
	fmt.Println(basics.Content[0])

	oos, err := client.GetOOS(ctx, 885213)
	if err != nil {
		panic(err)
	}
	fmt.Println(oos.Content[0].Oos)

	cargo, err := client.GetCargoCarried(ctx, 885213)
	if err != nil {
		panic(err)
	}
	fmt.Println(cargo.Content[0])

	carByDock, err := client.GetCarriersByDocket(ctx, 1515)
	if err != nil {
		panic(err)
	}
	fmt.Println(carByDock.Content[0].Carrier)

	dockets, err := client.GetDocketNumbers(ctx, 885213)
	if err != nil {
		panic(err)
	}
	fmt.Println(dockets.Content[0])

	searchRes, err := client.SearchCarriersByName(ctx, "werner", 1, 5)
	if err != nil {
		panic(err)
	}
	fmt.Println(searchRes.Content[0].Carrier)
}
