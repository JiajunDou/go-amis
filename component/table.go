package component

import (
	"fmt"
	"strings"
)

type Table struct {
	JsonStr []string
}

// Title of the table
// Displayed as string text in the first row of the table
// |title|
func (p Table) Title(s string) Table {
	return p._KeyVal("title", s)
}

// Data source, bind current environment variables
// Default value is ${items}
func (p Table) Source(s string) Table {
	return p._KeyVal("source", s)
}

func (p Table) Columns(newTableColumns ...string) Table {
	ss := strings.Join(newTableColumns, ",")
	p._NotEmpty("columns", newTableColumns)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"columns":[%s]`, ss))
	return p
}

// 悬浮行操作按钮组
func (p Table) ItemActions(newAction ...string) Table {
	ss := strings.Join(newAction, ",")
	p._NotEmpty("itemActions", newAction)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"itemActions":[%s]`, ss))
	return p
}

// 配置当前行是否可勾选的条件，要用 表达式
func (p Table) ItemCheckableOn(itemCheckableOn string) Table {
	return p._KeyVal("itemCheckableOn", itemCheckableOn)
}

// 配置当前行是否可拖拽的条件，要用 表达式
func (p Table) ItemDraggableOn(itemDraggableOn string) Table {
	return p._KeyVal("itemDraggableOn", itemDraggableOn)
}

// 点击数据行是否可以勾选当前行
func (p Table) CheckOnItemClick(checkOnItemClick bool) Table {
	return p._KeyVal("checkOnItemClick", checkOnItemClick)
}

// 是否固定表头
func (p Table) AffixHeader(s bool) Table {
	return p._KeyVal("affixHeader", s)
}

// 当没数据的时候的文字提示
func (p Table) Placeholder(s string) Table {
	return p._KeyVal("placeholder", s)
}

// 是否显示列筛选器
func (p Table) ColumnsTogglable(s bool) Table {
	return p._KeyVal("columnsTogglable", s)
}

// 列太多时，内容没办法全部显示完，可以让部分信息在底部显示，可以让用户展开查看详情。配置很简单，只需要开启 footable 属性，同时将想在底部展示的列加个 breakpoint 属性为 * 即可。
func (p Table) Footable(s bool) Table {
	return p._KeyVal("footable", s)
}

func (p Table) ClassName(s string) Table {
	return p._KeyVal("className", s)
}

func (p Table) TableClassName(s string) Table {
	return p._KeyVal("tableClassName", s)
}

func (p Table) HeaderClassName(s string) Table {
	return p._KeyVal("headerClassName", s)
}

func (p Table) FooterClassName(s string) Table {
	return p._KeyVal("footerClassName", s)
}

func (p Table) RowClassName(s string) Table {
	return p._KeyVal("rowClassName", s)
}

func (p Table) ToolbarClassName(s string) Table {
	return p._KeyVal("toolbarClassName", s)
}

// 可以通过配置rowClassNameExpr来为行添加 CSS 类，支持 模板 语法。
//
// 例如下例，"<%= data.id % 2 ? "bg-success" : "" %>" 表示当行数据的 id 变量为 不能被 2 整除时，给当前行添加bg-success CSS 类名，即绿色背景色
func (p Table) RowClassNameExpr(rowClassNameExpr string) Table {
	return p._KeyVal("rowClassNameExpr", rowClassNameExpr)
}

// 可以配置"isHead": true，来让当前列以表头的样式展示。应用场景是：
//
// 所有列label配置空字符串，不显示表头
// 配置combineNum，合并单元格，实现左侧表头的形式
// 列上配置"isHead": true，调整样式
func (p Table) CombineNum(combineNum int) Table {
	return p._KeyVal("combineNum", combineNum)
}

// The CRUD table has all of Table functions
func (p Table) Build() string {

	return "{" + strings.Join(p.JsonStr, ",") + "}"
}

func (p Table) _KeyVal(k string, v interface{}) Table {
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

func (p Table) _NotEmpty(py, v interface{}) {
	switch v := v.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			panic(fmt.Errorf("Property:%s,Cannot hold empty values", py))
		}
	}
}

func NewTable() Table {
	var c Table
	c.JsonStr = append(c.JsonStr, fmt.Sprintf(`"type":"%s"`, "table"))
	return c
}

func NewTableCrud() Table {
	var c Table
	c.JsonStr = append(c.JsonStr, fmt.Sprintf(`"type":"%s"`, "crud"))
	return c
}

type TableColums struct {
	JsonStr []string
}

func (p TableColums) Name(dataSourceKeyName string) TableColums {
	return p._KeyVal("name", dataSourceKeyName)
}

func (p TableColums) Label(columnLabelText string) TableColums {
	return p._KeyVal("label", columnLabelText)
}

// Tpl
func (p TableColums) TypeTpl(tpl string) TableColums {
	p._NotEmpty("tpl", tpl)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"type":"%s","tpl":"%s"`, "tpl", tpl))
	return p
}

// Text
func (p TableColums) TypeText() TableColums {
	return p._KeyVal("type", "text")
}

func (p TableColums) TypeImage() TableColums {
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"type":"%s"`, "image"))
	return p
}

func (p TableColums) TypeDate() TableColums {
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"type":"%s"`, "date"))
	return p
}

func (p TableColums) TypeProgress() TableColums {
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"type":"%s"`, "progress"))
	return p
}

