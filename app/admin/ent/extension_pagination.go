// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"slash-admin/app/admin/ent/sysapi"
	"slash-admin/app/admin/ent/sysdictionary"
	"slash-admin/app/admin/ent/sysdictionarydetail"
	"slash-admin/app/admin/ent/sysmenu"
	"slash-admin/app/admin/ent/sysmenuparam"
	"slash-admin/app/admin/ent/sysoauthprovider"
	"slash-admin/app/admin/ent/sysrole"
	"slash-admin/app/admin/ent/systoken"
	"slash-admin/app/admin/ent/sysuser"
)

const errInvalidPage = "INVALID_PAGE"

const (
	listField     = "list"
	pageNumField  = "pageNum"
	pageSizeField = "pageSize"
)

type PageDetails struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

// Cursor of an edge type.
type Cursor struct {
	ID    uint64
	Value Value
}

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

func (o OrderDirection) orderFunc(field string) OrderFunc {
	if o == OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

const errInvalidPagination = "INVALID_PAGINATION"

type sysapiPager struct {
	order  *SysApiOrder
	filter func(*SysApiQuery) (*SysApiQuery, error)
}

// SysApiPaginateOption enables pagination customization.
type SysApiPaginateOption func(*sysapiPager) error

// SysApiOrder defines the ordering of SysApi.
type SysApiOrder struct {
	Direction OrderDirection    `json:"direction"`
	Field     *SysApiOrderField `json:"field"`
}

// SysApiOrderField defines the ordering field of SysApi.
type SysApiOrderField struct {
	field    string
	toCursor func(*SysApi) Cursor
}

// DefaultSysApiOrder is the default ordering of SysApi.
var DefaultSysApiOrder = &SysApiOrder{
	Direction: OrderDirectionAsc,
	Field: &SysApiOrderField{
		field: sysapi.FieldID,
		toCursor: func(sa *SysApi) Cursor {
			return Cursor{ID: sa.ID}
		},
	},
}

func newSysApiPager(opts []SysApiPaginateOption) (*sysapiPager, error) {
	pager := &sysapiPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultSysApiOrder
	}
	return pager, nil
}

