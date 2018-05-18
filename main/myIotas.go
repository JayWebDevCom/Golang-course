package main

import "fmt"

// iotas increate in value by 1
const (
  aio = iota
  bio = iota
  cio = iota
)

const (
  _ = iota
  KB = 1 << (iota * 10) // 1 bit-left-shifted by (1 * 10)
  MB = 1 << (iota * 10) // 1 bit-left-shifted by (2 * 10) i.e. 10 again
  GB = 1 << (iota * 10) // 1 bit-left-shifted by (3 * 10) i.e. 10 again
)

func theIotas() {
  fmt.Println("%Binary\t%Decimal")
  fmt.Printf("%b\t", KB)
  fmt.Printf("%d\n", KB)
  fmt.Printf("%b\t", MB)
  fmt.Printf("%d\n", MB)
  fmt.Printf("%b\t", GB)
  fmt.Printf("%d\n", GB)
}
