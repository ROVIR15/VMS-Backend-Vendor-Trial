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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// VendorDetail is an object representing the database table.
type VendorDetail struct {
	ID                     int     `boil:"id" json:"id" toml:"id" yaml:"id"`
	VendorID               int     `boil:"vendor_id" json:"vendor_id" toml:"vendor_id" yaml:"vendor_id"`
	Commission             float32 `boil:"commission" json:"commission" toml:"commission" yaml:"commission"`
	IsCommissionIncludeTax bool    `boil:"is_commission_include_tax" json:"is_commission_include_tax" toml:"is_commission_include_tax" yaml:"is_commission_include_tax"`

	R *vendorDetailR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L vendorDetailL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var VendorDetailColumns = struct {
	ID                     string
	VendorID               string
	Commission             string
	IsCommissionIncludeTax string
}{
	ID:                     "id",
	VendorID:               "vendor_id",
	Commission:             "commission",
	IsCommissionIncludeTax: "is_commission_include_tax",
}

var VendorDetailTableColumns = struct {
	ID                     string
	VendorID               string
	Commission             string
	IsCommissionIncludeTax string
}{
	ID:                     "VendorDetail.id",
	VendorID:               "VendorDetail.vendor_id",
	Commission:             "VendorDetail.commission",
	IsCommissionIncludeTax: "VendorDetail.is_commission_include_tax",
}

// Generated where

type whereHelperfloat32 struct{ field string }

func (w whereHelperfloat32) EQ(x float32) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperfloat32) NEQ(x float32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperfloat32) LT(x float32) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperfloat32) LTE(x float32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperfloat32) GT(x float32) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperfloat32) GTE(x float32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperfloat32) IN(slice []float32) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperfloat32) NIN(slice []float32) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var VendorDetailWhere = struct {
	ID                     whereHelperint
	VendorID               whereHelperint
	Commission             whereHelperfloat32
	IsCommissionIncludeTax whereHelperbool
}{
	ID:                     whereHelperint{field: "\"VendorDetail\".\"id\""},
	VendorID:               whereHelperint{field: "\"VendorDetail\".\"vendor_id\""},
	Commission:             whereHelperfloat32{field: "\"VendorDetail\".\"commission\""},
	IsCommissionIncludeTax: whereHelperbool{field: "\"VendorDetail\".\"is_commission_include_tax\""},
}

// VendorDetailRels is where relationship names are stored.
var VendorDetailRels = struct {
	Vendor string
}{
	Vendor: "Vendor",
}

// vendorDetailR is where relationships are stored.
type vendorDetailR struct {
	Vendor *Vendor `boil:"Vendor" json:"Vendor" toml:"Vendor" yaml:"Vendor"`
}

// NewStruct creates a new relationship struct
func (*vendorDetailR) NewStruct() *vendorDetailR {
	return &vendorDetailR{}
}

func (r *vendorDetailR) GetVendor() *Vendor {
	if r == nil {
		return nil
	}
	return r.Vendor
}

// vendorDetailL is where Load methods for each relationship are stored.
type vendorDetailL struct{}

var (
	vendorDetailAllColumns            = []string{"id", "vendor_id", "commission", "is_commission_include_tax"}
	vendorDetailColumnsWithoutDefault = []string{"vendor_id"}
	vendorDetailColumnsWithDefault    = []string{"id", "commission", "is_commission_include_tax"}
	vendorDetailPrimaryKeyColumns     = []string{"id"}
	vendorDetailGeneratedColumns      = []string{}
)

