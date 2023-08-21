// @title           account_management_api
// @version         1.0
// @description     This is a REST API which can create an account and verify an account.
// @termsOfService  http://swagger.io/terms/

// @contact.name   	Leung Yan Tung
// @contact.url    	https://github.com/ushio0107
// @contact.email  	leungyantung0107@gmail.com

// @license.name  	Apache 2.0
// @license.url   	http://www.apache.org/licenses/LICENSE-2.0.html
package main

import (
	"log"
	"user_api/api"
	"user_api/db"

	_ "user_api/docs"
)

func main() {
	coll, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	api.NewApi(coll).Run()
}
