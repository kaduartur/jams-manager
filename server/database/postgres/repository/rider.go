package repository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"

	"github.com/BorsaTeam/jams-manager/server"
	"github.com/BorsaTeam/jams-manager/server/database"
)

var _ Rider = RiderRepo{}

type Rider interface {
	Save(rider RiderEntity) (string, error)
	FindOne(id string) (RiderEntity, error)
	Delete(id string) error
	FindAll(page server.PageRequest) ([]RiderEntity, error)
	Count() (int, error)
}

type RiderEntity struct {
	Id               string     `json:"id"`
	Name             string     `json:"name"`
	Age              int        `json:"age"`
	Gender           string     `json:"gender"`
	City             string     `json:"city"`
	Email            string     `json:"email"`
	PaidSubscription bool       `json:"paidSubscription"`
	Sponsors         string     `json:"sponsors"`
	CategoryId       string     `json:"categoryId"`
	CreateAt         time.Time  `json:"createAt"`
	UpdateAt         *time.Time `json:"updateAt,omitempty"`
}

type RiderRepo struct {
	database database.DbConnection
}

func NewRiderRepository(d database.DbConnection) RiderRepo {
	return RiderRepo{database: d}
}

func (r RiderRepo) Save(rider RiderEntity) (string, error) {

	statement := `INSERT INTO public.RIDERS
				  (rider_id, name, age, gender, city, cpf, paid_subscription, sponsors, category_id, created, updated)
				  VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11);`

	db := r.database.ConnectHandle()
	defer db.Close()

	id, err := uuid.NewRandom()

	if err != nil {
		return "", err
	}

	_, err = db.Exec(statement,
		id,
		rider.Name,
		rider.Age,
		rider.Gender,
		rider.City,
		rider.Email,
		rider.PaidSubscription,
		rider.Sponsors,
		rider.CategoryId,
		rider.CreateAt,
		rider.UpdateAt)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}
func (r RiderRepo) FindOne(id string) (RiderEntity, error) {
	statement := `SELECT * FROM RIDERS WHERE RIDER_ID=$1`
	db := r.database.ConnectHandle()
	defer db.Close()

	re := RiderEntity{}

	row := db.QueryRow(statement, id)
	err := row.Scan(
		&re.Id,
		&re.Name,
		&re.Age,
		&re.Gender,
		&re.City,
		&re.Email,
		&re.PaidSubscription,
		&re.Sponsors,
		&re.CategoryId,
		&re.CreateAt,
		&re.UpdateAt,
	)

	if err == sql.ErrNoRows {
		return RiderEntity{}, nil
	}

	if err != nil {
		return RiderEntity{}, err
	}
	return re, nil
}

func (r RiderRepo) Delete(id string) error {
	stmt := `DELETE FROM riders WHERE rider_id=$1`
	db := r.database.ConnectHandle()
	defer db.Close()

	if _, err := db.Exec(stmt, id); err != nil {
		return err
	}

	return nil
}

func (r RiderRepo) Count() (int, error) {
	stmt := `SELECT COUNT(*) FROM riders`
	db := r.database.ConnectHandle()
	defer db.Close()

	var result int

	row := db.QueryRow(stmt)
	if err := row.Scan(&result); err != nil {
		return 0, err
	}

	return result, nil
}

func (r RiderRepo) FindAll(page server.PageRequest) ([]RiderEntity, error) {
	stmt := `SELECT * FROM riders LIMIT $1 OFFSET $2`
	db := r.database.ConnectHandle()
	defer db.Close()

	offset := (page.Page - 1) * page.PerPage
	rows, err := db.Query(stmt, page.PerPage, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	riders := make([]RiderEntity, 0)
	for rows.Next() {
		re := RiderEntity{}
		err := rows.Scan(
			&re.Id,
			&re.Name,
			&re.Age,
			&re.Gender,
			&re.City,
			&re.Email,
			&re.PaidSubscription,
			&re.Sponsors,
			&re.CategoryId,
			&re.CreateAt,
			&re.UpdateAt,
		)

		if err != nil {
			return nil, err
		}

		riders = append(riders, re)
	}

	return riders, nil
}
