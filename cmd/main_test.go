package main

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ubombar/go-rest-api/pkg/generated/contract"
)

func TestMain(t *testing.T) {
	// Use a local node:
	cl, err := ethclient.Dial("http://localhost:8545")

	if err != nil {
		t.Fail()
		return
	}

	contractAddress := common.HexToAddress("0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9")

	cc, err := contract.NewContract(contractAddress, cl)

	if err != nil {
		t.Fail()
		return
	}

	bal, err := cc.SeeBalance(nil)

	if err != nil {
		fmt.Printf("wowowowoowowo: %v\n", err)
		t.Fail()
		return
	}

	fmt.Printf("Balance: %v\n", bal)

}
