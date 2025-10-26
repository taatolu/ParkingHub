import { useState, useEffect, useCallback, useMemo } from 'react';
import { Owner } from '../types/Owner';
import { getOwners } from '../services/OwnerService';

//カスタムフック：オーナー情報を取得・管理する
export const useOwnersList = (authToken?: string) => {
    const [owners, setOwners] = useState<Owner[]>([]); //オーナー一覧の状態管理
    const [loading, setLoading] = useState<boolean>(true); //読み込み状態の管理(初期値はtrueで読み込み中を示す)
    const [error, setError] = useState< string | null>(null); //エラー状態の管理(初期値はnullでエラーなしを示す)
    const [showExpiredOnly, setShowExpiredOnly] = useState(false); //免許期限切れフィルタの状態管理

    // データ取得のみを担当する関数
    const fetchOwners = useCallback(async () => {
        setLoading(true);
        setError(null);
        try {
            const response = await getOwners(authToken);
            if (response.success && response.data) {
                setOwners(response.data);
            }
        } catch (error) {
            setError(`オーナー一覧の取得に失敗しました: ${error}`);
        } finally {
            setLoading(false);
        }
    }, [authToken]); // authTokenが変わったら再生成

    // フィルタリングロジックを分離しメモ化
    const filteredOwners = useMemo(() => {
        if (!showExpiredOnly) return owners;
        
        const today = new Date(); // 今日の日付を取得
        return owners.filter((owner) => {
            const expirationDate = new Date(owner.license_expiration); // オーナーの免許期限日をDateオブジェクトに変換
            return expirationDate < today; // 免許期限が今日より前の場合にtrueを返す
        });
    }, [owners, showExpiredOnly]); // ownersまたはshowExpiredOnlyが変わったら再計算
    
    // 免許期限切れフィルタのトグル切替関数
    const toggleExpiredFilter = useCallback(() => {
        setShowExpiredOnly(prev => !prev);
    }, []);

    //コンポーネントのマウント時にオーナー情報を取得
    useEffect(() => {
        fetchOwners();
    }, [fetchOwners]);

    return { 
        owners: filteredOwners, // フィルタリング済みのデータを返す
        loading, 
        error, 
        fetchOwners,
        showExpiredOnly, 
        toggleExpiredFilter
    };
};
