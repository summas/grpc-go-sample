package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"github.com/simplesteph/grpc-go-course/mysql_sample/mysqlpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Aritcle(ctx context.Context, req *mysqlpb.ArticlesRequest) (*mysqlpb.ArticlesResponse, error) {
	fmt.Printf("Aritcle function was invoked with %v\n", req)
	category := req.GetCategory().GetCategory()
	err, title := getArticleTitle(category)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功test")
	}
	result := "category is " + title
	res := &mysqlpb.ArticlesResponse{
		Result: result,
	}

	return res, nil
}

func getArticleTitle(category_id string) (err error, title string) {
	db, err := sql.Open("mysql", "msuser:passwd@/knowledgeApp_development")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功test")
	}
	category_id_re := category_id + " LIMIT 1"
	testQuery := "SELECT * FROM articles WHERE id = " //category_id = "
	testQuery += category_id_re
	rows, err := db.Query(testQuery) //?", 1).Scan(&name)"SELECT * FROM articles WHERE id =
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功test")
	}
	testTitle := "testTitle"

	columns, err := rows.Columns() // カラム名を取得
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	//  rows.Scan は引数に `[]interface{}`が必要.

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	aritcles := make(map[int]map[string]string)
	for rows.Next() {

		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		i := 1

		aritcle := map[string]string{}
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			// fmt.Println(columns[i], ": ", value)
			aritcle[columns[i]] = value
		}
		aritcles[i] = make(map[string]string)
		aritcles[i] = aritcle
		i += 1
		fmt.Println("-----------------------------------")
	}

	defer db.Close()

	for _, aritcle := range aritcles {

		// var value string
		for index, col := range aritcle {
			// Here we can check if the value is nil (NULL value)
			// if col == nil {
			// 	value = "NULL"
			// } else {
			// 	value = string(col)
			// }
			fmt.Println(index, ": ", col)
			// aritcle[columns[i]] = value
		}

		fmt.Println("mapアウトプット-----------------------------------")
		fmt.Println("titleテスト:", aritcle["title"])
	}

	return err, testTitle

}

func main() {
	fmt.Println("Hello world Aritcle")

	// dbConnectTest()

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// opts := []grpc.ServerOption{}
	// tls := false
	// if tls {
	// 	certFile := "ssl/server.crt"
	// 	keyFile := "ssl/server.pem"
	// 	creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
	// 	if sslErr != nil {
	// 		log.Fatalf("Failed loading certificates: %v", sslErr)
	// 		return
	// 	}
	// 	opts = append(opts, grpc.Creds(creds))
	// }

	s := grpc.NewServer() //opts...
	mysqlpb.RegisterAritcleServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// _, err := sqlConnect()

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
	}
}

func dbConnectTest() (err error) {

	db, err := sql.Open("mysql", "msuser:passwd@/knowledgeApp_development")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功sss")
	}

	rows, err := db.Query("SELECT * FROM articles WHERE category_id = 1") //SELECT title FROM knowledgeApp_development.articles WHERE id = 1
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功sss")
	}

	columns, err := rows.Columns() // カラム名を取得
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	//  rows.Scan は引数に `[]interface{}`が必要.

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}

	defer db.Close()

	return err
}

// SQLConnect DB接続
// func sqlConnect() (database *gorm.DB, err error) {
// 	DBMS := "mysql"
// 	USER := "go_example"
// 	PASS := "12345!"
// 	PROTOCOL := "tcp(localhost:3306)"
// 	DBNAME := "go_example"

// 	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
// 	return gorm.Open(DBMS, CONNECT)
// }

// // db接続
// db, err := sqlConnect()
// if err != nil {
// 	panic(err.Error())
// }
// defer db.Close()
// addUserData(db)

// result := []*Users{}
// error := db.Find(&result).Error
// if error != nil || len(result) == 0 {
// 	return
// }
// for _, user := range result {
// 	fmt.Println(user.Name)
// }
// }
