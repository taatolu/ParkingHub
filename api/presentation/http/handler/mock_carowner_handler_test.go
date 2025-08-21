package handler

import(
	"github.com/taatolu/ParkingHub/api/domain/model"
	"github.com/taatolu/ParkingHub/api/mocks/usecase"
	"testing"
	"net/http/httptest"
	"net/http"
	"strings"
	"io"
	"io/ioutil"
	"bytes"
)

// このテストでは異常系として「URL不正（404）」は検証していません。
// 理由：GoのHTTPマルチプレクサがルートにマッチしない場合は自動で404レスポンスを返すため、
// ハンドラーレベルでその挙動をテストする意味が薄いからです。
// 代わりに、method不正やbody不正など、ハンドラーが直接扱う異常系のみをテスト対象としています。


func TestRegistCarOwner(t *testing.T){
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
            testname:   "異常系（method不正）",
            method:     "GET",
            url:        "/api/v1/car_owners",
            body:       bytes.NewBufferString(`{"id":"1"}`),
            wantError:  true,
        },
    }
    //テストケースをループで回す
    for _, tt := range tests {
        t.Run(tt.testname, func(t *testing.T){
            //CarOwnerUsecaseモックのインスタンス化
            mockUsecase := &usecase.MockCarOwnerUsecase{
                RegistCarOwnerFunc : func(owner *model.CarOwner) error{
                    return nil
                },
            }
            
            //ハンドラーのインスタンス生成
            handler := &CarOwnersHandler{Usecase: mockUsecase}
            
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

func TestFindByID(t *testing.T){
    //tableTest
    tests := []struct{
        testname    string
        method      string
        url         string
        wantError   bool
    }{
        //testcae作成
        {
            testname:   "正常系",
            method:     "GET",
            url:        "/api/v1/car_owners/1",
            wantError:  false,
        },
        {
            testname:   "異常系（メソッド不正）",
            method:     "POSTPOST",
            url:        "/api/v1/car_owners/",
            wantError:  true,
        },
        {
            testname:   "異常系（url不正：パスパラメータ無し）",
            method:     "GET",
            url:        "/api/v1/car_owners/",
            wantError:  true,
        }
    }
    //testをループ処理する
    for _, tt := range tests {
        t.Run(tt.testname, func(t *testing.T){
            
            //まずはハンドラにDIするためのCarOwnerUsecaseモックのインスタンスを生成
            mockUsecase := &usecase.MockCarOwnerUsecase{
                FindByIDFunc:   func(id int)(*model.CarOwner, error){
                    if id==1 && !tt.wantError {
                        return &model.CarOwner{
                            id:1,
                            FirstName:  "Test",
                            LastName:   "User",
                        }, nil
                    }
                    return nil,  fmt.Errorf("インフラストラクチャ層の実装をしたときにデータが返ってこなかった体（テイ）")
            }

            //handlerのインスタンス生成(上で作成したCarOwnerUsecaseのモックをDI)
            handler := &CarOwnerHandler{Usecase:mockUsecase}

            //http.NewRecorderでレスポンスを記録(テスト時にhttp.ResponseWriterの代わりで動くもの)
            rec := httptest.NewRecorder

            //http.NewRequest()でリクエスト作成
            req, err := http.NewRequest(tt.method, tt.url)
            if err != nil {
                t.Fatal(err)
            }
            //リクエストのボディに JSON データを含める場合には以下の設定が必要です。
            // しかし、FindByID のようにパスパラメータのみでリクエストし、ボディにデータを含めない場合は、Content-Type ヘッダーは不要
            // req.Header.Set("Content-Type", "application/json")
            //　※返却値（レスポンス）が JSON であれば、サーバー側が Content-Type: application/json を返しますが、リクエスト側で設定する必要はない
            //【そもそも】テストのリクエストが GET メソッドで、ボディが空なら Content-Type ヘッダーは省略しましょう。
            //         POST や PUT で JSON ボディを送る場合は必須です。

            //handler.ServeHttpで実行
            handler.ServeHttp(rec, req)

            //recorder.Resultでrec(http.ResponseWriterの代わりで動くもの)に帰ってきたレスポンスを検証
            resp := rec.Result()

        })
    }
}

