// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DailyTripsColumns holds the columns for the "daily_trips" table.
	DailyTripsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "day", Type: field.TypeInt32},
		{Name: "date", Type: field.TypeTime},
		{Name: "notes", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "trip_id", Type: field.TypeUUID},
	}
	// DailyTripsTable holds the schema information for the "daily_trips" table.
	DailyTripsTable = &schema.Table{
		Name:       "daily_trips",
		Columns:    DailyTripsColumns,
		PrimaryKey: []*schema.Column{DailyTripsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "daily_trips_trips_daily_trip",
				Columns:    []*schema.Column{DailyTripsColumns[6]},
				RefColumns: []*schema.Column{TripsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// DailyTripItemsColumns holds the columns for the "daily_trip_items" table.
	DailyTripItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "notes", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "daily_trip_id", Type: field.TypeUUID},
		{Name: "trip_id", Type: field.TypeUUID},
	}
	// DailyTripItemsTable holds the schema information for the "daily_trip_items" table.
	DailyTripItemsTable = &schema.Table{
		Name:       "daily_trip_items",
		Columns:    DailyTripItemsColumns,
		PrimaryKey: []*schema.Column{DailyTripItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "daily_trip_items_daily_trips_daily_trip_item",
				Columns:    []*schema.Column{DailyTripItemsColumns[4]},
				RefColumns: []*schema.Column{DailyTripsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "daily_trip_items_trips_daily_trip_item",
				Columns:    []*schema.Column{DailyTripItemsColumns[5]},
				RefColumns: []*schema.Column{TripsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// MediaColumns holds the columns for the "media" table.
	MediaColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "filename", Type: field.TypeString, Size: 255},
		{Name: "file_type", Type: field.TypeString, Size: 255},
		{Name: "storage_type", Type: field.TypeUint8},
		{Name: "path", Type: field.TypeString, Size: 255},
	}
	// MediaTable holds the schema information for the "media" table.
	MediaTable = &schema.Table{
		Name:       "media",
		Columns:    MediaColumns,
		PrimaryKey: []*schema.Column{MediaColumns[0]},
	}
	// RefreshTokensColumns holds the columns for the "refresh_tokens" table.
	RefreshTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "token", Type: field.TypeString, Size: 64},
		{Name: "expires_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// RefreshTokensTable holds the schema information for the "refresh_tokens" table.
	RefreshTokensTable = &schema.Table{
		Name:       "refresh_tokens",
		Columns:    RefreshTokensColumns,
		PrimaryKey: []*schema.Column{RefreshTokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "refresh_tokens_users_refresh_token",
				Columns:    []*schema.Column{RefreshTokensColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TripsColumns holds the columns for the "trips" table.
	TripsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeBool, Default: true},
		{Name: "title", Type: field.TypeString, Size: 50},
		{Name: "description", Type: field.TypeString, Size: 255},
		{Name: "start_date", Type: field.TypeTime},
		{Name: "end_date", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// TripsTable holds the schema information for the "trips" table.
	TripsTable = &schema.Table{
		Name:       "trips",
		Columns:    TripsColumns,
		PrimaryKey: []*schema.Column{TripsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "trips_users_trip",
				Columns:    []*schema.Column{TripsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeBool, Default: true},
		{Name: "username", Type: field.TypeString, Size: 50},
		{Name: "password", Type: field.TypeString, Size: 255},
		{Name: "email", Type: field.TypeString, Size: 255},
		{Name: "avatar_url", Type: field.TypeString, Nullable: true, Size: 255},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DailyTripsTable,
		DailyTripItemsTable,
		MediaTable,
		RefreshTokensTable,
		TripsTable,
		UsersTable,
	}
)

func init() {
	DailyTripsTable.ForeignKeys[0].RefTable = TripsTable
	DailyTripItemsTable.ForeignKeys[0].RefTable = DailyTripsTable
	DailyTripItemsTable.ForeignKeys[1].RefTable = TripsTable
	RefreshTokensTable.ForeignKeys[0].RefTable = UsersTable
	TripsTable.ForeignKeys[0].RefTable = UsersTable
}
