// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package queries

import (
	"context"
	"database/sql"
)

const getAllMarkets = `-- name: GetAllMarkets :many
SELECT market_id, provider_name, url, base_asset, quote_asset, active FROM market
`

func (q *Queries) GetAllMarkets(ctx context.Context) ([]Market, error) {
	rows, err := q.db.QueryContext(ctx, getAllMarkets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Market
	for rows.Next() {
		var i Market
		if err := rows.Scan(
			&i.MarketID,
			&i.ProviderName,
			&i.Url,
			&i.BaseAsset,
			&i.QuoteAsset,
			&i.Active,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertMarket = `-- name: InsertMarket :one
INSERT INTO market (
    provider_name,url,base_asset,quote_asset,active) VALUES (
             $1, $2, $3, $4, $5
    )
    RETURNING market_id, provider_name, url, base_asset, quote_asset, active
`

type InsertMarketParams struct {
	ProviderName string
	Url          string
	BaseAsset    string
	QuoteAsset   string
	Active       sql.NullBool
}

func (q *Queries) InsertMarket(ctx context.Context, arg InsertMarketParams) (Market, error) {
	row := q.db.QueryRowContext(ctx, insertMarket,
		arg.ProviderName,
		arg.Url,
		arg.BaseAsset,
		arg.QuoteAsset,
		arg.Active,
	)
	var i Market
	err := row.Scan(
		&i.MarketID,
		&i.ProviderName,
		&i.Url,
		&i.BaseAsset,
		&i.QuoteAsset,
		&i.Active,
	)
	return i, err
}

const updateActive = `-- name: UpdateActive :exec
UPDATE market set active = $1 where market_id = $2
`

type UpdateActiveParams struct {
	Active   sql.NullBool
	MarketID sql.NullInt32
}

func (q *Queries) UpdateActive(ctx context.Context, arg UpdateActiveParams) error {
	_, err := q.db.ExecContext(ctx, updateActive, arg.Active, arg.MarketID)
	return err
}
