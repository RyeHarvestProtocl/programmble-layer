package testutils

import (
	"fmt"
	"log"

	"github.com/RyeHarvestProtocol/programmable-layer/models"
)

func DropTablesAndReCreate(db *models.DB) error {
	tables := []string{
		"btc_user_deposits",
		"btc_user_withdraws",
		"scores",
		"invite_codes",
		"invite_code_rels",
		"stake_pools",
		"user_infos",
		"btc_prices",
		"rye_orders",
		"rye_utxos",
	}

	for _, table := range tables {
		err := db.Pg.Exec("DROP TABLE " + table).Error

		if err != nil {
			log.Printf("Failed to drop table %s: %v", table, err)
			return err
		} else {
			log.Printf("Table %s dropped successfully", table)
		}
	}

	db.Pg.AutoMigrate(&models.BTCNetworkInfo{})
	log.Printf("auto migrate success\n")
	return nil
}

func DeleteTables(db *models.DB) error {
	// List all the tables you want to delete
	tables := []interface{}{
		&models.BTCNetworkInfo{},
	}

	for _, model := range tables {
		// Using Delete method with a model type and no where clause
		// Passing a slice of model ensures that Delete affects all records
		err := db.Pg.Unscoped().Where("1 = 1").Delete(model).Error

		if err != nil {
			log.Printf("Failed to delete contents of table %T: %v", model, err)
			return err
		} else {
			log.Printf("Contents of table %T deleted successfully", model)
		}
	}

	return nil
}

func InsertData(db *models.DB) error {
	// Insert into user_infos
	twitterUserIds := []string{"1", "2", "3"} // Example user IDs

	// Start the SQL statement
	sqlStr := "INSERT INTO user_infos (twitter_user_id) VALUES "

	// Create a placeholder for each user ID
	vals := []interface{}{}
	for _, id := range twitterUserIds {
		sqlStr += "(?),"
		vals = append(vals, id)
	}

	// Trim the last comma
	sqlStr = sqlStr[0 : len(sqlStr)-1]

	// Execute the SQL statement
	err := db.Pg.Exec(sqlStr, vals...).Error
	if err != nil {
		return fmt.Errorf("failed to insert multiple user_ids into user_infos: %v", err)
	}

	// Insert into invite_codes
	inviteCodes := []struct {
		code   string
		userID int
	}{
		{"000000", 1},
		{"111111", 1},
		{"222222", 1},
		{"333333", 1},
		{"444444", 1},
		{"555555", 1},
		{"546546", 1},
		{"777777", 1},
		{"888888", 1},
		{"999999", 1},
		{"kG9FTy", 1},
		{"g20E2I", 1},
		{"8QDWHJ", 1},
		{"i9MDkR", 1},
		{"WhrMrI", 1},
		{"sUpC6s", 1},
		{"cPrFBQ", 1},
		{"NraL3S", 1},
		{"8JWCHA", 1},
		{"I4ZOpr", 1},
		{"vq57hV", 1},
		{"W6LVn9", 1},
		{"VysTxQ", 1},
		{"ogLcex", 1},
		{"4cRGJm", 1},
		{"dxLbxc", 1},
		{"Zl1dca", 1},
		{"Lzpq4Y", 1},
		{"cXb6DX", 1},
		{"zufc6X", 1},
		{"cJtPxs", 1},
		{"r7KFpc", 1},
		{"pA9Ia8", 1},
		{"uawC1X", 1},
		{"Rx0DjL", 1},
		{"rqrUt7", 1},
		{"0sYyTD", 1},
		{"UKxEw9", 1},
		{"FeI2n4", 1},
		{"DquhVf", 1},
		{"NaUWTu", 1},
		{"kN8haO", 1},
		{"kg58iy", 1},
		{"4QpLHs", 1},
		{"1TnHzU", 1},
		{"yDgdBq", 1},
		{"wSETZb", 1},
		{"TrCTUJ", 1},
		{"6uLV8X", 1},
		{"5E0E9T", 1},
		{"sxM5lH", 1},
		{"8R5MyU", 1},
		{"CquVEb", 1},
		{"g1SCe7", 1},
		{"i69iZl", 1},
		{"ZMcApv", 1},
		{"MpLPkx", 1},
		{"ZnDmXk", 1},
		{"6HGu7v", 1},
	}

	for _, v := range inviteCodes {
		err := db.Pg.Exec("INSERT INTO invite_codes (invite_code, inviter_user_id) VALUES (?, ?)", v.code, v.userID).Error
		if err != nil {
			return fmt.Errorf("failed to insert into invite_codes: %v", err)
		}
	}

	// type StakePool struct {
	// 	BaseModel
	// 	Name           string    `json:"name"`
	// 	HoldingEndTime time.Time `json:"holding_end_time"`
	// 	Boost          uint      `json:"boost"` // *1.1 : 110; *1.2: 120
	// }
	// insert to stakePools
	stakePools := []struct {
		name                     string
		holdingEndTime           string
		balanceToDailyPointRatio uint
	}{
		{"pool1", "2021-09-01", 10000},
	}

	for _, v := range stakePools {
		err := db.Pg.Exec("INSERT INTO stake_pools (name, holding_end_time, balance_to_daily_point_ratio) VALUES (?, ?, ?)", v.name, v.holdingEndTime, v.balanceToDailyPointRatio).Error
		if err != nil {
			return fmt.Errorf("failed to insert into stake_pools: %v", err)
		}
	}

	fmt.Println("Data inserted successfully")
	return nil
}
