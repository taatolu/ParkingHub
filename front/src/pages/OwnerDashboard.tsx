import React from "react";
// OwnerDashboardコンポーネント：オーナーダッシュボードのメインコンポーネント
import OwnerList from "../features/owner/components/OwnerList"

// ダミーデータ：実際にはAPIから取得したデータを使用するが、一旦は仮のデータを使用
const dummyOwners = [
    { id: 1, first_name: "山田", middle_name: "太郎", last_name: "一郎", license_expiration: "2025-01-01" },
    { id: 2, first_name: "佐藤", middle_name: "花子", last_name: "二郎", license_expiration: "2024-12-31" },
];

const OwnerDashboard: React.FC = () => {
    // Ownerの詳細ボタンクリック時の処理
    // この処理(関数)を以下のコンポーネントの中で、OwnerListのプロップスとして、渡す
    const onDetail = (id : number) => {
        // 詳細ページへ遷移する処理をここに書く
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
        </div>
    );
};

export default OwnerDashboard