package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := "postgresql://postgres.ctubnndjaexcdscmguqd:MarsyaAdinda123@aws-1-ap-southeast-2.pooler.supabase.com:6543/postgres?pgbouncer=true"

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error Connecting to Database:", err)
		panic(err)
	}

	fmt.Println("Database Connected Successfully!")
}