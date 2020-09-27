package response

// 菜单树信息响应, 字段含义见models.SysMenu
type MenuTreeResp struct {
	Id       uint           `json:"id"`
	ParentId uint           `json:"parentId"`
	Key      string         `json:"key"`
	Name     string         `json:"name"`
	Title    string         `json:"title"`
	Icon     string         `json:"icon"`
	Path     string         `json:"path"`
	Creator  string         `json:"creator"`
	Sort     int            `json:"sort"`
	Status   bool           `json:"status"`
	Hidden   bool           `json:"hidden"`
	Children []MenuTreeResp `json:"children"`
}
type MenuTreeRespList []MenuTreeResp

func (hs MenuTreeRespList) Len() int {
	return len(hs)
}
func (hs MenuTreeRespList) Less(i, j int) bool {
	return hs[i].Sort < hs[j].Sort // 按Sort从小到大排序
}

func (hs MenuTreeRespList) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

// 菜单树信息响应, 包含有权限访问的id列表
type MenuTreeWithAccessResp struct {
	DataList  []MenuTreeResp `json:"dataList"`
	AccessIds []uint         `json:"accessIds"`
}