import React, { useState, useEffect } from "react";
import styles from "../assets/css/Modal.module.css";

//編集対象のオーナー、情報の型を定義（useStateのジェネリクスなどとして使用）
type Owner = {
    id: number;
    first_name: string;
    middle_name: string;
    last_name: string;
    license_expiration: string;
};

//モーダルコンポーネントのプロップスの型を定義
type EditOwnerModalProps = {
    isOpen: boolean; //モーダルの表示状態
    owner: Owner | null; //編集対象のオーナー情報(nullの時は編集対象なし)
    onClose: () => void; //モーダルを閉じる関数
    onSave: (owner: Owner) => void; //保存ボタン押下時の関数(引数に編集されたオーナー情報を受け取る)
}

//EditOwnerModalコンポーネント: オーナー情報を編集するモーダル
const EditOwnerModal: React.FC< EditOwnerModalProps > = ({isOpen, owner, onClose, onSave}) => {
    //編集用の状態を管理（初期化）
    const [formData, setFormData] = useState<Owner>({
        id: 0,
        first_name: "",
        middle_name: "",
        last_name: "",
        license_expiration: "",
    });
    
    //EditOwnerModalの引数（owner）が変更されたときにformDataを更新
    useEffect(() => {
        if (owner) {
            //ownerがnullでない場合、formDataをownerの情報で更新
            setFormData({...owner});//オブジェクトの中でスプレッド演算子を使う
        } else {
            //ownerがnullの場合、formDataを初期化
            setFormData({
                id: 0,
                first_name: "",
                middle_name: "",
                last_name: "",
                license_expiration: "",
            });
        }
    }, [owner]);

    //入力フィールドの変更を処理する関数
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setFormData((prev) => ({ ...prev, [name]: value }));
    };

    //保存ボタン押下時の処理
    const handleSave = () => {
        onSave(formData); //親コンポーネントに編集されたオーナー情報を渡す
        onClose(); //モーダルを閉じる
    };
    
    //モーダルが閉じている場合は何も表示しない
    if (!isOpen) return null;

    //モーダルが開いている場合の表示内容
    return (
        <div className={styles.modal}>
            <div className={styles.modalContent}>
                <h2>オーナー情報の編集</h2>
                <form>
                    <label>
                        名:
                        <input
                            type="text"
                            name="first_name"
                            value={formData.first_name}
                            onChange={handleChange}
                        />
                    </label>
                    <label>
                        中間名:
                        <input
                            type="text"
                            name="middle_name"
                            value={formData.middle_name}
                            onChange={handleChange}
                        />
                    </label>
                    <label>
                        姓:
                        <input
                            type="text"
                            name="last_name"
                            value={formData.last_name}
                            onChange={handleChange}
                        />
                    </label>
                    <label>
                        免許証の有効期限:
                        <input
                            type="date"
                            name="license_expiration"
                            value={formData.license_expiration}
                            onChange={handleChange}
                        />
                    </label>
                    <button type="button" onClick={handleSave}>
                        保存
                    </button>
                    <button type="button" onClick={onClose}>
                        キャンセル
                    </button>
                </form>
            </div>
        </div>
    );
};

export default EditOwnerModal;
