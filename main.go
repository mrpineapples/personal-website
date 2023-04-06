package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrpineapples/personal-website/middleware"
	"github.com/mrpineapples/personal-website/models"
	"github.com/mrpineapples/personal-website/routes"
	"github.com/mrpineapples/personal-website/utils"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "michaelsite_dev"
)

func main() {
	// sqlStatement := `
	// 	INSERT INTO users (age, email, first_name, last_name)
	// 	VALUES ($1, $2, $3, $4)
	// 	RETURNING id
	// `
	// id := 0
	// err = db.QueryRow(sqlStatement, 26, "michaelmiranda161@gmail.com", "Michael", "Miranda").Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("New record ID is:", id)

	// sqlStatement := `
	// 	UPDATE users
	// 	SET first_name = $2, last_name = $3
	// 	WHERE id = $1;
	// `
	// _, err = db.Exec(sqlStatement, 2, "Adam", "Jameson")
	// if err != nil {
	// 	panic(err)
	// }

	// sqlStatement := `
	// 	DELETE FROM users
	// 	WHERE id = $1;
	// `
	// _, err = db.Exec(sqlStatement, 2)
	// if err != nil {
	// 	panic(err)
	// }

	// sqlStatement := `
	// 	UPDATE users
	// 	SET first_name = $2, last_name = $3
	// 	WHERE id = $1;
	// `

	// res, err := db.Exec(sqlStatement, 9, "NewFirst", "NewLast")
	// if err != nil {
	// 	panic(err)
	// }
	// count, err := res.RowsAffected()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(count)

	// sqlStatement := `
	// 	UPDATE users
	// 	SET first_name = $2, last_name = $3, email = $4
	// 	WHERE id = $1
	// 	RETURNING id, email;
	// `
	// var email string
	// var id int
	// err = db.QueryRow(sqlStatement, 9, "First", "Last", "test@test.com").Scan(&id, &email)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("id", id)
	// fmt.Println("email", email)

	// type User struct {
	// 	ID        int
	// 	Age       int
	// 	FirstName string
	// 	LastName  string
	// 	Email     string
	// }

	// sqlStatement := `SELECT * FROM users WHERE id=$1;`
	// var user User
	// row := db.QueryRow(sqlStatement, 1)
	// err = row.Scan(&user.ID, &user.Age, &user.FirstName, &user.LastName, &user.Email)
	// switch err {
	// case sql.ErrNoRows:
	// 	fmt.Println("No rows were returned!")
	// 	return
	// case nil:
	// 	fmt.Println(user)
	// default:
	// 	panic(err)
	// }

	// rows, err := pool.Query(ctx, "SELECT id, first_name FROM users LIMIT $1", 4)
	// if err != nil {
	// 	// handle this error better than this
	// 	panic(err)
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	var id int
	// 	var firstName string
	// 	err = rows.Scan(&id, &firstName)
	// 	if err != nil {
	// 		// handle this error
	// 		panic(err)
	// 	}
	// 	fmt.Println(id, firstName)
	// }
	// // get any error encountered during iteration
	// err = rows.Err()
	// if err != nil {
	// 	panic(err)
	// }

	r := gin.Default()
	r.Use(middleware.MethodOverride(r))
	r.Static("/public", "./public")
	r.HTMLRender = utils.LoadTemplates()
	routes.InitializeRoutes(r)

	models.InitDB()
	defer models.Pool.Close()

	// start the server on port 8080
	if err := r.Run(":8080"); err != nil {
		// handle the error if the server fails to start
		panic(err)
	}
}
