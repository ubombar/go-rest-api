package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

	num, err := cl.BlockNumber(context.TODO())

	if err != nil {
		t.Fail()
		return
	}

	fmt.Printf("num: %v\n", num)

	address := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	balance, err := cl.BalanceAt(context.TODO(), address, nil)

	if err != nil {
		t.Fail()
		return
	}

	fmt.Printf("balance.String(): %v\n", balance.String())

	contractAddress := common.HexToAddress("0x8464135c8f25da09e49bc8782676a84730c318bc")

	cc, err := contract.NewContract(contractAddress, cl)

	if err != nil {
		t.Fail()
		return
	}

	callOpts := &bind.CallOpts{Context: context.TODO(), Pending: false}

	adminAddress, err := cc.GetAdmin(callOpts)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		t.Fail()
		return
	}

	fmt.Printf("adminAddress: %v\n", adminAddress)

}
