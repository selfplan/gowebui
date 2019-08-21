package gowebui

import (
	"fmt"
	"runtime"
	"sync"
	"syscall"
	"unsafe"
)

type JsExecState int32
type WkeWebView int32
type JsValue string //使用文本来存长整型数字
type HWnd int32
type WkeString int32

var (
	_dllCommand chan dllAPI
	_dllReturn  chan interface{}
	_mhWnd      HWnd                //主窗口句柄
	_allhWnd    map[HWnd]WkeWebView //存放所有句柄， 键：WIN窗口句柄，值：webView句柄
	Exit        chan bool           //所有操作完成，退出进程
	_lc         sync.Mutex          //操作回调函数记数

	_callBackNum uint32 //记录了运行的回调函数。（如果不为0，表示有操作是在回调函数里的，调用DLL时，需要判断线程）
	_mThreadId   uint32 //主线程ID
)

type dllAPI struct {
	hWebView WkeWebView
	apiStr   string
	arg1     interface{}
	arg2     interface{}
	arg3     interface{}
	arg4     interface{}
	arg5     interface{}
	arg6     interface{}
}

type WebView struct {
	hWnd     HWnd
	hWebView WkeWebView
}

//初始化，整个程序内只能在最开始时调用且只能调用一次
func Initialize(mbPath, gonodePath string) bool {
	loadMBDLL(mbPath, gonodePath)
	_dllCommand = make(chan dllAPI)
	_dllReturn = make(chan interface{})
	Exit = make(chan bool)
	_allhWnd = make(map[HWnd]WkeWebView)
	go func() {
		runtime.LockOSThread()
		_mThreadId = getCurrentThreadId()
		for {

			select {
			case dllCommand := <-_dllCommand:
				switch dllCommand.apiStr {
				case "wkeOnWindowDestroy":
					f1 := syscall.NewCallbackCDecl(dllCommand.arg1.(func(uintptr) uintptr))
					ret, _, _ := syscall.Syscall(dllWkeOnWindowDestroy, 3, uintptr(dllCommand.hWebView), uintptr(unsafe.Pointer(f1)), 0)
					_dllReturn <- ret
				case "wkeJsBindFunction":
					f1 := syscall.NewCallbackCDecl(dllCommand.arg2.(func(JsExecState, uintptr) uintptr))
					ret, _, _ := syscall.Syscall6(dllWkeJsBindFunction, 4, strToPtr(dllCommand.arg1.(string)), f1, dllCommand.arg3.(uintptr), uintptr(dllCommand.arg4.(int32)), 0, 0)
					_dllReturn <- ret
				case "wkeOnURLChanged":
					f1 := syscall.NewCallbackCDecl(dllCommand.arg1.(func(WkeWebView, int32, WkeString) uintptr))
					ret, _, _ := syscall.Syscall(dllWkeOnURLChanged, 3, uintptr(dllCommand.hWebView), f1, int32ToPtr(dllCommand.arg2.(int32)))
					_dllReturn <- ret
				case "wkeOnAlertBox":
					f1 := syscall.NewCallbackCDecl(dllCommand.arg1.(func(WkeWebView, int32, WkeString) uintptr))
					ret, _, _ := syscall.Syscall(dllWkeOnAlertBox, 3, uintptr(dllCommand.hWebView), f1, int32ToPtr(dllCommand.arg2.(int32)))
					_dllReturn <- ret
				case "wkeOnTitleChanged":
					f1 := syscall.NewCallbackCDecl(dllCommand.arg1.(func(WkeWebView, int32, WkeString) uintptr))
					ret, _, _ := syscall.Syscall(dllWkeOnTitleChanged, 3, uintptr(dllCommand.hWebView), f1, int32ToPtr(dllCommand.arg2.(int32)))
					_dllReturn <- ret
				case "wkeOnNavigation":
					f1 := syscall.NewCallback(dllCommand.arg1.(func(WkeWebView, int32, int32, WkeString) uintptr))
					ret, _, _ := syscall.Syscall(GOBindNavigation, 3, uintptr(dllCommand.hWebView), f1, int32ToPtr(dllCommand.arg2.(int32)))
					_dllReturn <- ret
				case "wkeOnCreateView":
					f1 := syscall.NewCallback(dllCommand.arg1.(func(WkeWebView, int32, int32, WkeString, int32) uintptr))
					ret, _, _ := syscall.Syscall(GOBindCreateView, 3, uintptr(dllCommand.hWebView), f1, int32ToPtr(dllCommand.arg2.(int32)))
					_dllReturn <- ret
				case "wkeOnDocumentReady":
					f1 := syscall.NewCallbackCDecl(dllCommand.arg1.(func(WkeWebView, int32) uintptr))
					ret, _, _ := syscall.Syscall(dllWkeOnDocumentReady, 3, uintptr(dllCommand.hWebView), f1, int32ToPtr(dllCommand.arg2.(int32)))
					_dllReturn <- ret
				case "wkeOnDocumentReady2":
					f1 := syscall.NewCallbackCDecl(dllCommand.arg1.(func(WkeWebView, int32, int32) uintptr))
					ret, _, _ := syscall.Syscall(dllWkeOnDocumentReady2, 3, uintptr(dllCommand.hWebView), f1, int32ToPtr(dllCommand.arg2.(int32)))
					_dllReturn <- ret
				case "wkeOnLoadUrlBegin":
					f1 := syscall.NewCallback(dllCommand.arg1.(func(WkeWebView, int32, uintptr, uintptr) uintptr))
					ret, _, _ := syscall.Syscall(GOBindLoadUrlBegin, 3, uintptr(dllCommand.hWebView), f1, int32ToPtr(dllCommand.arg2.(int32)))
					_dllReturn <- ret
				case "wkeOnLoadUrlEnd":
					f1 := syscall.NewCallbackCDecl(dllCommand.arg1.(func(WkeWebView, int32, uintptr, uintptr, uintptr, uintptr) uintptr))
					ret, _, _ := syscall.Syscall(dllWkeOnLoadUrlEnd, 3, uintptr(dllCommand.hWebView), f1, int32ToPtr(dllCommand.arg2.(int32)))
					_dllReturn <- ret

				case "Exit":
					postQuitMessage(0)
				default:
					_dllReturn <- miniblinkCommand(dllCommand.hWebView, dllCommand.apiStr, dllCommand.arg1, dllCommand.arg2, dllCommand.arg3, dllCommand.arg4, dllCommand.arg5, dllCommand.arg6)
				}
			default:
				var msg msg
				if peekMessage(&msg, 0, 0, 0, 1) {
					translateMessage(&msg)
					dispatchMessage(&msg)
					if msg.Message == WM_QUIT {
						goto Loop
					}
				}
			}
		}

	Loop:
		fmt.Println("Go OUT!")
		freeMBDLL() //卸载DLL
		Exit <- true

	}()

	var api dllAPI
	api.apiStr = "wkeIsInitialize"
	_dllCommand <- api
	ret := <-_dllReturn

	if ret != 1 {
		api.apiStr = "wkeInitialize"
		_dllCommand <- api
		ret = <-_dllReturn

	}
	api.apiStr = "wkeIsInitialize"
	_dllCommand <- api
	ret = <-_dllReturn
	fmt.Println(ret)
	return (ret == 1)
}

