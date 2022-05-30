package utils

import (
	"context"
	"database/sql"
	"time"
)

type CommonService interface {
	TimeNowUTC() time.Time
	GetPrepareStatement(ctx context.Context, key, value string, db *sql.DB) (*sql.Stmt, error)
	GetQuery(key string) string
}

type RealCommonService struct {
	queryMap    map[string]*sql.Stmt
	queryMapStr map[string]string
}

func NewCommonService() CommonService {
	queryMap := make(map[string]*sql.Stmt)
	queryMapStr := make(map[string]string)
	return &RealCommonService{
		queryMap:    queryMap,
		queryMapStr: queryMapStr,
	}
}

func (c *RealCommonService) TimeNowUTC() time.Time {
	return time.Now().UTC()
}

func (c *RealCommonService) GetPrepareStatement(ctx context.Context, key, value string, db *sql.DB) (*sql.Stmt, error) {
	if val, ok := c.queryMap[key]; ok {
		return val, nil
	}

	p, err := db.PrepareContext(ctx, value)
	if err != nil {
		return nil, err
	}

	c.queryMap[key] = p
	c.queryMapStr[key] = value
	return p, nil
}

func (c *RealCommonService) GetQuery(key string) string {
	return c.queryMapStr[key]
}
