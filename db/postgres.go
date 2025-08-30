// Package db provides DB drivers, configurations, and initialization logic.
package db

import (
	"database/sql"
	"fmt"
	"os"
	"taas/config"
	"taas/pkg/tlog"

	_ "github.com/lib/pq" // pg driver
	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

// InitDB initializes the PostgreSQL database connection using GORM and runs goose migrations.
func InitDB(cfg *config.DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		cfg.Host, cfg.UserName, cfg.Password, cfg.Name, cfg.Port, cfg.SSLMode,
	)

	// Initialize raw SQL connection for goose migrations.
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		tlog.Fatal("Failed to open SQL DB: %v", err)
	}

	// Configure and run goose migrations.
	goose.SetBaseFS(os.DirFS("."))
	if err := goose.SetDialect("postgres"); err != nil {
		tlog.Fatal("Failed to set goose dialect: %v", err)
	}

	migrationsDir := "db/migrations"
	if err := goose.Up(sqlDB, migrationsDir); err != nil {
		tlog.Fatal("Goose migration failed: %v", err)
	}

	tlog.Info("Goose migrations completed successfully")

	// Wrap sql.DB into a GORM DB instance.
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		tlog.Fatal("Failed to connect to GORM DB: %v", err)
	}

	// Register GORM multi-tenant callbacks.
	if err := registerTenantCallback(gormDB); err != nil {
		tlog.Fatal("Failed to register tenant callbacks: %v", err)
	}

	return gormDB
}

// registerTenantCallback adds tenant-based filtering and setting hooks to GORM.
func registerTenantCallback(db *gorm.DB) error {
	// Applies tenant filter (WHERE tenant_id = ?) before queries.
	tenantFilter := func(tx *gorm.DB) {
		ctx := tx.Statement.Context
		tenantID, ok := ctx.Value("tenant_id").(uint)
		if ok && tenantID != 0 {
			tx.Statement.AddClause(clause.Where{Exprs: []clause.Expression{
				clause.Eq{Column: "tenant_id", Value: tenantID},
			}})
		}
	}

	// Sets tenant_id field before creating records.
	tenantSetter := func(tx *gorm.DB) {
		ctx := tx.Statement.Context
		tenantID, ok := ctx.Value("tenant_id").(uint)
		if ok && tenantID != 0 && tx.Statement.Schema != nil {
			if field := tx.Statement.Schema.LookUpField("TenantID"); field != nil {
				_ = field.Set(ctx, tx.Statement.ReflectValue, tenantID)
			}
		}
	}

	// Register tenant filter callbacks.
	if err := db.Callback().Query().Before("gorm:query").Register("tenant_filter_query", tenantFilter); err != nil {
		return err
	}
	if err := db.Callback().Delete().Before("gorm:delete").Register("tenant_filter_delete", tenantFilter); err != nil {
		return err
	}
	if err := db.Callback().Update().Before("gorm:update").Register("tenant_filter_update", tenantFilter); err != nil {
		return err
	}
	if err := db.Callback().Row().Before("gorm:row").Register("tenant_filter_row", tenantFilter); err != nil {
		return err
	}
	if err := db.Callback().Raw().Before("gorm:raw").Register("tenant_filter_raw", tenantFilter); err != nil {
		return err
	}
	if err := db.Callback().Create().Before("gorm:create").Register("tenant_setter_create", tenantSetter); err != nil {
		return err
	}

	return nil
}