//将当前窗口句柄设置为主窗口句柄，设置之后，如果此窗口被销毁，则所有窗口都被销毁。
func (mb *WebView) SetMainHWND() {
	_mhWnd = mb.hWnd
}

//窗口将被销毁回调
func (mb *WebView) wkeOnWindowDestroy(wkeWebView uintptr) uintptr {
	StartCallBack()
	defer EndCallBack()
	fmt.Println("将被销毁，", mb.hWnd, _mhWnd, mb.hWebView)
	if mb.hWnd == _mhWnd {
		//为主窗口被销毁

		var api dllAPI
		api.apiStr = "Exit"
		callDLLAPI(api)
	} else {
		//普通窗口被销毁
		delete(_allhWnd, mb.hWnd)
	}
	return 0
}

//获取网页标题
func (mb *WebView) GetWebTitle() string {
	var api dllAPI
	api.hWebView = mb.hWebView
	api.apiStr = "wkeGetTitle"
	ret := callDLLAPI(api)

	return ret.(string)
}

//设置WIN窗口标题
func (mb *WebView) SetWindowTitle(title string) {
	var api dllAPI
	api.apiStr = "wkeSetWindowTitle"
	api.hWebView = mb.hWebView
	api.arg1 = title
	callDLLAPI(api)
}

