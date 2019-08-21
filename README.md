# 关于本项目
基于github.com/weolar/miniblink49 封装的GOLANG使用的WebUI。

直接调用DLL，未使用CGO。

仅支持32位DLL。


# 使用前准备
dll文件夹内有两个DLL文件，一个是 miniblink 的32位接口DLL。

大家可以自行到 https://github.com/weolar/miniblink49/releases 下载最新编译后的 dll 替换就行。（注意，名称一定得是 node.dll，后有说明）

另一个名称为 gonode.dll 是我自己编译的，主要用于解决GO和mb的DLL之前的调用问题，其中的问题有：

1、mb的DLL在返回值或者是回调函数里的参数值类型不为 int32 时，会导致GO接收后的结果不确定。（比如逻辑值应该是0和1，但GO接收到的都是一串数值）

2、mb需要GO回调函数返回值为bool型时，GO无论返回0或者1都是同样的效果。

对此，我猜想可能是GO调用DLL无论参数和返回类型都只能使用 uintptr 的原因。

如果有朋友知道原理并解决，望告知，谢谢！

所以，gonode.dll的作用就是为了处理一些GO和mbDLL之间传值的问题。

然后，应该没有其它要准备的了。

简而言之就是，把两个DLL放在同目录，再在你的代码中设置它们的位置。


另外，如果你的系统是 64位的，在首次编译或者构建时记得请先：
set GOARCH=386

# 如何使用
以下是一个简单且完整的显示一个浏览器窗口的代码：

package main

import (
	"gowebui"  //引入包
	"strconv"
)

var mb gowebui.WebView

func main() {

	gowebui.Initialize("node.dll", "gonode.dll")  //设置两个DLL的位置，我这里是两个DLL放在此demo源码文件内。
	gowebui.BindJsFunction("showalert", abcC, 99, 3) //绑定一个和网页JS交互的函数

	mb.CreateWebWindow("测试窗口", 0, 0, 0, 0, 700, 400) //创建一个标题为“测试窗口”,宽为700，高为400的窗口
	mb.ShowWindow(true)

	mb.LoadHTML(`<html><head><title>测试窗口</title></head><body>
	<a href="https://www.baidu.com">点击打开百度</a><br>
	<a href="javascript:showalert('一1一',2,true)">点击显示alert</a>
	</body></html>`)
  //加载一段HTML代码

	<-gowebui.Exit  等待MB窗口被关闭
}

//下面是被绑定的函数
func abcC(es gowebui.JsExecState, param uintptr) uintptr {
	gowebui.StartCallBack()
	defer gowebui.EndCallBack()

	mb.RunJS("alert('链接被点击了，第1个参数为：" + mb.GetJsString(es, mb.GetJsValueFromArg(es, 0)) + "')")
	mb.RunJS("alert('链接被点击了，第2个参数为：" + strconv.Itoa(int(mb.GetJsInt(es, mb.GetJsValueFromArg(es, 1)))) + "')")

	if mb.GetJsBool(mb.GetJsValueFromArg(es, 2)) == true {
		mb.RunJS("alert('链接被点击了，第3个参数为：true')")
	} else {
		mb.RunJS("alert('链接被点击了，第3个参数为：false')")
	}
	return 0
}

# 目前已实现公开接口

公开类型：

type JsExecState int32

type WkeWebView int32

type JsValue string //使用文本来存长整型数字

type HWnd int32

type WkeString int32


type WebView struct {
	hWnd     HWnd
	hWebView WkeWebView
} //此结构代表浏览窗口，见DEMO代码

公开包级接口：

Initialize(mbPath, gonodePath string) bool //初始化，整个程序内只能在最开始时调用且只能调用一次

StartCallBack() //回调函数开始时调用

EndCallBack() //回调函数结束时调，可见上方DEMO代码


WebView类型的方法：

func (mb *WebView) SetMainHWND() //将当前窗口句柄设置为主窗口句柄，设置之后，如果此窗口被销毁，则所有窗口都被销毁。

