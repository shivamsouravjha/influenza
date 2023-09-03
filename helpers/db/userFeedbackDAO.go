package db

import (
	"context"
	"errors"
	"fmt"

	db "github.com/shivamsouravjha/influenza/helpers/sql"
)

func FetchFeedback(ctx context.Context) error {
	var matchingSearch int
	sqlString := fmt.Sprintf("SELECT * from user where id=  \"%v\" ")
	err := db.Dbmap.SelectOne(&matchingSearch, sqlString)
	if err != nil || matchingSearch <= 0 {
		return errors.New("no user found")
	}
	// images := strings.Join(userProductData.ProductImages, ",")
	// createdProduct, err := db.Dbmap.Exec("INSERT INTO products (product_name, product_description,  product_price, product_images,created_at,updated_at,compressed_product_images) VALUES(?,?,?,?,?,?,?) ", userProductData.ProductName, userProductData.ProductDescription, userProductData.ProductPrice, images, time.Now(), time.Now(), "")
	// if err != nil {
	// 	return err
	// }
	// productID, _ := createdProduct.LastInsertId()
	return nil
}
