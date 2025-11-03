package model

import(
    "time"
    "gorm.io/gorm"
    )

// Car構造体（車両情報）
// OwnerIDはCarOwner.IDのみ参照可能な外部キーです。
// Ownerフィールドで所有者情報を取得できます。
type Car struct {
    gorm.Model
    Region              string      `json:"region" gorm:"column:region;not null"`              //地域
    VehicleClass        uint        `json:"vehicle_class" gorm:"column:vehicle_class"`        //原付は分類番号がない場合がある
    UsageKana           string      `json:"usage_kana" gorm:"column:usage_kana;size:3;not null"`  //ひらがな
    SerialNumber1       uint        `json:"serial1" gorm:"column:serial1;not null"`     //ナンバー（-の前2数字）
    SerialNumber2       uint        `json:"serial2" gorm:"column:serial2;not null"`     //ナンバー（-の後2数字）
    ShakenExpiration    time.Time   `json:"shaken_exp" gorm:"column:shaken_exp"`                //車検期限(原付は車検なし)
    InsuranceExpiration time.Time   `json:"insurance_exp" gorm:"column:insurance_exp;not null"`//任意保険期限
    OwnerID             uint        `json:"owner_id" gorm:"column:owner_id;not null"`
    Owner               CarOwner    `json:"owner" gorm:"foreignKey:OwnerID;references:ID"`  //Car構造体のOwnerIDはOwnerのIDにを参照する
    Note                string      `json:"note" gorm:"column:note;size:255"`               //備考欄
}

// 車検の期限切れ確認
func (c *Car) IsShakenExpired () bool {
    //現在時刻がc.ShakenExpirationより後かどうか（現在時刻の方が後ならtrueを返す）
    return time.Now().After(c.ShakenExpiration)
}

// 任意保険の期限切れ確認
func (c *Car) IsInsuranceExpired () bool {
    //現在時刻がc.IsInsuranceExpirationより後かどうか（現在時刻の方が後ならtrueを返す）
    return time.Now().After(c.InsuranceExpiration)
}