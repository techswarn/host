package middleware

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"os"
	_ "github.com/lib/pq"
	"github.com/techswarn/host/models"
    "encoding/json"
)

func init(){
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("POSTGRES_URL"))

}

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConnection() *sql.DB {

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		fmt.Println(err)
		log.Fatal("error connecting to database")
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to postgresSQL successfully")
	return db
}

func CreateStocks(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	if r.Body == nil {
		log.Fatalln("Request object is empty")
	}
	err := json.NewDecoder( r.Body).Decode(&stock)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	insertId := insertDB(stock)
	
	res := &response{
		ID : insertId,
		Message: "Inserted stock successfully",
	}

	json.NewEncoder(w).Encode(res)

}

func GetAllstock(w http.ResponseWriter, r *http.Request) {

	var stock models.Stock
	
	stocks, err := getAllstock(stock)

	if err != nil {
		log.Fatalf("GetAllStock: %v\n", err)
	}

	json.NewEncoder(w).Encode(stocks)
}




//SQL methods

//Insert stock
func insertDB(stock models.Stock) int64{
	// db := CreateConnection()
	// fmt.Printf("%v", stock)

    // result, err := db.Exec("INSERT INTO stocks (name, price, company ) VALUES (?, ?, ?)", stock.Name, stock.Price. stock.Company)
    // if err != nil {
    //     log.Fatalf("Error while adding stock to db: %v \n", err)
    // }

	// id, err := result.LastInsertId()
	// if err != nil {
    //     log.Fatalf("Error while fetching ID: %v \n", err)
    // }
    // return id

	db := createConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO stocks (name, price, company) VALUES ($1, $2, $3) RETURNING stockid`

	var id int64

	err := db.QueryRow(sqlStatement, stock.Name, stock.Price, stock.Company).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)
	return id
}

//getAllstock

func getAllstock(stock models.Stock) ([]models.Stock, error) {
	db := createConnection()
	defer db.Close()
	//Define a slice of stock
	var stocks []models.Stock
	sqlStatement := `SELECT * FROM stocks`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Error while running GetAllStock statement %v \n", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(stock.StockID, stock.Name, stock.Price, stock.Company)

		stocks = append(stocks, stock)
	}
	return stocks, nil
}