// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

	"emperror.dev/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
	"github.com/volatiletech/sqlboiler/types"
)

// CoreConfig is an object representing the database table.
type CoreConfig struct {
	GuildID                 int64            `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	AllowedReadOnlyRoles    types.Int64Array `boil:"allowed_read_only_roles" json:"allowed_read_only_roles,omitempty" toml:"allowed_read_only_roles" yaml:"allowed_read_only_roles,omitempty"`
	AllowedWriteRoles       types.Int64Array `boil:"allowed_write_roles" json:"allowed_write_roles,omitempty" toml:"allowed_write_roles" yaml:"allowed_write_roles,omitempty"`
	AllowAllMembersReadOnly bool             `boil:"allow_all_members_read_only" json:"allow_all_members_read_only" toml:"allow_all_members_read_only" yaml:"allow_all_members_read_only"`
	AllowNonMembersReadOnly bool             `boil:"allow_non_members_read_only" json:"allow_non_members_read_only" toml:"allow_non_members_read_only" yaml:"allow_non_members_read_only"`

	R *coreConfigR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L coreConfigL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CoreConfigColumns = struct {
	GuildID                 string
	AllowedReadOnlyRoles    string
	AllowedWriteRoles       string
	AllowAllMembersReadOnly string
	AllowNonMembersReadOnly string
}{
	GuildID:                 "guild_id",
	AllowedReadOnlyRoles:    "allowed_read_only_roles",
	AllowedWriteRoles:       "allowed_write_roles",
	AllowAllMembersReadOnly: "allow_all_members_read_only",
	AllowNonMembersReadOnly: "allow_non_members_read_only",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpertypes_Int64Array struct{ field string }

func (w whereHelpertypes_Int64Array) EQ(x types.Int64Array) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpertypes_Int64Array) NEQ(x types.Int64Array) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpertypes_Int64Array) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpertypes_Int64Array) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpertypes_Int64Array) LT(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_Int64Array) LTE(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_Int64Array) GT(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_Int64Array) GTE(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var CoreConfigWhere = struct {
	GuildID                 whereHelperint64
	AllowedReadOnlyRoles    whereHelpertypes_Int64Array
	AllowedWriteRoles       whereHelpertypes_Int64Array
	AllowAllMembersReadOnly whereHelperbool
	AllowNonMembersReadOnly whereHelperbool
}{
	GuildID:                 whereHelperint64{field: "\"core_configs\".\"guild_id\""},
	AllowedReadOnlyRoles:    whereHelpertypes_Int64Array{field: "\"core_configs\".\"allowed_read_only_roles\""},
	AllowedWriteRoles:       whereHelpertypes_Int64Array{field: "\"core_configs\".\"allowed_write_roles\""},
	AllowAllMembersReadOnly: whereHelperbool{field: "\"core_configs\".\"allow_all_members_read_only\""},
	AllowNonMembersReadOnly: whereHelperbool{field: "\"core_configs\".\"allow_non_members_read_only\""},
}

// CoreConfigRels is where relationship names are stored.
var CoreConfigRels = struct {
}{}

// coreConfigR is where relationships are stored.
type coreConfigR struct {
}

// NewStruct creates a new relationship struct
func (*coreConfigR) NewStruct() *coreConfigR {
	return &coreConfigR{}
}

// coreConfigL is where Load methods for each relationship are stored.
type coreConfigL struct{}

var (
	coreConfigAllColumns            = []string{"guild_id", "allowed_read_only_roles", "allowed_write_roles", "allow_all_members_read_only", "allow_non_members_read_only"}
	coreConfigColumnsWithoutDefault = []string{"guild_id", "allowed_read_only_roles", "allowed_write_roles", "allow_all_members_read_only", "allow_non_members_read_only"}
	coreConfigColumnsWithDefault    = []string{}
	coreConfigPrimaryKeyColumns     = []string{"guild_id"}
)

type (
	// CoreConfigSlice is an alias for a slice of pointers to CoreConfig.
	// This should generally be used opposed to []CoreConfig.
	CoreConfigSlice []*CoreConfig

	coreConfigQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	coreConfigType                 = reflect.TypeOf(&CoreConfig{})
	coreConfigMapping              = queries.MakeStructMapping(coreConfigType)
	coreConfigPrimaryKeyMapping, _ = queries.BindMapping(coreConfigType, coreConfigMapping, coreConfigPrimaryKeyColumns)
	coreConfigInsertCacheMut       sync.RWMutex
	coreConfigInsertCache          = make(map[string]insertCache)
	coreConfigUpdateCacheMut       sync.RWMutex
	coreConfigUpdateCache          = make(map[string]updateCache)
	coreConfigUpsertCacheMut       sync.RWMutex
	coreConfigUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single coreConfig record from the query using the global executor.
func (q coreConfigQuery) OneG(ctx context.Context) (*CoreConfig, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single coreConfig record from the query.
func (q coreConfigQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CoreConfig, error) {
	o := &CoreConfig{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.WrapIf(err, "models: failed to execute a one query for core_configs")
	}

	return o, nil
}

// AllG returns all CoreConfig records from the query using the global executor.
func (q coreConfigQuery) AllG(ctx context.Context) (CoreConfigSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all CoreConfig records from the query.
func (q coreConfigQuery) All(ctx context.Context, exec boil.ContextExecutor) (CoreConfigSlice, error) {
	var o []*CoreConfig

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.WrapIf(err, "models: failed to assign all query results to CoreConfig slice")
	}

	return o, nil
}

// CountG returns the count of all CoreConfig records in the query, and panics on error.
func (q coreConfigQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all CoreConfig records in the query.
func (q coreConfigQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to count core_configs rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q coreConfigQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q coreConfigQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.WrapIf(err, "models: failed to check if core_configs exists")
	}

	return count > 0, nil
}

// CoreConfigs retrieves all the records using an executor.
func CoreConfigs(mods ...qm.QueryMod) coreConfigQuery {
	mods = append(mods, qm.From("\"core_configs\""))
	return coreConfigQuery{NewQuery(mods...)}
}

// FindCoreConfigG retrieves a single record by ID.
func FindCoreConfigG(ctx context.Context, guildID int64, selectCols ...string) (*CoreConfig, error) {
	return FindCoreConfig(ctx, boil.GetContextDB(), guildID, selectCols...)
}

// FindCoreConfig retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCoreConfig(ctx context.Context, exec boil.ContextExecutor, guildID int64, selectCols ...string) (*CoreConfig, error) {
	coreConfigObj := &CoreConfig{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"core_configs\" where \"guild_id\"=$1", sel,
	)

	q := queries.Raw(query, guildID)

	err := q.Bind(ctx, exec, coreConfigObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.WrapIf(err, "models: unable to select from core_configs")
	}

	return coreConfigObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *CoreConfig) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CoreConfig) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no core_configs provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(coreConfigColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	coreConfigInsertCacheMut.RLock()
	cache, cached := coreConfigInsertCache[key]
	coreConfigInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			coreConfigAllColumns,
			coreConfigColumnsWithDefault,
			coreConfigColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(coreConfigType, coreConfigMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(coreConfigType, coreConfigMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"core_configs\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"core_configs\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.WrapIf(err, "models: unable to insert into core_configs")
	}

	if !cached {
		coreConfigInsertCacheMut.Lock()
		coreConfigInsertCache[key] = cache
		coreConfigInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single CoreConfig record using the global executor.
// See Update for more documentation.
func (o *CoreConfig) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the CoreConfig.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CoreConfig) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	coreConfigUpdateCacheMut.RLock()
	cache, cached := coreConfigUpdateCache[key]
	coreConfigUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			coreConfigAllColumns,
			coreConfigPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update core_configs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"core_configs\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, coreConfigPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(coreConfigType, coreConfigMapping, append(wl, coreConfigPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to update core_configs row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to get rows affected by update for core_configs")
	}

	if !cached {
		coreConfigUpdateCacheMut.Lock()
		coreConfigUpdateCache[key] = cache
		coreConfigUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q coreConfigQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q coreConfigQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to update all for core_configs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to retrieve rows affected for core_configs")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o CoreConfigSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CoreConfigSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), coreConfigPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"core_configs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, coreConfigPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to update all in coreConfig slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to retrieve rows affected all in update all coreConfig")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *CoreConfig) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CoreConfig) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no core_configs provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(coreConfigColumnsWithDefault, o)

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

	coreConfigUpsertCacheMut.RLock()
	cache, cached := coreConfigUpsertCache[key]
	coreConfigUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			coreConfigAllColumns,
			coreConfigColumnsWithDefault,
			coreConfigColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			coreConfigAllColumns,
			coreConfigPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert core_configs, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(coreConfigPrimaryKeyColumns))
			copy(conflict, coreConfigPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"core_configs\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(coreConfigType, coreConfigMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(coreConfigType, coreConfigMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.WrapIf(err, "models: unable to upsert core_configs")
	}

	if !cached {
		coreConfigUpsertCacheMut.Lock()
		coreConfigUpsertCache[key] = cache
		coreConfigUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single CoreConfig record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *CoreConfig) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single CoreConfig record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CoreConfig) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CoreConfig provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), coreConfigPrimaryKeyMapping)
	sql := "DELETE FROM \"core_configs\" WHERE \"guild_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to delete from core_configs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to get rows affected by delete for core_configs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q coreConfigQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no coreConfigQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to delete all from core_configs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to get rows affected by deleteall for core_configs")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o CoreConfigSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CoreConfigSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), coreConfigPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"core_configs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, coreConfigPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to delete all from coreConfig slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to get rows affected by deleteall for core_configs")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *CoreConfig) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no CoreConfig provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CoreConfig) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCoreConfig(ctx, exec, o.GuildID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CoreConfigSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty CoreConfigSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CoreConfigSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CoreConfigSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), coreConfigPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"core_configs\".* FROM \"core_configs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, coreConfigPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.WrapIf(err, "models: unable to reload all in CoreConfigSlice")
	}

	*o = slice

	return nil
}

// CoreConfigExistsG checks if the CoreConfig row exists.
func CoreConfigExistsG(ctx context.Context, guildID int64) (bool, error) {
	return CoreConfigExists(ctx, boil.GetContextDB(), guildID)
}

// CoreConfigExists checks if the CoreConfig row exists.
func CoreConfigExists(ctx context.Context, exec boil.ContextExecutor, guildID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"core_configs\" where \"guild_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, guildID)
	}

	row := exec.QueryRowContext(ctx, sql, guildID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.WrapIf(err, "models: unable to check if core_configs exists")
	}

	return exists, nil
}
