package main

import (
  "bufio"
  "fmt"
  "os"
  "github.com/tchap/go-patricia/patricia"
)

func readLines(file string) ([]string, error) {
  f, err := os.Open(file)
  if err != nil {
    return nil, err
  }
  defer f.Close()

  scanner := bufio.NewScanner(f)
  words := []string{}
  for scanner.Scan() {
    words = append(words, scanner.Text())
  }
  if err = scanner.Err(); err != nil {
    return nil, err
  }
  return words, nil
}

func compound(s string, trie *patricia.Trie) bool {
  for i := 1; i < len(s)-1; i++ {
    if trie.Get([]byte(s[:i])) != nil {
      if trie.Get([]byte(s[i:])) != nil || compound(s[i:], trie){
        return true
      }
    }
  }
  return false
}

func main() {
  var fname = "word.list"
  if len(os.Args) > 1 {
    fname = os.Args[1]
  }
  words, err := readLines(fname)
  if err != nil {
    fmt.Println(err)
    return
  }
  trie := patricia.NewTrie()
  for k := 0; k < len(words); k++ {
    trie.Insert([]byte(words[k]), 1)
  }

  longest := ""
  for k := len(words)-1; k >=0; k-- {
    if len(words[k]) > len(longest) && compound(words[k], trie) {
      longest = words[k]
    }
  }

  fmt.Printf("Longest compound string is %q\n", longest)
}