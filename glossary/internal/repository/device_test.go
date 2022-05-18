package repository

// func TestAddDevice(t *testing.T){
// 	dbConn, mock := NewGORMMock()
// 	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
// 		Conn:                      dbConn,
// 		SkipInitializeWithVersion: true,
// 	}})
// 	deviceRepo := NewDeviceRepoDriver(db)
// 	defer dbConn.Close()
//
// 	tableName := GetGORMTableName(db, &model.Device{})
//
// 	query := fmt.Sprintf("INSERT INTO `%s`", tableName)
//
//
// }
