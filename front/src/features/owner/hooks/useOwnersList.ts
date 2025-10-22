import { useState, useEffect, useCallback } from 'react';
import { Owner } from '../types/Owner';
import { getOwners } from '../services/OwnerService';

//カスタムフック：オーナー情報を取得・管理する
export const useOwnersList = (authToken?: string) => {
    const [owners, setOwners] = useState<Owner[]>([]); //オーナー一覧の状態管理
    const [loading, setLoading] = useState<boolean>(true); //読み込み状態の管理(初期値はtrueで読み込み中を示す)
    const [error, setError] = useState< string | null>(null); //エラー状態の管理(初期値はnullでエラーなしを示す)

    //オーナー情報を取得する関数
    const fetchOwners = useCallback( async () => {
        setLoading(true); //読み込み開始
        setError(null); //エラー状態をリセット
        try {
            const response = await getOwners(authToken); //作成したAPI（getOwners）を呼び出し
            if (response.success && response.data) {
                setOwners(response.data); //取得したオーナー情報を状態にセット
                setLoading(false); //読み込み完了
            }
        } catch (err) {
            setError(` オーナー情報の取得に失敗しました: ${err}`); //エラー発生時にエラーメッセージをセット
            setLoading(false); //読み込み完了
        }
    }, [authToken]);

    //コンポーネントのマウント時にオーナー情報を取得
    useEffect(() => {
        fetchOwners();
    }, [fetchOwners]);
    return { owners, loading, error, fetchOwners };
};
