package repository

import (
	"linebot/internal/model"

	"gorm.io/gorm"
)

//新建營地
func CreateNewCamp(camp *model.Camp) error {
	return BeginTranscation(db, func(tx *gorm.DB) error {
		return tx.Create(&camp).Error
	})

}

func QueryCampByCampName(campName string) (model.Camp, error) {
	var camp model.Camp
	err := db.Limit(1).Where("Name=?", campName).Find(&camp).Error
	return camp, err

}
func InsertTest() {
	// if _, err := db.Exec("CREATE TABLE IF NOT EXISTS ticks (tick timestamp)"); err != nil {

	// 	fmt.Sprintf("Error creating database table: %q", err)
	// 	return
	// }

	// if _, err := db.Exec("INSERT INTO ticks VALUES (now())"); err != nil {
	// 	fmt.Sprintf("Error incrementing tick: %q", err)
	// 	return
	// }

	// rows, err := db.Query("SELECT tick FROM ticks")
	// if err != nil {
	// 	fmt.Sprintf("Error reading ticks: %q", err)
	// 	return
	// }

	// defer rows.Close()
	// for rows.Next() {
	// 	var tick time.Time
	// 	if err := rows.Scan(&tick); err != nil {
	// 		fmt.Sprintf("Error scanning ticks: %q", err)
	// 		return
	// 	}
	// }

}
