// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// HostContract is an object representing the database table.
type HostContract struct {
	ID            int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	HostPartnerID int         `boil:"host_partner_id" json:"host_partner_id" toml:"host_partner_id" yaml:"host_partner_id"`
	URL           null.String `boil:"url" json:"url,omitempty" toml:"url" yaml:"url,omitempty"`
	CreatedAt     null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt     null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	DeletedAt     null.Time   `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	StartPeriod   time.Time   `boil:"start_period" json:"start_period" toml:"start_period" yaml:"start_period"`
	EndPeriod     time.Time   `boil:"end_period" json:"end_period" toml:"end_period" yaml:"end_period"`

	R *hostContractR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L hostContractL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var HostContractColumns = struct {
	ID            string
	HostPartnerID string
	URL           string
	CreatedAt     string
	UpdatedAt     string
	DeletedAt     string
	StartPeriod   string
	EndPeriod     string
}{
	ID:            "id",
	HostPartnerID: "host_partner_id",
	URL:           "url",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
	StartPeriod:   "start_period",
	EndPeriod:     "end_period",
}

var HostContractTableColumns = struct {
	ID            string
	HostPartnerID string
	URL           string
	CreatedAt     string
	UpdatedAt     string
	DeletedAt     string
	StartPeriod   string
	EndPeriod     string
}{
	ID:            "HostContract.id",
	HostPartnerID: "HostContract.host_partner_id",
	URL:           "HostContract.url",
	CreatedAt:     "HostContract.created_at",
	UpdatedAt:     "HostContract.updated_at",
	DeletedAt:     "HostContract.deleted_at",
	StartPeriod:   "HostContract.start_period",
	EndPeriod:     "HostContract.end_period",
}

// Generated where

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var HostContractWhere = struct {
	ID            whereHelperint
	HostPartnerID whereHelperint
	URL           whereHelpernull_String
	CreatedAt     whereHelpernull_Time
	UpdatedAt     whereHelpernull_Time
	DeletedAt     whereHelpernull_Time
	StartPeriod   whereHelpertime_Time
	EndPeriod     whereHelpertime_Time
}{
	ID:            whereHelperint{field: "\"HostContract\".\"id\""},
	HostPartnerID: whereHelperint{field: "\"HostContract\".\"host_partner_id\""},
	URL:           whereHelpernull_String{field: "\"HostContract\".\"url\""},
	CreatedAt:     whereHelpernull_Time{field: "\"HostContract\".\"created_at\""},
	UpdatedAt:     whereHelpernull_Time{field: "\"HostContract\".\"updated_at\""},
	DeletedAt:     whereHelpernull_Time{field: "\"HostContract\".\"deleted_at\""},
	StartPeriod:   whereHelpertime_Time{field: "\"HostContract\".\"start_period\""},
	EndPeriod:     whereHelpertime_Time{field: "\"HostContract\".\"end_period\""},
}

// HostContractRels is where relationship names are stored.
var HostContractRels = struct {
	HostPartner string
}{
	HostPartner: "HostPartner",
}

// hostContractR is where relationships are stored.
type hostContractR struct {
	HostPartner *Host `boil:"HostPartner" json:"HostPartner" toml:"HostPartner" yaml:"HostPartner"`
}

// NewStruct creates a new relationship struct
func (*hostContractR) NewStruct() *hostContractR {
	return &hostContractR{}
}

func (r *hostContractR) GetHostPartner() *Host {
	if r == nil {
		return nil
	}
	return r.HostPartner
}

// hostContractL is where Load methods for each relationship are stored.
type hostContractL struct{}

var (
	hostContractAllColumns            = []string{"id", "host_partner_id", "url", "created_at", "updated_at", "deleted_at", "start_period", "end_period"}
	hostContractColumnsWithoutDefault = []string{"host_partner_id", "start_period", "end_period"}
	hostContractColumnsWithDefault    = []string{"id", "url", "created_at", "updated_at", "deleted_at"}
	hostContractPrimaryKeyColumns     = []string{"id"}
	hostContractGeneratedColumns      = []string{}
)

type (
	// HostContractSlice is an alias for a slice of pointers to HostContract.
	// This should almost always be used instead of []HostContract.
	HostContractSlice []*HostContract
	// HostContractHook is the signature for custom HostContract hook methods
	HostContractHook func(context.Context, boil.ContextExecutor, *HostContract) error

	hostContractQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	hostContractType                 = reflect.TypeOf(&HostContract{})
	hostContractMapping              = queries.MakeStructMapping(hostContractType)
	hostContractPrimaryKeyMapping, _ = queries.BindMapping(hostContractType, hostContractMapping, hostContractPrimaryKeyColumns)
	hostContractInsertCacheMut       sync.RWMutex
	hostContractInsertCache          = make(map[string]insertCache)
	hostContractUpdateCacheMut       sync.RWMutex
	hostContractUpdateCache          = make(map[string]updateCache)
	hostContractUpsertCacheMut       sync.RWMutex
	hostContractUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var hostContractAfterSelectMu sync.Mutex
var hostContractAfterSelectHooks []HostContractHook

var hostContractBeforeInsertMu sync.Mutex
var hostContractBeforeInsertHooks []HostContractHook
var hostContractAfterInsertMu sync.Mutex
var hostContractAfterInsertHooks []HostContractHook

var hostContractBeforeUpdateMu sync.Mutex
var hostContractBeforeUpdateHooks []HostContractHook
var hostContractAfterUpdateMu sync.Mutex
var hostContractAfterUpdateHooks []HostContractHook

var hostContractBeforeDeleteMu sync.Mutex
var hostContractBeforeDeleteHooks []HostContractHook
var hostContractAfterDeleteMu sync.Mutex
var hostContractAfterDeleteHooks []HostContractHook

var hostContractBeforeUpsertMu sync.Mutex
var hostContractBeforeUpsertHooks []HostContractHook
var hostContractAfterUpsertMu sync.Mutex
var hostContractAfterUpsertHooks []HostContractHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *HostContract) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostContractAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *HostContract) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostContractBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *HostContract) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostContractAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *HostContract) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostContractBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *HostContract) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostContractAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *HostContract) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostContractBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *HostContract) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostContractAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *HostContract) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostContractBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *HostContract) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostContractAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddHostContractHook registers your hook function for all future operations.
func AddHostContractHook(hookPoint boil.HookPoint, hostContractHook HostContractHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		hostContractAfterSelectMu.Lock()
		hostContractAfterSelectHooks = append(hostContractAfterSelectHooks, hostContractHook)
		hostContractAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		hostContractBeforeInsertMu.Lock()
		hostContractBeforeInsertHooks = append(hostContractBeforeInsertHooks, hostContractHook)
		hostContractBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		hostContractAfterInsertMu.Lock()
		hostContractAfterInsertHooks = append(hostContractAfterInsertHooks, hostContractHook)
		hostContractAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		hostContractBeforeUpdateMu.Lock()
		hostContractBeforeUpdateHooks = append(hostContractBeforeUpdateHooks, hostContractHook)
		hostContractBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		hostContractAfterUpdateMu.Lock()
		hostContractAfterUpdateHooks = append(hostContractAfterUpdateHooks, hostContractHook)
		hostContractAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		hostContractBeforeDeleteMu.Lock()
		hostContractBeforeDeleteHooks = append(hostContractBeforeDeleteHooks, hostContractHook)
		hostContractBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		hostContractAfterDeleteMu.Lock()
		hostContractAfterDeleteHooks = append(hostContractAfterDeleteHooks, hostContractHook)
		hostContractAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		hostContractBeforeUpsertMu.Lock()
		hostContractBeforeUpsertHooks = append(hostContractBeforeUpsertHooks, hostContractHook)
		hostContractBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		hostContractAfterUpsertMu.Lock()
		hostContractAfterUpsertHooks = append(hostContractAfterUpsertHooks, hostContractHook)
		hostContractAfterUpsertMu.Unlock()
	}
}

// One returns a single hostContract record from the query.
func (q hostContractQuery) One(ctx context.Context, exec boil.ContextExecutor) (*HostContract, error) {
	o := &HostContract{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for HostContract")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all HostContract records from the query.
func (q hostContractQuery) All(ctx context.Context, exec boil.ContextExecutor) (HostContractSlice, error) {
	var o []*HostContract

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to HostContract slice")
	}

	if len(hostContractAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all HostContract records in the query.
func (q hostContractQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count HostContract rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q hostContractQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if HostContract exists")
	}

	return count > 0, nil
}

// HostPartner pointed to by the foreign key.
func (o *HostContract) HostPartner(mods ...qm.QueryMod) hostQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.HostPartnerID),
	}

	queryMods = append(queryMods, mods...)

	return Hosts(queryMods...)
}

// LoadHostPartner allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (hostContractL) LoadHostPartner(ctx context.Context, e boil.ContextExecutor, singular bool, maybeHostContract interface{}, mods queries.Applicator) error {
	var slice []*HostContract
	var object *HostContract

	if singular {
		var ok bool
		object, ok = maybeHostContract.(*HostContract)
		if !ok {
			object = new(HostContract)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeHostContract)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeHostContract))
			}
		}
	} else {
		s, ok := maybeHostContract.(*[]*HostContract)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeHostContract)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeHostContract))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &hostContractR{}
		}
		args[object.HostPartnerID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &hostContractR{}
			}

			args[obj.HostPartnerID] = struct{}{}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`Host`),
		qm.WhereIn(`Host.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Host")
	}

	var resultSlice []*Host
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Host")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for Host")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for Host")
	}

	if len(hostAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.HostPartner = foreign
		if foreign.R == nil {
			foreign.R = &hostR{}
		}
		foreign.R.HostPartnerHostContracts = append(foreign.R.HostPartnerHostContracts, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.HostPartnerID == foreign.ID {
				local.R.HostPartner = foreign
				if foreign.R == nil {
					foreign.R = &hostR{}
				}
				foreign.R.HostPartnerHostContracts = append(foreign.R.HostPartnerHostContracts, local)
				break
			}
		}
	}

	return nil
}

