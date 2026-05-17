package handler

import (
	"bytes"
	"fmt"
	"github.com/taatolu/ParkingHub/api/domain/model"
	"github.com/taatolu/ParkingHub/api/mocks/usecase"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// このテストでは異常系として「URL不正（404）」は検証していません。
// 理由：GoのHTTPマルチプレクサがルートにマッチしない場合は自動で404レスポンスを返すため、
// ハンドラーレベルでその挙動をテストする意味が薄いからです。
// 代わりに、method不正やbody不正など、ハンドラーが直接扱う異常系のみをテスト対象としています。

func TestRegistCarOwner(t *testing.T) {
	//tableTest
	tests := []struct {
		testname  string
		method    string
		url       string
		body      io.Reader
		wantError bool
	}{
		//テストケースの作成
		{
			testname: "正常系",
			method:   "POST",
			url:      "/api/v1/car_owners",
			body: bytes.NewBufferString(`{"id":1,
                                                "first_name":"test",
                                                "middle_name":"山田",
                                                "last_name":"太郎",
                                                "license_expiration":"2030-12-31"}`),
			wantError: false,
		},
	}
	//テストケースをループで回す
	for _, tt := range tests {
		tt := tt
		t.Run(tt.testname, func(t *testing.T) {
			//CarOwnerUsecaseモックのインスタンス化
			mockUsecase := &usecase.MockCarOwnerUsecase{
				RegistCarOwnerFunc: func(owner *model.CarOwner) error {
					return nil
				},
			}

			//ハンドラーのインスタンス生成
			handler := &CarOwnerHandler{Usecase: mockUsecase}

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
			bodyBytes, _ := io.ReadAll(resp.Body)
			bodyString := string(bodyBytes)

			if tt.wantError {
				//wantErrorがtrue　=　異常系だったら
				if resp.StatusCode == http.StatusCreated {
					t.Errorf("error:異常系なのに201が返っています")
				}
				if !strings.Contains(bodyString, "error") && resp.StatusCode != http.StatusNotFound {
					t.Errorf("error: エラーメッセージが含まれていません: %s", bodyString)
				}
			} else {
				//wantErrorがfalse　=　正常系だったら
				if resp.StatusCode != http.StatusCreated {
					t.Errorf("error: GotStatus=%d WantStatus=%d", resp.StatusCode, http.StatusCreated)
				}
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	//tableTest
	tests := []struct {
		testname	string
		method		string
		url			string
		wantError	bool
	}{
		//testCase作成
		{
			testname:	"正常系",
			method:		"GET",
			url:		"/api/v1/car_owners",
			wantError:	false,
		},
	}
	//testCaseのループ処理
	for _, tt := range tests {
		tt := tt
		t.Run(tt.testname, func(t *testing.T){
			//handlerにDIするためのUsecase層のモックをインスタンス化
			mockUsecase := &usecase.MockCarOwnerUsecase{
				GetAllFunc:	func() ([]*model.CarOwner, error){
					if !tt.wantError {
						return []*model.CarOwner{
							{
								ID:        1,
								FirstName: "Test",
								LastName:  "User",
							},
						}, nil
					}
					return nil, fmt.Errorf("error:取得に失敗しました")
				},
			}
			//handlerのインスタンス生成(上で作成したUsecase層のモックをDI)
			handler := &CarOwnerHandler{Usecase: mockUsecase}

			//http.NewRecorderでレスポンスを記録(テスト時にhttp.ResponseWriterの代わりで動くもの)
			rec := httptest.NewRecorder()
			//http.NewRequest()でリクエスト作成
			req, err := http.NewRequest(tt.method, tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}
			//handler.ServeHttpで実行
			handler.ServeHTTP(rec, req)
			//recorder.Resultでrec(http.ResponseWriterの代わりで動くもの)に帰ってきたレスポンスを検証
			resp := rec.Result()
			defer resp.Body.Close()

			//respからBodyの値を取得
			bodyBytes, _ := io.ReadAll(resp.Body)
			bodyString := string(bodyBytes)

			if tt.wantError {
				//tt.wantErrorがtrue=異常値の場合
				if resp.StatusCode == http.StatusOK {
					t.Errorf("異常系なのに200番が返っています")
				}
				if resp.StatusCode != http.StatusNotFound && !strings.Contains(bodyString, "error") {
					t.Error("エラーメッセージが含まれていません")
				}
			} else {
				//tt.wantErrorがfalse=正常値の場合
				if resp.StatusCode != http.StatusOK {
					t.Errorf("Got:%d Want:%d", resp.StatusCode, http.StatusOK)
				}
			}
		})
	}
}


func TestFindByID(t *testing.T) {
	//tableTest
	tests := []struct {
		testname  string
		method    string
		url       string
		wantError bool
	}{
		//testcae作成
		{
			testname:  "正常系",
			method:    "GET",
			url:       "/api/v1/car_owners/1",
			wantError: false,
		},
		{
			testname:  "異常系（メソッド不正）",
			method:    "POST",
			url:       "/api/v1/car_owners/",
			wantError: true,
		},
		{
			testname:  "異常系（url不正：パスパラメータ無し）",
			method:    "GET",
			url:       "/api/v1/car_owners/",
			wantError: true,
		},
	}
	//testをループ処理する
	for _, tt := range tests {
		tt := tt
		t.Run(tt.testname, func(t *testing.T) {

			//まずはハンドラにDIするためのCarOwnerUsecaseモックのインスタンスを生成
			mockUsecase := &usecase.MockCarOwnerUsecase{
				FindByIDFunc: func(id uint) (*model.CarOwner, error) {
					if id == 1 && !tt.wantError {
						return &model.CarOwner{
							ID:        1,
							FirstName: "Test",
							LastName:  "User",
						}, nil
					}
					return nil, fmt.Errorf("インフラストラクチャ層の実装をしたときにデータが返ってこなかった体（テイ）")
				},
			}

			//handlerのインスタンス生成(上で作成したCarOwnerUsecaseのモックをDI)
			handler := &CarOwnerHandler{Usecase: mockUsecase}

			//http.NewRecorderでレスポンスを記録(テスト時にhttp.ResponseWriterの代わりで動くもの)
			rec := httptest.NewRecorder()

			//http.NewRequest()でリクエスト作成
			req, err := http.NewRequest(tt.method, tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}
			//リクエストのボディに JSON データを含める場合には以下の設定が必要。
			// だが、FindByID のようにパスパラメータのみでリクエストし、ボディにデータを含めない場合は、Content-Type ヘッダーは不要
			// req.Header.Set("Content-Type", "application/json")
			//　※返却値（レスポンス）が JSON であれば、サーバー側が Content-Type: application/json を返しますが、リクエスト側で設定する必要はない
			//【そもそも】テストのリクエストが GET メソッドで、ボディが空なら Content-Type ヘッダーは省略しましょう。
			//         POST や PUT で JSON ボディを送る場合は必須です。

			//handler.ServeHttpで実行
			handler.ServeHTTP(rec, req)

			//recorder.Resultでrec(http.ResponseWriterの代わりで動くもの)に帰ってきたレスポンスを検証
			resp := rec.Result()
			defer resp.Body.Close()

			//respからBodyの値を取得
			bodyBytes, _ := io.ReadAll(resp.Body)
			bodyString := string(bodyBytes)

			if tt.wantError {
				//tt.wantErrorがtrue=異常値の場合
				if resp.StatusCode == http.StatusOK {
					t.Errorf("異常系なのに200番が返っています")
				}
				if resp.StatusCode != http.StatusNotFound && !strings.Contains(bodyString, "error") {
					t.Error("エラーメッセージが含まれていません")
				}
			} else {
				//tt.wantErrorがfalse=正常値の場合
				if resp.StatusCode != http.StatusOK {
					t.Errorf("Got:%d Want:%d", resp.StatusCode, http.StatusOK)
				}
			}
		})
	}
}


func TestFindByName(t *testing.T) {
	//tableTest
	tests := []struct{
		testname	string
		method		string
		url			string
		expectError	error
	}{
		//testCaseの作成
		{
			testname:	"正常系",
			method:		"GET",
			url:		"/api/v1/car_owners/田",
			expectError:	nil,
		},
		{
			testname:	"異常系（method不正）",
			method:		"POST",
			url:		"/api/v1/car_owners/田",
			expectError:	fmt.Errorf("error:methodが不正です"),
		},
		{
			testname:	"異常系（パラメータ不正）",
			method:		"GET",
			url:		"/api/v1/car_owners/",
			expectError:	fmt.Errorf("error:パラメータが空です"),
		},
	}
	//TestCaseをループ処理
	for _, tt := range tests {
		tt := tt
		t.Run(tt.testname, func(t *testing.T) {
			// モックユースケースの用意
			mock := &usecase.MockCarOwnerUsecase{
				FindByNameFunc: func(name string) ([]*model.CarOwner, error) {
					if tt.expectError != nil {
						return nil, tt.expectError
					}
					return []*model.CarOwner{
						{ID: 1, FirstName: "山田", MiddleName: "jon", LastName: "健二", LicenseExpiration: time.Now().AddDate(1, 0, 0)},
						{ID: 2, FirstName: "井本", MiddleName: "有田", LastName: "", LicenseExpiration: time.Now().AddDate(1, 0, 0)},
					}, nil
				},
			}

			//handlerのインスタンス生成(UsecaseのモックをDI)
			handler := &CarOwnerHandler{Usecase: mock}

			//httptest.NewRecorder()でレスポンスを記録（テストの時にhttp.ResposseWriterの代わりになるもの）
			rec := httptest.NewRecorder()

			// http.NewRequest() でリクエスト作成
			req, err := http.NewRequest(tt.method, tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			// handler.ServeHTTPで実行
			handler.ServeHTTP(rec, req)
			// recorder.Result() でレスポンス検証
			// 受け取ったrecのBodyを取得
			resp := rec.Result()
			defer resp.Body.Close()

			//上記Bodyの中身を取得
			bodyBytes, err := io.ReadAll(resp.Body)
			bodyString := string(bodyBytes)

			if tt.expectError != nil {
				if resp.StatusCode == http.StatusOK {
					t.Errorf("異常系なのに200番が返っています")
				}
				if resp.StatusCode != http.StatusNotFound && !strings.Contains(bodyString, "error") {
					t.Errorf("エラーメッセージが含まれていません")
				}
			} else {
				if resp.StatusCode != http.StatusOK {
					t.Errorf("エラーになってしまった: got: %d  want: %d ", resp.StatusCode, http.StatusOK)
				}
			}
		})
	}
}

// Updateのテスト(Usecaseのモックを使うのでMockテストといいます)
func TestUpdate_mock(t *testing.T) {
	//tableTest
	tests := []struct {
		testname   string
		method     string
		url        string
		body       io.Reader
		wantError bool
	}{
		//テストケースの作成
		{
			testname: "正常系",
			method:   "PUT",
			url:      "/api/v1/car_owners/1",
			body: bytes.NewBufferString(`{"first_name":"test", "middle_name":"山田", "last_name":"太郎", "license_expiration":"2030-12-31"}`),
			wantError: false,
		},
		{
			testname:  "異常系（method不正）",
			method:    "GET",
			url:       "/api/v1/car_owners/1",
			body:      bytes.NewBufferString(`{"first_name":"test"}`),
			wantError: true,
		},
		{
			testname:  "異常系（body不正）",
			method:    "PUT",
			url:       "/api/v1/car_owners/1",
			// 不正なJSON
			body:      bytes.NewBufferString(`{"first_name":`),
			wantError: true,
		},
	}
	//テストケースをループで回す
	for _, tt := range tests {
		tt := tt
		t.Run(tt.testname, func(t *testing.T) {
			//CarOwnerUsecaseモックのインスタンス化
			mockUsecase := &usecase.MockCarOwnerUsecase{
				UpdateFunc: func(owner *model.CarOwner) error {
					return nil
				},
			}
			
			//ハンドラーのインスタンス生成
			handler := &CarOwnerHandler{Usecase: mockUsecase}
		
			// httptest.NewRecorder() でレスポンス記録
			rec := httptest.NewRecorder()

			// http.NewRequest()でリクエストの作成
			req, err := http.NewRequest(tt.method, tt.url, tt.body)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")

			// handler.ServeHTTPで実行
			handler.ServeHTTP(rec, req)

			// recorder.Result()でレスポンス情報を検証
			resp := rec.Result()
			defer resp.Body.Close()

			// respで受けとたったBodyの値を取得
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}
			bodyString := string(bodyBytes)

			if tt.wantError {
				//wantErrorがtrue = 異常系だったら
				if resp.StatusCode == http.StatusOK {
					t.Errorf("異常系であるべきなのに200が返っています")
				}
				// Bodyに”error”が含まれているか、かつステータスコードが404でないことを確認
				if !strings.Contains(bodyString, "error") && resp.StatusCode != http.StatusNotFound {
					t.Errorf("エラーメッセージが含まれていません: %v", bodyString)
				}
			} else {
				//wantErrorがfalse = 正常系だったら
				if resp.StatusCode != http.StatusOK {
					t.Errorf("GotStatus=%d WantStatus=%d", resp.StatusCode, http.StatusOK)
				}
				// Bodyに"error"が含まれていないことを確認
				if strings.Contains(bodyString, "error") {
					t.Errorf("正常系なのにエラーメッセージが含まれています: %v", bodyString)
				}
			}
		})
	}
}

// Deleteのテスト(Usecaseのモックを使う)
func TestDelete_mock(t *testing.T) {
	//テーブルテスト
	tests := []struct {
		testname  string
		method    string
		url       string
		wantError bool
	}{
		//テストケースの作成
		{
			testname:	"正常系",
			method:		"DELETE",
			url:		"/api/v1/car_owners/1",
			wantError:	false,
		},
		{
			testname:	"異常系(method不正)",
			method:		"GET",
			url:		"/api/v1/car_owners/1",
			wantError:	true,
		},
		{
			testname:	"異常系(パスパラメータ無し)",
			method:		"DELETE",
			url:		"/api/v1/car_owners/",
			wantError:	true,
		},
	}
	//テストケースをループで回す
	for _, tt := range tests {
		tt := tt
		t.Run(tt.testname, func(t *testing.T) {
			//CarOwnerUsecaseモックのインスタンス化
			//tt.wantErrorがtrueのときはエラーを返すようにする
			mockUsecase := &usecase.MockCarOwnerUsecase{
				DeleteFunc: func(id uint) error {
					if tt.wantError {
						return fmt.Errorf("インフラ層の実装をしたときにデータが返ってこなかった体（テイ）")
					}
					return nil
				},
			}
			//test対象のhandlerのインスタンス生成
			handler := &CarOwnerHandler{Usecase: mockUsecase}
			//httptest.NewRecorder()でレスポンスを記録(テスト時にhttp.ResponseWriterの代わりで動くもの)
			rec := httptest.NewRecorder()

			//http.NewRequest()でリクエスト作成
			req, err := http.NewRequest(tt.method, tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			//handler.ServeHTTPで実行
			handler.ServeHTTP(rec, req)

			//recorder.Result()でレスポンス検証
			resp := rec.Result()
			defer resp.Body.Close()

			//respからBodyの値を取得
			bodyBytes, _ := io.ReadAll(resp.Body)
			bodyString := string(bodyBytes)

			if tt.wantError {
				//wantErrorがtrue = 異常系だったら
				if resp.StatusCode == http.StatusNoContent {
					t.Errorf("異常系なのに204が返っています")
				}
				if resp.StatusCode != http.StatusNotFound && !strings.Contains(bodyString, "error") {
					t.Errorf("エラーメッセージが含まれていません: %s", bodyString)
				}
			} else {
				//wantErrorがfalse = 正常系だったら
				if resp.StatusCode != http.StatusNoContent {
					t.Errorf("GotStatus=%d WantStatus=%d", resp.StatusCode, http.StatusNoContent)
				}
			}
		})

	}
}