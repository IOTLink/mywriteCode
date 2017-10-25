
package main

import (
"crypto"
"crypto/rand"
"crypto/rsa"
_ "crypto/sha256"
"fmt"
"github.com/cryptoballot/fdh"
"github.com/cryptoballot/rsablind"
)

func main() {
	message := []byte("ATTACKATDAWN")

	keysize := 2048
	hashize := 1536

	// We do a SHA256 full-domain-hash expanded to 1536 bits (3/4 the key size)
	hashed := fdh.Sum(crypto.SHA256, hashize, message)
	fmt.Printf("hashed %x\n", hashed)
	// Generate a key
	key, _ := rsa.GenerateKey(rand.Reader, keysize)

	// Blind the hashed message
	blinded, unblinder, err := rsablind.Blind(&key.PublicKey, hashed)
	if err != nil {
		panic(err)
	}

	// Blind sign the blinded message
	sig, err := rsablind.BlindSign(key, blinded)
	if err != nil {
		panic(err)
	}

	// Unblind the signature
	unblindedSig := rsablind.Unblind(&key.PublicKey, sig, unblinder)
	fmt.Printf("unblindedSig %x\n", unblindedSig)

	// Verify the original hashed message against the unblinded signature
	if err := rsablind.VerifyBlindSignature(&key.PublicKey, hashed, unblindedSig); err != nil {
		panic("failed to verify signature")
	} else {
		fmt.Println("ALL IS WELL")
	}


	//my test
	sign ,err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hashed)
	if err != nil {
		fmt.Errorf("error %s",err.Error())
		return
	}
	fmt.Printf("sign %x\n", sign)


}


//https://github.com/cryptoballot/rsablind
/*
https://www.enjoysign.com/news/185.html

盲签名
盲签名是一种特殊的数字签名方法,相对于一般的数字签名而言还应当具有下列3 个特性:
        1) 签名者不能看到明文消息;
        2) 认证者不能看到明文消息,只能通过签名来确认文件的合法性;
        3) 无论是签名者,还是认证者,都不能将盲签名与盲消息对应起来;
        在上面提到的“收方不可否认数字签名”方案中,也需要由可信第三方对双方的通信进行担保,但又不希望其获知通信的具体内容,这时,就需要用到盲签名。总之,盲签名具有消息内容的保密性以及盲签名与原消息的概率无关等特征,较好的保护了消息通信的隐私,具有较为广泛的应用,目前,它主要用于基于Internet 的匿名金融交易,如匿名电子现金支付系统、匿名电子拍卖系统等系统中。
        在某些情况下,需要由一组用户来进行数字签名,这种情况下的数字签名被称为“门限签名”或“群签名” 。门限签名的生成必须由多个成员合作才能完成,但验证只需要知道群体的公开密钥即可进行。

http://www.jianshu.com/p/839333eb5a4d

 */