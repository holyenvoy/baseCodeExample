package util

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	if Sum(1, 2, 3) != 6 {
		t.Fatal("sum error")
	}
}

func TestAbs(t *testing.T) {
	if Abs(5) != 5 {
		t.Fatal("abs error, except:5, result:", Abs(5))
	} else {
		fmt.Printf("test Abs is ok value=%v\n", Abs(5))
	}
}
