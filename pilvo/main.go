/*This Exercise is done to fullfill the requirement of PILVO SDE Assignment
 */
package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

//A structure which represents a record
type Phonebook struct {
	Name    string
	Email   string
	Phone   int
	Address string
}

//Using the Phonebook struct as inhouse DB, since DB connection would be costly at this point of time
type DB []Phonebook

var db = DB{}

//using a map to maintain the uniqness of the records with email as key
var ContactMap map[string]Phonebook

/*Function to create an entry with email as unique parameter
Parameters taken: user details as a struct
*/
func Create(s Phonebook) string {
	_, ok := ContactMap[s.Email]
	if !ok {
		ContactMap[s.Email] = s
		db = append(db, s)
		return "contact created with name " + s.Name
	}
	return "contact already present please use update service"
}

/*Function to find an entry based on the name of the contact
Parameters taken: name
*/
func Findname(name string) ([]Phonebook, bool) {
	var result_array []Phonebook
	for _, r := range db {
		if strings.Contains(r.Name, name) {
			fmt.Println("found: ", r.Name)
			result_array = append(result_array, r)
		}
	}
	if len(result_array) == 0 {
		fmt.Println("Not found")
		return nil, false
	}
	return result_array, true
}

/*Function to find an entry based on the email id
Parameters taken: email
*/
func Findemail(email string) (int, bool) {
	for i, r := range db {
		if strings.EqualFold(r.Email, email) {
			fmt.Println("found: ", r.Name)
			return i, true
		}
	}
	fmt.Println("Not found")
	return 0, false
}

/*Function to update an entry based on the email id
Parameters taken: userdetails as a struct
*/
func Update(s Phonebook) string {
	if i, found := Findemail(s.Email); found {
		db[i] = s
		ContactMap[s.Email] = s
		fmt.Println("updated contact is: ", s)
		return "contact updated with name " + s.Name
	} else {
		return "contact was not found in the phone book"
	}
}

/*Function to delete an entry based on the email id
Parameters taken: email
*/
func Delete(email string) string {
	if i, found := Findemail(email); found {
		copy(db[i:], db[i+1:])
		db = db[0 : len(db)-1]
		delete(ContactMap, email)
		return "contact deleted successfully"
	}
	return "contact was not found"
}

func main() {
	e := echo.New()
	ContactMap = make(map[string]Phonebook)
	e.POST("/login", authControl)
	//	e.GET("/users/:id", getUser)
	e.POST("/createcontact", createContact)
	e.POST("/findcontact", findcontact)
	e.POST("/updatecontact", updatecontact)
	e.POST("/deletecontact", deletecontact)
	//HTML template render to send an html response
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("*.html")),
	}
	e.Renderer = renderer
	//Specifying the Css, in this example internal Css is used
	e.Static("/static", "*.css")
	/*These are the Get routers which will return th HTML files with no data*/
	e.GET("/login", func(c echo.Context) error {
		return c.Render(http.StatusOK, "login.html", nil)
	}).Name = "login"
	e.GET("/create", func(c echo.Context) error {
		return c.Render(http.StatusOK, "create.html", nil)
	}).Name = "create"
	e.GET("/search", func(c echo.Context) error {
		return c.Render(http.StatusOK, "search.html", nil)
	}).Name = "search"
	e.GET("/update", func(c echo.Context) error {
		return c.Render(http.StatusOK, "update.html", nil)
	}).Name = "update"
	e.GET("/delete", func(c echo.Context) error {
		return c.Render(http.StatusOK, "delete.html", nil)
	}).Name = "delete"
	e.Logger.Fatal(e.Start(":8080"))
}

/*router to get the user details
Currently the authentication is hardcoded for admin admin
*/
func authControl(c echo.Context) error {
	if strings.EqualFold(c.FormValue("username"), "admin") && strings.EqualFold(c.FormValue("password"), "admin") {
		fmt.Println("Auth successfull")
		return c.Render(http.StatusOK, "dashboard.html", nil)
	} else {
		fmt.Println("username password mismatch")
		return c.Render(http.StatusOK, "login.html", map[string]interface{}{"message": "Authentication Failed Please Try Again"})
	}
	fmt.Println("Login Process")
	return nil
}

/*function to create the contact user
Create an user only if it does not exist
*/
func createContact(c echo.Context) error {
	var newContact Phonebook
	// Get name and email
	newContact.Name = c.FormValue("name")
	newContact.Email = c.FormValue("email")
	newContact.Phone, _ = strconv.Atoi(c.FormValue("phone"))
	newContact.Address = c.FormValue("address")
	message := Create(newContact)
	for _, r := range db {
		fmt.Println(r.Name, r.Email)
	}
	return c.Render(http.StatusOK, "create.html", map[string]interface{}{"message": message})
}

/*function to find the contact user
Find a contact either by name or by email, combination of both is not implemented
*/
func findcontact(c echo.Context) error {
	//	email := c.FormValue("email")
	if len(c.FormValue("email")) != 0 {
		fmt.Println("inside find email")
		var contact []Phonebook
		result_single, ok := ContactMap[c.FormValue("email")]
		if ok {
			fmt.Println(result_single)
			contact = append(contact, result_single)
			return c.Render(http.StatusOK, "search.html", map[string]interface{}{"message": contact})
		} else {
			return c.Render(http.StatusNotFound, "search.html", nil)
		}
	}
	if len(c.FormValue("name")) != 0 {
		result, _ := Findname(c.FormValue("name"))
		return c.Render(http.StatusOK, "search.html", map[string]interface{}{"message": result})
	}
	return c.Render(http.StatusNotFound, "search.html", nil)

}

/*function to find and update the contact user
Update a contact only if the email Id entered exists or else show error
*/
func updatecontact(c echo.Context) error {
	var updateContact Phonebook
	// Get name and email
	updateContact.Name = c.FormValue("name")
	updateContact.Email = c.FormValue("email")
	updateContact.Phone, _ = strconv.Atoi(c.FormValue("phone"))
	updateContact.Address = c.FormValue("address")
	message := Update(updateContact)
	for _, r := range db {
		fmt.Println("------------")
		fmt.Println(r.Name, r.Email)
	}
	return c.Render(http.StatusOK, "update.html", map[string]interface{}{"message": message})
}

/*function to delete contact user
Check for the email id and if it exists delete the same
*/
func deletecontact(c echo.Context) error {
	message := Delete(c.FormValue("email"))
	for _, r := range db {
		fmt.Println("------delete------")
		fmt.Println(r.Name, r.Email)
	}
	return c.Render(http.StatusOK, "delete.html", map[string]interface{}{"message": message})
}
