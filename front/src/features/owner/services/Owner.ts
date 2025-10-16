//CarOwner関連のAPIを作成
import { API_BASE_URL, API_HEADERS, ApiError, ApiResponse } from  "./apiConfig";
import { Owner } from "../types/Owner";


//ヘッダー作成のヘルパー関数
//必要に応じてBearer トークンを追加可能
// 三項演算子でauthTokenが付与されているか判断
const makeHeaders = (authToken?:string) => {
    return authToken
        ?{...API_HEADERS, Authorization: `Bearer${authToken}`}
        :{...API_HEADERS};
};

/*Fetchレスポンスを共通処理するヘルパー関数
　成功時は ApiResponse<T> を返し、失敗時は ApiError をthrowします。
*/
async function handleResponse<T>(res: Response):Promise<ApiResponse<T>>{
    const text await res.text();
    let json: any = null;
    
    if (text) {
        try {
            json = JSON.parse(text)
        } catch (err) {
            //JSONパースに失敗した場合
            throw new ApiError("Invalid JSON response", res.status);
        }
    }
    if (res.ok) {
        return { succcess: true, data:json as T };
    }
    
    const message = json?.error || json?.message || text || "Unknown error";
    throw new ApiError(message, res.status);
}