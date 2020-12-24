package component

import (
	"encoding/json"
	"fmt"
	"strings"
)

type CRUD struct {
	JsonStr []string
}

// "table" 、 "cards" 或者 "list"
func (p CRUD) Mode(tableCardsList string) CRUD {
	return p._KeyVal("mode", tableCardsList)
}

// 和表格标题重复
func (p CRUD) Title(title string) CRUD {
	return p._KeyVal("title", title)
}

// 表格外层 Dom 的类名
func (p CRUD) ClassName(className string) CRUD {
	return p._KeyVal("className", className)
}

// CRUD 用来获取列表数据的 api。
func (p CRUD) Api(listApiUrl string) CRUD {
	return p._KeyVal("api", listApiUrl)
}

// 是否一次性加载所有数据（前端分页）
func (p CRUD) LoadDataOnce(loadDataOnce bool) CRUD {
	return p._KeyVal("loadDataOnce", loadDataOnce)
}

// 在开启 loadDataOnce 时，filter 时是否去重新请求 api
func (p CRUD) LoadDataOnceFetchOnFilter(loadDataOnceFetchOnFilter bool) CRUD {
	return p._KeyVal("loadDataOnceFetchOnFilter", loadDataOnceFetchOnFilter)
}

// 数据映射接口返回某字段的值，不设置会默认把接口返回的items或者rows填充进mode区域
// 这里也和Table重复
func (p CRUD) Source(source string) CRUD {
	return p._KeyVal("source", source)
}

// 设置过滤器，当该表单提交后，会把数据带给当前 mode 刷新列表。
// 这里的过滤器就是一个Form
func (p CRUD) Filter(newForm string) CRUD {
	return p._KeyVal("filter", newForm)
}

// 是否可显隐过滤器
func (p CRUD) FilterTogglable(filterTogglable bool) CRUD {
	return p._KeyVal("filterTogglable", filterTogglable)
}

// 设置过滤器默认是否可见。
func (p CRUD) FilterDefaultVisible(filterDefaultVisible bool) CRUD {
	return p._KeyVal("filterDefaultVisible", filterDefaultVisible)
}

// 是否初始化的时候拉取数据, 只针对有 filter 的情况, 没有 filter 初始都会拉取数据。
func (p CRUD) InitFetch(initFetch bool) CRUD {
	return p._KeyVal("initFetch", initFetch)
}

// 刷新时间(最低 3000)
func (p CRUD) Interval(interval int) CRUD {
	return p._KeyVal("interval", interval)
}

// 配置刷新时是否隐藏加载动画
func (p CRUD) SilentPolling(silentPolling bool) CRUD {
	return p._KeyVal("silentPolling", silentPolling)
}

// 通过表达式来配置停止刷新的条件
func (p CRUD) StopAutoRefreshWhen(stopAutoRefreshWhen string) CRUD {
	return p._KeyVal("stopAutoRefreshWhen", stopAutoRefreshWhen)
}

// 当有弹框时关闭自动刷新，关闭弹框又恢复
func (p CRUD) StopAutoRefreshWhenModalIsOpen(stopAutoRefreshWhenModalIsOpen bool) CRUD {
	return p._KeyVal("stopAutoRefreshWhenModalIsOpen", stopAutoRefreshWhenModalIsOpen)
}

// 是否将过滤条件的参数同步到地址栏
func (p CRUD) SyncLocation(isFilterSyncLocation bool) CRUD {
	return p._KeyVal("syncLocation", isFilterSyncLocation)
}

// 是否可通过拖拽排序
func (p CRUD) Draggable(draggable bool) CRUD {
	return p._KeyVal("draggable", draggable)
}

// 用表达式来配置是否可拖拽排序
// 和Table重复
func (p CRUD) ItemDraggableOn(itemDraggableOn string) CRUD {
	return p._KeyVal("itemDraggableOn", itemDraggableOn)
}

// 保存排序的 api。
func (p CRUD) SaveOrderApi(saveSortApi string) CRUD {
	return p._KeyVal("saveOrderApi", saveSortApi)
}

// 快速编辑后用来批量保存的 API。
func (p CRUD) QuickSaveApi(quickSaveApi string) CRUD {
	return p._KeyVal("quickSaveApi", quickSaveApi)
}

