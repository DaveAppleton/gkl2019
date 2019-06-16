package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"./contracts"

	"github.com/DaveAppleton/etherUtils"
	"github.com/DaveAppleton/ether_go/ethKeys"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type goldData struct {
	Buy       float64
	Sell      float64
	SpotPrice float64 `json:"spot_price"`
}
type goldPrice struct {
	Result string
	Data   goldData
}

func main() {
	var gp goldPrice
	url := "https://cws.hellogold.com/api/v2/spot_price.json"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("1", err)
	}
	err = json.NewDecoder(resp.Body).Decode(&gp)
	if err != nil {
		log.Fatal("2", err)
	}
	s64 := fmt.Sprint(gp.Data.SpotPrice)
	fmt.Println("Gold Price : ", s64)
	val, ok := etherUtils.StrToEther(s64)
	if !ok {
		log.Fatal("3 Cannot get val from ", s64)
	}
	client, err := ethclient.Dial("https://kovan.infura.io/v3/2c954d83bea043f48d5ac6b70fa3b7b0")
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress("0xe07e7d16ef1d26b3dfe7493af6833bdf184860f3")
	light, err := contracts.NewLighthouse(contractAddress, client)
	if err != nil {
		log.Fatal("4", err)
	}
	payer := ethKeys.NewKey("mykey")
	err = payer.RestoreOrCreate()
	if err != nil {
		log.Fatal("5", err)
	}
	log.Println("handler ", payer.PublicKeyAsHexString())
	ownerTx := bind.NewKeyedTransactor(payer.GetKey())
	tx, err := light.Write(ownerTx, val)
	if err != nil {
		log.Fatal("6", err)
	}
	fmt.Println("https://kovan.etherscan.io/tx/" + tx.Hash().Hex())
}