func (p TableColums) TypeStatus() TableColums {
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"type":"%s"`, "status"))
	return p
}

func (p TableColums) TypeSwitch() TableColums {
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"type":"%s"`, "switch"))
	return p
}

func (p TableColums) TypeColor() TableColums {
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"type":"%s"`, "color"))
	return p
}

func (p TableColums) CustomTypeImage(newImage string) TableColums {
	p._NotEmpty("image", newImage)
	p.JsonStr = append(p.JsonStr, newImage)
	return p
}

func (p TableColums) CustomTypeDate(newDate string) TableColums {
	p._NotEmpty("date", newDate)
	p.JsonStr = append(p.JsonStr, newDate)
	return p
}

func (p TableColums) CustomTypeProgress(newProgress string) TableColums {
	p._NotEmpty("progress", newProgress)
	p.JsonStr = append(p.JsonStr, newProgress)
	return p
}

func (p TableColums) CustomTypeStatus(newStatus string) TableColums {
	p._NotEmpty("status", newStatus)
	p.JsonStr = append(p.JsonStr, newStatus)
	return p
}

func (p TableColums) CustomTypeSwitch(newSwitch string) TableColums {
	p._NotEmpty("switch", newSwitch)
	p.JsonStr = append(p.JsonStr, newSwitch)
	return p
}

func (p TableColums) CustomTypeLink(newLink string) TableColums {
	p._NotEmpty("link", newLink)
	p.JsonStr = append(p.JsonStr, newLink)
	return p
}

func (p TableColums) CustomTypeColor(newColor string) TableColums {
	p._NotEmpty("color", newColor)
	p.JsonStr = append(p.JsonStr, newColor)
	return p
}

// 列宽像素
func (p TableColums) WidthInPixel(pixel int) TableColums {
	return p._KeyVal("width", pixel)
}

// 列宽百分比值
func (p TableColums) WidthInPercentage(percentage string) TableColums {
	return p._KeyVal("width", percentage)
}

// 固定列宽
// TODO 在新的表格演示中演示
var (
	TableColumsFixedLeft  = "left"
	TableColumsFixedRight = "right"
)

func (p TableColums) Fixed(fixedThisColumnAtLeftOrRight string) TableColums {
	return p._KeyVal("fixed", fixedThisColumnAtLeftOrRight)
}

// 复制该列的值到剪贴板
func (p TableColums) Copyable(isCopyThisColumnValueToClipboard bool) TableColums {
	return p._KeyVal("copyable", isCopyThisColumnValueToClipboard)
}

// 复制该列的值到剪贴板,自定义复制的提示语
func (p TableColums) CopyableWithContent(copyableWithContent string) TableColums {
	p._NotEmpty("copyableWithContent", copyableWithContent)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"copyable":{"content":"%s"}`, copyableWithContent))
	return p
}

// 这里只做一个默认的Tips提示
// 其他的弹出dialog，drawer对相对静态的表格作用不大
// popOver 配置详情：
// mode 可配置成 popOver、dialog 或者 drawer。 默认为 popOver。
// size 当配置成 dialog 或者 drawer 的时候有用。
// position 配置弹出位置，只有 popOver 模式有用。可选参数：center、left-top、right-top、left-bottom、right-bottom、fixed-center、fixed-left-top、fixed-right-top、fixed-left-bottom、fixed-right-bottom。
// offset 默认 {top: 0, left: 0}，如果要来一定的便宜请设置这个。
// trigger 触发弹出的条件。可配置为 click 或者 hover。默认为 click。
// showIcon 是否显示图标。默认会有个放大形状的图标出现在列里面。如果配置成 false，则触发事件出现在列上就会触发弹出。
// title 弹出框的标题。
// body 弹出框的内容。
func (p TableColums) PopOverTips(defaultPopOverTips string) TableColums {
	p._NotEmpty("defaultPopOverTips", defaultPopOverTips)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"popover":{"trigger": "hover","position": "left-top","showIcon": false,"body": {"type": "tpl","tpl": "%s"}}`, defaultPopOverTips))
	return p
}

// 可以配置"isHead": true，来让当前列以表头的样式展示。应用场景是：
//
// 所有列label配置空字符串，不显示表头
// 配置combineNum，合并单元格，实现左侧表头的形式
// 列上配置"isHead": true，调整样式
func (p TableColums) IsHead(isHead bool) TableColums {
	return p._KeyVal("isHead", isHead)
}

// 列太多时，内容没办法全部显示完，可以让部分信息在底部显示，可以让用户展开查看详情。配置很简单，只需要开启 footable 属性，同时将想在底部展示的列加个 breakpoint 属性为 * 即可。
func (p TableColums) Breakpoint() TableColums {
	return p._KeyVal("breakpoint", "*")
}

// 超级表头意思是，表头还可以再一次进行分组。额外添加个 groupName 属性即可。
func (p TableColums) GroupName(groupName string) TableColums {
	return p._KeyVal("groupName", groupName)
}

// 是否显示该列
func (p TableColums) Toggled(isShowThisColumn bool) TableColums {
	return p._KeyVal("toggled", isShowThisColumn)
}

func (p TableColums) Build() string {
	return "{" + strings.Join(p.JsonStr, ",") + "}"
}

func (p TableColums) _KeyVal(k string, v interface{}) TableColums {
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

func (p TableColums) _NotEmpty(py, v interface{}) {
	switch v := v.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			panic(fmt.Errorf("Property:%s,Cannot hold empty values", py))
		}
	}
}

func NewTableColumns() TableColums {
	var c TableColums
	return c
}
