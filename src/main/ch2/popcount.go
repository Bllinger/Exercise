package main

import (
	"fmt"
	"os"
)

var pc [256]byte

func init() {
	for i, _ := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

var cwd string

func init() {
	cwd, err := os.Getwd()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v:%v", cwd, err)
	}
}

func PopCount(x uint64) int {
	return int(pc[(x>>0)&0xff] +
		pc[(x>>1*8)&0xff] +
		pc[(x>>2*8)&0xff] +
		pc[(x>>3*8)&0xff] +
		pc[(x>>4*8)&0xff] +
		pc[(x>>5*8)&0xff] +
		pc[(x>>6*8)&0xff] +
		pc[(x>>7*8)&0xff])
}
