package domain

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rokafela/udemy-banking/errs"
	"github.com/rokafela/udemy-banking/logger"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error
	customers := make([]Customer, 0)

	statusCond := ""
	if status == "active" {
		statusCond = "WHERE status = 1"
	} else if status == "inactive" {
		statusCond = "WHERE status = 0"
	}

	findAllSql := "SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customers " + statusCond
	err = d.client.Select(&customers, findAllSql)

	// rows, err := d.client.Query(findAllSql)
	if err != nil {
		logger.Error("Error querying customers table" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	// err = sqlx.StructScan(rows, &customers)
	// if err != nil {
	// 	logger.Error("Error scanning customers" + err.Error())
	// 	return nil, errs.NewUnexpectedError("unexpected database error")
	// }

	return customers, nil
}

func (d CustomerRepositoryDb) FindById(id string) (*Customer, *errs.AppError) {
	var c Customer

	findByIdSql := "SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customers WHERE customer_id = ?;"
	// row := d.client.QueryRow(findByIdSql, id)
	// err := row.Scan(&c.Id, &c.Name, &c.DateOfBirth, &c.City, &c.Zipcode, &c.Status)
	err := d.client.Get(&c, findByIdSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error scanning customer" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}
