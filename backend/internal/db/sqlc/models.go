// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type ClimateNormal struct {
	ID      int32       `json:"id"`
	PlaceID pgtype.Int4 `json:"place_id"`
	// Range: 1 to 12
	Month       int32              `json:"month"`
	Tavg        pgtype.Numeric     `json:"tavg"`
	Tmin        pgtype.Numeric     `json:"tmin"`
	Tmax        pgtype.Numeric     `json:"tmax"`
	Prcp        pgtype.Numeric     `json:"prcp"`
	Wspd        pgtype.Numeric     `json:"wspd"`
	Pres        pgtype.Numeric     `json:"pres"`
	Tsun        pgtype.Int4        `json:"tsun"`
	LastUpdated pgtype.Timestamptz `json:"last_updated"`
}

type Place struct {
	ID        int32          `json:"id"`
	Name      string         `json:"name"`
	Country   string         `json:"country"`
	Latitude  pgtype.Numeric `json:"latitude"`
	Longitude pgtype.Numeric `json:"longitude"`
	Altitude  pgtype.Int4    `json:"altitude"`
}
