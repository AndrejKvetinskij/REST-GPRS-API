package postgresql

import (
	"NewApiProd/internal/author/model"
	"NewApiProd/internal/author/storage"
	"NewApiProd/pkg/api/filter"
	"NewApiProd/pkg/client/postgresql"
	"NewApiProd/pkg/logging"
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgconn"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

// Create implements author.Repository
func (r *repository) Create(ctx context.Context, author model.Author) error {
	q := `
		INSERT INTO author (name)
		VALUES ($1)
		RETURNING id
	`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	rows := r.client.QueryRow(ctx, q, author.Name)
	err := rows.Scan(&author.ID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, pgErr) {
			pgErr = err.(*pgconn.PgError)
			x := fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState)
			newErr := fmt.Errorf(x)
			r.logger.Error(newErr)
			return nil
		}
		return err
	}

	return nil
}

// Delete implements author.Repository
func (r *repository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// FindAll implements author.Repository
func (r *repository) FindAll(ctx context.Context, paginationOptions storage.PaginationOptions, filterOptions storage.FilterOptions, sortOptions storage.SortOptions) (a []model.Author, err error) {

	qb := squirrel.Select("id, name, age, is_alive, created_at").From("public.author")

	for _, v := range filterOptions.GetDateForQuery() {

		if v.Type != filter.DataTypeDate {
			s1 := ""
			switch v.Operator {
			case filter.OperatorEq:
				s1 = "="
			case filter.OperatorNotEq:
				s1 = "!="
			case filter.OperatorLowerThan:
				s1 = "<"
			case filter.OperatorLowerThanEq:
				s1 = "<="
			case filter.OperatorGreaterThan:
				s1 = ">"
			case filter.OperatorGreaterThanEq:
				s1 = ">="
			case filter.OperatorLike:
				s1 = "LIKE"
				v.Value = fmt.Sprintf("'%s'", v.Value)
			case filter.OperatorBetween:
				s1 = ""
			}
			s := fmt.Sprintf("%s %s %s", v.Name, s1, v.Value)
			qb = qb.Where(s)
		} else {
			s := ""
			s1 := ">="
			s2 := "<="
			switch v.Operator {
			case filter.OperatorEq:

				s = fmt.Sprintf("%s %s '%s 00:00:00.000000' and %s %s '%s 23:59:59.000000'", v.Name, s1, v.Value, v.Name, s2, v.Value)

			case filter.OperatorBetween:

				splited := strings.Split(v.Value, ":")
				value1 := splited[0]
				value2 := splited[1]

				s = fmt.Sprintf("%s %s '%s 00:00:00.000000' and %s %s '%s 23:59:59.000000'", v.Name, s1, value1, v.Name, s2, value2)
				fmt.Println(s)
			}
			qb = qb.Where(s)

		}

	}
	qb = qb.Limit(uint64(filterOptions.GetLimit()))
	if sortOptions != nil {
		qb = qb.OrderBy(sortOptions.GetOrderBy())
	}

	li := filterOptions.GetLimit()
	if li != 0 {
		pt := paginationOptions.GetPtoken() * li
		qb = qb.Offset(uint64(pt))
	}

	sql, i, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(sql)))

	rows, err := r.client.Query(ctx, sql, i...)
	if err != nil {
		return nil, err
	}

	authors := make([]model.Author, 0)

	for rows.Next() {
		var ath model.Author

		err := rows.Scan(&ath.ID, &ath.Name, &ath.Age, &ath.IsAlive, &ath.CreatedAt)
		if err != nil {

			return nil, err
		}

		authors = append(authors, ath)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return authors, nil
}

// FindOne implements author.Repository
func (r *repository) FindOne(ctx context.Context, id string) (model.Author, error) {
	q := `
		SELECT id, name FROM public.author WHERE id = $1
	`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	rows := r.client.QueryRow(ctx, q)

	var ath model.Author

	err := rows.Scan(&ath.ID, &ath.Name)
	if err != nil {
		return model.Author{}, err
	}
	return ath, nil
}

// Update implements author.Repository
func (r *repository) Update(ctx context.Context, author model.Author) error {
	panic("unimplemented")
}

func NewRepository(client postgresql.Client, logger *logging.Logger) storage.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
