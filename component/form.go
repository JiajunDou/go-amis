package component

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Form struct {
	JsonStr []string
}

// 设置一个名字后，方便其他组件与其通信
func (p *Form) Name(newCRUD string) *Form {
	p._NotEmpty("CRUD", newCRUD)
	p.JsonStr = append(p.JsonStr, newCRUD)
	return p
}

const (
	FormModeNormal     = "normal"
	FormModeHorizontal = "horizontal"
	FormModeInline     = "inline"
)

// 表单展示方式，可以是：normal、horizontal 或者 inline
func (p *Form) Mode(formMode string) *Form {
	return p._KeyVal("mode", formMode)
}

// 当 mode 为 horizontal 时有用，用来控制 label
// {"left":"col-sm-2", "right":"col-sm-10", "offset":"col-sm-offset-2"}
func (p *Form) Horizontal(leftRightOffserCss string) *Form {
	return p._KeyVal("Horizontal", leftRightOffserCss)
}

// Form 的标题
// 默认"表单"
func (p *Form) Title(formTitle string) *Form {
	return p._KeyVal("title", formTitle)
}

// 默认的提交按钮名称，如果设置成空，则可以把默认按钮去掉。
// 默认"提交"
func (p *Form) SubmitText(submitText string) *Form {
	return p._KeyVal("submitText", submitText)
}

// 外层 Dom 的类名
func (p *Form) ClassName(s string) *Form {
	return p._KeyVal("className", s)
}

// Form 表单项集合
func (p *Form) Controls(newFormControls ...string) *Form {
	ss := strings.Join(newFormControls, ",")
	p._NotEmpty("controls", ss)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"controls":[%s]`, ss))
	return p
}

// Form 表单项集合 Controls的别名
// 更容易理解 表单项
func (p *Form) FormItems(newFormItems ...string) *Form {
	ss := strings.Join(newFormItems, ",")
	p._NotEmpty("controls", ss)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"controls":[%s]`, ss))
	return p
}

