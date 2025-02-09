package repository

import (
	"fmt"
	"strings"
	"taskserver/dbserver/internal/entity"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ContainersPostgres struct {
	db *sqlx.DB
}

func NewContainersPostgres(db *sqlx.DB) *ContainersPostgres {
	return &ContainersPostgres{db: db}
}

func (r *ContainersPostgres) Update(item entity.UpdateContainers) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	setValues = append(setValues, fmt.Sprintf("dateofping=$%d", argId))
	args = append(args, item.DateOfPing)
	argId++

	setValues = append(setValues, fmt.Sprintf("timeofping=$%d", argId))
	args = append(args, item.TimeOfPing)
	argId++

	args = append(args, item.Port)

	// date=$1, time=$2
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE port=$3",
		"containers", setQuery)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	res, err := r.db.Exec(query, args...)
	k, _ := res.RowsAffected()
	if k == 0 {
		query := fmt.Sprintf("INSERT INTO %s (port, dateofping, timeofping) VALUES ($1, $2, $3)", "containers")
		_, err = r.db.Exec(query, item.Port, item.DateOfPing, item.TimeOfPing)
	}

	return err
}

func (r *ContainersPostgres) GetAll() ([]entity.Container, error) {
	var items []entity.Container
	query := fmt.Sprintf(`SELECT id, port, timeofping, dateofping FROM %s`,
		"containers")
	if err := r.db.Select(&items, query); err != nil {
		return nil, err
	}

	return items, nil
}
