package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	errs "github.com/rokafela/udemy-banking/helpers"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	statusCond := ""
	if status == "active" {
		statusCond = "WHERE status = 1"
	} else if status == "inactive" {
		statusCond = "WHERE status = 0"
	}

	findAllSql := "SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customers " + statusCond
	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("Error querying customers table" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.DateOfBirth, &c.City, &c.Zipcode, &c.Status)
		if err != nil {
			log.Println("Error scanning customers" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func (d CustomerRepositoryDb) FindById(id string) (*Customer, *errs.AppError) {
	findByIdSql := "SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customers WHERE customer_id = ?;"
	row := d.client.QueryRow(findByIdSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.DateOfBirth, &c.City, &c.Zipcode, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			log.Println("Error scanning customer" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
