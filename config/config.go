package config

import (
  "os"
  "io/ioutil"
  "encoding/json"
)

type SectionsT map[string]MessagesT
type MessagesT map[string]string

const (
  categoriesLocation = "config/messages.json"
)

var (
  Messages SectionsT
)

func Load() (err error) {
  file, err := os.Open(categoriesLocation)
  if err != nil {
    return
  }

  bytes, err := ioutil.ReadAll(file)
  if err != nil {
    return
  }

  defer file.Close()

  err = json.Unmarshal(bytes, &Messages)
  return
}
