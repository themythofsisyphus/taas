package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"taas/config"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func InitDB(config *config.DatabaseConfig) *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		config.Host, config.UserName, config.Password, config.Name, config.Port, config.SSLMode)

	// Raw sql.DB connection for goose
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to open SQL DB:", err)
	}

	// Run Goose migrations
	goose.SetBaseFS(os.DirFS("."))
	goose.SetDialect("postgres")
	migrationsDir := "db/migrations"

	if err := goose.Up(sqlDB, migrationsDir); err != nil {
		log.Fatalf("Goose migration failed: %v", err)
	}

	log.Println("Goose migrations completed successfully")

	// Wrap into GORM
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to GORM DB:", err)
	}

	// Register GORM hooks or callbacks
	registerTenantCallback(gormDB)

	return gormDB
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
