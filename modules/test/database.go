package test

import (
	dbmock "github.com/selvatico/go-mocket"
)

//LoadDBTest LoadDBTest
func LoadDBTest() {
	setupDB(getTestData())
}

func setupDB(commonReply map[string][]map[string]interface{}, commonQuery map[string]string) {
	// Important: Use database files here (snake_case) and not struct variables (CamelCase)
	// eg: first_name, last_name, date_of_birth NOT FirstName, LastName or DateOfBirth
	fakers := make([]*dbmock.FakeResponse, 0)
	for index, value := range commonReply {
		faker := dbmock.FakeResponse{
			Pattern:  commonQuery[index],
			Response: value,
			Once:     false, // could be done via chaining .OneTime()
		}
		fakers = append(fakers, &faker)
	}
	dbmock.Catcher.Attach(fakers)
}

func getTestData() (map[string][]map[string]interface{}, map[string]string) {
	dummyTransaction := []map[string]interface{}{
		{
			"id":          1,
			"product_id":  "1",
			"customer_id": "081234567890",
		},
	}
	dummyBiller := []map[string]interface{}{
		{
			"id":          1,
			"label":       "fbs_test_1",
			"description": "Biller Test",
		},
	}

	commonReply := map[string][]map[string]interface{}{
		"transaction": dummyTransaction,
		"biller":      dummyBiller,
	}

	// Define database mock query
	commonQuery := map[string]string{
		"transaction": `SELECT * FROM "transactions"`,
		"biller":      `SELECT * FROM "billers"`,
	}

	return commonReply, commonQuery
}
