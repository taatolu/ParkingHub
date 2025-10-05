import React from "react";

// owner型定義：オーナー１人分の情報を格納するための型
type Owner = {
    id: number;
    first_name: string;
    middle_name: string;
    last_name: string;
    license_expiration: string;
}

// OwnerListコンポーネントで使用するprops（引数）の型定義
type OwnerListProps = {
    owners: Owner[]; // Owner型の配列
}

//OwnerListコンポーネント：オーナーテーブルで表示するテーブルの本体
const OwnerList: React.FC<OwnerListProps> = ({ owners }) => {
    // ownersを使って処理を書く
    return (
        <table>
            <thead>
                <tr>
                    <th>ID</th>
                    <th>姓</th>
                    <th>ミドルネーム</th>
                    <th>名</th>
                    <th>免許証期限</th>
                </tr>
            </thead>
            <tbody>
                {owners.map((owner) => (
                    <tr key={owner.id}>
                        <td>{owner.id}</td>
                        <td>{owner.first_name}</td>
                        <td>{owner.middle_name}</td>
                        <td>{owner.last_name}</td>
                        <td>{owner.license_expiration}</td>
                    </tr>
                ))}
            </tbody>
        </table>
    );
};

export default OwnerList;