import { useState, useCallback} from  'react';
import { Owner} from '../types/Owner';
import { getOwnerByID, updateOwner, createOwner } from '../services/OwnerService';
import { create } from 'domain';

/*カスタムフック：オーナー詳細情報を取得・管理する
このフックは、特定のオーナーの詳細情報の取得と更新の機能を提供します。
*/
export const useOwnerDetails = (authToken?: string) => {
    const [selectedOwner, setSelectedOwner] = useState<Owner | null>(null); //選択されたオーナーの状態管理
    const [loading, setLoading] = useState<boolean>(true); //読み込み状態の管理
    const [error, setError] = useState< string | null>(null); //エラー状態の管理

    //オーナーの詳細情報を取得する関数
    const fetchOwnerDetails = async (ownerID: number) => {
        setLoading(true); //読み込み開始
        setError(null); //エラー状態をリセット
        setSelectedOwner(null); //前回の選択をリセット
        try {
            const response = await getOwnerByID(ownerID, authToken); //作成したサービス関数を呼び出し
            if (response.success && response.data) {
                setSelectedOwner(response.data); //選択されたオーナーの情報を状態にセット
                setLoading(false); //読み込み完了
                return response.data; //データを直接返す
            }
        } catch (error) {
            setError(`オーナー情報の取得に失敗しました: ${error}`); //エラー発生時にエラーメッセージをセット
            throw error;
        } finally {
            setLoading(false); //読み込み完了
        }
    }
    
    //オーナー情報を更新する関数
    const updateOwnerDetails = async (ownerID: number, ownerData: Partial<Owner>) => {
        setError(null); //エラー状態をリセット
        setLoading(true); //読み込み開始
        try {
            const response = await updateOwner(ownerID, ownerData, authToken);
            if (response.success && response.data) {
                setSelectedOwner(response.data); //更新されたオーナーの情報を状態にセット
                setLoading(false); //読み込み完了
                return response.data; //データを直接返す
            }
        } catch (error) {
            setError( `オーナー情報の更新に失敗しました: ${error}`); //エラー発生時にエラーメッセージをセット
            throw error;
        } finally {
            setLoading(false); //読み込み完了
        }
    }

    // オーナー情報の新規作成
    const createOwnerDetails = async (ownerDetails: Owner) => {
        setError(null); //エラー状態をリセット
        setLoading(true); //読み込み開始
        try {
            const response = await createOwner(ownerDetails, authToken);
            if (response.success && response.data) {
                setSelectedOwner(response.data); //作成されたオーナーの情報を状態にセット
                setLoading(false); //読み込み完了
                return response.data; //データを直接返す
            }
        } catch (error) {
            setError( `オーナー情報の作成に失敗しました: ${error}`); //エラー発生時にエラーメッセージをセット
            throw error;
        } finally {
            setLoading(false); //読み込み完了
        }
    }
    
    // フックが返すオブジェクト
    return { selectedOwner, loading, error, fetchOwnerDetails, updateOwnerDetails, createOwnerDetails };
}