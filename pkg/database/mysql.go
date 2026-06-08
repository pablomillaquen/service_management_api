package database

import (
	"fmt"
	"time"

	"github.com/pablomillaquen/speckit_golang_api/configs"
	"github.com/pablomillaquen/speckit_golang_api/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

type tabler interface {
	TableName() string
}

func Connect(cfg configs.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(cfg.MaxLifetime)

	DB = db
	logger.Info("Database connected successfully")
	return db, nil
}

func Ping(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}

func IsReady(db *gorm.DB) bool {
	sqlDB, err := db.DB()
	if err != nil {
		return false
	}
	if err := sqlDB.Ping(); err != nil {
		return false
	}
	return true
}

type Migration struct {
	ID      string
	Up      func(db *gorm.DB) error
	Down    func(db *gorm.DB) error
}

type MigrationRunner struct {
	db         *gorm.DB
	migrations []Migration
	tableName  string
}

func NewMigrationRunner(db *gorm.DB, tableName string) *MigrationRunner {
	return &MigrationRunner{
		db:        db,
		tableName: tableName,
	}
}

func (r *MigrationRunner) Add(m Migration) {
	r.migrations = append(r.migrations, m)
}

func (r *MigrationRunner) Run() error {
	if err := r.ensureTable(); err != nil {
		return fmt.Errorf("failed to ensure migrations table: %w", err)
	}

	for _, m := range r.migrations {
		applied, err := r.isApplied(m.ID)
		if err != nil {
			return err
		}
		if applied {
			continue
		}
		logger.Info("Running migration: %s", m.ID)
		if err := m.Up(r.db); err != nil {
			return fmt.Errorf("migration %s failed: %w", m.ID, err)
		}
		if err := r.markApplied(m.ID); err != nil {
			return err
		}
	}
	return nil
}

func (r *MigrationRunner) ensureTable() error {
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id VARCHAR(255) PRIMARY KEY,
		applied_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`, r.tableName)
	return r.db.Exec(query).Error
}

func (r *MigrationRunner) isApplied(id string) (bool, error) {
	var count int64
	err := r.db.Table(r.tableName).Where("id = ?", id).Count(&count).Error
	return count > 0, err
}

func (r *MigrationRunner) markApplied(id string) error {
	return r.db.Table(r.tableName).Exec(
		"INSERT INTO ? (id, applied_at) VALUES (?, ?)",
		gorm.Expr(r.tableName), id, time.Now(),
	).Error
}

func AutoMigrate(db *gorm.DB, models ...interface{}) error {
	logger.Info("Running auto migrations")
	return db.AutoMigrate(models...)
}
