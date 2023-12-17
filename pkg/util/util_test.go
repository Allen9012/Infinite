package util

import (
	"testing"
	"time"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/10/26
  @desc:
  @modified by:
**/

// TestEndOfDay
//
//	@Description:
//	@param t
func TestEndOfDay(t *testing.T) {
	now := EndOfDay(time.Now())
	t.Log("func EndOfDay:", now)
}

func TestRandomNumeric(t *testing.T) {
	var numeric string
	for i := 0; i < 10; i++ {
		numeric = RandomNumeric(10)
		t.Log("func RandomNumeric:", numeric)
	}
}
