package access

import (
	"fmt"
	"time"

	"github.com/DemoLiang/go-micro/basic/config"
	"github.com/dgrijalva/jwt-go"
)

func (s *service)createTokenClaims(subject *Subject)(m *jwt.StandardClaims,err error){
	now:=time.Now()
	m = &jwt.StandardClaims{
		ExpiresAt: now.Add(tokenExpiredDate).Unix(),
		NotBefore: now.Unix(),
		Id:subject.Id,
		IssuedAt:now.Unix(),
		Subject:subject.Id,
	}
	return
}


func (s *service)saveTokenCache(subject *Subject,val string)(err error){
	if err=ca.Set(tokenIDKeyPrefix+subject.Id,val,tokenExpiredDate).Err();err!=nil{
		return fmt.Errorf("[saveTokenCache] 保存token到缓存失败:",err.Error())
	}
	return
}

func (s *service)delTokenFromCache(subject *Subject)(err error){
	if err=ca.Del(tokenIDKeyPrefix+subject.Id).Err();err!=nil{
		return fmt.Errorf("[delTokenFromCache]晴空token错误",err.Error())
	}
	return
}

func (s *service)getTokenFromCache(subject *Subject)(token string,err error){
	tokenCached,err:=ca.Get(tokenIDKeyPrefix+subject.Id).Result()
	if err!=nil{
		return token,fmt.Errorf("[getTokenFromCache]token get err",err.Error())
	}
	return string(tokenCached),nil
}

func (s *service)parseToken(tk string)(c *jwt.StandardClaims,err error){
	token,err:=jwt.Parse(tk, func(token *jwt.Token) (i interface{}, e error) {
		_,ok:=token.Method.(*jwt.SigningMethodHMAC)
		if !ok{
			return nil,fmt.Errorf("不合法的Token格式:%v",token.Header["alg"])
		}
		return []byte(config.GetJwtConfig().GetSecretKey()),nil
	})
	if err!=nil{
		switch e:=err.(type) {
		case *jwt.ValidationError:
			switch e.Errors {
			case jwt.ValidationErrorExpired:
				return nil,fmt.Errorf("[parseToken]过期的token")
			default:
				break
			}
			break
		default:
			break
		}
		return nil,fmt.Errorf("[parseToken]不合法的token:",err.Error())
	}
	claims,ok:=token.Claims.(jwt.MapClaims)
	if !ok||!token.Valid{
		return nil,fmt.Errorf("[parserToken]不合法的token")
	}
	return mapClaimToJwClaim(claims),nil
}
func mapClaimToJwClaim(claims jwt.MapClaims)*jwt.StandardClaims{
	jC:=&jwt.StandardClaims{
		Subject:claims["sub"].(string),
	}
	return jC
}