// Form 提交按钮，成员为 Action
func (p *Form) FormSubmitActions(newFormSubmitActions ...string) *Form {
	ss := strings.Join(newFormSubmitActions, ",")
	p._NotEmpty("actions", ss)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"actions":[%s]`, ss))
	return p
}

// 是否让 Form 用 panel 包起来，设置为 false 后，actions 将无效。
func (p *Form) WrapWithPanel(isWrapWithPanel bool) *Form {
	return p._KeyVal("wrapWithPanel", isWrapWithPanel)
}

// 外层 panel 的类名
func (p *Form) PanelClassName(panelClassName string) *Form {
	return p._KeyVal("panelClassName", panelClassName)
}

// Form 用来保存数据的 api。
func (p *Form) Api(api string) *Form {
	return p._KeyVal("api", api)
}

// 	Form 用来获取初始数据的 api。
func (p *Form) InitApi(initApi string) *Form {
	return p._KeyVal("initApi", initApi)
}

// 刷新时间(最低 3000)
func (p *Form) Interval(interval int) *Form {
	return p._KeyVal("interval", interval)
}

// 通过表达式来配置停止刷新的条件
func (p *Form) StopAutoRefreshWhen(stopAutoRefreshWhen string) *Form {
	return p._KeyVal("stopAutoRefreshWhen", stopAutoRefreshWhen)
}

// 	Form 用来获取初始数据的 api,与 initApi 不同的是，会一直轮询请求该接口，直到返回 finished 属性为 true 才 结束。
func (p *Form) InitAsyncApi(initAsyncApi string) *Form {
	return p._KeyVal("initApi", initAsyncApi)
}

// 设置了 initApi 或者 initAsyncApi 后，默认会开始就发请求，设置为 false 后就不会起始就请求接口
func (p *Form) InitFetch(initFetch bool) *Form {
	return p._KeyVal("initFetch", initFetch)
}

// 用表达式来配置
func (p *Form) InitFetchOn(initFetchOn string) *Form {
	return p._KeyVal("initFetchOn", initFetchOn)
}

// 设置了 initAsyncApi 以后，默认拉取的时间间隔(最低 3000)
func (p *Form) InitCheckInterval(initCheckInterval int) *Form {
	return p._KeyVal("initCheckInterval", initCheckInterval)
}

// 	设置了 initAsyncApi 后，默认会从返回数据的 data.finished 来判断是否完成，也可以设置成其他的 xxx，就会从 data.xxx 中获取
func (p *Form) InitFinishedField(initFinishedField string) *Form {
	return p._KeyVal("initFinishedField", initFinishedField)
}

// 设置此属性后，表单提交发送保存接口后，还会继续轮询请求该接口，直到返回 finished 属性为 true 才 结束。
func (p *Form) AsyncApi(asyncApi string) *Form {
	return p._KeyVal("asyncApi", asyncApi)
}

// 轮询请求的时间间隔，默认为 3 秒。设置 asyncApi 才有效,刷新时间
func (p *Form) CheckInterval(checkinterval int) *Form {
	return p._KeyVal("checkinterval", checkinterval)
}

// 如果决定结束的字段名不是 finished 请设置此属性，比如 is_success
func (p *Form) FinishedField(finishedField string) *Form {
	return p._KeyVal("finishedField", finishedField)
}

// 表单修改即提交
func (p *Form) SubmitOnChange(submitOnChange bool) *Form {
	return p._KeyVal("submitOnChange", submitOnChange)
}

// 初始就提交一次
func (p *Form) SubmitOnInit(submitOnInit bool) *Form {
	return p._KeyVal("submitOnInit", submitOnInit)
}

// 提交后是否重置表单
func (p *Form) ResetAfterSubmit(resetAfterSubmit bool) *Form {
	return p._KeyVal("resetAfterSubmit", resetAfterSubmit)
}

// 设置主键 id, 当设置后，检测表单是否完成时（asyncApi），只会携带此数据。
func (p *Form) PrimaryField(primaryField string) *Form {
	return p._KeyVal("primaryField", primaryField)
}

// 默认表单提交自己会通过发送 api 保存数据，
// 但是也可以设定另外一个 form 的 name 值，
// 或者另外一个 CRUD 模型的 name 值。
// 如果 target 目标是一个 Form ，则目标 Form 会重新触发 initApi，api 可以拿到当前 form 数据。
// 如果目标是一个 CRUD 模型，则目标模型会重新触发搜索，
// 参数为当前 Form 数据。当目标是 window 时，会把当前表单的数据附带到页面地址上。
func (p *Form) Target(target string) *Form {
	return p._KeyVal("target", target)
}

// 设置此属性后，Form 保存成功后，自动跳转到指定页面。支持相对地址，和绝对地址（相对于组内的）。
func (p *Form) Redirect(redirect string) *Form {
	return p._KeyVal("redirect", redirect)
}

// 操作完后刷新目标对象。请填写目标组件设置的 name 值，如果填写为 window 则让当前页面整体刷新。
func (p *Form) Reload(reload string) *Form {
	return p._KeyVal("reload", reload)
}

// 是否自动聚焦。
func (p *Form) AutoFocus(autoFocus bool) *Form {
	return p._KeyVal("autoFocus", autoFocus)
}

// 指定是否可以自动获取上层的数据并映射到表单项上
func (p *Form) CanAccessSuperData(canAccessSuperData bool) *Form {
	return p._KeyVal("canAccessSuperData", canAccessSuperData)
}

// 指定表单是否开启本地缓存
func (p *Form) PersistData(persistData bool) *Form {
	return p._KeyVal("persistData", persistData)
}

// 指定表单提交成功后是否清除本地缓存
func (p *Form) ClearPersistDataAfterSubmit(clearPersistDataAfterSubmit bool) *Form {
	return p._KeyVal("clearPersistDataAfterSubmit", clearPersistDataAfterSubmit)
}

// trim 当前表单项的每一个值
func (p *Form) TrimValues(trimValues bool) *Form {
	return p._KeyVal("submitOnInit", trimValues)
}

// The CRUD table has all of *Form functions
func (p *Form) Build() string {
	return "{" + strings.Join(p.JsonStr, ",") + "}"
}

func (p *Form) _KeyVal(k string, v interface{}) *Form {
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

func (p *Form) _NotEmpty(py, v interface{}) {
	switch v := v.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			panic(fmt.Errorf("Property:%s,Cannot hold empty values", py))
		}
	}
}

func NewForm() *Form {
	c := new(Form)
	c.JsonStr = append(c.JsonStr, fmt.Sprintf(`"type":"%s"`, "table"))
	return c
}

type FormMessages struct {
	// 获取成功时提示
	FetchSuccessMsg string `json:"fetchSuccess,omitempty"`
	// 获取失败时提示
	FetchFailedMsg string `json:"fetchFailed,omitempty"`
	// 保存成功时提示
	SaveSuccessMsg string `json:"saveSuccess,omitempty"`
	// 保存失败时提示
	SaveFailedMsg string `json:"saveFailed,omitempty"`
}

func (m *FormMessages) FetchFailed(fetchFailed string) *FormMessages {
	m.FetchFailedMsg = fetchFailed
	return m
}
func (m *FormMessages) FetchSuccess(fetchSuccess string) *FormMessages {
	m.FetchSuccessMsg = fetchSuccess
	return m
}
func (m *FormMessages) SaveOrderSuccess(saveSuccess string) *FormMessages {
	m.SaveSuccessMsg = saveSuccess
	return m
}
func (m *FormMessages) SaveFailed(saveFailed string) *FormMessages {
	m.SaveFailedMsg = saveFailed
	return m
}

func (m *FormMessages) Build() string {
	s, _ := json.Marshal(&m)
	return string(s)
}

func NewFormMessages() *FormMessages {
	m := new(FormMessages)
	return m
}

type FormItems struct {
	JsonStr []string
}

// 内联模式
func (p *FormItems) Inline() *FormItems {
	return p._KeyVal("mode", "inline")
}

const (
	FormItemsSizeXs = "xs"
	// 极小
	FormItemsSizeSm = "sm"
	// 小
	FormItemsSizeMd = "md"
	// 大
	FormItemsSizeLg = "lg"
	// 占满
	FormItemsSizeFull = "full"
)

// 可以配置size，来调整表单项的尺寸，支持'xs' | 'sm' | 'md' | 'lg' | 'full'，如下：
func (p *FormItems) Size(size string) *FormItems {
	return p._KeyVal("size", size)
}

// 	表单控制器类名
func (p *FormItems) InputClassName(s string) *FormItems {
	return p._KeyVal("inputClassName", s)
}

// label 的类名
func (p *FormItems) LabelClassName(s string) *FormItems {
	return p._KeyVal("labelClassName", s)
}

// 设置label属性来配置表单项标签。
//
// 当表单为水平布局时，左边即便是不设置label为了保持对齐也会留空，如果想要去掉空白，请设置成false。
func (p *FormItems) HorizontalLabelDisplayNone() *FormItems {
	return p._KeyVal("label", false)
}

// 表单项标签提示
// 配置labelRemark可以实现标签描述提示
func (p *FormItems) LabelRemark(newRemarkNoType string) *FormItems {
	p._NotEmpty("labelRemark", newRemarkNoType)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"labelRemark":%s`, newRemarkNoType))
	return p
}

