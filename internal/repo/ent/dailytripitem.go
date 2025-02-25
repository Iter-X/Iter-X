// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/iter-x/iter-x/internal/repo/ent/dailytrip"
	"github.com/iter-x/iter-x/internal/repo/ent/dailytripitem"
	"github.com/iter-x/iter-x/internal/repo/ent/trip"
)

// DailyTripItem is the model entity for the DailyTripItem schema.
type DailyTripItem struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// TripID holds the value of the "trip_id" field.
	TripID uuid.UUID `json:"trip_id,omitempty"`
	// DailyTripID holds the value of the "daily_trip_id" field.
	DailyTripID uuid.UUID `json:"daily_trip_id,omitempty"`
	// Notes holds the value of the "notes" field.
	Notes string `json:"notes,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DailyTripItemQuery when eager-loading is set.
	Edges        DailyTripItemEdges `json:"edges"`
	selectValues sql.SelectValues
}

// DailyTripItemEdges holds the relations/edges for other nodes in the graph.
type DailyTripItemEdges struct {
	// Trip holds the value of the trip edge.
	Trip *Trip `json:"trip,omitempty"`
	// DailyTrip holds the value of the daily_trip edge.
	DailyTrip *DailyTrip `json:"daily_trip,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TripOrErr returns the Trip value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DailyTripItemEdges) TripOrErr() (*Trip, error) {
	if e.Trip != nil {
		return e.Trip, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: trip.Label}
	}
	return nil, &NotLoadedError{edge: "trip"}
}

// DailyTripOrErr returns the DailyTrip value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DailyTripItemEdges) DailyTripOrErr() (*DailyTrip, error) {
	if e.DailyTrip != nil {
		return e.DailyTrip, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: dailytrip.Label}
	}
	return nil, &NotLoadedError{edge: "daily_trip"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DailyTripItem) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case dailytripitem.FieldNotes:
			values[i] = new(sql.NullString)
		case dailytripitem.FieldCreatedAt, dailytripitem.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case dailytripitem.FieldID, dailytripitem.FieldTripID, dailytripitem.FieldDailyTripID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DailyTripItem fields.
func (dti *DailyTripItem) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dailytripitem.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				dti.ID = *value
			}
		case dailytripitem.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				dti.CreatedAt = value.Time
			}
		case dailytripitem.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				dti.UpdatedAt = value.Time
			}
		case dailytripitem.FieldTripID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field trip_id", values[i])
			} else if value != nil {
				dti.TripID = *value
			}
		case dailytripitem.FieldDailyTripID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field daily_trip_id", values[i])
			} else if value != nil {
				dti.DailyTripID = *value
			}
		case dailytripitem.FieldNotes:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field notes", values[i])
			} else if value.Valid {
				dti.Notes = value.String
			}
		default:
			dti.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the DailyTripItem.
// This includes values selected through modifiers, order, etc.
func (dti *DailyTripItem) Value(name string) (ent.Value, error) {
	return dti.selectValues.Get(name)
}

// QueryTrip queries the "trip" edge of the DailyTripItem entity.
func (dti *DailyTripItem) QueryTrip() *TripQuery {
	return NewDailyTripItemClient(dti.config).QueryTrip(dti)
}

// QueryDailyTrip queries the "daily_trip" edge of the DailyTripItem entity.
func (dti *DailyTripItem) QueryDailyTrip() *DailyTripQuery {
	return NewDailyTripItemClient(dti.config).QueryDailyTrip(dti)
}

// Update returns a builder for updating this DailyTripItem.
// Note that you need to call DailyTripItem.Unwrap() before calling this method if this DailyTripItem
// was returned from a transaction, and the transaction was committed or rolled back.
func (dti *DailyTripItem) Update() *DailyTripItemUpdateOne {
	return NewDailyTripItemClient(dti.config).UpdateOne(dti)
}

// Unwrap unwraps the DailyTripItem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dti *DailyTripItem) Unwrap() *DailyTripItem {
	_tx, ok := dti.config.driver.(*txDriver)
	if !ok {
		panic("ent: DailyTripItem is not a transactional entity")
	}
	dti.config.driver = _tx.drv
	return dti
}

// String implements the fmt.Stringer.
func (dti *DailyTripItem) String() string {
	var builder strings.Builder
	builder.WriteString("DailyTripItem(")
	builder.WriteString(fmt.Sprintf("id=%v, ", dti.ID))
	builder.WriteString("created_at=")
	builder.WriteString(dti.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(dti.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("trip_id=")
	builder.WriteString(fmt.Sprintf("%v", dti.TripID))
	builder.WriteString(", ")
	builder.WriteString("daily_trip_id=")
	builder.WriteString(fmt.Sprintf("%v", dti.DailyTripID))
	builder.WriteString(", ")
	builder.WriteString("notes=")
	builder.WriteString(dti.Notes)
	builder.WriteByte(')')
	return builder.String()
}

// DailyTripItems is a parsable slice of DailyTripItem.
type DailyTripItems []*DailyTripItem
