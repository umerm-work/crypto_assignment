package postgres

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/umerm-work/crypto_assignment/config"
	"github.com/umerm-work/crypto_assignment/domain"
	"strings"
)

type repository struct {
	db *gorm.DB
	//logger log.go.Logger
}

func (r *repository) Fetch(ctx context.Context, fromVirtualCurrency, toPhysicalCurrency string) ([]domain.CurrencyConversions, error) {
	var res []domain.CurrencyConversions
	err := r.db.Where(`"virtualcurrencyname" IN (?) AND "physicalcurrencyname" IN (?)`,
		strings.Split(fromVirtualCurrency, ","),
		strings.Split(toPhysicalCurrency, ",")).
		Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *repository) Insert(ctx context.Context, in domain.CurrencyConversions) error {
	logrus.Println("data", in)
	var currentDate domain.CurrencyConversions
	err := r.db.Where(`"virtualcurrencyname" = ? AND "physicalcurrencyname" = ?`, in.VIRTUALCURRENCYNAME, in.PHYSICALCURRENCYNAME).Find(&currentDate).Error
	if err != nil {
		logrus.Println("db error", err)
	}
	id := currentDate.ID
	currentDate = in
	currentDate.ID = id
	saveErr := r.db.Model(domain.CurrencyConversions{}).Save(&currentDate).Error
	if saveErr != nil {
		logrus.Println("db error", saveErr)
		return err
	}
	return nil
}

func New(config config.Config) domain.PriceRepository {
	arg := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Database,
	)
	db, err := gorm.Open("postgres", arg)
	if err != nil {
		panic(err.Error())
	}
	//Migrate the schema
	db.AutoMigrate(&domain.CurrencyConversions{})
	db.LogMode(config.DB.Debug)
	return &repository{
		db: db,
	}
}
