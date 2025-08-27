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
        wantBool    bool
    }{
        //testケースの作成
        {
            testname:   "正常系",
            id:         1,
            wantBool:   true,
        },
        {
            testname:   "異常系（値が0）",
            id:         0,
            wantBool:   false,
        },
        {
            testname:   "異常系（値が負の数）",
            id:         -1,
            wantBool:   false,
        },
    }
    //テストケースをループで回す
    for _, tt := range tests {
        t.Run(tt.testname, func(t *testing.T){
            c := &CarOwner{
                ID: tt.id,
            }
            result := c.IsIDPositive()
            assert.Equal(t, tt.wantBool, result)
        })
    }
}

func TestIsContainsName(t *testing.T) {
    //Create tableTest
    tests := []struct {
        testname    string
        owner       *CarOwner
        wantString  string
        wantContains   bool
    }{
        //testCaseの作成
        {
            testname:   "検索ヒット",
            owner:      &CarOwner{ID:1, FirstName:"yanada", MiddleName:"taro", LastName:"一世"},
            wantString: "taro",
            wantContains:  true,
        },
        {
            testname:   "検索ヒットしない：大文字",
            owner:      &CarOwner{ID:1, FirstName:"yanada", MiddleName:"taro", LastName:"一世"},
            wantString: "Taro",
            wantContains:  false,
        },
        {
            testname:   "検索ヒットしない",
            owner:      &CarOwner{ID:1, FirstName:"yanada", MiddleName:"taro", LastName:"一世"},
            wantString: "Naro",
            wantContains:  false,
        },
    }
    //testCaseをループ処理
    for _, tt := range tests {
        t.Run(tt.testname, func(t *testing.T){
            IsContains := tt.owner.IsContainsName(tt.wantString)
            assert.Equal(t, tt.wantContains, IsContains) 
        })
    }
}