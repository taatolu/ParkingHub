package model

import (
    "time"
    "testing"
    "github.com/stretchr/testify/assert"
    )

func TestIsLicenseExpired(t *testing.T){
    tests := []struct{
        testname    string
        licenseExpiration   time.Time
        wantExpired   bool
    }{
        //testCase作成
        {
            testname:   "期限内",
            licenseExpiration:  time.Now().Add(24 * time.Hour), //1日後
            wantExpired:  false,
        },
        {
            testname:   "期限切れ",
            licenseExpiration:  time.Now().Add(-24 * time.Hour), //1日前
            wantExpired:  true,
        }
    }
    //testCaseをループで回す
    for _, tt := range tests{
        t.Run(tt.testname, func(t *testing.T){
            //CarOwner構造体のインスタンスを作成（LicenseExpirationの値だけ設定）
            c := &CarOwner{
                LicenseExpiration:  tt.licenseExpiration,
            }
            result := c.IsLicenseExpired()
            assert.Equal(t, tt.wantExpired, result, "テストケース: %s", tt.testname)

        })
    }
}