//显示窗口
func (mb *WebView) ShowWindow(visible bool) {
	var api dllAPI
	api.hWebView = mb.hWebView
	api.apiStr = "wkeShowWindow"
	if visible == true {
		api.arg1 = uintptr(1)
	} else {
		api.arg1 = uintptr(0)
	}

	callDLLAPI(api)
}

//加载网页地址
func (mb *WebView) LoadURL(url string) {
	var api dllAPI
	api.apiStr = "wkeLoadURL"
	api.hWebView = mb.hWebView
	api.arg1 = url
	callDLLAPI(api)
}

//加载HTML
func (mb *WebView) LoadHTML(html string) {
	var api dllAPI
	api.apiStr = "wkeLoadHTML"
	api.hWebView = mb.hWebView
	api.arg1 = html
	callDLLAPI(api)
}

//取网页地址
func (mb *WebView) GetURL() string {
	var api dllAPI
	api.apiStr = "wkeGetURL"
	api.hWebView = mb.hWebView
	ret := callDLLAPI(api)
	return ret.(string)
}

//取WIN窗口句柄
func (mb *WebView) GetHWND() HWnd {
	return mb.hWnd
}

//设置窗口宽高
func (mb *WebView) SetSize(w, h uintptr) {
	var api dllAPI
	api.apiStr = "wkeResize"
	api.hWebView = mb.hWebView
	api.arg1 = w
	api.arg2 = h
	callDLLAPI(api)
}

//将窗口居中
func (mb *WebView) MoveToCenter() {
	var api dllAPI
	api.apiStr = "wkeMoveToCenter"
	api.hWebView = mb.hWebView
	callDLLAPI(api)
}

//获取浏览器UA
func (mb *WebView) GetUserAgent() string {
	var api dllAPI
	api.apiStr = "wkeGetUserAgent"
	api.hWebView = mb.hWebView
	ret := callDLLAPI(api)
	return ret.(string)
}

//获取页面COOKIE
func (mb *WebView) GetCookie() string {
	var api dllAPI
	api.apiStr = "wkeGetCookie"
	api.hWebView = mb.hWebView
	ret := callDLLAPI(api)
	return ret.(string)
}

//DOM文档结构是否加载完成
func (mb *WebView) DOMReady() bool {
	var api dllAPI
	api.apiStr = "wkeIsDocumentReady"
	api.hWebView = mb.hWebView
	ret := callDLLAPI(api)
	return ret.(bool)
}

//WkeString 转 string
func (mb *WebView) WkeStringToString(s WkeString) string {
	var api dllAPI
	api.apiStr = "wkeGetString"
	api.arg1 = s
	ret := callDLLAPI(api)
	return ret.(string)

}

//设置点a标签后是在本窗口跳转还是新窗口跳转，相关可见 BindWillCreateWebWindow()
func (mb *WebView) SetNavigationToNewWindow(y bool) {
	var api dllAPI
	api.apiStr = "wkeSetNavigationToNewWindowEnable"
	api.hWebView = mb.hWebView
	if y {
		api.arg1 = uintptr(1)
	} else {
		api.arg1 = uintptr(0)
	}
	callDLLAPI(api)
}

//是否可后退
func (mb *WebView) CanGoBack() bool {
	var api dllAPI
	api.apiStr = "wkeCanGoBack"
	api.hWebView = mb.hWebView
	ret := callDLLAPI(api)
	return ret.(bool)
}

//是否可前进
func (mb *WebView) CanGoForward() bool {
	var api dllAPI
	api.apiStr = "wkeCanGoForward"
	api.hWebView = mb.hWebView
	ret := callDLLAPI(api)
	return ret.(bool)
}

//后退
func (mb *WebView) GoBack() {
	var api dllAPI
	api.apiStr = "wkeGoBack"
	api.hWebView = mb.hWebView
	callDLLAPI(api)
}

//前进
func (mb *WebView) GoForward() {
	var api dllAPI
	api.apiStr = "wkeGoForward"
	api.hWebView = mb.hWebView
	callDLLAPI(api)
}

