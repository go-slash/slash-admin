// Code generated by goctl. DO NOT EDIT.
package types

// The basic response without data | 基础不带数据信息
// swagger:response BaseMsgResp
type BaseMsgResp struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

// The SimpleMsgResp | 简单信息Resp
// swagger:response SimpleMsgResp
type SimpleMsgResp struct {
	// Message | 信息
	Msg string `json:"msg"`
}

// The page request parameters | 列表请求参数
// swagger:model PageReq
type PageReq struct {
	// Page number | 第几页
	// Minimum: 1
	PageNo uint64 `json:"pageNo,optional,default=1" validate:"number,min=1"`
	// Page size | 单页数据行数
	// Maximum: 100000
	PageSize uint64 `json:"pageSize,optional,default=10," validate:"number,max=100"`
}

// Basic id request | 基础id参数请求
// swagger:model IDReq
// swagger:model IDReq
type IDReq struct {
	// ID
	// Required: true
	ID uint `json:"id" validate:"number"`
}

// Basic id request | 基础id参数请求
// swagger:model IDPathReq
type IDPathReq struct {
	// ID
	// Required: true
	ID uint `path:"id"`
}

// Basic UUID request | 基础UUID参数请求
// swagger:model UUIDReq
type UUIDReq struct {
	// UUID
	// Required: true
	// Max length: 36
	UUID string `json:"UUID" validate:"len=36"`
}

// The base response data | 基础信息
// swagger:model BaseInfo
type BaseInfo struct {
	// ID
	ID uint `json:"id"`
	// Create date | 创建日期
	CreatedAt int64 `json:"createdAt"`
	// Update date | 更新日期
	UpdatedAt int64 `json:"updatedAt"`
	// Delete date | 删除日期
	DeletedAt int64 `json:"deletedAt"`
}

// The request params of setting boolean status | 设置状态参数
// swagger:model SetBooleanStatusReq
type SetBooleanStatusReq struct {
	// ID
	// Required: true
	Id uint64 `json:"id" validate:"number"`
	// Status code | 状态码
	// Required: true
	Status uint32 `json:"status" validate:"number"`
}

// Create or update role information params | 创建或更新角色信息参数
// swagger:model RoleInfo
type RoleInfo struct {
	// Role ID | 角色 ID
	// Required : true
	// Maximum: 1000
	Id uint64 `json:"id" validate:"number,max=1000"`
	// Role Name | 角色名
	// Required : true
	// Max length: 20
	Name string `json:"name" validate:"max=20"`
	// Role value | 角色值
	// Required : true
	// Max length: 10
	Value string `json:"value" validate:"max=10"`
	// Role's default page | 角色默认管理页面
	// Required : true
	// Max length: 20
	DefaultRouter string `json:"defaultRouter" validate:"max=50"`
	// Role status | 角色状态
	// Maximum: 10
	Status uint8 `json:"status,default=0" validate:"number,max=10"`
	// Role remark | 角色备注
	// Max length: 200
	Remark string `json:"remark,default=''" validate:"omitempty,max=200"`
	// Role's sorting number | 角色排序
	// Required : true
	// Maximum: 1000
	OrderNo uint32 `json:"orderNo" validate:"number,max=1000"`
}

// The response data of role list | 角色列表数据
// swagger:response RoleListResp
// swagger:response RoleListResp
type RoleListResp struct {
	// The total number of data | 数据总数
	Total uint64 `json:"total"`
	// The role list data | 角色列表数据
	// in: body
	Data []*RoleInfo `json:"data"`
}

// The request params of setting role status | 设置角色状态参数
// swagger:model SetStatusReq
type SetStatusReq struct {
	// ID
	// Required: true
	// Maximum: 1000
	Id uint64 `json:"id" validate:"number,max=1000"`
	// Status code | 状态码
	// Required: true
	// Maximum: 10
	Status uint8 `json:"status" validate:"number,max=10"`
}
