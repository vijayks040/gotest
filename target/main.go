package main

import (
	"fmt"
	"net/http"

	"strconv"

	"sbgoclient/target/mongo_package"
	"sbgoclient/target/postgre_package"

	"github.com/labstack/echo"
	//	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("Starting PRODUCT MANAGEMENT")
	e := echo.New()
	/*Incase of https use the following with certificates
	//	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	//e.Use(middleware.Recover())
	//e.Use(middleware.Logger())
	//e.Logger.Fatal(e.StartAutoTLS(":443"))
	*/
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusCreated, "Welcome to Product Management system\ncheck following API's \n"+
			"1)localhost:3000/createproduct POST input-type: FORM values\n"+
			"2)localhost:3000/updateproduct PUT input-type: JSON format\n"+
			"3)localhost:3000/findproduct POST input-type: FORM value\n"+
			"4)localhost:3000/deleteproduct POST input-type: FORM value\n"+
			"For form input fields contact @vijay")
	})
	e.POST("/createproduct", createProduct)
	e.POST("/findproduct", findProduct)
	e.PUT("/updateproduct", updateproduct)
	e.POST("/deleteproduct", deleteproduct)
	e.Logger.Fatal(e.Start(":3000"))
}

/*function to create the Product
Create a product only if it does not exist
*/
func createProduct(c echo.Context) error {
	var product_new mongo_package.Product
	product_new.ProductId = c.FormValue("productid")
	product_new.ProductName = c.FormValue("name")
	product_new.Description = c.FormValue("description")
	product_new.Current_price, _ = strconv.ParseFloat(c.FormValue("price"), 2)
	product_new.Currency = c.FormValue("currency")

	err := postgre_package.InsertProduct(mongo_package.Product{ProductId: product_new.ProductId, ProductName: product_new.ProductName, Currency: product_new.Currency, Current_price: product_new.Current_price})

	if err != nil {
		fmt.Println("error ", err)
	} else {
		mongo_package.Createproduct(product_new)
		fmt.Println("Product Added successfully with name " + c.FormValue("name"))
		return c.String(http.StatusCreated, "Product Added successfully with name "+c.FormValue("name"))
	}
	return err
}

/*function to find the product
 */
func findProduct(c echo.Context) error {
	fmt.Println("inside find")

	productid, err := strconv.Atoi(c.FormValue("productid"))
	if err != nil {
		fmt.Println("error ", err)
		return c.String(http.StatusNotFound, err.Error())
	}
	result_product, err := postgre_package.FindProduct(productid)

	if err != nil {
		fmt.Println("error ", err)
		return c.String(http.StatusNotFound, err.Error())
	}
	encodedJSON := []byte(fmt.Sprintf("%v", result_product))
	return c.JSONBlob(http.StatusOK, encodedJSON)
}

/*function to create the Product
update a product only if it does not exist
*/
func updateproduct(c echo.Context) error {
	var product_new mongo_package.Product
	if err := c.Bind(&product_new); err != nil {
		return err
	}

	err := postgre_package.UpdateProduct(mongo_package.Product{ProductId: product_new.ProductId, ProductName: product_new.ProductName, Currency: product_new.Currency, Current_price: product_new.Current_price})

	if err != nil {
		fmt.Println("error ", err)
	} else {
		mongo_package.Updateproduct(product_new)
		fmt.Println("Product Updated successfully with name " + c.FormValue("name"))
		return c.String(http.StatusCreated, "Product Updated successfully with name "+c.FormValue("name"))
	}
	return err
}

/*function to delete the product
 */
func deleteproduct(c echo.Context) error {
	fmt.Println("inside find")

	productid, err := strconv.Atoi(c.FormValue("productid"))
	if err != nil {
		fmt.Println("error ", err)
		return c.String(http.StatusNotFound, err.Error())
	}
	err = postgre_package.DeleteProduct(productid)
	if err != nil {
		fmt.Println("error ", err)
		return c.String(http.StatusNotFound, err.Error())
	} else {
		err = mongo_package.Deleteproduct(c.FormValue("productid"))
		if err != nil {
			return c.String(http.StatusNotFound, err.Error())
		}
	}
	return c.String(http.StatusAccepted, "Product Deleted successfully")
}