//刷新
func (mb *WebView) Reload() {
	var api dllAPI
	api.apiStr = "wkeReload"
	api.hWebView = mb.hWebView
	callDLLAPI(api)
}

//创建浏览器窗口
//参数分别为 标题，类型（0带边框的可调窗口，1透明窗口，2子窗口；为2时，parent 必须设置），父窗口句柄，左边，顶边，宽，高
func (mb *WebView) CreateWebWindow(title string, wkeWindowType, parent, x, y, width, height int32) WkeWebView {
	var api dllAPI
	api.apiStr = "wkeCreateWebWindow"
	api.arg1 = wkeWindowType
	api.arg2 = parent
	api.arg3 = x
	api.arg4 = y
	api.arg5 = width
	api.arg6 = height

	ret := callDLLAPI(api)
	mb.hWebView = ret.(WkeWebView)
	api.apiStr = "wkeGetWindowHandle"
	api.hWebView = mb.hWebView
	mb.hWnd = callDLLAPI(api).(HWnd)

	api.apiStr = "wkeOnWindowDestroy"
	api.hWebView = mb.hWebView
	api.arg1 = mb.wkeOnWindowDestroy
	callDLLAPI(api)

	_allhWnd[mb.hWnd] = mb.hWebView
	if _mhWnd == 0 {
		_mhWnd = mb.hWnd //默认将第一个窗口设为主窗口
	}

	api.apiStr = "wkeSetWindowTitle"
	api.hWebView = mb.hWebView
	api.arg1 = title
	callDLLAPI(api)
	return ret.(WkeWebView)
}

//以下为 NET 相关接口

//HookRequest，在BindLoadUrlBegin 回调函数中使用，使用之后，如果对job设置了NetHookRequest，
//则表示WebView会缓存获取到的网络数据，并在这次网络请求结束后调用BindLoadUrlEnd设置的回调函数，同时传递缓存的数据。在此期间，mb不会处理网络数据。
func (mb *WebView) NetHookRequest(job uintptr) {
	var api dllAPI
	api.apiStr = "wkeNetHookRequest"
	api.arg1 = job
	callDLLAPI(api)
}

//在BindLoadUrlBegin回调里调用，设置后，此请求将被取消。
//参见 BindLoadUrlBegin 反返回值说明，个人尝试，返回值无效，需要取消连接的话，直接调用此方法
func (mb *WebView) NetCancelRequest(job uintptr) {
	var api dllAPI
	api.apiStr = "wkeNetCancelRequest"
	api.arg1 = job
	callDLLAPI(api)
}

//以下为有回调函数的事件
//绑定alert回调
func (mb *WebView) BindAlertBox(f func(webView WkeWebView, param int32, url WkeString) uintptr, callbackParam int32) {
	var api dllAPI
	api.apiStr = "wkeOnAlertBox"
	api.hWebView = mb.hWebView
	api.arg1 = f
	api.arg2 = callbackParam
	callDLLAPI(api)
}

//绑定URL改变回调
func (mb *WebView) BindURLChanged(f func(webView WkeWebView, param int32, url WkeString) uintptr, callbackParam int32) {
	var api dllAPI
	api.apiStr = "wkeOnURLChanged"
	api.hWebView = mb.hWebView
	api.arg1 = f
	api.arg2 = callbackParam
	callDLLAPI(api)
}

//绑定标题变化回调函数
func (mb *WebView) BindTitleChanged(f func(webView WkeWebView, param int32, title WkeString) uintptr, callbackParam int32) {
	var api dllAPI
	api.apiStr = "wkeOnTitleChanged"
	api.hWebView = mb.hWebView
	api.arg1 = f
	api.arg2 = callbackParam
	callDLLAPI(api)
}

//网页准备浏览时触发此回调。
//第三个回调的参数：0，表示点击a标签触发；1,点击form触发;2,前进后退触发;3,重新加载触发;4,表单重新提交；5，其它方式触发
//回调函数返回 0 表示阻止本次浏览，1表示继续进行浏览 (见 NetCancelRequest() 方法说明)
func (mb *WebView) BindNavigation(f func(webView WkeWebView, param int32, wkeNavigationType int32, url WkeString) uintptr, callbackParam int32) {
	var api dllAPI
	api.apiStr = "wkeOnNavigation"
	api.hWebView = mb.hWebView
	api.arg1 = f
	api.arg2 = callbackParam
	callDLLAPI(api)
}

