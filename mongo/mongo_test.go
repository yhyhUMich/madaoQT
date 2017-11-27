package mongo

import (
	"testing"
	"time"
)

func _TestLoadingCharts(t *testing.T) {
	mongo := new(Charts);
	err := mongo.Connect();
	if err == nil {
		mongo.LoadCharts("Poloniex", "USDT-ETH", 15);
	}
}

func TestInsertTradeRecord(t *testing.T) {
	tradesDB := new(Trades);
	err := tradesDB.Connect();
	if err == nil {
		record := &TradesRecord{
			Time: time.Now(),
			Oper: "buy",
			Exchange: "okex",
			Coin: "btc",
			Quantity: 123.45,
		}
		tradesDB.Insert(record)
	}	
}