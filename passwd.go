package tsx

import "golang.org/x/crypto/bcrypt"

/*
	密码存储，采用单向加密。
1. 普通：通过有md5,sha等。但是安全性差。 rainbow table（攻击者可以将所有密码的常见组合进行单向哈希，得到一个摘要组合，然后与数据库中的摘要进行比对即可获得对应的密码）很容易反查出来
2. 高级：使用md5+盐。即：密码先md5后，在加上一个值(比如用户创建的时间)再二次md5。只要盐未泄漏，基本上不可能推倒出密码。
3. 专家级：采用golang.org/x/crypto/bcrypt模块。它相对上面更安全，原因是让计算出rainbow table耗费的资源和时间基本不可能达到。
*/

func BcryptPasswd(passwd string) (string, error) {
	encodePw, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	return string(encodePw), err
}

func BcryptValidPasswd(encodePw string, passwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(encodePw), []byte(passwd))
}
