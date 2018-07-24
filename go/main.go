package main

import (
  "context"
  "fmt"
  "log"
  "math/big"
  "strings"

  "github.com/ethereum/go-ethereum"
  "github.com/ethereum/go-ethereum/accounts/abi"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/crypto"
  "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
  contractToWatch := "0x5EC1607870cbE46EFcF79b8DE4DBe20C1D1EbDA1"

  client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")
  if err != nil {
    log.Fatal(err)
  }

  contractAddress := common.HexToAddress(contractToWatch)
  query := ethereum.FilterQuery{
    FromBlock: big.NewInt(2691975), // 2691975
    ToBlock:   big.NewInt(2692808), // 2692808
    Addresses: []common.Address{
      contractAddress,
    },
  }

  logs, err := client.FilterLogs(context.Background(), query)
  if err != nil {
    log.Fatal(err)
  }

  contractAbi, err := abi.JSON(strings.NewReader(string(ABCtokenABI)))
  if err != nil {
    log.Fatal(err)
  }

  logsFound := 0

  for _, vLog := range logs {
    logsFound++

    // anonymous struct, can add more fields
    // if the event argument is `indexed`, no need to put here
    event := struct { Value *big.Int }{}

    err := contractAbi.Unpack(&event, "Transfer", vLog.Data)
    if err != nil {
      log.Fatal(err)
    }
    val := event.Value.String()

    var topics [3]string // 3 `indexed` fields = event topic, from, to
    for i := range vLog.Topics {
      if i == 0 {
        // @return 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef
        h := vLog.Topics[i] // Hash <-- the event topic
        topics[i] = h.String()
      } else {
        // @return 0x627306090abaB3A6e1400e9345bC60c78a8BEf57
        a := common.BytesToAddress(vLog.Topics[i].Bytes())
        topics[i] = a.String()
      }
    }

    fmt.Printf("\nEvent topic:\n↳ %s", topics[0])
    fmt.Printf("\nFrom:\n↳ %s", topics[1])
    fmt.Printf("\nTo:\n↳ %s", topics[2])
    fmt.Printf("\nValue:\n↳ %s", val)
    fmt.Println("")
  }

  fmt.Println("")
  printLineBreak()
  if logsFound == 0 {
    fmt.Println("No logs found")
  } else {
    fmt.Printf("%d logs found\n", logsFound)
  }
  printLineBreak()

  // Just for confirmation, let's compare the event topic.
  eventSignature := []byte("Transfer(address,address,uint256)")
  hash := crypto.Keccak256Hash(eventSignature)
  eventTopic := hash.Hex()
  fmt.Printf("\nExpected event topic:\n%s\n\n", eventTopic)
}

func printLineBreak() {
  fmt.Printf("%s\n", strings.Repeat("-", 42))
}
