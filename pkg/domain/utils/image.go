package utils

import (
  "bufio"
  "github.com/sirupsen/logrus"
  "os"
)

func ReadBytesFromImage(path string) []byte  {

  file, err := os.Open(path)

  if err != nil {
    logrus.Error(err)
  }

  defer file.Close()

  info, _ := file.Stat()
  bytes := make([]byte, info.Size())

  buffer := bufio.NewReader(file)
  _, err = buffer.Read(bytes)

  return bytes
}
