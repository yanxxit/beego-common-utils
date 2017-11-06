package crypto

import (
	//"encoding/base64"
	"fmt"
	"testing"
	"github.com/shawflying/beego-common-utils/utils/comutil"
)

func TestAesStringEncrypt(t *testing.T) {
	for i := 0; i < 1000; i++ {
		res, err := AesStringEncrypt(comutil.TransInterfaceToString(i), "shdxmohoo@189.cn")
		fmt.Println("加密：", res, err) //ybXYkem1X9drSV9Bs1iN3YNeSjKY/SvnjOZYwHi3qdo=

		ares, aerr := AesStringDecrypt(res, "shdxmohoo@189.cn")

		fmt.Println("解密：", ares, aerr) //ybXYkem1X9drSV9Bs1iN3YNeSjKY/SvnjOZYwHi3qdo=
		fmt.Println("结束。。。。。。。。。。。。。。。。。。。。。。")
	}

}

func TestAesStringDecrypt(t *testing.T) {
	//res, err := AesStringDecrypt("3Ov/I3LlCu//+PmOp8DNzw==", "shdxmohoo@189.cn")
	////res, err := AesStringDecrypt("123", "shdxmohoo@189.cn")
	//
	//fmt.Println("解密：", res, err) //ybXYkem1X9drSV9Bs1iN3YNeSjKY/SvnjOZYwHi3qdo=
}

//func TestAes(t *testing.T) {
//	// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
//	key := []byte("wechat-mobile-17")
//	result, err := AesEncrypt([]byte("oKXUCj1MOddnp-sCpGi1J1dg3TyM"), key)
//	if err != nil {
//		panic(err)
//	}
//	jm := base64.StdEncoding.EncodeToString(result);
//	jiemi, _ := base64.StdEncoding.DecodeString(jm);
//	origData, err := AesDecrypt(jiemi, key)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(string(origData))
//}
