package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"server/dao"
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
		if res := dao.DB.Where("name = ?", b.Name).Find(&b); res.Error != nil || res.RowsAffected != 0 {
			return errors.New("exist buyer error")
		}
		if err = dao.DB.Create(&b).Error; err != nil {
			return err
		}
		dao.RDB.Set("Buyername", b.Name, 0)
	} else if s.Name != "" {
		temp, passwordErr := utils.Crypto(s.Password)
		s.Password = temp
		if passwordErr != nil {
			log.Fatal(err)
			return err
		}
		if res := dao.DB.Where("name = ?", s.Name).Find(&s); res.Error != nil || res.RowsAffected != 0 {
			return errors.New("exist seller error")
		}
		if err = dao.DB.Create(&s).Error; err != nil {
			return err
		}
		dao.RDB.Set("SellerName", s.Name, 0)
	}
	return nil
}
func SetToken(s *flashkill.Seller, b *flashkill.Buyer) (str string, err error) {
	var cnt = time.Now().Format("2006-01-02 15:04:05")
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
	} else if s.Name != "" {
		res := dao.DB.Model(&s).Where("name = ?", s.Name).Update("token", str)
		if res.Error != nil {
			return res.Error
		}
		dao.RDB.Set("SellerToken", s.Token, 0)
	}
	return nil
}
func Login(s *flashkill.Seller, b *flashkill.Buyer) (err error) {
	if b.Name != "" {
		var token string
		var test model.Buyers
		pwd, _ := utils.Crypto(b.Password)
		res := dao.DB.Model(&model.Buyers{}).Where("name = ? and password = ?", b.Name, pwd).Find(&test)
		if res.Error != nil || res.RowsAffected == 0 {
			return errors.New("buyer not exist")
		}
		token, err = dao.RDB.Get("BuyerToken").Result()
		if err != nil {
			return err
		}
		if token != test.Token {
			return errors.New("buyer token error")
		}
	} else if s.Name != "" {
		var token string
		var test model.Sellers
		pwd, _ := utils.Crypto(s.Password)
		res := dao.DB.Model(&model.Sellers{}).Where("name = ? and password = ?", s.Name, pwd).First(&test)
		if res.Error != nil || res.RowsAffected == 0 {
			return errors.New("seller not exist:" + err.Error())
		}
		token, err = dao.RDB.Get("SellerToken").Result()
		if err != nil {
			return err
		}
		if token != test.Token {
			return errors.New("seller token error")
		}
	}
	return nil
}