func (mb *WebView) GetWebTitle() string { //获取网页标题

func (mb *WebView) SetWindowTitle(title string) { //设置WIN窗口标题

func (mb *WebView) ShowWindow(visible bool) {

func (mb *WebView) LoadURL(url string) { //加载网页地址

func (mb *WebView) LoadHTML(html string) { //加载HTML

func (mb *WebView) GetURL() string { //取网页地址

func (mb *WebView) GetHWND() HWnd { //取WIN窗口句柄

func (mb *WebView) SetSize(w, h uintptr) { //设置窗口宽高

func (mb *WebView) MoveToCenter() { //将窗口居中

func (mb *WebView) GetUserAgent() string { //取浏览器UA

func (mb *WebView) GetCookie() string { //获取页面COOKIE

func (mb *WebView) DOMReady() bool { //DOM文档结构是否加载完成

func (mb *WebView) WkeStringToString(s WkeString) string { //WkeString 转 string

func (mb *WebView) SetNavigationToNewWindow(y bool) { //设置点a标签后是在本窗口跳转还是新窗口跳转，相关可见 BindWillCreateWebWindow()

func (mb *WebView) CanGoBack() bool { //是否可后退

func (mb *WebView) CanGoForward() bool { //是否可前进

func (mb *WebView) GoBack() { //后退

func (mb *WebView) GoForward() { //前进

func (mb *WebView) Reload() { //刷新


//创建浏览器窗口 

//参数分别为 标题，类型（0带边框的可调窗口，1透明窗口，2子窗口；为2时，parent 必须设置），父窗口句柄，左边，顶边，宽，高

func (mb *WebView) CreateWebWindow(title string, wkeWindowType, parent, x, y, width, height int32) WkeWebView {


//HookRequest，在BindLoadUrlBegin 回调函数中使用，使用之后，如果对job设置了NetHookRequest，

//则表示WebView会缓存获取到的网络数据，并在这次网络请求结束后调用BindLoadUrlEnd设置的回调函数，同时传递缓存的数据。在此期间，mb不会处理网络数据。

func (mb *WebView) NetHookRequest(job uintptr) {


//在BindLoadUrlBegin回调里调用，设置后，此请求将被取消。

//参见 BindLoadUrlBegin 反返回值说明，个人尝试，返回值无效，需要取消连接的话，直接调用此方法

func (mb *WebView) NetCancelRequest(job uintptr) {


//绑定alert回调

func (mb *WebView) BindAlertBox(f func(webView WkeWebView, param int32, url WkeString) uintptr, callbackParam int32) {


//绑定URL改变回调

func (mb *WebView) BindURLChanged(f func(webView WkeWebView, param int32, url WkeString) uintptr, callbackParam int32) {


//绑定标题变化回调函数

func (mb *WebView) BindTitleChanged(f func(webView WkeWebView, param int32, title WkeString) uintptr, callbackParam int32) {


//网页准备浏览时触发此回调。

//第三个回调的参数：0，表示点击a标签触发；1,点击form触发;2,前进后退触发;3,重新加载触发;4,表单重新提交；5，其它方式触发

//回调函数返回 0 表示阻止本次浏览，1表示继续进行浏览 (见 NetCancelRequest() 方法说明)

func (mb *WebView) BindNavigation(f func(webView WkeWebView, param int32, wkeNavigationType int32, url WkeString) uintptr, callbackParam int32) {


//即将创建新浏览窗口

//第一个参数为回调函数，第二个为处定义数据，将会出现在回调函数的 param 这个参数中。

//回调函数的第3个参数见 BindNavigation() 说明。

//返回值为 WkeWebView,表示使用该窗口来显示将创建的新的网页内容,返回自己的WkeWebView或者是关闭SetNavigationToNewWindow（），则使用本窗口加载新窗口内容

func (mb *WebView) BindWillCreateWebWindow(f func(webView WkeWebView, param int32, wkeNavigationType int32, url WkeString, windowFeatures int32) uintptr, callbackParam int32) {


//对应 对应js里的body onload事件

//第一个参数为回调函数，第二个参数为自定义数据，将会出现在回调函数里的 param

func (mb *WebView) BindDocumentReady(f func(webView WkeWebView, param int32) uintptr, callbackParam int32) {


//对应 对应js里的body onload事件，不同之处是回调函数的第三个参数会传入当前 onload 事件的框架句柄

//第一个参数为回调函数，第二个参数为自定义数据，将会出现在回调函数里的 param

func (mb *WebView) BindDocumentReady2(f func(webView WkeWebView, param int32, frameId int32) uintptr, callbackParam int32) {


//任何网络请求发起前会触发此回调，见 NetHookRequest() 说明

func (mb *WebView) BindLoadUrlBegin(f func(webView WkeWebView, param int32, url_char, job uintptr) uintptr, callbackParam int32) {


//见 NetHookRequest() 说明

//第三个参数为当前数据的URL， 它是ANSI 编码的 CHAR（通常，如果里面只有ASCII字符的话，可直接 PtrToString() 转成GO的STRING型），第五个参数为数据指针位置，第六个为数据长度

func (mb *WebView) BindLoadUrlEnd(f func(webView WkeWebView, param int32, url_char, job, buf, bufLen uintptr) uintptr, callbackParam int32) {


//绑定一个全局函数到主frame的window上，必须在创建窗口前绑定，也就是在 miniblink.Initialize和mb.CreateWebWindow之间

//第一个参数为JSfunc的名称；第二个为回调函数；第三个为自定义参数，可通过回调函数的第2个参数获取； 第四个为JSfunc的参数数量。

func BindJsFunction(jsFuncName string, f func(es JsExecState, param uintptr) uintptr, param uintptr, argCount int32) {


//用于JSfunc的回调函数，根据参数索引取得JSfunc传过来的值的 jsValue

//第一个参数 使用回调函数的es参数，第二个参数为索引，从0开始

func (mb *WebView) GetJsValueFromArg(es JsExecState, argIdx int32) JsValue {


//获取页面主frame的jsExecState

func (mb *WebView) GetExecState() JsExecState {


//通过ES和JsValue将JS传过来的值转为string

//es通过 GetExecState() 获取或者回调函数的es参数，jsValue通过 GetJsValue…… 之类的获取或者RunJS()返回值

func (mb *WebView) GetJsString(es JsExecState, v JsValue) string {


//通过ES和JsValue将JS传过来的值转为整数，如果v是个整形或者浮点，返回相应值（如果是浮点，返回取整后的值）。如果是其他类型，返回0（这里注意）

//es通过 GetExecState() 获取或者回调函数的es参数，jsValue通过 GetJsValue…… 之类的获取或者RunJS()返回值

func (mb *WebView) GetJsInt(es JsExecState, v JsValue) int32 {


//通过ES和JsValue将JS传过来的值转为浮点型，如果v是个浮点型，返回相应值。如果是其他类型，返回0.0（这里注意）

//es通过 GetExecState() 获取或者回调函数的es参数，jsValue通过 GetJsValue…… 之类的获取或者RunJS()返回值

func (mb *WebView) GetJsFloat64(es JsExecState, v JsValue) float64 {


//通过JsValue将JS传过来的值转为BOOL型，如果v本身是个布尔值，返回对应的true或者false；如果是个对象（JSTYPE_OBJECT），返回false（这里注意）

//jsValue通过 GetJsValue…… 之类的获取或者RunJS()返回值

func (mb *WebView) GetJsBool(v JsValue) bool {


//让窗口执行JS代码，此代码是在一个 function(){} 中执行，如果要取返回值的话，需要加 return ；如 RunJS("return $('title').text()")

//参数为JS脚本代码，返回值为 JsValue

func (mb *WebView) RunJS(JsCode string) JsValue {



