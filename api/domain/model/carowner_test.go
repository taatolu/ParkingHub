package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIsLicenseExpired(t *testing.T) {
	tests := []struct {
		testname          string
		licenseExpiration time.Time
		wantExpired       bool
	}{
		//testCase作成
		{
			testname:          "期限内",
			licenseExpiration: time.Now().Add(24 * time.Hour), //1日後
			wantExpired:       false,
		},
		{
			testname:          "期限切れ",
			licenseExpiration: time.Now().Add(-24 * time.Hour), //1日前
			wantExpired:       true,
		},
	}
	//testCaseをループで回す
	for _, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
			//CarOwner構造体のインスタンスを作成（LicenseExpirationの値だけ設定）
			c := &CarOwner{
				LicenseExpiration: tt.licenseExpiration,
			}
			result := c.IsLicenseExpired()
			assert.Equal(t, tt.wantExpired, result, "テストケース: %s", tt.testname)

		})
	}
}

func TestIsIDPositive(t *testing.T) {
    tests := []struct {
        testname    string
        id          int
        wantError   bool
    }{
        //testケースの作成
        {
            testname:   "正常系",
            id:         1,
            wantError:  false,
        },
        {
            testname:   "異常系（値が0）",
            id:         0,
            wantError:  true,
        },
        {
            testname:   "異常系（値が負の数）",
            id:         -1,
            wantError:  true,
        },
    }
    //テストケースrをループで回す
    for _, tt := range tests {
        t.Run(tt.testname, func(t *testing.T){
            if tt.wantError{
                //wantError=true(異常系の場合)
                result := IsIDPositive(tt.id)
                assert.Error(t, result, "エラーを期待していたのにエラーが返らない")
            } else {
                //wantError=false(正常系の場合)
                result := IsIDPositive(tt.id)
                assert.NoError(t, result, "エラになってしまった")
            }
        })
    }
}
