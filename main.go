package main

import (
	"fmt"
	"log"
	"time"

	"github.com/rider-cantuaria-eng/price-alerter/utils"
)

type Thresholds struct {
	Min float64
	Max float64
}

func main() {

	BTCThreshold := Thresholds{
		Min: 64000,
		Max: 73000,
	}

	ATOMThreshold := Thresholds{
		Min: 7,
		Max: 8,
	}

	for {
		// for btc
		btc, err := utils.FetchCotationBTC()
		if err != nil {
			log.Fatalf("Error getting price of BTC: %v", err)
		}

		if btc.Buy < BTCThreshold.Min {
			utils.SendDesktopNotification("Good to buy", fmt.Sprintf("BTC price to buy: %.6f USD", btc.Buy))
		} else if btc.Sell > BTCThreshold.Max {
			utils.SendDesktopNotification("Good to sell", fmt.Sprintf("BTC price to sell: %.6f USD", btc.Sell))
		} else {
			log.Printf("BTC price is: %.6f USD", btc.Last)
		}

		// for atom
		atom, err := utils.FetchCotationATOM()
		if err != nil {
			log.Fatalf("Error getting price of ATOM: %v", err)
		}

		if atom.Last < ATOMThreshold.Min {
			utils.SendDesktopNotification("Good to buy", fmt.Sprintf("ATOM price to buy: %.6f USD", atom.Last))
		} else if atom.Last > ATOMThreshold.Max {
			utils.SendDesktopNotification("Good to sell", fmt.Sprintf("ATOM price to sell: %.6f USD", atom.Last))
		} else {
			log.Printf("ATOM price is: %.6f USD", atom.Last)
		}

		time.Sleep(10 * time.Minute)
	}
}