func (p *sysapiPager) applyFilter(query *SysApiQuery) (*SysApiQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

// SysApiPageList is SysApi PageList result.
type SysApiPageList struct {
	List        []*SysApi    `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (sa *SysApiQuery) Page(
	ctx context.Context, pageNum int, pageSize int, opts ...SysApiPaginateOption,
) (*SysApiPageList, error) {

	pager, err := newSysApiPager(opts)
	if err != nil {
		return nil, err
	}

	if sa, err = pager.applyFilter(sa); err != nil {
		return nil, err
	}

	ret := &SysApiPageList{}

	ret.PageDetails = &PageDetails{
		Page:  pageNum,
		Limit: pageSize,
	}

	count, err := sa.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = count

	direction := pager.order.Direction
	sa = sa.Order(direction.orderFunc(pager.order.Field.field))
	if pager.order.Field != DefaultSysApiOrder.Field {
		sa = sa.Order(direction.orderFunc(DefaultSysApiOrder.Field.field))
	}

	sa = sa.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	list, err := sa.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type sysdictionaryPager struct {
	order  *SysDictionaryOrder
	filter func(*SysDictionaryQuery) (*SysDictionaryQuery, error)
}

// SysDictionaryPaginateOption enables pagination customization.
type SysDictionaryPaginateOption func(*sysdictionaryPager) error

// SysDictionaryOrder defines the ordering of SysDictionary.
type SysDictionaryOrder struct {
	Direction OrderDirection           `json:"direction"`
	Field     *SysDictionaryOrderField `json:"field"`
}

// SysDictionaryOrderField defines the ordering field of SysDictionary.
type SysDictionaryOrderField struct {
	field    string
	toCursor func(*SysDictionary) Cursor
}

// DefaultSysDictionaryOrder is the default ordering of SysDictionary.
var DefaultSysDictionaryOrder = &SysDictionaryOrder{
	Direction: OrderDirectionAsc,
	Field: &SysDictionaryOrderField{
		field: sysdictionary.FieldID,
		toCursor: func(sd *SysDictionary) Cursor {
			return Cursor{ID: sd.ID}
		},
	},
}

func newSysDictionaryPager(opts []SysDictionaryPaginateOption) (*sysdictionaryPager, error) {
	pager := &sysdictionaryPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultSysDictionaryOrder
	}
	return pager, nil
}

func (p *sysdictionaryPager) applyFilter(query *SysDictionaryQuery) (*SysDictionaryQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

// SysDictionaryPageList is SysDictionary PageList result.
type SysDictionaryPageList struct {
	List        []*SysDictionary `json:"list"`
	PageDetails *PageDetails     `json:"pageDetails"`
}

func (sd *SysDictionaryQuery) Page(
	ctx context.Context, pageNum int, pageSize int, opts ...SysDictionaryPaginateOption,
) (*SysDictionaryPageList, error) {

	pager, err := newSysDictionaryPager(opts)
	if err != nil {
		return nil, err
	}

	if sd, err = pager.applyFilter(sd); err != nil {
		return nil, err
	}

	ret := &SysDictionaryPageList{}

	ret.PageDetails = &PageDetails{
		Page:  pageNum,
		Limit: pageSize,
	}

	count, err := sd.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = count

	direction := pager.order.Direction
	sd = sd.Order(direction.orderFunc(pager.order.Field.field))
	if pager.order.Field != DefaultSysDictionaryOrder.Field {
		sd = sd.Order(direction.orderFunc(DefaultSysDictionaryOrder.Field.field))
	}

	sd = sd.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	list, err := sd.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type sysdictionarydetailPager struct {
	order  *SysDictionaryDetailOrder
	filter func(*SysDictionaryDetailQuery) (*SysDictionaryDetailQuery, error)
}

// SysDictionaryDetailPaginateOption enables pagination customization.
type SysDictionaryDetailPaginateOption func(*sysdictionarydetailPager) error

// SysDictionaryDetailOrder defines the ordering of SysDictionaryDetail.
type SysDictionaryDetailOrder struct {
	Direction OrderDirection                 `json:"direction"`
	Field     *SysDictionaryDetailOrderField `json:"field"`
}

// SysDictionaryDetailOrderField defines the ordering field of SysDictionaryDetail.
type SysDictionaryDetailOrderField struct {
	field    string
	toCursor func(*SysDictionaryDetail) Cursor
}

// DefaultSysDictionaryDetailOrder is the default ordering of SysDictionaryDetail.
var DefaultSysDictionaryDetailOrder = &SysDictionaryDetailOrder{
	Direction: OrderDirectionAsc,
	Field: &SysDictionaryDetailOrderField{
		field: sysdictionarydetail.FieldID,
		toCursor: func(sdd *SysDictionaryDetail) Cursor {
			return Cursor{ID: sdd.ID}
		},
	},
}

func newSysDictionaryDetailPager(opts []SysDictionaryDetailPaginateOption) (*sysdictionarydetailPager, error) {
	pager := &sysdictionarydetailPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultSysDictionaryDetailOrder
	}
	return pager, nil
}

func (p *sysdictionarydetailPager) applyFilter(query *SysDictionaryDetailQuery) (*SysDictionaryDetailQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

// SysDictionaryDetailPageList is SysDictionaryDetail PageList result.
type SysDictionaryDetailPageList struct {
	List        []*SysDictionaryDetail `json:"list"`
	PageDetails *PageDetails           `json:"pageDetails"`
}

func (sdd *SysDictionaryDetailQuery) Page(
	ctx context.Context, pageNum int, pageSize int, opts ...SysDictionaryDetailPaginateOption,
) (*SysDictionaryDetailPageList, error) {

	pager, err := newSysDictionaryDetailPager(opts)
	if err != nil {
		return nil, err
	}

	if sdd, err = pager.applyFilter(sdd); err != nil {
		return nil, err
	}

	ret := &SysDictionaryDetailPageList{}

	ret.PageDetails = &PageDetails{
		Page:  pageNum,
		Limit: pageSize,
	}

	count, err := sdd.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = count

	direction := pager.order.Direction
	sdd = sdd.Order(direction.orderFunc(pager.order.Field.field))
	if pager.order.Field != DefaultSysDictionaryDetailOrder.Field {
		sdd = sdd.Order(direction.orderFunc(DefaultSysDictionaryDetailOrder.Field.field))
	}

	sdd = sdd.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	list, err := sdd.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type sysmenuPager struct {
	order  *SysMenuOrder
	filter func(*SysMenuQuery) (*SysMenuQuery, error)
}

// SysMenuPaginateOption enables pagination customization.
type SysMenuPaginateOption func(*sysmenuPager) error

// SysMenuOrder defines the ordering of SysMenu.
type SysMenuOrder struct {
	Direction OrderDirection     `json:"direction"`
	Field     *SysMenuOrderField `json:"field"`
}

// SysMenuOrderField defines the ordering field of SysMenu.
type SysMenuOrderField struct {
	field    string
	toCursor func(*SysMenu) Cursor
}

// DefaultSysMenuOrder is the default ordering of SysMenu.
var DefaultSysMenuOrder = &SysMenuOrder{
	Direction: OrderDirectionAsc,
	Field: &SysMenuOrderField{
		field: sysmenu.FieldID,
		toCursor: func(sm *SysMenu) Cursor {
			return Cursor{ID: sm.ID}
		},
	},
}

func newSysMenuPager(opts []SysMenuPaginateOption) (*sysmenuPager, error) {
	pager := &sysmenuPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultSysMenuOrder
	}
	return pager, nil
}

func (p *sysmenuPager) applyFilter(query *SysMenuQuery) (*SysMenuQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

// SysMenuPageList is SysMenu PageList result.
type SysMenuPageList struct {
	List        []*SysMenu   `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (sm *SysMenuQuery) Page(
	ctx context.Context, pageNum int, pageSize int, opts ...SysMenuPaginateOption,
) (*SysMenuPageList, error) {

	pager, err := newSysMenuPager(opts)
	if err != nil {
		return nil, err
	}

	if sm, err = pager.applyFilter(sm); err != nil {
		return nil, err
	}

	ret := &SysMenuPageList{}

	ret.PageDetails = &PageDetails{
		Page:  pageNum,
		Limit: pageSize,
	}

	count, err := sm.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = count

	direction := pager.order.Direction
	sm = sm.Order(direction.orderFunc(pager.order.Field.field))
	if pager.order.Field != DefaultSysMenuOrder.Field {
		sm = sm.Order(direction.orderFunc(DefaultSysMenuOrder.Field.field))
	}

	sm = sm.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	list, err := sm.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type sysmenuparamPager struct {
	order  *SysMenuParamOrder
	filter func(*SysMenuParamQuery) (*SysMenuParamQuery, error)
}

// SysMenuParamPaginateOption enables pagination customization.
type SysMenuParamPaginateOption func(*sysmenuparamPager) error

// SysMenuParamOrder defines the ordering of SysMenuParam.
type SysMenuParamOrder struct {
	Direction OrderDirection          `json:"direction"`
	Field     *SysMenuParamOrderField `json:"field"`
}

// SysMenuParamOrderField defines the ordering field of SysMenuParam.
type SysMenuParamOrderField struct {
	field    string
	toCursor func(*SysMenuParam) Cursor
}

// DefaultSysMenuParamOrder is the default ordering of SysMenuParam.
var DefaultSysMenuParamOrder = &SysMenuParamOrder{
	Direction: OrderDirectionAsc,
	Field: &SysMenuParamOrderField{
		field: sysmenuparam.FieldID,
		toCursor: func(smp *SysMenuParam) Cursor {
			return Cursor{ID: smp.ID}
		},
	},
}

func newSysMenuParamPager(opts []SysMenuParamPaginateOption) (*sysmenuparamPager, error) {
	pager := &sysmenuparamPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultSysMenuParamOrder
	}
	return pager, nil
}

func (p *sysmenuparamPager) applyFilter(query *SysMenuParamQuery) (*SysMenuParamQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

// SysMenuParamPageList is SysMenuParam PageList result.
type SysMenuParamPageList struct {
	List        []*SysMenuParam `json:"list"`
	PageDetails *PageDetails    `json:"pageDetails"`
}

func (smp *SysMenuParamQuery) Page(
	ctx context.Context, pageNum int, pageSize int, opts ...SysMenuParamPaginateOption,
) (*SysMenuParamPageList, error) {

	pager, err := newSysMenuParamPager(opts)
	if err != nil {
		return nil, err
	}

	if smp, err = pager.applyFilter(smp); err != nil {
		return nil, err
	}

	ret := &SysMenuParamPageList{}

	ret.PageDetails = &PageDetails{
		Page:  pageNum,
		Limit: pageSize,
	}

	count, err := smp.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = count

	direction := pager.order.Direction
	smp = smp.Order(direction.orderFunc(pager.order.Field.field))
	if pager.order.Field != DefaultSysMenuParamOrder.Field {
		smp = smp.Order(direction.orderFunc(DefaultSysMenuParamOrder.Field.field))
	}

	smp = smp.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	list, err := smp.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type sysoauthproviderPager struct {
	order  *SysOauthProviderOrder
	filter func(*SysOauthProviderQuery) (*SysOauthProviderQuery, error)
}

// SysOauthProviderPaginateOption enables pagination customization.
type SysOauthProviderPaginateOption func(*sysoauthproviderPager) error

// SysOauthProviderOrder defines the ordering of SysOauthProvider.
type SysOauthProviderOrder struct {
	Direction OrderDirection              `json:"direction"`
	Field     *SysOauthProviderOrderField `json:"field"`
}

// SysOauthProviderOrderField defines the ordering field of SysOauthProvider.
type SysOauthProviderOrderField struct {
	field    string
	toCursor func(*SysOauthProvider) Cursor
}

// DefaultSysOauthProviderOrder is the default ordering of SysOauthProvider.
var DefaultSysOauthProviderOrder = &SysOauthProviderOrder{
	Direction: OrderDirectionAsc,
	Field: &SysOauthProviderOrderField{
		field: sysoauthprovider.FieldID,
		toCursor: func(sop *SysOauthProvider) Cursor {
			return Cursor{ID: sop.ID}
		},
	},
}

func newSysOauthProviderPager(opts []SysOauthProviderPaginateOption) (*sysoauthproviderPager, error) {
	pager := &sysoauthproviderPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultSysOauthProviderOrder
	}
	return pager, nil
}

func (p *sysoauthproviderPager) applyFilter(query *SysOauthProviderQuery) (*SysOauthProviderQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

// SysOauthProviderPageList is SysOauthProvider PageList result.
type SysOauthProviderPageList struct {
	List        []*SysOauthProvider `json:"list"`
	PageDetails *PageDetails        `json:"pageDetails"`
}

func (sop *SysOauthProviderQuery) Page(
	ctx context.Context, pageNum int, pageSize int, opts ...SysOauthProviderPaginateOption,
) (*SysOauthProviderPageList, error) {

	pager, err := newSysOauthProviderPager(opts)
	if err != nil {
		return nil, err
	}

	if sop, err = pager.applyFilter(sop); err != nil {
		return nil, err
	}

	ret := &SysOauthProviderPageList{}

	ret.PageDetails = &PageDetails{
		Page:  pageNum,
		Limit: pageSize,
	}

	count, err := sop.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = count

	direction := pager.order.Direction
	sop = sop.Order(direction.orderFunc(pager.order.Field.field))
	if pager.order.Field != DefaultSysOauthProviderOrder.Field {
		sop = sop.Order(direction.orderFunc(DefaultSysOauthProviderOrder.Field.field))
	}

	sop = sop.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	list, err := sop.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type sysrolePager struct {
	order  *SysRoleOrder
	filter func(*SysRoleQuery) (*SysRoleQuery, error)
}

// SysRolePaginateOption enables pagination customization.
type SysRolePaginateOption func(*sysrolePager) error

// SysRoleOrder defines the ordering of SysRole.
type SysRoleOrder struct {
	Direction OrderDirection     `json:"direction"`
	Field     *SysRoleOrderField `json:"field"`
}

// SysRoleOrderField defines the ordering field of SysRole.
type SysRoleOrderField struct {
	field    string
	toCursor func(*SysRole) Cursor
}

// DefaultSysRoleOrder is the default ordering of SysRole.
var DefaultSysRoleOrder = &SysRoleOrder{
	Direction: OrderDirectionAsc,
	Field: &SysRoleOrderField{
		field: sysrole.FieldID,
		toCursor: func(sr *SysRole) Cursor {
			return Cursor{ID: sr.ID}
		},
	},
}

func newSysRolePager(opts []SysRolePaginateOption) (*sysrolePager, error) {
	pager := &sysrolePager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultSysRoleOrder
	}
	return pager, nil
}

func (p *sysrolePager) applyFilter(query *SysRoleQuery) (*SysRoleQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

// SysRolePageList is SysRole PageList result.
type SysRolePageList struct {
	List        []*SysRole   `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (sr *SysRoleQuery) Page(
	ctx context.Context, pageNum int, pageSize int, opts ...SysRolePaginateOption,
) (*SysRolePageList, error) {

	pager, err := newSysRolePager(opts)
	if err != nil {
		return nil, err
	}

	if sr, err = pager.applyFilter(sr); err != nil {
		return nil, err
	}

	ret := &SysRolePageList{}

	ret.PageDetails = &PageDetails{
		Page:  pageNum,
		Limit: pageSize,
	}

	count, err := sr.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = count

	direction := pager.order.Direction
	sr = sr.Order(direction.orderFunc(pager.order.Field.field))
	if pager.order.Field != DefaultSysRoleOrder.Field {
		sr = sr.Order(direction.orderFunc(DefaultSysRoleOrder.Field.field))
	}

	sr = sr.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	list, err := sr.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type systokenPager struct {
	order  *SysTokenOrder
	filter func(*SysTokenQuery) (*SysTokenQuery, error)
}

// SysTokenPaginateOption enables pagination customization.
type SysTokenPaginateOption func(*systokenPager) error

// SysTokenOrder defines the ordering of SysToken.
type SysTokenOrder struct {
	Direction OrderDirection      `json:"direction"`
	Field     *SysTokenOrderField `json:"field"`
}

// SysTokenOrderField defines the ordering field of SysToken.
type SysTokenOrderField struct {
	field    string
	toCursor func(*SysToken) Cursor
}

// DefaultSysTokenOrder is the default ordering of SysToken.
var DefaultSysTokenOrder = &SysTokenOrder{
	Direction: OrderDirectionAsc,
	Field: &SysTokenOrderField{
		field: systoken.FieldID,
		toCursor: func(st *SysToken) Cursor {
			return Cursor{ID: st.ID}
		},
	},
}

func newSysTokenPager(opts []SysTokenPaginateOption) (*systokenPager, error) {
	pager := &systokenPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultSysTokenOrder
	}
	return pager, nil
}

func (p *systokenPager) applyFilter(query *SysTokenQuery) (*SysTokenQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

// SysTokenPageList is SysToken PageList result.
type SysTokenPageList struct {
	List        []*SysToken  `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (st *SysTokenQuery) Page(
	ctx context.Context, pageNum int, pageSize int, opts ...SysTokenPaginateOption,
) (*SysTokenPageList, error) {

	pager, err := newSysTokenPager(opts)
	if err != nil {
		return nil, err
	}

	if st, err = pager.applyFilter(st); err != nil {
		return nil, err
	}

	ret := &SysTokenPageList{}

	ret.PageDetails = &PageDetails{
		Page:  pageNum,
		Limit: pageSize,
	}

	count, err := st.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = count

	direction := pager.order.Direction
	st = st.Order(direction.orderFunc(pager.order.Field.field))
	if pager.order.Field != DefaultSysTokenOrder.Field {
		st = st.Order(direction.orderFunc(DefaultSysTokenOrder.Field.field))
	}

	st = st.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	list, err := st.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type sysuserPager struct {
	order  *SysUserOrder
	filter func(*SysUserQuery) (*SysUserQuery, error)
}

// SysUserPaginateOption enables pagination customization.
type SysUserPaginateOption func(*sysuserPager) error

// SysUserOrder defines the ordering of SysUser.
type SysUserOrder struct {
	Direction OrderDirection     `json:"direction"`
	Field     *SysUserOrderField `json:"field"`
}

// SysUserOrderField defines the ordering field of SysUser.
type SysUserOrderField struct {
	field    string
	toCursor func(*SysUser) Cursor
}

// DefaultSysUserOrder is the default ordering of SysUser.
var DefaultSysUserOrder = &SysUserOrder{
	Direction: OrderDirectionAsc,
	Field: &SysUserOrderField{
		field: sysuser.FieldID,
		toCursor: func(su *SysUser) Cursor {
			return Cursor{ID: su.ID}
		},
	},
}

func newSysUserPager(opts []SysUserPaginateOption) (*sysuserPager, error) {
	pager := &sysuserPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultSysUserOrder
	}
	return pager, nil
}

func (p *sysuserPager) applyFilter(query *SysUserQuery) (*SysUserQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

// SysUserPageList is SysUser PageList result.
type SysUserPageList struct {
	List        []*SysUser   `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (su *SysUserQuery) Page(
	ctx context.Context, pageNum int, pageSize int, opts ...SysUserPaginateOption,
) (*SysUserPageList, error) {

	pager, err := newSysUserPager(opts)
	if err != nil {
		return nil, err
	}

	if su, err = pager.applyFilter(su); err != nil {
		return nil, err
	}

	ret := &SysUserPageList{}

	ret.PageDetails = &PageDetails{
		Page:  pageNum,
		Limit: pageSize,
	}

	count, err := su.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = count

	direction := pager.order.Direction
	su = su.Order(direction.orderFunc(pager.order.Field.field))
	if pager.order.Field != DefaultSysUserOrder.Field {
		su = su.Order(direction.orderFunc(DefaultSysUserOrder.Field.field))
	}

	su = su.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	list, err := su.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}
