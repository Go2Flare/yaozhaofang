package model

import (
	"log"
	"testing"
)
func Init(){
	InitConfig()
	InitDb()
}

func TestCheckMobile(t *testing.T){
	err := CheckMobile("13790887214")
	if err != nil {
		log.Fatalf("CheckMobile err : %v", err)
	}
}
