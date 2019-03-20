package main

import (
	"fmt"
  "os"
  "strings"
  "github.com/cosmos/cosmos-sdk/client"
  "github.com/cosmos/cosmos-sdk/crypto/keys"
  "./wordlist"
  bip39 "github.com/cosmos/go-bip39"
)

func main() {
  buf := client.BufferStdin()
  bip39Message := "Enter your bip39 mnemonic"
  addressMessage := "Enter your cosmos address from the fundraiser"
	mnemonic, _ := client.GetString(bip39Message, buf)
  address, _ := client.GetString(addressMessage, buf)
  words := strings.Fields(mnemonic)
  list := strings.Fields(wordlist.WL)
  for i,_ := range words {
    for j,_ := range list {
      exchangeWordAndCheck(words,i,list[j], address)
    }
  }
}
func exchangeWordAndCheck(words []string,i int,list string, address string) {
	tempWords := make([]string, len(words))
	copy(tempWords,words)
  tempWords[i] = list
  temp := strings.Join(tempWords, " ")
  found := checkMnemonic(temp, address)
  if found {
    os.Exit(1)
  }
}

func checkMnemonic(mnemonic string,address string) bool {
  var kb keys.Keybase
  kb = keys.NewInMemory()
  if !bip39.IsMnemonicValid(mnemonic) {
    fmt.Println("Invalid")
		return false
	} else {
    info, _ := kb.CreateAccount("Test", mnemonic, "", "", 0, 0)
    if info != nil {
      //fmt.Println("Valid")
      fmt.Println(mnemonic)
      out, err := keys.Bech32KeyOutput(info)
      if err != nil {
        return false
      } else {
        if address == out.Address {
          fmt.Println("Found it!",mnemonic)
          return true
        }
      }
    }
  }
  return false
}
