package repository

import (
	"context"
	"database/sql"
	"privy-test/model"
	"privy-test/param/cake"
	"privy-test/utils"
)

var (
	cakeFindByID     = "FindByID"
	cakeFindByFilter = "FindByFilter"
	cakeCreate       = "Create"
	cakeDelete       = "Delete"
	cakeUpdate       = "Update"
	cakesQueries     = map[string]string{
		cakeFindByID:     `SELECT * FROM cakes WHERE id = ? LIMIT 1`,
		cakeCreate:       `INSERT INTO cakes (title, description, rating, image) values (?,?,?,?)`,
		cakeDelete:       `DELETE FROM cakes WHERE id = ?`,
		cakeUpdate:       `UPDATE cakes set title = ?, description = ?, rating = ?, image = ? where id = ?`,
		cakeFindByFilter: `SELECT * FROM cakes WHERE title like ? and description like ? and rating between ? and ? order by rating desc, title asc`,
	}
)

type cakeRepository struct {
	db *sql.DB
	cm utils.CommonService
}

func NewCakeRepository(db *sql.DB, cm utils.CommonService) CakeRepository {
	return &cakeRepository{db, cm}
}

// Statement Get registered prepared statement.
func (c *cakeRepository) Statement(ctx context.Context, name string) (*sql.Stmt, error) {
	v, err := c.cm.GetPrepareStatement(ctx, name, cakesQueries[name], c.db)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (c cakeRepository) GetDetail(ctx context.Context, cakeID int64) (*model.Cake, error) {
	var result *model.Cake

	stmt, err := c.Statement(ctx, cakeFindByID)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, cakeID)
	if err != nil {
		return nil, err
	}

	return result, utils.LoadOne(ctx, rows, &result)
}

func (c *cakeRepository) Create(ctx context.Context, cakeModel *model.Cake) (*model.Cake, error) {
	stmt, err := c.Statement(ctx, cakeCreate)
	if err != nil {
		return nil, err
	}

	res, err := stmt.ExecContext(ctx, cakeModel.Title, cakeModel.Description, cakeModel.Rating, cakeModel.Image)
	if err != nil {
		return nil, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return c.GetDetail(ctx, lastID)
}

func (c *cakeRepository) Delete(ctx context.Context, cakeID int64) error {
	stmt, err := c.Statement(ctx, cakeDelete)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, cakeID)
	if err != nil {
		return err
	}

	return nil
}

func (c *cakeRepository) Update(ctx context.Context, cakeID int64, cakeModel *model.Cake) (*model.Cake, error) {
	stmt, err := c.Statement(ctx, cakeUpdate)
	if err != nil {
		return nil, err
	}

	_, err = stmt.ExecContext(ctx, cakeModel.Title, cakeModel.Description, cakeModel.Rating, cakeModel.Image, cakeID)
	if err != nil {
		return nil, err
	}

	return c.GetDetail(ctx, cakeID)
}

func (c *cakeRepository) GetList(ctx context.Context, request *cake.FindAllRequest) ([]model.Cake, error) {
	var result []model.Cake
	stmt, err := c.Statement(ctx, cakeFindByFilter)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, request.Title, request.Description, request.MinRating, request.MaxRating)
	if err != nil {
		return nil, err
	}

	_, err = utils.LoadToStruct(ctx, rows, &result)
	return result, err
}
