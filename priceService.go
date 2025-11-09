package services

import (
	"math/rand"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	mu         sync.RWMutex
	priceCache = map[string]float64{}
	stocks     = []string{"RELIANCE", "TCS", "INFY"}
)

func StartPriceUpdater() {
	go func() {
		for {
			mu.Lock()
			for _, s := range stocks {
				priceCache[s] = 2000 + rand.Float64()*1000 // random mock price
			}
			mu.Unlock()
			logrus.Info("💰 Updated stock prices")
			time.Sleep(1 * time.Hour)
		}
	}()
}

func GetPrice(stock string) float64 {
	mu.RLock()
	defer mu.RUnlock()
	if price, ok := priceCache[stock]; ok {
		return price
	}
	return 0.0
}
