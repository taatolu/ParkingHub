package model

import(
    "github.com/stretchr/testify/assert"
	"testing"
	"time"
    )

//車検証の期限切確認
func TestIsShakenExpired(t *testing.T) {
    //tableTest
    tests := []struct {
        testname    string
        shakenExp   time.Time
        wantError   bool
    }{
        {
            testname: "正常系",
            shakenExp:  time.Now().AddDate(1, 0, 0),
            wantError:  false,
        },
        {
            testname: "異常系（車検期限切れ）",
            shakenExp:  time.Now().AddDate(-1, 0, 0),
            wantError:  true,
        },
    }
    for _, tt := range tests {
        t.Run(tt.testname, func(t *testing.T){
            car := &Car{ShakenExpiration: tt.shakenExp}
            result := car.IsShakenExpired()
            assert.Equal(t, tt.wantError, result)
        })
    }
}

//任意保険の期限切れ確認