type (
	// VendorDetailSlice is an alias for a slice of pointers to VendorDetail.
	// This should almost always be used instead of []VendorDetail.
	VendorDetailSlice []*VendorDetail
	// VendorDetailHook is the signature for custom VendorDetail hook methods
	VendorDetailHook func(context.Context, boil.ContextExecutor, *VendorDetail) error

	vendorDetailQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	vendorDetailType                 = reflect.TypeOf(&VendorDetail{})
	vendorDetailMapping              = queries.MakeStructMapping(vendorDetailType)
	vendorDetailPrimaryKeyMapping, _ = queries.BindMapping(vendorDetailType, vendorDetailMapping, vendorDetailPrimaryKeyColumns)
	vendorDetailInsertCacheMut       sync.RWMutex
	vendorDetailInsertCache          = make(map[string]insertCache)
	vendorDetailUpdateCacheMut       sync.RWMutex
	vendorDetailUpdateCache          = make(map[string]updateCache)
	vendorDetailUpsertCacheMut       sync.RWMutex
	vendorDetailUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var vendorDetailAfterSelectMu sync.Mutex
var vendorDetailAfterSelectHooks []VendorDetailHook

var vendorDetailBeforeInsertMu sync.Mutex
var vendorDetailBeforeInsertHooks []VendorDetailHook
var vendorDetailAfterInsertMu sync.Mutex
var vendorDetailAfterInsertHooks []VendorDetailHook

var vendorDetailBeforeUpdateMu sync.Mutex
var vendorDetailBeforeUpdateHooks []VendorDetailHook
var vendorDetailAfterUpdateMu sync.Mutex
var vendorDetailAfterUpdateHooks []VendorDetailHook

var vendorDetailBeforeDeleteMu sync.Mutex
var vendorDetailBeforeDeleteHooks []VendorDetailHook
var vendorDetailAfterDeleteMu sync.Mutex
var vendorDetailAfterDeleteHooks []VendorDetailHook

var vendorDetailBeforeUpsertMu sync.Mutex
var vendorDetailBeforeUpsertHooks []VendorDetailHook
var vendorDetailAfterUpsertMu sync.Mutex
var vendorDetailAfterUpsertHooks []VendorDetailHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *VendorDetail) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vendorDetailAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *VendorDetail) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vendorDetailBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *VendorDetail) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vendorDetailAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *VendorDetail) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vendorDetailBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *VendorDetail) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vendorDetailAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *VendorDetail) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vendorDetailBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *VendorDetail) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vendorDetailAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *VendorDetail) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vendorDetailBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *VendorDetail) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vendorDetailAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddVendorDetailHook registers your hook function for all future operations.
func AddVendorDetailHook(hookPoint boil.HookPoint, vendorDetailHook VendorDetailHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		vendorDetailAfterSelectMu.Lock()
		vendorDetailAfterSelectHooks = append(vendorDetailAfterSelectHooks, vendorDetailHook)
		vendorDetailAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		vendorDetailBeforeInsertMu.Lock()
		vendorDetailBeforeInsertHooks = append(vendorDetailBeforeInsertHooks, vendorDetailHook)
		vendorDetailBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		vendorDetailAfterInsertMu.Lock()
		vendorDetailAfterInsertHooks = append(vendorDetailAfterInsertHooks, vendorDetailHook)
		vendorDetailAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		vendorDetailBeforeUpdateMu.Lock()
		vendorDetailBeforeUpdateHooks = append(vendorDetailBeforeUpdateHooks, vendorDetailHook)
		vendorDetailBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		vendorDetailAfterUpdateMu.Lock()
		vendorDetailAfterUpdateHooks = append(vendorDetailAfterUpdateHooks, vendorDetailHook)
		vendorDetailAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		vendorDetailBeforeDeleteMu.Lock()
		vendorDetailBeforeDeleteHooks = append(vendorDetailBeforeDeleteHooks, vendorDetailHook)
		vendorDetailBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		vendorDetailAfterDeleteMu.Lock()
		vendorDetailAfterDeleteHooks = append(vendorDetailAfterDeleteHooks, vendorDetailHook)
		vendorDetailAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		vendorDetailBeforeUpsertMu.Lock()
		vendorDetailBeforeUpsertHooks = append(vendorDetailBeforeUpsertHooks, vendorDetailHook)
		vendorDetailBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		vendorDetailAfterUpsertMu.Lock()
		vendorDetailAfterUpsertHooks = append(vendorDetailAfterUpsertHooks, vendorDetailHook)
		vendorDetailAfterUpsertMu.Unlock()
	}
}

// One returns a single vendorDetail record from the query.
func (q vendorDetailQuery) One(ctx context.Context, exec boil.ContextExecutor) (*VendorDetail, error) {
	o := &VendorDetail{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for VendorDetail")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all VendorDetail records from the query.
func (q vendorDetailQuery) All(ctx context.Context, exec boil.ContextExecutor) (VendorDetailSlice, error) {
	var o []*VendorDetail

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to VendorDetail slice")
	}

	if len(vendorDetailAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all VendorDetail records in the query.
func (q vendorDetailQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count VendorDetail rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q vendorDetailQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if VendorDetail exists")
	}

	return count > 0, nil
}

// Vendor pointed to by the foreign key.
func (o *VendorDetail) Vendor(mods ...qm.QueryMod) vendorQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.VendorID),
	}

	queryMods = append(queryMods, mods...)

	return Vendors(queryMods...)
}

