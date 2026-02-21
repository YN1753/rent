package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

type Rsa interface {
	GetPublicKey() (*rsa.PublicKey, error)
	GetPrivateKey() (*rsa.PrivateKey, error)
	RSADecrypt(privateKey *rsa.PrivateKey, cryptedText string) (string, error)
	GenRSAKeyPair(bits int) error
}

type RsaCrypt struct {
}

func (c *RsaCrypt) GetPublicKey() (*rsa.PublicKey, error) {
	PublicKeyBytes, err := os.ReadFile("publicKey.pem")
	if err != nil {
		return nil, errors.New("读取私钥失败")
	}
	block, _ := pem.Decode(PublicKeyBytes)
	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, errors.New("x509解密失败")
	}
	return publicKey, nil
}

func (c *RsaCrypt) GetPrivateKey() (*rsa.PrivateKey, error) {
	PrivateKeyBytes, err := os.ReadFile("privateKey.pem")
	if err != nil {
		return nil, errors.New("读取公钥失败")
	}
	block, _ := pem.Decode(PrivateKeyBytes)
	PrivateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.New("x509解密失败")
	}
	return PrivateKey, nil
}

func (c *RsaCrypt) RSADecrypt(privateKey *rsa.PrivateKey, cryptedText string) (string, error) {
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, []byte(cryptedText))
	if err != nil {
		return "", errors.New("rsa解密失败")
	}
	return string(plainText), nil
}

//func RSAEncrypt(publicKey *rsa.PublicKey, plainText string) (string, error) {
//	cryptedText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(plainText))
//	if err != nil {
//		return "", err
//	}
//	return string(cryptedText), nil
//}

func NewRsa() Rsa {
	return &RsaCrypt{}
}
func (*RsaCrypt) GenRSAKeyPair(bits int) (err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return
	}
	privateBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privatePem := pem.Block{
		Type:  "private key",
		Bytes: privateBytes,
	}
	privateFile, err := os.Create("privateKey.pem")
	if err != nil {
		return
	}
	defer func(privateFile *os.File) {
		err := privateFile.Close()
		if err != nil {
			return
		}
	}(privateFile)
	if err := pem.Encode(privateFile, &privatePem); err != nil {
		return err
	}
	publicKey := &privateKey.PublicKey
	publicBytes := x509.MarshalPKCS1PublicKey(publicKey)
	publicPem := pem.Block{
		Type:  "private key",
		Bytes: publicBytes,
	}
	publicFile, err := os.Create("publicKey.pem")
	if err != nil {
		return
	}
	defer func(publicFile *os.File) {
		err := publicFile.Close()
		if err != nil {
			return
		}
	}(publicFile)
	if err := pem.Encode(publicFile, &publicPem); err != nil {
		return err
	}
	return nil
}

//func main() {
//	Rsa := NewRsa()
//	Rsa.GenRSAKeyPair(2048)
//	privateKey, _ := Rsa.GetPrivateKey()
//	fmt.Println(privateKey)
//	publicKey, _ := Rsa.GetPublicKey()
//	fmt.Println(publicKey)
//	crypted, err := RSAEncrypt(publicKey, "你好")
//	if err != nil {
//		log.Fatalln(err)
//	}
//	fmt.Println(crypted)
//	plainText, err := Rsa.RSADecrypt(privateKey, crypted)
//	fmt.Println(plainText)
//}
