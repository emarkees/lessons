package main

import (
	"fmt"
	"testing"
)


func TestCountCharacter(t *testing.T) {
	tests := []struct {
		char string
		target rune
		want int
	}{
		{"hello", 'l', 2},
	}
	for _, tx := range tests {
		t.Run(fmt.Sprintf("(%v, %v)",tx.char, tx.target), func(t *testing.T){
			got := CountCharacter(tx.char, tx.target)
			if got != tx.want {
				t.Fatalf("CountCharacter() = %v; want %v", got, tx.want)
			}
		})
	}
}