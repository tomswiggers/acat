package main

import (
    "bufio"
    "fmt"
    "flag"
    "os"
    "strconv"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  c := 1

  filename := flag.String("filename", "test.md", "file to cat")
  start := flag.String("start", "0", "start from line")
  end := flag.String("end", "10", "end line")
  flag.Parse()

  startInt, err := strconv.Atoi(*start)
  endInt, err := strconv.Atoi(*end)

  fp, err := os.Open(*filename)
  check(err)

  scanner := bufio.NewScanner(fp)

  for scanner.Scan() {
    buffer := scanner.Text()

    if c >= startInt && c <= endInt {
      fmt.Println(buffer)
    }

    c++
  }
}
