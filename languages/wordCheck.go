// Special thanks to this package for some inspiration: https://pkg.go.dev/gitlab.com/ffe4/exercism-go/scrabble-score#section-readme

package language
import (
  "fmt"
  "os"
  "io"
)

type alphabet struct{
  listOfWords = make(map[string]int)
  }

type word struct{
  spelling string
  score int
}

var activeAlphabet = new(alphabet)

func parseAlphabet string{
  jsonFile, err := os.Open("englishAlphabet.json")
  if err != nil {
    fmt.Println(err)
  } else {
    alphabet.listOfWords[]
  }

  defer jsonFile.Close()
}

var path_to_dict = string
