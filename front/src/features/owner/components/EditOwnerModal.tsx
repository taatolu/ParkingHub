import React, { useState, useEffect } from "react";
import { Owner } from "../types/Owner";  //Ownerの型定義
import styles from "../assets/css/Modal.module.css";


//モーダルコンポーネントのプロップスの型を定義
type EditOwnerModalProps = {
    isOpen: boolean; //モーダルの表示状態
    owner: Owner | null; //編集対象のオーナー情報(nullの時は編集対象なし)
    onClose: () => void; //モーダルを閉じる関数
    onSave: (owner: Owner) => void; //保存ボタン押下時の関数(引数に編集されたオーナー情報を受け取る)
}

//EditOwnerModalコンポーネント: オーナー情報を編集するモーダル
export const EditOwnerModal: React.FC< EditOwnerModalProps > = ({isOpen, owner, onClose, onSave}) => {
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

    // 日付を "yyyy-MM-dd" 形式に変換する関数
    const formatDate = (dateString: string) => {
        if (!dateString) return '';
        return dateString.split('T')[0];
    };

    //入力フィールドの変更を処理する関数
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        // id フィールドの場合は数値に変換
        if (name === 'id') {
            setFormData((prev) => ({ ...prev, [name]: parseInt(value) || 0 }));
        } else {
            setFormData((prev) => ({ ...prev, [name]: value }));
        }
    };

    //保存ボタン押下時の処理
    const handleSave = () => {
        onSave(formData); //親コンポーネントで定義されたonSave関数を呼び出して保存処理を実行(PropsとしてformDataを渡す)
        onClose(); //親コンポーネントで定義されたonClose関数を呼び出してモーダルを閉じる
    };
    
    //モーダルが閉じている場合は何も表示しない
    if (!isOpen) return null;

    //モーダルが開いている場合の表示内容
    return (
        <>
            {isOpen && <div className={styles.overlay} onClick={onClose}></div>}
            {isOpen && (
                <div className={styles.modal}>
                    <div className={styles.modalContent}>
                        <div className={styles.modalHeader}>
                            <h2>オーナー情報の編集</h2>
                        </div>
                        <form>
                            <div className={styles.formGroup}>
                                <label htmlFor="id">社員ID</label>
                                <input
                                    type="number"
                                    id="id"
                                    name="id"
                                    value={formData.id}
                                    onChange={handleChange}
                                    required    //社員IDは必須入力
                                />
                            </div>
                            <div className={styles.formGroup}>
                                <label htmlFor="first_name">姓</label>
                                <input
                                    type="text"
                                    id="first_name"
                                    name="first_name"
                                    value={formData.first_name}
                                    onChange={handleChange}
                                />
                            </div>
                            <div className={styles.formGroup}>
                                <label htmlFor="middle_name">ミドルネーム</label>
                                <input
                                    type="text"
                                    id="middle_name"
                                    name="middle_name"
                                    value={formData.middle_name}
                                    onChange={handleChange}
                                />
                            </div>
                            <div className={styles.formGroup}>
                                <label htmlFor="last_name">名</label>
                                <input
                                    type="text"
                                    id="last_name"
                                    name="last_name"
                                    value={formData.last_name}
                                    onChange={handleChange}
                                />
                            </div>
                            <div className={styles.formGroup}>
                                <label htmlFor="license_expiration">免許証期限</label>
                                <input
                                    type="date"
                                    id="license_expiration"
                                    name="license_expiration"
                                    value={formatDate(formData.license_expiration)}
                                    onChange={handleChange}
                                />
                            </div>
                            <div className={styles.modalFooter}>
                                <button 
                                    type="button" 
                                    className={`${styles.button} ${styles.cancelButton}`}
                                    onClick={onClose}
                                >
                                    キャンセル
                                </button>
                                <button 
                                    type="button" 
                                    className={`${styles.button} ${styles.saveButton}`}
                                    onClick={handleSave}
                                >
                                    保存
                                </button>
                            </div>
                        </form>
                    </div>
                </div>
            )}
        </>
    );
};


