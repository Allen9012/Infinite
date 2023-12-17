package encrypt

import (
	"testing"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/10/26
  @desc:
  @modified by:
**/

// TestEncryptMobile
//
//	@Description: 字符串加密util
//	@param t
func TestEncryptMobile(t *testing.T) {
	mobile := "13800138000"
	encryptedMobile, err := EncMobile(mobile)
	if err != nil {
		t.Fatal(err)
	}
	decryptedMobile, err := DecMobile(encryptedMobile)
	if err != nil {
		t.Fatal(err)
	}
	if mobile != decryptedMobile {
		t.Fatalf("expected %s, but got %s", mobile, decryptedMobile)
	}
}