//即将创建新浏览窗口
//第一个参数为回调函数，第二个为处定义数据，将会出现在回调函数的 param 这个参数中。
//回调函数的第3个参数见 BindNavigation() 说明。
//返回值为 WkeWebView,表示使用该窗口来显示将创建的新的网页内容,返回自己的WkeWebView或者是关闭SetNavigationToNewWindow（），则使用本窗口加载新窗口内容
func (mb *WebView) BindWillCreateWebWindow(f func(webView WkeWebView, param int32, wkeNavigationType int32, url WkeString, windowFeatures int32) uintptr, callbackParam int32) {
	var api dllAPI
	api.apiStr = "wkeOnCreateView"
	api.hWebView = mb.hWebView
	api.arg1 = f
	api.arg2 = callbackParam
	callDLLAPI(api)
}

//对应 对应js里的body onload事件
//第一个参数为回调函数，第二个参数为自定义数据，将会出现在回调函数里的 param
func (mb *WebView) BindDocumentReady(f func(webView WkeWebView, param int32) uintptr, callbackParam int32) {
	var api dllAPI
	api.apiStr = "wkeOnDocumentReady"
	api.hWebView = mb.hWebView
	api.arg1 = f
	api.arg2 = callbackParam
	callDLLAPI(api)
}

//对应 对应js里的body onload事件，不同之处是回调函数的第三个参数会传入当前 onload 事件的框架句柄
//第一个参数为回调函数，第二个参数为自定义数据，将会出现在回调函数里的 param
func (mb *WebView) BindDocumentReady2(f func(webView WkeWebView, param int32, frameId int32) uintptr, callbackParam int32) {
	var api dllAPI
	api.apiStr = "wkeOnDocumentReady2"
	api.hWebView = mb.hWebView
	api.arg1 = f
	api.arg2 = callbackParam
	callDLLAPI(api)
}

//任何网络请求发起前会触发此回调，见 NetHookRequest() 说明
func (mb *WebView) BindLoadUrlBegin(f func(webView WkeWebView, param int32, url_char, job uintptr) uintptr, callbackParam int32) {
	var api dllAPI
	api.apiStr = "wkeOnLoadUrlBegin"
	api.hWebView = mb.hWebView
	api.arg1 = f
	api.arg2 = callbackParam
	callDLLAPI(api)
}

//见 NetHookRequest() 说明
//第三个参数为当前数据的URL， 它是ANSI 编码的 CHAR（通常，如果里面只有ASCII字符的话，可直接 PtrToString() 转成GO的STRING型），第五个参数为数据指针位置，第六个为数据长度
func (mb *WebView) BindLoadUrlEnd(f func(webView WkeWebView, param int32, url_char, job, buf, bufLen uintptr) uintptr, callbackParam int32) {
	var api dllAPI
	api.apiStr = "wkeOnLoadUrlEnd"
	api.hWebView = mb.hWebView
	api.arg1 = f
	api.arg2 = callbackParam
	callDLLAPI(api)
}

//以下为JS相关接口
//绑定一个全局函数到主frame的window上，必须在创建窗口前绑定，也就是在 miniblink.Initialize和mb.CreateWebWindow之间
//第一个参数为JSfunc的名称；第二个为回调函数；第三个为自定义参数，可通过回调函数的第2个参数获取； 第四个为JSfunc的参数数量。
func BindJsFunction(jsFuncName string, f func(es JsExecState, param uintptr) uintptr, param uintptr, argCount int32) {
	var api dllAPI
	api.apiStr = "wkeJsBindFunction"
	api.arg1 = jsFuncName
	api.arg2 = f
	api.arg3 = param
	api.arg4 = argCount
	callDLLAPI(api)

}

//用于JSfunc的回调函数，根据参数索引取得JSfunc传过来的值的 jsValue
//第一个参数 使用回调函数的es参数，第二个参数为索引，从0开始
func (mb *WebView) GetJsValueFromArg(es JsExecState, argIdx int32) JsValue {
	var api dllAPI
	api.apiStr = "jsArg"
	api.arg1 = es
	api.arg2 = argIdx
	ret := callDLLAPI(api)
	return ret.(JsValue)

}

