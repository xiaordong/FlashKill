package service

import (
	"log"
	"server/model"
)

func Register(s model.Sellers, b model.Buyers) error {
	if b.Name != "" {
		err := b.New()
		if err != nil {
			log.Fatal(err)
			return err
		}
	} else if s.Name != "" {
		err := s.New()
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}
func Login(s model.Sellers, b model.Buyers) {
	if b.Name != "" {
	} else if s.Name != "" {
	}
}
