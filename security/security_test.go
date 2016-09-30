package security

import (
	"testing"
	//"fmt"
)


func TestMd5(t *testing.T)  {
	var key, data, enData string

	key, data = "123456", "前面liufuling123456我@#哈哈哈，您是逗逼吗？"
	
	t.Logf("key:%v\ndata:%v\n", key, data)
	
	enData = Hmacmd5(data, key)
	t.Logf("HmacMd5:%v\n", enData)

	enData = Md5(data)
	t.Logf("Md5:%v\n", enData)

	enData = Md5_16(data)
	t.Logf("Md5_16:%v\n-----------------\n", enData)
}



func TestPassword(t *testing.T)  {
	pwd, errpwd := "W1@1236&*n123Q", "W1@1236&*n123q"
	saltpwd := SaltPassword(pwd)
	
	t.Logf("pwd:%v\nerrpwd:%v\n", pwd, errpwd)
	t.Logf("saltPwd:%v\n", saltpwd)
	// hashValues := strings.Split(saltpwd, ":")
	// t.Logf("values: %v, len:%v", hashValues, len(hashValues))
	//判断校验是否正确
	if !VerifySaltPassword(pwd, saltpwd){
		t.Errorf("pwd verify fail => pwd:%v, saltPwd:%v", pwd, saltpwd)
	}
	if VerifySaltPassword(errpwd, saltpwd){
		t.Fatalf("errpwd verify success => errpwd:v%, saltPwd:%v", pwd, saltpwd)
	}
}

func TestAes(t *testing.T)  {
	key, data := "123456", "前面liufuling123456我@#哈哈哈，您是逗逼吗？"
	enData, err := AesEncode(data, key)
	if err != nil{
		t.Fatalf("AesEncode err %v", err)
	}
	deData, err := AesDecode(enData, key)
	if err != nil{
		t.Fatalf("AesDecode err %v", err)
	}
	t.Logf("\n-----------------\nkey: %v\ndata: %v\nencode: %v\ndecode: %v\n-----------------\n", key, data, enData, deData)
}
