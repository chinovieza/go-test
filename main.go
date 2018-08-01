package main

import (
	"encoding/json"
	"fmt"
)

//JSONToNan struct ของ json ที่จะส่งให้ nan
type JSONToNan struct {
	TVSID          string   `json:"tvs_id"`
	CustomerID     string   `json:"customer_id"`
	SubscriptionID int64    `json:"subscription_id"`
	AlacartePack   []string `json:"alacarte_pack"`
}

func main() {

	strFromDB := `["TVS_COMBO","TVS_TSS_PACK"]`

	//ทำ json string ที่ได้จาก DB มาทำให้เป็น array(slice) ก่อน
	var f []string
	_ = json.Unmarshal([]byte(strFromDB), &f)

	// fmt.Println(f[0])
	// fmt.Println(f[1])

	//สร้าง json ที่จะเอาไว้ส่งให้ แนน
	jToNan := JSONToNan{}

	jToNan.TVSID = "2468"
	jToNan.CustomerID = "8876"
	jToNan.SubscriptionID = 1234
	jToNan.AlacartePack = f

	jsonProfile, _ := json.MarshalIndent(jToNan, "", "  ")
	fmt.Println(string(jsonProfile))

}