// 快速编辑后用来批量保存的 API。
func (p CRUD) QuickSaveItemApi(quickSaveItemApi string) CRUD {
	return p._KeyVal("quickSaveItemApi", quickSaveItemApi)
}

// 批量操作列表，配置后，表格可进行选中操作。
func (p CRUD) BulkActions(newButton ...string) CRUD {
	ss := strings.Join(newButton, ",")
	p._NotEmpty("bulkActions", ss)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"bulkActions":[%s]`, ss))
	return p
}

// 当可批量操作时，默认是否全部勾选。
func (p CRUD) DefaultChecked(isMultiChecked bool) CRUD {
	return p._KeyVal("defaultChecked", isMultiChecked)
}

// 覆盖消息提示，如果不指定，将采用 api 返回的 message。
func (p CRUD) Messages(newMessages string) CRUD {
	p._NotEmpty("messages", newMessages)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"messages":[%s]`, newMessages))
	return p
}

// 设置 ID 字段名。
func (p CRUD) PrimaryField(primaryIDField string) CRUD {
	return p._KeyVal("primaryField", primaryIDField)
}

// 设置默认 filter 默认参数，会在查询的时候一起发给后端
func (p CRUD) DefaultParams(additionalFilterParams map[string]interface{}) CRUD {
	params, _ := json.Marshal(&additionalFilterParams)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"defaultParams":%s`, string(params)))
	return p
}

// 设置分页页码字段名。
func (p CRUD) PageField(pageField string) CRUD {
	return p._KeyVal("pageField", pageField)
}

// 设置一页显示多少条数据下拉框可选条数。
// [5, 10, 20, 50, 100]
func (p CRUD) PerPageAvailable(perPageAvailable []int) CRUD {
	if len(perPageAvailable) < 1 {
		return p
	}
	pp := struct {
		PerPageAvailable []int `json:"perPageAvailable"`
	}{
		PerPageAvailable: perPageAvailable,
	}
	params, _ := json.Marshal(&pp)
	p.JsonStr = append(p.JsonStr, string(params))
	return p
}

// 设置用来确定位置的字段名，设置后新的顺序将被赋值到该字段中。
func (p CRUD) OrderField(orderField string) CRUD {
	return p._KeyVal("orderField", orderField)
}

// 隐藏顶部快速保存提示
func (p CRUD) HideQuickSaveBtn(hideQuickSaveBtn bool) CRUD {
	return p._KeyVal("hideQuickSaveBtn", hideQuickSaveBtn)
}

// 当切分页的时候，是否自动跳顶部。
func (p CRUD) AutoJumpToTopOnPagerChange(autoJumpToTopOnPagerChange bool) CRUD {
	return p._KeyVal("autoJumpToTopOnPagerChange", autoJumpToTopOnPagerChange)
}

// 将返回数据同步到过滤器上。
func (p CRUD) SyncResponse2Query(syncResponse2Query bool) CRUD {
	return p._KeyVal("syncResponse2Query", syncResponse2Query)
}

// 保留条目选择，默认分页、搜素后，用户选择条目会被清空，开启此选项后会保留用户选择，可以实现跨页面批量操作。
func (p CRUD) KeepItemSelectionOnPageChange(keepItemSelectionOnPageChange bool) CRUD {
	return p._KeyVal("keepItemSelectionOnPageChange", keepItemSelectionOnPageChange)
}

// 单条描述模板，keepItemSelectionOnPageChange设置为true后会把所有已选择条目列出来，此选项可以用来定制条目展示文案。
func (p CRUD) LabelTpl(labelTpl string) CRUD {
	return p._KeyVal("labelTpl", labelTpl)
}

var (
	// 添加bulkActions字符串，并且在 crud 上配置bulkActions行为按钮数组，可以实现选中表格项并批量操作的功能。
	// 需要设置primaryField用于标识选中状态，配置当前行数据中的某一唯一标识字段，例如id，否则可能会出现无法选中的问题
	CRUDToolbarBulkActions = "bulkActions"
	// 添加statistics字符串，可以实现简单的数据统计功能
	CRUDToolbarStatistics = "statistics"
	// 添加pagination字符串，并且在数据源接口中返回了数据总数total，即可以渲染分页组件
	CRUDToolbarPagination = "pagination"
	// 添加switch-per-page字符串，可以渲染切换每页条数组件
	CRUDToolbarSwitchPerPage = "switch-per-page"
	// 添加load-more字符串，可以实现点击加载更多功能
	CRUDToolbarLoadMore = "load-more"
	// 添加export-csv字符串，可以实现点击下载 CSV 的功能，注意这里只包括当前分页的数据，要下载全部数据需要通过后端 API 实现
	CRUDToolbarExportCsv = "export-csv"
	// 添加export-excel字符串，可以实现点击下载 Excel 的功能，和导出 CSV 一样只包括当前分页的数据，但它们有明显区别：
	// 导出 CSV 是将 api 返回数据导出，表头是数据里的 key，而 Excel 的表头使用的是 label。
	// 导出 Excel 更重视展现一致，支持合并单元格、链接、mapping 映射、图片（需要加跨域 Header）。
	// 导出 Excel 只在 mode 为 table 时能用。
	CRUDToolbarExportExcel = "export-excel"
	// 添加filter-toggler字符串，并且在 crud 中配置"filterTogglable": true后，可以渲染一个可以切换显示查询表单的功能按钮
	CRUDToolbarFilterToggler = "filter-toggler"
)

// 顶部工具栏配置
// 分页
// 在headerToolbar或者footerToolbar数组中
func (p CRUD) HeaderToolbar(headerToolbar ...string) CRUD {
	ss := strings.Join(headerToolbar, ",")
	p._NotEmpty("headerToolbar", ss)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"headerToolbar":[%s]`, ss))
	return p
}

