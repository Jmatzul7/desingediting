package main

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"path/filepath"
	"github.com/nu7hatch/gouuid"
)
var port =1433
var user ="NuevoUsuario"
var password ="#Test.00.Test#"
var database="ReactCRUD"
var db *sql.DB

type Book struct{
	Id string
	Title string
	Description string
}
func getAllBooks(c *gin.Context){

	var Books []Book
	db, err:=getDbConection()
	if err!=nil{
		val:=fmt.Sprintf("%s", err)
		println("Error getting a db conection: ", val)
	}else{
		rows, err:=db.Query("select * from ReactCRUD.dbo.Books")
		if err!=nil{
			println("Error in the query", err)
		}else{
			for rows.Next(){
				var book Book
				err= rows.Scan(&book.Id, &book.Title, &book.Description)
				if err!=nil{
					println("Error scanning")
				}
				Books = append(Books, book)
			}
		}
	}
	c.JSON(200, Books)

}

func getDbConection() (*sql.DB, error){
	query := url.Values{}
	query.Add("app name", database)

	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(user, password),
		Host:     fmt.Sprintf("%s:%d", "localhost", port),
		// Path:  instance, // if connecting to an instance instead of a port
		RawQuery: query.Encode(),
	}
	var err error
	db, err = sql.Open("sqlserver", u.String())
	if err!=nil{
		println("Error connecting to this db: ",err)
		return nil, err
	}else{
		return db, nil
	}
}
func addNewBook(c *gin.Context){
	u, err := uuid.NewV4()
	fmt.Println("UUID: ", u)
	var book Book
	err:= c.ShouldBind(&book)

	if err==nil{
		fmt.Println("Book0: ", book)
		db, err:=getDbConection()
		if err==nil{
			cSQL := "INSERT INTO ReactCRUD.dbo.Books (Title,Description) VALUES(@p1,@p2)"
			if _, err := db.Exec(cSQL,book.Title,book.Description); err != nil {
				fmt.Println("Error inserting the data: ", err)
			}
		}
	}else{
		fmt.Println("Book: ", book)
	}
}

func deleteBook(c *gin.Context){
	Id, Success:= c.GetQuery("Id")
	db, err:=getDbConection()
	if err==nil{
		if Success{
			cSQL := "DELETE FROM ReactCRUD.dbo.Books WHERE Id=@p1"
			if _, err := db.Exec(cSQL, Id); err != nil {
				fmt.Println("Error deleting record: ", err)
			}
		}else{
			fmt.Println("Error: ")
		}
	}
}
func saveFileHandler(c *gin.Context) {
	file, err := c.FormFile("file")

	// The file cannot be received.
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	// Retrieve file information
	extension := filepath.Ext(file.Filename)
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := uuid.New().String() + extension

	// The file is received, so let's save it
	if err := c.SaveUploadedFile(file, "C:/Users/HA975MF/Desktop/Go/" + newFileName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	// File saved successfully. Return proper result
	c.JSON(http.StatusOK, gin.H{
		"message": "Your file has been successfully uploaded.",
	})
}

func FileDownload(c *gin.Context){
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "0037cf40-5a07-4a12-9b3d-bfb37db77a36.jpg"))//fmt.Sprintf("attachment; filename=%s", filename) Downloaded file renamed
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File("C:/Users/HA975MF/Desktop/Go/0037cf40-5a07-4a12-9b3d-bfb37db77a36.jpg")
}
func main(){
	server := gin.Default()
	server.Use(cors.Default())
	server.GET("/getAllBooks", getAllBooks)
	server.POST("/addNewBook", addNewBook)
	server.DELETE("/deleteBook", deleteBook)
	server.POST("/saveFileHandler",saveFileHandler)
	server.GET("/FileDownload",FileDownload)

	server.Run(":8081")
}