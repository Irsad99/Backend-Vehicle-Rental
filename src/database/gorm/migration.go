package gorm

import (
	"log"

	"BackendGo/src/database/gorm/models"
	"BackendGo/src/database"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "run database migration",
	RunE:  dbMigrate,
}

var migUp bool
var migDown bool

func init() {
	MigrateCmd.Flags().BoolVarP(&migUp, "up", "u", false, "run miggration up")
	MigrateCmd.Flags().BoolVarP(&migDown, "down", "d", false, "run miggration down")
}

func dbMigrate(cmd *cobra.Command, args []string) error {
	db, err := database.New()
	if err != nil {
		return err
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "001",
			Migrate: func(d *gorm.DB) error {
				return d.AutoMigrate(&models.User{})
			},
			Rollback: func(d *gorm.DB) error {
				return d.Migrator().DropColumn(&models.User{}, "description")
			},
		},
	})

	if migUp {
		if err = m.Migrate(); err != nil {
			return err
		}
		log.Fatal("migration did run successfully")
		return nil
	}

	if migDown {
		if err = m.RollbackLast(); err != nil {
			return err
		}
		log.Fatal("Rollback did run successfully")
		return nil
	}

	m.InitSchema(func(tx *gorm.DB) error {
		err := tx.AutoMigrate(
			&models.User{},
			&models.History{},
			&models.Vehicle{},
		)

		if err != nil {
			return err
		}

		return nil
	})

	if err = m.Migrate(); err != nil {
		return err
	}

	log.Fatal("init schema successfully")
	return nil

}