//apiのベースURLを定義
export const API_BASE_URL = "http://localhost:3000/api/v1";

//apiの共通ヘッダーを定義
export const API_HEADERS =  {
    "Content-Type" : "application/json",
    //必要に応じて認証トークンなどを追加
};

//apiレスポンスの共通処理を定義
export type ApiResponse<T> = {
    success: boolean;
    data?: T;
    error?: string;
};

//apiエラーハンドリングの共通処理を定義
export class ApiError extends Error {
    statusCode?: number;

    constructor(message: string, statusCode?: number) {
        super(message);
        this.statusCode = statusCode;
        this.name = "ApiError";
    }
};