// 底部工具栏配置
// 分页
// 在headerToolbar或者footerToolbar数组中添加pagination字符串，并且在数据源接口中返回了数据总数count，即可以渲染分页组件；添加switch-per-page字符串，可以渲染切换每页条数组件
func (p CRUD) FooterToolbar(footerToolbar ...string) CRUD {
	ss := strings.Join(footerToolbar, ",")
	p._NotEmpty("footerToolbar", ss)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"footerToolbar":[%s]`, ss))
	return p
}

func (p CRUD) Build() string {
	return strings.Join(p.JsonStr, ",")
}

func (p CRUD) _KeyVal(k string, v interface{}) CRUD {
	p._NotEmpty(k, v)
	switch v := v.(type) {
	case string:
		p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"%s":"%s"`, k, v))
	case int:
		p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"%s":%d`, k, v))
	case bool:
		p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"%s":%t`, k, v))
	}
	return p
}

func (p CRUD) _NotEmpty(py, v interface{}) {
	switch v := v.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			panic(fmt.Errorf("Property:%s,Cannot hold empty values", py))
		}
	}
}

func NewCRUD() CRUD {
	var c CRUD
	return c
}

type CRUDColumns struct {
	JsonStr []string
}

// 可以在列上配置"sortable": true，该列表头右侧会渲染一个可点击的排序图标，可以切换正序和倒序。
func (p CRUDColumns) Sortable(sortable bool) CRUDColumns {
	return p._KeyVal("sortable", sortable)
}

// 可以在列上配置"searchable": true，该列表头右侧会渲染一个可点击的搜索图标，点击可以输入关键字进行该列的搜索
func (p CRUDColumns) Searchable(searchable bool) CRUDColumns {
	return p._KeyVal("searchable", searchable)
}

// 可以在列上配置filterable属性，该列表头右侧会渲染一个可点击的过滤图标，点击显示下拉框，选中进行过滤
// amis 只负责生成下拉选择器组件，并将搜索参数传递给接口，而不会在前端对数据进行搜索处理
func (p CRUDColumns) Filterable(filterableOptions []interface{}) CRUDColumns {
	if len(filterableOptions) < 1 {
		return p
	}
	pp := struct {
		Options []interface{} `json:"options"`
	}{
		Options: filterableOptions,
	}
	params, _ := json.Marshal(&pp)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"filterable":{%s}`, string(params)))
	return p
}

// 可以通过给列配置："quickEdit":true和quickSaveApi 可以实现表格内快速编辑并批量保存的功能。
func (p CRUDColumns) QuickEdit(quickEdit bool) CRUDColumns {
	return p._KeyVal("quickEdit", quickEdit)
}

// QuickEdit内联模式
func (p CRUDColumns) QuickEditInLine() CRUDColumns {
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"quickEdit":{"mode":"%s"}`, "inline"))
	return p
}

