package http

import "net/http"

// CORSミドルウェアの実装
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// どのオリジンからのアクセスを許可するか
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
        // どのHTTPメソッドを許可するか
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        // どんなヘッダーを許可するか
        w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
        
		// プリフライトリクエストに対する応答
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// 次のハンドラーを呼び出す
		next.ServeHTTP(w, r)
	})
}