import React, { useState } from "react";
// OwnerDashboardコンポーネント：オーナーダッシュボードのメインコンポーネント
import OwnerList from "../features/owner/components/OwnerList"
import EditOwnerModal from "../features/owner/components/EditOwnerModal";

// ダミーデータ：実際にはAPIから取得したデータを使用するが、一旦は仮のデータを使用
const dummyOwners = [
    { id: 1, first_name: "山田", middle_name: "太郎", last_name: "一郎", license_expiration: "2025-01-01" },
    { id: 2, first_name: "佐藤", middle_name: "花子", last_name: "二郎", license_expiration: "2024-12-31" },
];

const OwnerDashboard: React.FC = () => {
    // Owner型の定義（下で出てくるselectedOwnerやopenModalの型として使用）
    type Owner = {
        id: number;
        first_name: string;
        middle_name: string;
        last_name: string;
        license_expiration: string;
    };
    //modal関連
    // モーダルの表示状態を管理するstateの定義
    const [isOpen, setIsOpen] = useState(false);
    // 編集対象のオーナー情報を管理するstateの定義
    //↓useStateは<null | { ... }>のように、nullかオーナー情報のオブジェクトを持つことができる
    const [selectedOwner, setSelectedOwner] = useState<null | Owner>(null);

    // モーダルを開く関数
    const openModal = (owner: Owner) => {
        setSelectedOwner(owner);
        setIsOpen(true);
    };


    // Button関連
    // Ownerの詳細ボタンクリック時の処理
    // この処理(関数)を以下のコンポーネントの中で、OwnerListのプロップスとして、渡す
    const onDetail = (id : number) => {
        // 詳細ページへ遷移する処理をここに書く(EditOwnerModalを開く)
        const owner = dummyOwners.find((o) => o.id === id);
        if (owner) {
            openModal(owner);
        }
    };
    // Ownerの削除ボタンクリック時の処理
    // この処理(関数)を以下のコンポーネントの中で、OwnerListのプロップスとして、渡す
    const onDelete = (id : number) => {
        // 削除処理をここに書く
    };

    // オーナー一覧を表示するためのコンポーネント
    return (
        <div>
            <h1>オーナー一覧</h1>
            <OwnerList
                owners={dummyOwners}    //OwnerListコンポーネントのpropsにダミーデータを渡す
                onDetail={onDetail}  //OwnerListコンポーネントのpropsにonDetail関数を渡す
                onDelete={onDelete}  //OwnerListコンポーネントのpropsにonDelete関数を渡す
            />
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

export default OwnerDashboard