// 配置禁用
// 静态配置
// 通过配置"disabled": true来禁用表单项
func (p *FormItems) Disabled() *FormItems {
	return p._KeyVal("disabled", true)
}

//
func (p *FormItems) DisabledOn(disabledOnExp string) *FormItems {
	return p._KeyVal("disabledOn", disabledOnExp)
}

// 是否隐藏
func (p *FormItems) Hidden() *FormItems {
	return p._KeyVal("hidden", true)
}

// 通过条件配置显隐
// 你也通过表达式配置hiddenOn，来实现在某个条件下禁用当前表单
func (p *FormItems) HiddenOn(hiddenOnExp string) *FormItems {
	return p._KeyVal("hiddenOn", hiddenOnExp)
}

// 是否显示
func (p *FormItems) Visible() *FormItems {
	return p._KeyVal("visible", true)
}

// 通过配置"required": true来标识该表单项为必填。
func (p *FormItems) Required() *FormItems {
	return p._KeyVal("required", true)
}

// 你也通过表达式配置requiredOn，来实现在某个条件下使当前表单项必填。
func (p *FormItems) RequiredOn(requiredOnExp string) *FormItems {
	return p._KeyVal("requiredOn", requiredOnExp)
}

func (p *FormItems) Validations(newValidations string) *FormItems {
	p._NotEmpty("validations", newValidations)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"validations":%s`, newValidations))
	return p
}

// 是否该表单项值发生变化时就提交当前表单。
func (p *FormItems) SubmitOnChange() *FormItems {
	return p._KeyVal("submitOnChange", true)
}

func (p *FormItems) Build() string {
	return "{" + strings.Join(p.JsonStr, ",") + "}"
}

func (p *FormItems) _KeyVal(k string, v interface{}) *FormItems {
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

func (p *FormItems) _NotEmpty(py, v interface{}) {
	switch v := v.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			panic(fmt.Errorf("Property:%s,Cannot hold empty values", py))
		}
	}
}

func NewFormItemsAdditional() *FormItems {
	c := new(FormItems)
	return c
}

type Validations struct {
	JsonStr []string
}

func (p *Validations) ValidateOnChange() *Validations {
	return p._KeyVal("validateOnChange", true)
}

// 必须是空白字符。注意！ 该格式校验是值，校验空白字符，而不是当前表单项是否为空，想校验是否为空，请配置 必填校验
func (p *Validations) IsEmptyString() *Validations {
	return p._KeyVal("isEmptyString", true)
}

func (p *Validations) NotEmptyString() *Validations {
	return p._KeyVal("notEmptyString", true)
}

func (p *Validations) IsEmail() *Validations {
	return p._KeyVal("isEmail", true)
}

func (p *Validations) IsUrl() *Validations {
	return p._KeyVal("isUrl", true)
}

func (p *Validations) IsUrlPath() *Validations {
	return p._KeyVal("isUrlPath", true)
}

func (p *Validations) IsNumeric() *Validations {
	return p._KeyVal("isNumeric", true)
}

func (p *Validations) IsAlpha() *Validations {
	return p._KeyVal("isAlpha", true)
}

func (p *Validations) IsAlphanumeric() *Validations {
	return p._KeyVal("isAlphanumeric", true)
}

func (p *Validations) IsInt() *Validations {
	return p._KeyVal("isInt", true)
}

func (p *Validations) isFloat() *Validations {
	return p._KeyVal("isFloat", true)
}

func (p *Validations) IsLength(length int) *Validations {
	return p._KeyVal("isLength", length)
}

func (p *Validations) IsJson() *Validations {
	return p._KeyVal("isJson", true)
}

func (p *Validations) MinLength(length int) *Validations {
	return p._KeyVal("minLength", length)
}

func (p *Validations) MaxLength(length int) *Validations {
	return p._KeyVal("maxLength", length)
}

func (p *Validations) Maximum(number int) *Validations {
	return p._KeyVal("maximum", number)
}

func (p *Validations) Minimum(number int) *Validations {
	return p._KeyVal("minimum", number)
}

func (p *Validations) Equals(someValue interface{}) *Validations {
	return p._KeyVal("equals", someValue)
}

func (p *Validations) EqualsField(fieldName string) *Validations {
	return p._KeyVal("equalsField", fieldName)
}

func (p *Validations) MatchRegexp(matchRegexp string) *Validations {
	return p._KeyVal("matchRegexp", matchRegexp)
}
func (p *Validations) MatchRegexp1(matchRegexp string) *Validations {
	return p._KeyVal("matchRegexp1", matchRegexp)
}

func (p *Validations) MatchRegexp2(matchRegexp string) *Validations {
	return p._KeyVal("matchRegexp2", matchRegexp)
}

func (p *Validations) MatchRegexp3(matchRegexp string) *Validations {
	return p._KeyVal("matchRegexp3", matchRegexp)
}

func (p *Validations) MatchRegexp4(matchRegexp string) *Validations {
	return p._KeyVal("matchRegexp4", matchRegexp)
}

func (p *Validations) Description(description string) *Validations {
	return p._KeyVal("description", description)
}

const (
	FormValidationIsEmptyString  = "isEmptyString"
	FormValidationIsEmail        = "isEmail"
	FormValidationIsUrl          = "isUrl"
	FormValidationIsNumeric      = "isNumeric"
	FormValidationIsAlpha        = "isAlpha"
	FormValidationIsAlphanumeric = "isAlphanumeric"
	FormValidationIsInt          = "isInt"
	FormValidationIsFloat        = "isFloat"
	FormValidationIsLength       = "isLength"
	FormValidationMinLength      = "minLength"
	FormValidationMaxLength      = "maxLength"
	FormValidationMaximum        = "maximum"
	FormValidationMinimum        = "minimum"
	FormValidationEquals         = "equals"
	FormValidationEqualsField    = "equalsField"
	FormValidationIsJson         = "isJson"
	FormValidationNotEmptyString = "notEmptyString"
	FormValidationIsUrlPath      = "isUrlPath"
	FormValidationMatchRegexp    = "matchRegexp"
	FormValidationMatchRegexp1   = "matchRegexp1"
	FormValidationMatchRegexp2   = "matchRegexp2"
	FormValidationMatchRegexp3   = "matchRegexp3"
	FormValidationMatchRegexp4   = "matchRegexp4"
)

func (p *Validations) ValidationErrors(formValidationType, errorText string) *Validations {
	p._NotEmpty("ValidationErrors", formValidationType)
	p._NotEmpty("ValidationErrorsText", errorText)
	p.JsonStr = append(p.JsonStr, fmt.Sprintf(`"validationErrors":{"%s":"%s"}`, formValidationType, errorText))
	return p
}

func (p *Validations) Build() string {
	return strings.Join(p.JsonStr, ",")
}

func (p *Validations) _KeyVal(k string, v interface{}) *Validations {
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

func (p *Validations) _NotEmpty(py, v interface{}) {
	switch v := v.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			panic(fmt.Errorf("Property:%s,Cannot hold empty values", py))
		}
	}
}

func NewValidations() *Validations {
	c := new(Validations)
	return c
}