// QuickEdit自定义表单项模式
func (p CRUDColumns) QuickEditCustom(newFormItems string) CRUDColumns {
	p._NotEmpty("quickEdit", newFormItems)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"quickEdit":{%s}`, newFormItems))
	return p
}

// QuickEdit内联自定义表单项模式
func (p CRUDColumns) QuickEditCustomInLine(newFormItems string) CRUDColumns {
	p._NotEmpty("quickEdit", newFormItems)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"quickEdit":{"mode":"%s",%s}`, "inline", newFormItems))
	return p
}

// 可以在列上配置"sortable": true，该列表头右侧会渲染一个可点击的排序图标，可以切换正序和倒序。
func (p CRUDColumns) SaveImmediately(saveImmediately bool) CRUDColumns {
	return p._KeyVal("saveImmediately", saveImmediately)
}

// 在saveImmediately中配置 api，实现即时保存
func (p CRUDColumns) SaveImmediatelyApi(saveImmediatelyApiUrl string) CRUDColumns {
	p._NotEmpty("saveImmediately", saveImmediatelyApiUrl)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"saveImmediately":{"api":"%s"}`, saveImmediatelyApiUrl))
	return p
}

// CRUDColumns.Button.Build
// 这个是一个独立的Button
func (p CRUDColumns) Button(newButton string) CRUDColumns {
	p._NotEmpty("newButton", newButton)
	p.JsonStr = append(p.JsonStr, newButton)
	return p
}

func (p CRUDColumns) Build() string {
	return strings.Join(p.JsonStr, ",")
}

func (p CRUDColumns) _KeyVal(k string, v interface{}) CRUDColumns {
	p._NotEmpty(k, v)
	switch v := v.(type) {
	case string:
		p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"%s":"%s"`, k, v))
	case int:
		p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"%s":%d`, k, v))
	case bool:
		p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"%s":%t`, k, v))
	}
	return p
}

func (p CRUDColumns) _NotEmpty(py, v interface{}) {
	switch v := v.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			panic(fmt.Errorf("Property:%s,Cannot hold empty values", py))
		}
	}
}

func NewCRUDColumns() CRUDColumns {
	var c CRUDColumns
	return c
}

type Messages struct {
	// 获取失败时提示
	FetchFailedMsg string `json:"fetchFailed,omitempty"`
	// 保存顺序失败提示
	SaveOrderFailedMsg string `json:"saveOrderFailed,omitempty"`
	// 保存顺序成功提示
	SaveOrderSuccessMsg string `json:"saveOrderSuccess,omitempty"`
	// 快速保存失败提示
	QuickSaveFailedMsg string `json:"quickSaveFailed,omitempty"`
	// 快速保存成功提示
	QuickSaveSuccessMsg string `json:"quickSaveSuccess,omitempty"`
}

func (m Messages) FetchFailed(fetchFailed string) Messages {
	m.FetchFailedMsg = fetchFailed
	return m
}
func (m Messages) SaveOrderFailed(saveOrderFailed string) Messages {
	m.SaveOrderFailedMsg = saveOrderFailed
	return m
}
func (m Messages) SaveOrderSuccess(saveOrderSuccess string) Messages {
	m.SaveOrderSuccessMsg = saveOrderSuccess
	return m
}
func (m Messages) QuickSaveFailed(quickSaveFailed string) Messages {
	m.QuickSaveFailedMsg = quickSaveFailed
	return m
}
func (m Messages) QuickSaveSuccess(quickSaveSuccess string) Messages {
	m.QuickSaveSuccessMsg = quickSaveSuccess
	return m
}
func (m Messages) Build() string {
	s, _ := json.Marshal(&m)
	return string(s)
}

func NewMessages() Messages {
	var m Messages
	return m
}