//获取页面主frame的jsExecState
func (mb *WebView) GetExecState() JsExecState {
	var api dllAPI
	api.apiStr = "wkeGlobalExec"
	api.hWebView = mb.hWebView
	return callDLLAPI(api).(JsExecState)
}

//通过ES和JsValue将JS传过来的值转为string
//es通过 GetExecState() 获取或者回调函数的es参数，jsValue通过 GetJsValue…… 之类的获取或者RunJS()返回值
func (mb *WebView) GetJsString(es JsExecState, v JsValue) string {
	var api dllAPI
	api.apiStr = "jsToString"
	api.arg1 = es
	api.arg2 = v
	ret := callDLLAPI(api)
	return ret.(string)
}

//通过ES和JsValue将JS传过来的值转为整数，如果v是个整形或者浮点，返回相应值（如果是浮点，返回取整后的值）。如果是其他类型，返回0（这里注意）
//es通过 GetExecState() 获取或者回调函数的es参数，jsValue通过 GetJsValue…… 之类的获取或者RunJS()返回值
func (mb *WebView) GetJsInt(es JsExecState, v JsValue) int32 {
	var api dllAPI
	api.apiStr = "jsToInt"
	api.arg1 = es
	api.arg2 = v
	ret := callDLLAPI(api)
	return ret.(int32)
}

//通过ES和JsValue将JS传过来的值转为浮点型，如果v是个浮点型，返回相应值。如果是其他类型，返回0.0（这里注意）
//es通过 GetExecState() 获取或者回调函数的es参数，jsValue通过 GetJsValue…… 之类的获取或者RunJS()返回值
func (mb *WebView) GetJsFloat64(es JsExecState, v JsValue) float64 {
	var api dllAPI
	api.apiStr = "jsToDouble"
	api.arg1 = es
	api.arg2 = v
	ret := callDLLAPI(api)
	return ret.(float64)
}

//通过JsValue将JS传过来的值转为BOOL型，如果v本身是个布尔值，返回对应的true或者false；如果是个对象（JSTYPE_OBJECT），返回false（这里注意）
//jsValue通过 GetJsValue…… 之类的获取或者RunJS()返回值
func (mb *WebView) GetJsBool(v JsValue) bool {
	var api dllAPI
	api.apiStr = "jsIsTrue"
	api.arg1 = v
	ret := callDLLAPI(api)
	return ret.(bool)
}

//让窗口执行JS代码，此代码是在一个 function(){} 中执行，如果要取返回值的话，需要加 return ；如 RunJS("return $('title').text()")
//参数为JS脚本代码，返回值为 JsValue
func (mb *WebView) RunJS(JsCode string) JsValue {
	var api dllAPI
	api.apiStr = "wkeRunJS"
	api.hWebView = mb.hWebView
	api.arg1 = JsCode
	ret := callDLLAPI(api)
	return ret.(JsValue)
}

//内部支持
//调用DLL，主要解决在回调函数里操作时，channel 死锁的问题
func callDLLAPI(api dllAPI) interface{} {
	_lc.Lock()
	callBackNum := _callBackNum
	_lc.Unlock()
	if callBackNum > 0 {
		//表示有操作是在回调函数里进行的
		if _mThreadId == getCurrentThreadId() {
			//当前在主线程，表示此操作是在回调里进行的
			switch api.apiStr {
			case "Exit":
				postQuitMessage(0)
				return 0
			default:
				return miniblinkCommand(api.hWebView, api.apiStr, api.arg1, api.arg2, api.arg3, api.arg4, api.arg5, api.arg6)
			}
		} else {
			//非回调里进行的
			_dllCommand <- api

			return <-_dllReturn
		}
	} else {
		//非回调里进行的
		_dllCommand <- api
		return <-_dllReturn
	}

}

//回调函数开始时调用
func StartCallBack() {
	_lc.Lock()
	defer _lc.Unlock()
	_callBackNum++
}

//回调函数结束时调用
func EndCallBack() {
	_lc.Lock()
	defer _lc.Unlock()
	_callBackNum--
}
