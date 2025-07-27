package db

import (
	"fmt"
	"log"
	"taas/config"
	"taas/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func InitDB(config *config.DatabaseConfig) *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		config.Host, config.UserName, config.Password, config.Name, config.Port, config.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	if err := db.AutoMigrate(&model.Tag{}); err != nil {
		log.Fatal("Failed to migrate Tag table:", err)
	}

	if err := db.AutoMigrate(&model.TagMapping{}); err != nil {
		log.Fatal("Failed to migrate Tag Mapping table:", err)
	}

	if err := db.AutoMigrate(&model.Entity{}); err != nil {
		log.Fatal("Failed to migrate Entity table:", err)
	}

	if err := db.AutoMigrate(&model.Tenant{}); err != nil {
		log.Println("Attempting to migrate Tenant table...")
		if err := db.AutoMigrate(&model.Tenant{}); err != nil {
			log.Fatal("Failed to migrate Tenant table:", err)
		}
		log.Println("Tenant table migrated successfully.")
		log.Fatal("Failed to migrate Tenant table:", err)
	}

	registerTenantCallback(db)

	return db
}

func registerTenantCallback(db *gorm.DB) {
	tenantFilter := func(tx *gorm.DB) {
		ctx := tx.Statement.Context
		tenantID, ok := ctx.Value("tenant_id").(uint)
		if ok && tenantID != 0 {
			tx.Statement.AddClause(clause.Where{Exprs: []clause.Expression{
				clause.Eq{Column: "tenant_id", Value: tenantID},
			}})
		}
	}

	tenantSetter := func(tx *gorm.DB) {
		ctx := tx.Statement.Context
		tenantID, ok := ctx.Value("tenant_id").(uint)
		if ok && tenantID != 0 && tx.Statement.Schema != nil {
			if field := tx.Statement.Schema.LookUpField("TenantID"); field != nil {
				_ = field.Set(tx.Statement.Context, tx.Statement.ReflectValue, tenantID)
			}
		}
	}

	db.Callback().Query().Before("gorm:query").Register("tenant_filter_query", tenantFilter)
	db.Callback().Delete().Before("gorm:delete").Register("tenant_filter_delete", tenantFilter)
	db.Callback().Update().Before("gorm:update").Register("tenant_filter_update", tenantFilter)
	db.Callback().Row().Before("gorm:row").Register("tenant_filter_row", tenantFilter)
	db.Callback().Raw().Before("gorm:raw").Register("tenant_filter_raw", tenantFilter)
	db.Callback().Create().Before("gorm:create").Register("tenant_setter_create", tenantSetter)
}
