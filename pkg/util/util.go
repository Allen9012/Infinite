package util

import (
	"math/rand"
	"strconv"
	"time"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/10/21
  @desc:
  @modified by:
**/

// RandomNumeric
//
//	@Description: generate one string of a num with len == size randomly
//	@param size
//	@return string
func RandomNumeric(size int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if size <= 0 {
		panic("{ size : " + strconv.Itoa(size) + " } must be more than 0 ")
	}
	value := ""
	for index := 0; index < size; index++ {
		value += strconv.Itoa(r.Intn(10))
	}

	return value
}

// EndOfDay
//
//	@Description: convert time type 2023-10-26 23:59:59 +0800 CST
//	@param t
//	@return time.Time
func EndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, t.Location())
}
