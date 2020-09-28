package response

import (
	"anew-server/models"
)

// 角色返回权限信息
type RolePermsResp struct {
	Id    uint             `json:"id"`
	Name  string           `json:"name"`
	Menus []models.SysMenu `json:"menus"`
	Apis  []models.SysApi  `json:"apis"`
}

// 角色返回字符串权限信息
type RolePermsStrResp struct {
	Id    uint     `json:"id"`
	Name  string   `json:"name"`
	Menus []string `json:"menus"`
	Apis  []string `json:"apis"`
}

// 角色信息响应, 字段含义见models
type RoleListResp struct {
	Id        uint             `json:"id"`
	Key       string           `json:"key"`
	Name      string           `json:"name"`
	Title     string           `json:"title"`
	Keyword   string           `json:"keyword"`
	Desc      string           `json:"desc"`
	Status    *bool            `json:"status"`
	Creator   string           `json:"creator"`
	CreatedAt models.LocalTime `json:"createdAt"`
	Menus     []string         `json:"menus"`
	Apis      []string         `json:"apis"`
}
