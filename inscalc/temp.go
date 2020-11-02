package inscalc

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func gormInitForTest() (*gorm.DB, error) {
	dsn := "host=192.168.5.11 user=fzzl password=fzzl032003 dbname=lpr port=15432 sslmode=disable TimeZone=Asia/Shanghai"
	gormv2, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//	TablePrefix:   "t_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	if err != nil {
		return nil, err
	}
	return gormv2, err

}
