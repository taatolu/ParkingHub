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
        const res = await fetch( `${API_BASE_URL}/car_owners`,{
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

//オーナー詳細取得(ID指定)
export async function getOwnerByID(id: number, authToken?:string):Promise<ApiResponse<Owner>>{
    try {
        const res = await fetch( `${API_BASE_URL}/car_owners/${id}`,{
            method: "GET",
            headers: makeHeaders(authToken)
        });
        //レスポンスを共通処理
        const result = await handleResponse<Owner>(res);
        return result;
    //Try内でエラーが発生した場合catchで捕捉(error処理)
    } catch (error) {
        console.error("オーナー詳細取得エラー", error);
        throw error;
    }
}

//オーナー情報更新
export async function updateOwner(id:number, ownerData:Partial<Owner>, authToken?:string):Promise<ApiResponse<Owner>>{
    try {
        const res = await fetch( `${API_BASE_URL}/car_owners/${id}`, {
            method: "PUT",
            headers: makeHeaders(authToken),
            body: JSON.stringify(ownerData),
        });
        //レスポンスを共通処理
        const result = await handleResponse<Owner>(res);
        return result;
    //Try内でエラーが発生した場合catchで捕捉(error処理)
    } catch (error) {
        console.error("オーナー情報更新エラー", error);
        throw error;
    }
}

//オーナー作成
export async function createOwner (ownerData:Owner, authToken?:string):Promise<ApiResponse<Owner>>{
    try {
        const res = await fetch( `${API_BASE_URL}/car_owners`, {
            method: "POST",
            headers: makeHeaders(authToken),
            body: JSON.stringify(ownerData),
        });
        //レスポンスを共通処理
        const result = await handleResponse<Owner>(res);
        return result;
    //Try内でエラーが発生した場合catchで捕捉(error処理)
    } catch (error) {
        console.error("オーナー作成エラー", error);
        throw error;
    }
}

//オーナー削除
export async function deleteOwner (id:number, authToken?:string):Promise<ApiResponse<Owner>>{
    try {
        const res = await fetch( `${API_BASE_URL}/car_owners/${id}`, {
            method: "DELETE",
            headers: makeHeaders(authToken),
        });
        //レスポンスを共通処理
        const result = await handleResponse<Owner>(res);
        return result;
    } catch (error) {
        console.error("オーナーの削除エラー", error);
        throw error;
    }
}

//オーナー名前検索 (APIエンドポイントがある場合)
export async function searchOwnersByName(searchText: string, authToken?: string): Promise<ApiResponse<Owner[]>> {
    try {
        // バックエンドの実装に合わせて、パスパラメータとして検索テキストを追加
        const res = await fetch(`${API_BASE_URL}/car_owners/${searchText}`, {
            method: "GET",
            headers: makeHeaders(authToken)
        });
        //レスポンスを共通処理
        const result = await handleResponse<Owner[]>(res);
        return result;
    } catch (error) {
        console.error("オーナー名前検索エラー", error);
        throw error;
    }
}