package service

import (
	"log"
	"server/dao"

	"github.com/dgrijalva/jwt-go"
	"server/model"
	flashkill "server/rpc/kitex_gen/FlashKill"
	"server/utils"
	"time"
)

func Register(s *flashkill.Seller, b *flashkill.Buyer) (err error) {
	if b.Name != "" {
		temp, passwordErr := utils.Crypto(b.Password)
		b.Password = temp
		if passwordErr != nil {
			log.Fatal(err)
			return err
		}
		if err = dao.DB.Create(&b).Error; err != nil {
			return err
		}
		dao.RDB.Set("Buyername", b.Name, 0)
		dao.RDB.Set("BuyerPassword", b.Password, 0)
	} else if s.Name != "" {
		temp, passwordErr := utils.Crypto(s.Password)
		s.Password = temp
		if passwordErr != nil {
			log.Fatal(err)
			return err
		}
		if err = dao.DB.Create(&s).Error; err != nil {
			return err
		}
		dao.RDB.Set("Sellername", s.Name, 0)
		dao.RDB.Set("SellerPassword", s.Password, 0)
	}
	return nil
}
func Login(s model.Sellers, b model.Buyers) {
	if b.Name != "" {
	} else if s.Name != "" {
	}
}
func SetToken(s *flashkill.Seller, b *flashkill.Buyer) (str string, err error) {
	var cnt string = time.Now().Format("2006-01-02 15:04:05")
	key := []byte(cnt)
	if b.Name != "" {
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"name":   b.Name,
			"passwd": b.Password,
		})
		str, err = claims.SignedString(key)
	} else if s.Name != "" {
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"name":   s.Name,
			"passwd": b.Password,
		})
		str, err = claims.SignedString(key)
	}
	return
}
func WriteToken(b *flashkill.Buyer, s *flashkill.Seller, str string) (err error) {
	if b.Name != "" {
		res := dao.DB.Model(&b).Where("name = ?", b.Name).Update("token", str)
		if res.Error != nil {
			return res.Error
		}
		dao.RDB.Set("BuyerToken", b.Token, 0)
	} else if b.Name != "" {
		res := dao.DB.Model(&s).Where("name = ?", s.Name).Update("token", str)
		if res.Error != nil {
			return res.Error
		}
		dao.RDB.Set("SellerToken", s.Token, 0)
	}
	return nil
}
