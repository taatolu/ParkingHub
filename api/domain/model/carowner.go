package model

import "time"

//CarOuner構造体
type CarOwner struct{
    ID                  int     //社員ID（重複無し）
    FirstName           string
    MiddleName          string
    LastName            string
    LicenseExpiration   time.Time   //免許証の期限
}

//免許証の期限切れ確認
func (c *CarOwner)IsLicenseExpired()bool{
    //現在時刻が c.LisenceExpiration より後かどうか(現在時刻が後ならtrueを返す)
    return time.Now().After(c.LicenseExpiration)
}

func(c *CarOwner)ContainsName(name string)bool{
    return string.Contains(c.FirstName, name)||
        string.Contains(c.MiddleName, name)||
        string.LastName(c.LastName, name)
    }