package graph

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/highlight-run/highlight/backend/model"
)

var DB *gorm.DB

func createAndMigrateTestDB(dbName string) (*gorm.DB, error) {
	log.Println("host", os.Getenv("PSQL_HOST"))
	psqlConf := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s sslmode=disable",
		os.Getenv("PSQL_HOST"),
		os.Getenv("PSQL_PORT"),
		os.Getenv("PSQL_USER"),
		os.Getenv("PSQL_PASSWORD"))
	// Open the database object without an actual db_name.
	db, err := gorm.Open(postgres.Open(psqlConf))
	if err != nil {
		return nil, errors.Wrap(err, "error opening test db")
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.Wrap(err, "error retrieving test db")
	}
	defer sqlDB.Close()
	// Attempt to create the database.
	db = db.Exec(fmt.Sprintf("CREATE DATABASE %v;", dbName))
	return model.SetupDB(dbName)
}

// Gets run once; M.run() calls the tests in this file.
func TestMain(m *testing.M) {
	var err error
	DB, err = createAndMigrateTestDB("highlight_testing_db")
	if err != nil {
		log.Fatalf("error creating testdb: %v", err)
	}
	code := m.Run()
	os.Exit(code)
}

func TestHideViewedSessions(t *testing.T) {
	// insert data
	// sessionsToInsert := []model.Session{
	// 	{Viewed: &model.T},
	// 	{Viewed: &model.F},
	// }
	// err := DB.Create(&sessionsToInsert)
	// if err != nil {
	// 	t.Fatalf("error inserting sessions: %v", err)
	// }
	// // test
	// r := &queryResolver{Resolver: &Resolver{DB: DB}}
	// params := &modelInputs.SearchParamsInput{HideViewed: &model.T}
	// sessions, err := r.Sessions(context.Background(), 1, 5, modelInputs.SessionLifecycleAll, false, params)
	// if err != nil {
	// 	t.Fatalf("error querying sessions: %v", err)
	// }

	// expected := &model.SessionResults{
	// 	Sessions: []model.Session{
	// 		model.Session{Viewed: &model.F},
	// 	},
	// 	TotalCount: 2,
	// }
	// if diff := deep.Equal(sessions, expected); diff != nil {
	// 	t.Fatalf("received sessions and expected sessions not equal: %v", diff)
	// }

	// log.Println("hello")
}

func TestSessionsOther(t *testing.T) {
	log.Println("hillo")
}
