// Code generated by ent, DO NOT EDIT.

package ent

import (
	"slash-admin/pkg/types"
	"time"
)

type CreateSysApiInput struct {
	ID          *uint64
	Path        *string
	Description *string
	APIGroup    *string
	Method      *string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}

type CreateSysDictionaryInput struct {
	ID        *uint64
	Title     *string
	Name      *string
	Status    *types.Status
	Desc      *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type CreateSysDictionaryDetailInput struct {
	ID           *uint64
	Title        *string
	Key          *string
	Value        *string
	Status       *types.Status
	DictionaryID *uint64
	Remark       *string
	OrderNo      *uint32
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	DeletedAt    *time.Time
}

type CreateSysMenuInput struct {
	ID        *uint64
	ParentID  *uint64
	MenuLevel *uint8
	MenuType  *uint8
	Path      *string
	Name      *string
	Redirect  *string
	Component *string
	OrderNo   *uint8
	Disabled  *bool
	Meta      *types.MenuMeta
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type CreateSysMenuParamInput struct {
	ID        *uint64
	MenuID    *uint64
	Type      *string
	Key       *string
	Value     *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type CreateSysOauthProviderInput struct {
	ID           *uint64
	Name         *string
	ClientID     *string
	ClientSecret *string
	RedirectURL  *string
	Scopes       *string
	AuthURL      *string
	TokenURL     *string
	AuthStyle    *uint8
	InfoURL      *string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	DeletedAt    *time.Time
}

type CreateSysRoleInput struct {
	ID            *uint64
	Name          *string
	Value         *string
	DefaultRouter *string
	Status        *types.Status
	Remark        *string
	OrderNo       *uint32
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	DeletedAt     *time.Time
}

type CreateSysTokenInput struct {
	ID        *uint64
	UUID      *string
	Token     *string
	Source    *string
	Status    *types.Status
	ExpiredAt *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type CreateSysUserInput struct {
	ID          *uint64
	UUID        *string
	Username    *string
	Password    *string
	Nickname    *string
	SideMode    *string
	Avatar      *string
	BaseColor   *string
	ActiveColor *string
	RoleID      *uint64
	Mobile      *string
	Email       *string
	Status      *types.Status
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}
