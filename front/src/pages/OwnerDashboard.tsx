import React, { useState, useMemo} from "react";
import styles from "./css/OwnerDashboard.module.css"; // CSSをインポート
// OwnerDashboardコンポーネント：オーナーダッシュボードのメインコンポーネント
import { Owner} from "../features/owner/types/Owner"
import { OwnerList } from "../features/owner/components/OwnerList"
import { EditOwnerModal } from "../features/owner/components/EditOwnerModal";
import { useOwnersList } from "../features/owner/hooks/useOwnersList";
import { useOwnerDetails } from "../features/owner/hooks/useOwnersDetail";

export const OwnerDashboard: React.FC = () => {
    // フック関連(作成したカスタムフックからの受け取りをここに記述)
    // オーナー一覧取得フック（useOwnersList)がreturnする値を分割代入で受け取り
    const { owners, loading: listLoading, error: listError, fetchOwners, showExpiredOnly, toggleExpiredFilter } = useOwnersList();
    // オーナー詳細取得フック（useOwnerDetail)がreturnする値を分割代入で受け取り
    const { selectedOwner, loading: detailLoading, error: detailError, fetchOwnerDetails, createOwnerDetails, updateOwnerDetails } = useOwnerDetails();

    //modal関連
    // モーダルの表示状態を管理するstateの定義
    const [isOpen, setIsOpen] = useState(false);

    // Button関連の関数（処理を定義）
    // Ownerの詳細ボタンクリック時の処理
    // この処理(関数)を以下のコンポーネントの中で、OwnerListのプロップスとして、渡す
    const onDetail = (id : number) => {
        // 詳細ページへ遷移する処理をここに書く(EditOwnerModalを開く)
        fetchOwnerDetails(id);
        setIsOpen(true);
    };
    // Ownerの保存（onSave）ボタンクリック時の処理
    // この処理(関数)を以下のコンポーネントの中で、EditOwnerModalのプロップスとして、渡す
    const onSave = (owner : Owner) => {
        // 保存処理をここに書く
        createOwnerDetails(owner);
    }
    // Ownerの削除ボタンクリック時の処理
    // この処理(関数)を以下のコンポーネントの中で、OwnerListのプロップスとして、渡す
    const onDelete = (id : number) => {
        // 削除処理をここに書く
    };

    // オーナー一覧を表示するためのコンポーネント
    return (
        <div className={styles.container}>
            <div className={styles.header}>
                <h1 className={styles.title}>オーナー一覧</h1>
                {/* 免許期限切れフィルタのトグルボタン */}
                <div className={styles.toggleContainer}>
                    <label className={styles.toggleLabel}>
                        <input
                            type="checkbox"
                            checked={showExpiredOnly}
                            onChange={toggleExpiredFilter}
                        />
                        免許期限切れのみ表示
                    </label>
                </div>
                {/* 新規作成ボタン */}
                <button className={styles.createButton} onClick={() => {
                    setIsOpen(true);
                }}>
                    新規作成
                </button>
            </div>
            {/* OwnerListコンポーネントを表示 */}
            <div className={styles.tableContainer}>
                <OwnerList
                    owners={owners}    //OwnerListコンポーネントのpropsに絞り込んだオーナーリストを渡す
                    onDetail={onDetail}  //OwnerListコンポーネントのpropsにonDetail関数を渡す
                    onDelete={onDelete}  //OwnerListコンポーネントのpropsにonDelete関数を渡す
                />
            </div>
        {/* モーダルコンポーネントを追加 */}
        <EditOwnerModal
            isOpen={isOpen}
            owner={selectedOwner}
            onClose={() => setIsOpen(false)}
            onSave={(updatedOwner) => {
                // 保存処理を実装
                console.log('更新されたオーナー:', updatedOwner);
                setIsOpen(false);
            }}
        />
        </div>
    );
};

