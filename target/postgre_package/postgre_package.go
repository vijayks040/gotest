/*package consisting of all the postgre db operations*/

package postgre_package

import (
	"database/sql"
	"fmt"
	"sbgoclient/target/mongo_package"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "test"
)

var database *sql.DB
var err error

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	database, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//	defer db.Close()

	err = database.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}

/*Function to insert the Product*/
func InsertProduct(product mongo_package.Product) error {
	fmt.Println("package postgre", product.ProductName)
	sqlStatement := fmt.Sprintf(`INSERT INTO public.product ("Productname", "currentPrice", "productId", "currency") VALUES ('%s', '%f', '%s', '%s')`, product.ProductName,
		product.Current_price, product.ProductId, product.Currency)
	fmt.Println("query formed is ", sqlStatement)
	_, err = database.Exec(sqlStatement)
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

/*Function to find the Produc based on product Idt*/

func FindProduct(productId int) (mongo_package.Product, error) {
	fmt.Println("FindProduct()")
	var Result_product mongo_package.Product
	query_string := fmt.Sprintf(`SELECT "Productname", "currentPrice", "productId", "currency" FROM public.product where "productId"=%d`, productId)
	rows := database.QueryRow(query_string)
	err = rows.Scan(&Result_product.ProductName, &Result_product.Current_price, &Result_product.ProductId, &Result_product.Currency)
	if err != nil {
		// handle this error
		fmt.Println("error happend")
		fmt.Println(err)
		return Result_product, err
	}
	if Result_product.ProductName == "" {
		fmt.Println("product not found")
		return Result_product, err
	}
	fmt.Println(len(Result_product.ProductId))
	fmt.Println("product found", Result_product)
	// get any error encountered during iteration
	return Result_product, nil
}

/*Function to update the Product based on product Id*/
func UpdateProduct(product mongo_package.Product) error {
	sqlStatement := `
UPDATE public.product
SET "Productname" = $1, "currentPrice" = $2, "productId"=$3, "currency"=$4
WHERE "productId" = $3;`
	fmt.Println("query formed is ", sqlStatement)
	_, err = database.Exec(sqlStatement, product.ProductName, product.Current_price, product.ProductId, product.Currency)
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

/*Function to Delete the Product based on product Id*/
func DeleteProduct(product_id int) error {
	sqlStatement := `
DELETE FROM public.product
WHERE "productId" = $1;`
	fmt.Println("query formed is ", sqlStatement)
	_, err = database.Exec(sqlStatement, product_id)
	if err != nil {
		panic(err)
		return err
	}
	return nil
}
