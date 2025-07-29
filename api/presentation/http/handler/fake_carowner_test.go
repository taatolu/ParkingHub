package handler

import(
    "testing"
	"net/http/httptest"
	"net/http"
	"strings"
	"io"
	"io/ioutil"
	"bytes"
    "github.com/taatolu/ParkingHub/api/mocks/usecase"
    "github.com/taatolu/ParkingHub/api/domain/model"
    )

func TestRegistCarOwner_FakeUsecase(t *testing.T){
    //tableTest
    tests := [] struct {
        testname    string
        method      string
        url         string
        body        io.Reader
        wantError   bool
    }{
        //テストケースの作成
        {
            testname:   "正常系",
            method:     "POST",
            url:        "/api/v1/car_owners",
            body:       bytes.NewBufferString(`{"id":"1",
                                                "first_name":"test",
                                                "middle_name":"山田",
                                                "last_name":"太郎",
                                                "license_expiration":"2030-12-31"}`),
            wantError:  false,
        },
        {
            testname:   "異常系（name入力不足）",
            method:     "POST",
            url:        "/api/v1/car_owners",
            body:       bytes.NewBufferString(`{"id":"1",
                                                "first_name":"",
                                                "middle_name":"",
                                                "last_name":"太郎",
                                                "license_expiration":"2030-12-31"}`),
            wantError:  true,
        },
        {
            testname:   "異常系（免許期限切れ）",
            method:     "POST",
            url:        "/api/v1/car_owners",
            body:       bytes.NewBufferString(`{"id":"1",
                                                "first_name":"test",
                                                "middle_name":"山田",
                                                "last_name":"太郎",
                                                "license_expiration":"2020-12-31"}`),
            wantError:  true,
        },
    }
    //テストケースをループで回す
    for _, tt := range tests {
        t.Run(tt.testname, func(t *testing.T){
            //CarOwnerUsecaseモックのインスタンス化
            fakeUsecase := &usecase.FakeCarOwnerUsecase{
                RegistCarOwnerFunc : func(owner *model.CarOwner) error{
                    return nil
                },
            }
            
            //ハンドラーのインスタンス生成
            handler := &CarOwnerHandler{Usecase: fakeUsecase}
            
            // httptest.NewRecorder() でレスポンス記録
            rec := httptest.NewRecorder()
            
            // http.NewRequest() でリクエスト作成
        
            req, err := http.NewRequest(tt.method, tt.url, tt.body)
            if err != nil {
                t.Fatal(err)
            }
            req.Header.Set("Content-Type", "application/json")
            
            // handler.ServeHTTPで実行
            handler.ServeHTTP(rec, req)
            
            // recorder.Result() でレスポンス検証
            resp := rec.Result()
            defer resp.Body.Close()
            
            //respからBodyの値を取得
            bodyBytes, _ := ioutil.ReadAll(resp.Body)
            bodyString := string(bodyBytes)
            
            if tt.wantError{
                //wantErrorがtrue　=　異常系だったら
                if resp.StatusCode == http.StatusCreated {
                    t.Errorf("異常系なのに201が返っています")
                }
                if !strings.Contains(bodyString, "error") && resp.StatusCode != http.StatusNotFound {
                    t.Errorf("エラーメッセージが含まれていません: %s", bodyString)
                }
            }else{
                //wantErrorがfalse　=　正常系だったら
                if resp.StatusCode != http.StatusCreated {
                    t.Errorf("GotStatus=%d WantStatus=%d", resp.StatusCode, http.StatusCreated)
                }
            }
            
            
        })
    }
}