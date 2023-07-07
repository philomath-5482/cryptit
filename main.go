package main

import (
  "github.com/philomath-5482/cryptit/encrypt"
  "fmt"
  "github.com/philomath-5482/cryptit/decrypt"
)


func main() {
 encryptedStr := encrypt.Nimbus("kodekloud")
 fmt.Println(encryptedStr)
 fmt.Println(decrypt.Nimbus(encryptedStr))
}
