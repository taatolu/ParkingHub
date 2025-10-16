//CarOwner関連のAPIを作成
import { API_BASE_URL, API_HEADERS, ApiError, ApiResponse } from  "./apiConfig";
import { Owner } from "../types/Owner";


//ヘッダー作成のヘルパー関数
//必要に応じてBearer トークンを追加可能
// 三項演算子でauthTokenが付与されているか判断
const makeHeaders = (authToken?:string) => {
    // authTokenが文字列であり、空白でない場合にAuthorizationヘッダーを追加
    return typeof authToken === 'string' && authToken.trim() !== ''
        ?{...API_HEADERS, Authorization: `Bearer ${authToken}`}
        :{...API_HEADERS};
};

/*Fetchレスポンスを共通処理するヘルパー関数
　成功時は ApiResponse<T> を返し、失敗時は ApiError をthrowします。
*/
async function handleResponse<T>(res: Response):Promise<ApiResponse<T>>{
    const text = await res.text();
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
        return { success: true, data:json as T };
    }

    /*エラーメッセージを抽出
    エラー属性があればその値を使用 => jsonオブジェクトのメッセージ属性を使用 => textを使用 => Unknown error
    の順で評価*/
    const message = json?.error || json?.message || text || "Unknown error";
    throw new ApiError(message, res.status);
}

//APIサービス関数群
//オーナー一覧取得
export async function getOwners(authToken?:string):Promise<ApiResponse<Owner[]>>{
    try {
        const res = await fetch( `${API_BASE_URL}/owners`,{
            method: "GET",
            headers: makeHeaders(authToken)
        });
        //レスポンスを共通処理
        /* 直前のfetchでとってきたレスポンスを共通のレスポンス処理関数で処理 */
        const result = await handleResponse<Owner[]>(res);
        return result;
    //Try内でエラーが発生した場合catchで捕捉(error処理)
    } catch (error) {
        console.error("オーナー一覧取得エラー", error);
        throw error;
    }
}