// SetHostPartner of the hostContract to the related item.
// Sets o.R.HostPartner to related.
// Adds o to related.R.HostPartnerHostContracts.
func (o *HostContract) SetHostPartner(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Host) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"HostContract\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"host_partner_id"}),
		strmangle.WhereClause("\"", "\"", 2, hostContractPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.HostPartnerID = related.ID
	if o.R == nil {
		o.R = &hostContractR{
			HostPartner: related,
		}
	} else {
		o.R.HostPartner = related
	}

	if related.R == nil {
		related.R = &hostR{
			HostPartnerHostContracts: HostContractSlice{o},
		}
	} else {
		related.R.HostPartnerHostContracts = append(related.R.HostPartnerHostContracts, o)
	}

	return nil
}

// HostContracts retrieves all the records using an executor.
func HostContracts(mods ...qm.QueryMod) hostContractQuery {
	mods = append(mods, qm.From("\"HostContract\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"HostContract\".*"})
	}

	return hostContractQuery{q}
}

// FindHostContract retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindHostContract(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*HostContract, error) {
	hostContractObj := &HostContract{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"HostContract\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, hostContractObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from HostContract")
	}

	if err = hostContractObj.doAfterSelectHooks(ctx, exec); err != nil {
		return hostContractObj, err
	}

	return hostContractObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *HostContract) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no HostContract provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(hostContractColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	hostContractInsertCacheMut.RLock()
	cache, cached := hostContractInsertCache[key]
	hostContractInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			hostContractAllColumns,
			hostContractColumnsWithDefault,
			hostContractColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(hostContractType, hostContractMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(hostContractType, hostContractMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"HostContract\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"HostContract\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into HostContract")
	}

	if !cached {
		hostContractInsertCacheMut.Lock()
		hostContractInsertCache[key] = cache
		hostContractInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the HostContract.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *HostContract) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	hostContractUpdateCacheMut.RLock()
	cache, cached := hostContractUpdateCache[key]
	hostContractUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			hostContractAllColumns,
			hostContractPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update HostContract, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"HostContract\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, hostContractPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(hostContractType, hostContractMapping, append(wl, hostContractPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update HostContract row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for HostContract")
	}

	if !cached {
		hostContractUpdateCacheMut.Lock()
		hostContractUpdateCache[key] = cache
		hostContractUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q hostContractQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for HostContract")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for HostContract")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o HostContractSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), hostContractPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"HostContract\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, hostContractPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in hostContract slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all hostContract")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *HostContract) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no HostContract provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(hostContractColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	hostContractUpsertCacheMut.RLock()
	cache, cached := hostContractUpsertCache[key]
	hostContractUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			hostContractAllColumns,
			hostContractColumnsWithDefault,
			hostContractColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			hostContractAllColumns,
			hostContractPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert HostContract, could not build update column list")
		}

		ret := strmangle.SetComplement(hostContractAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(hostContractPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert HostContract, could not build conflict column list")
			}

			conflict = make([]string, len(hostContractPrimaryKeyColumns))
			copy(conflict, hostContractPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"HostContract\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(hostContractType, hostContractMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(hostContractType, hostContractMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert HostContract")
	}

	if !cached {
		hostContractUpsertCacheMut.Lock()
		hostContractUpsertCache[key] = cache
		hostContractUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single HostContract record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *HostContract) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no HostContract provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), hostContractPrimaryKeyMapping)
	sql := "DELETE FROM \"HostContract\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from HostContract")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for HostContract")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q hostContractQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no hostContractQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from HostContract")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for HostContract")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o HostContractSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(hostContractBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), hostContractPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"HostContract\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, hostContractPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from hostContract slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for HostContract")
	}

	if len(hostContractAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *HostContract) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindHostContract(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *HostContractSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := HostContractSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), hostContractPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"HostContract\".* FROM \"HostContract\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, hostContractPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in HostContractSlice")
	}

	*o = slice

	return nil
}

// HostContractExists checks if the HostContract row exists.
func HostContractExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"HostContract\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if HostContract exists")
	}

	return exists, nil
}

// Exists checks if the HostContract row exists.
func (o *HostContract) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return HostContractExists(ctx, exec, o.ID)
}