// LoadVendor allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (vendorDetailL) LoadVendor(ctx context.Context, e boil.ContextExecutor, singular bool, maybeVendorDetail interface{}, mods queries.Applicator) error {
	var slice []*VendorDetail
	var object *VendorDetail

	if singular {
		var ok bool
		object, ok = maybeVendorDetail.(*VendorDetail)
		if !ok {
			object = new(VendorDetail)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeVendorDetail)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeVendorDetail))
			}
		}
	} else {
		s, ok := maybeVendorDetail.(*[]*VendorDetail)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeVendorDetail)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeVendorDetail))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &vendorDetailR{}
		}
		args[object.VendorID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &vendorDetailR{}
			}

			args[obj.VendorID] = struct{}{}

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
		qm.From(`Vendors`),
		qm.WhereIn(`Vendors.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Vendor")
	}

	var resultSlice []*Vendor
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Vendor")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for Vendors")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for Vendors")
	}

	if len(vendorAfterSelectHooks) != 0 {
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
		object.R.Vendor = foreign
		if foreign.R == nil {
			foreign.R = &vendorR{}
		}
		foreign.R.VendorVendorDetail = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.VendorID == foreign.ID {
				local.R.Vendor = foreign
				if foreign.R == nil {
					foreign.R = &vendorR{}
				}
				foreign.R.VendorVendorDetail = local
				break
			}
		}
	}

	return nil
}

// SetVendor of the vendorDetail to the related item.
// Sets o.R.Vendor to related.
// Adds o to related.R.VendorVendorDetail.
func (o *VendorDetail) SetVendor(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Vendor) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"VendorDetail\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"vendor_id"}),
		strmangle.WhereClause("\"", "\"", 2, vendorDetailPrimaryKeyColumns),
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

	o.VendorID = related.ID
	if o.R == nil {
		o.R = &vendorDetailR{
			Vendor: related,
		}
	} else {
		o.R.Vendor = related
	}

	if related.R == nil {
		related.R = &vendorR{
			VendorVendorDetail: o,
		}
	} else {
		related.R.VendorVendorDetail = o
	}

	return nil
}

// VendorDetails retrieves all the records using an executor.
func VendorDetails(mods ...qm.QueryMod) vendorDetailQuery {
	mods = append(mods, qm.From("\"VendorDetail\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"VendorDetail\".*"})
	}

	return vendorDetailQuery{q}
}

// FindVendorDetail retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindVendorDetail(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*VendorDetail, error) {
	vendorDetailObj := &VendorDetail{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"VendorDetail\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, vendorDetailObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from VendorDetail")
	}

	if err = vendorDetailObj.doAfterSelectHooks(ctx, exec); err != nil {
		return vendorDetailObj, err
	}

	return vendorDetailObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *VendorDetail) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no VendorDetail provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(vendorDetailColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	vendorDetailInsertCacheMut.RLock()
	cache, cached := vendorDetailInsertCache[key]
	vendorDetailInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			vendorDetailAllColumns,
			vendorDetailColumnsWithDefault,
			vendorDetailColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(vendorDetailType, vendorDetailMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(vendorDetailType, vendorDetailMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"VendorDetail\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"VendorDetail\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into VendorDetail")
	}

	if !cached {
		vendorDetailInsertCacheMut.Lock()
		vendorDetailInsertCache[key] = cache
		vendorDetailInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the VendorDetail.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *VendorDetail) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	vendorDetailUpdateCacheMut.RLock()
	cache, cached := vendorDetailUpdateCache[key]
	vendorDetailUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			vendorDetailAllColumns,
			vendorDetailPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update VendorDetail, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"VendorDetail\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, vendorDetailPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(vendorDetailType, vendorDetailMapping, append(wl, vendorDetailPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update VendorDetail row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for VendorDetail")
	}

	if !cached {
		vendorDetailUpdateCacheMut.Lock()
		vendorDetailUpdateCache[key] = cache
		vendorDetailUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q vendorDetailQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for VendorDetail")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for VendorDetail")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o VendorDetailSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), vendorDetailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"VendorDetail\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, vendorDetailPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in vendorDetail slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all vendorDetail")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *VendorDetail) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no VendorDetail provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(vendorDetailColumnsWithDefault, o)

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

	vendorDetailUpsertCacheMut.RLock()
	cache, cached := vendorDetailUpsertCache[key]
	vendorDetailUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			vendorDetailAllColumns,
			vendorDetailColumnsWithDefault,
			vendorDetailColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			vendorDetailAllColumns,
			vendorDetailPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert VendorDetail, could not build update column list")
		}

		ret := strmangle.SetComplement(vendorDetailAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(vendorDetailPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert VendorDetail, could not build conflict column list")
			}

			conflict = make([]string, len(vendorDetailPrimaryKeyColumns))
			copy(conflict, vendorDetailPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"VendorDetail\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(vendorDetailType, vendorDetailMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(vendorDetailType, vendorDetailMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert VendorDetail")
	}

	if !cached {
		vendorDetailUpsertCacheMut.Lock()
		vendorDetailUpsertCache[key] = cache
		vendorDetailUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single VendorDetail record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *VendorDetail) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no VendorDetail provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), vendorDetailPrimaryKeyMapping)
	sql := "DELETE FROM \"VendorDetail\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from VendorDetail")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for VendorDetail")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q vendorDetailQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no vendorDetailQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from VendorDetail")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for VendorDetail")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o VendorDetailSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(vendorDetailBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), vendorDetailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"VendorDetail\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, vendorDetailPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from vendorDetail slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for VendorDetail")
	}

	if len(vendorDetailAfterDeleteHooks) != 0 {
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
func (o *VendorDetail) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindVendorDetail(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *VendorDetailSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := VendorDetailSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), vendorDetailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"VendorDetail\".* FROM \"VendorDetail\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, vendorDetailPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in VendorDetailSlice")
	}

	*o = slice

	return nil
}

// VendorDetailExists checks if the VendorDetail row exists.
func VendorDetailExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"VendorDetail\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if VendorDetail exists")
	}

	return exists, nil
}

// Exists checks if the VendorDetail row exists.
func (o *VendorDetail) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return VendorDetailExists(ctx, exec, o.ID)
}
