package services

import (
	"database/sql"
	"zentools-mysql/models"
)

func (s *Service) GetCustomers(db *sql.DB) ([]models.Customer, error) {
	sqlQuery := "select customerNumber, customerName, phone, creditLimit from customers limit 10"
	result, err := db.Query(sqlQuery)
	checkError(err)
	defer result.Close()

	var customers []models.Customer
	for result.Next() {
		var c models.Customer

		err := result.Scan(&c.CustomerNumber, &c.CustomerName, &c.Phone, &c.CreditLimit)

		checkError(err)

		customers = append(customers, c)
	}

	return customers, nil
}

func (s *Service) GetCustomerByNumber(db *sql.DB, custNo int) (models.Customer, error) {
	sqlQuery := "select customerNumber, customerName, phone, " +
		"creditLimit from customers where customerNumber = ?"

	var c models.Customer
	err := db.QueryRow(sqlQuery, custNo).Scan(&c.CustomerNumber, &c.CustomerName, &c.Phone, &c.CreditLimit)

	if err != nil {
		return c, err
	}

	return c, nil
}
