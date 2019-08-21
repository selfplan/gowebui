package gowebui

import (
	"syscall"
)

const (
	WKE_LBUTTON = 1
	WKE_RBUTTON = 2
	WKE_SHIFT   = 4
	WKE_CONTROL = 8
	WKE_MBUTTON = 16

	WM_PAINT                = 15
	WM_SIZE                 = 5
	WM_QUIT                 = 18
	WM_KEYDOWN              = 256
	WM_KEYUP                = 257
	WM_CHAR                 = 258
	WM_LBUTTONDOWN          = 513
	WM_MBUTTONDOWN          = 519
	WM_RBUTTONDOWN          = 516
	WM_LBUTTONDBLCLK        = 515
	WM_MBUTTONDBLCLK        = 521
	WM_RBUTTONDBLCLK        = 518
	WM_LBUTTONUP            = 514
	WM_MBUTTONUP            = 520
	WM_RBUTTONUP            = 517
	WM_MOUSEMOVE            = 512
	MK_CONTROL              = 8
	MK_SHIFT                = 4
	MK_LBUTTON              = 1
	MK_MBUTTON              = 16
	MK_RBUTTON              = 2
	WM_CONTEXTMENU          = 123
	WM_MOUSEWHEEL           = 522
	WM_SETFOCUS             = 7
	WM_KILLFOCUS            = 8
	WM_IME_STARTCOMPOSITION = 269
	CFS_EXCLUDE             = 128
	WS_CHILD                = 1073741824
	WS_VISIBLE              = 268435456
	SS_NOTIFY               = 256
	WS_EX_LAYERED           = 524288
	WM_NCHITTEST            = 132
	WM_GETMINMAXINFO        = 36
	HTTOP                   = 12
	HTLEFT                  = 10
	HTRIGHT                 = 11
	HTBOTTOM                = 15
	HTTOPLEFT               = 13
	HTTOPRIGHT              = 14
	HTBOTTOMLEFT            = 16
	HTBOTTOMRIGHT           = 17
	SPI_GETWORKAREA         = 48
	WM_DESTROY              = 2
	WKE_EXTENDED            = 256
	WKE_REPEAT              = 16384
	GWL_EXSTYLE             = -20
	SRCCOPY                 = 13369376
	WM_TIMER                = 275
	WM_ERASEBKGND           = 20
	WM_SETCURSOR            = 32
	OBJ_BITMAP              = 7
	CAPTUREBLT              = 1073741824
	GA_ROOT                 = 2
	WM_NCDESTROY            = 130
	WM_CLOSE                = 16
	CP_UTF8                 = 65001
	CP_936                  = 936
	CFS_POINT               = 2
	CFS_FORCE_POSITION      = 32
)

var (
	kr32                  syscall.Handle
	dllGetCurrentThreadId uintptr

	u32                 syscall.Handle
	dllPostQuitMessage  uintptr
	dllDispatchMessage  uintptr
	dllTranslateMessage uintptr
	dllPeekMessage      uintptr

	mb syscall.Handle

	dllWkeCreateWebWindow                   uintptr
	dllWkeDestroyWebWindow                  uintptr
	dllWkeGetWindowHandle                   uintptr
	dllWkeShowWindow                        uintptr
	dllWkeEnableWindow                      uintptr
	dllWkeMoveWindow                        uintptr
	dllWkeMoveToCenter                      uintptr
	dllWkeResizeWindow                      uintptr
	dllWkeSetWindowTitleW                   uintptr
	dllWkeInitialize                        uintptr
	dllWkeIsInitialize                      uintptr
	dllWkeInitializeEx                      uintptr
	dllWkeFinalize                          uintptr
	dllWkeUpdate                            uintptr
	dllWkeGetVersion                        uintptr
	dllWkeGetVersionString                  uintptr
	dllWkeOnWindowClosing                   uintptr
	dllWkeOnWindowDestroy                   uintptr
	dllWkeSetNavigationToNewWindowEnable    uintptr
	dllWkeLoadW                             uintptr
	dllWkeSetProxy                          uintptr
	dllWkeSetViewProxy                      uintptr
	dllWkeGetName                           uintptr
	dllWkeSetName                           uintptr
	dllWkeSetHandle                         uintptr
	dllWkeSetHandleOffset                   uintptr
	dllWkeSetUserAgentW                     uintptr
	dllWkeSetUserAgent                      uintptr
	dllWkeLoadURLW                          uintptr
	dllWkeLoadURL                           uintptr
	dllWkeLoadHTMLW                         uintptr
	dllWkeLoadHTML                          uintptr
	dllWkeLoadFileW                         uintptr
	dllWkePostURLW                          uintptr
	dllWkePostURL                           uintptr
	dllWkeGetURL                            uintptr
	dllWkeIsLoading                         uintptr
	dllWkeIsLoadingSucceeded                uintptr
	dllWkeIsLoadingFailed                   uintptr
	dllWkeIsLoadingCompleted                uintptr
	dllWkeIsDocumentReady                   uintptr
	dllWkeStopLoading                       uintptr
	dllWkeReload                            uintptr
	dllWkeGetTitleW                         uintptr
	dllWkeGetTitle                          uintptr
	dllWkeResize                            uintptr
	dllWkeGetWidth                          uintptr
	dllWkeGetHeight                         uintptr
	dllWkeGetContentWidth                   uintptr
	dllWkeGetContentHeight                  uintptr
	dllWkeRepaintIfNeeded                   uintptr
	dllWkeSetDirty                          uintptr
	dllWkeIsDirty                           uintptr
	dllWkeAddDirtyArea                      uintptr
	dllWkeLayoutIfNeeded                    uintptr
	dllWkePaint2                            uintptr
	dllWkePaint                             uintptr
	dllWkeGetViewDC                         uintptr
	dllWkeGC                                uintptr
	dllWkeGetHostHWND                       uintptr
	dllWkeGetSource                         uintptr
	dllWkeCanGoBack                         uintptr
	dllWkeGoBack                            uintptr
	dllWkeCanGoForward                      uintptr
	dllWkeGoForward                         uintptr
	dllWkeEditorSelectAll                   uintptr
	dllWkeEditorUnSelect                    uintptr
	dllWkeEditorCopy                        uintptr
	dllWkeEditorCut                         uintptr
	dllWkeEditorPaste                       uintptr
	dllWkeEditorDelete                      uintptr
	dllWkeEditorUndo                        uintptr
	dllWkeEditorRedo                        uintptr
	dllWkeGetCookieW                        uintptr
	dllWkeGetCookie                         uintptr
	dllWkeSetCookie                         uintptr
	dllWkeSetCookieEnabled                  uintptr
	dllWkeIsCookieEnabled                   uintptr
	dllWkeSetCookieJarPath                  uintptr
	dllWkeSetCookieJarFullPath              uintptr
	dllWkeSetMediaVolume                    uintptr
	dllWkeGetMediaVolume                    uintptr
	dllWkeSetFocus                          uintptr
	dllWkeKillFocus                         uintptr
	dllWkeRunJSW                            uintptr
	dllWkeRunJS                             uintptr
	dllWkeGlobalExec                        uintptr
	dllWkeSleep                             uintptr
	dllWkeWake                              uintptr
	dllWkeIsAwake                           uintptr
	dllWkeSetTransparent                    uintptr
	dllWkeSetZoomFactor                     uintptr
	dllWkeGetZoomFactor                     uintptr
	dllWkeSetEditable                       uintptr
	dllWkeGetStringW                        uintptr
	dllWkeGetString                         uintptr
	dllWkeSetStringW                        uintptr
	dllWkeGetWebViewForCurrentContext       uintptr
	dllWkeSetUserKeyValue                   uintptr
	dllWkeGetUserKeyValue                   uintptr
	dllWkeSetDragFiles                      uintptr
	dllWkeOnTitleChanged                    uintptr
	dllWkeOnURLChanged                      uintptr
	dllWkeOnURLChanged2                     uintptr
	dllWkeOnAlertBox                        uintptr
	dllWkeOnConfirmBox                      uintptr
	dllWkeOnPromptBox                       uintptr
	dllWkeOnNavigation                      uintptr
	dllWkeOnCreateView                      uintptr
	dllWkeOnDocumentReady                   uintptr
	dllWkeOnDocumentReady2                  uintptr
	dllWkeOnLoadingFinish                   uintptr
	dllWkeOnDownload                        uintptr
	dllWkeOnConsole                         uintptr
	dllWkeOnPaintUpdated                    uintptr
	dllWkeOnWillMediaLoad                   uintptr
	dllWkeFireMouseEvent                    uintptr
	dllWkeFireContextMenuEvent              uintptr
	dllWkeFireMouseWheelEvent               uintptr
	dllWkeFireKeyUpEvent                    uintptr
	dllWkeFireKeyDownEvent                  uintptr
	dllWkeFireKeyPressEvent                 uintptr
	dllWkeFireWindowsMessage                uintptr
	dllWkeIsMainFrame                       uintptr
	dllWkeIsWebRemoteFrame                  uintptr
	dllWkeWebFrameGetMainFrame              uintptr
	dllWkeNetSetMIMEType                    uintptr
	dllWkeNetSetHTTPHeaderField             uintptr
	dllWkeNetGetHTTPHeaderField             uintptr
	dllWkeNetSetURL                         uintptr
	dllWkeNetSetData                        uintptr
	dllWkeGetCursorInfoType                 uintptr
	dllWkeSetCspCheckEnable                 uintptr
	dllWkeSetViewNetInterface               uintptr
	dllWkeNetHookRequest                    uintptr
	dllWkeOnLoadUrlBegin                    uintptr
	dllWkeOnLoadUrlEnd                      uintptr
	dllWkeOnDidCreateScriptContext          uintptr
	dllWkeOnWillReleaseScriptContext        uintptr
	dllWkeNetOnResponse                     uintptr
	dllWkeNetGetMIMEType                    uintptr
	dllWkeWebFrameGetMainWorldScriptContext uintptr
	dllWkeGetBlinkMainThreadIsolate         uintptr
	dllWkeCreateStringW                     uintptr
	dllWkeDeleteString                      uintptr
	dllWkeSetLocalStorageFullPath           uintptr
	dllWkeJsBindFunction                    uintptr
	dllJsBindFunction                       uintptr
	dllWkeJsBindGetter                      uintptr
	dllJsBindGetter                         uintptr
	dllWkeJsBindSetter                      uintptr
	dllJsBindSetter                         uintptr
	dllJsArgCount                           uintptr
	dllJsArgType                            uintptr
	dllJsArg                                uintptr
	dllJsTypeOf                             uintptr
	dllJsIsNumber                           uintptr
	dllJsIsString                           uintptr
	dllJsIsBoolean                          uintptr
	dllJsIsObject                           uintptr
	dllJsIsFunction                         uintptr
	dllJsIsUndefined                        uintptr
	dllJsIsNull                             uintptr
	dllJsIsArray                            uintptr
	dllJsIsTrue                             uintptr
	dllJsIsFalse                            uintptr
	dllJsToInt                              uintptr
	dllJsToFloat                            uintptr
	dllJsToDouble                           uintptr
	dllJsToBoolean                          uintptr
	dllJsToTempStringW                      uintptr
	dllJsToTempString                       uintptr
	dllJsInt                                uintptr
	dllJsFloat                              uintptr
	dllJsDouble                             uintptr
	dllJsBoolean                            uintptr
	dllJsUndefined                          uintptr
	dllJsNull                               uintptr
	dllJsTrue                               uintptr
	dllJsFalse                              uintptr
	dllJsStringW                            uintptr
	dllJsEmptyObject                        uintptr
	dllJsEmptyArray                         uintptr
	dllJsObject                             uintptr
	dllJsFunction                           uintptr
	dllJsGetData                            uintptr
	dllJsGet                                uintptr
	dllJsSet                                uintptr
	dllJsGetAt                              uintptr
	dllJsSetAt                              uintptr
	dllJsGetLength                          uintptr
	dllJsSetLength                          uintptr
	dllJsGlobalObject                       uintptr
	dllJsGetWebView                         uintptr
	dllJsEvalW                              uintptr
	dllJsEvalExW                            uintptr
	dllJsCall                               uintptr
	dllJsCallGlobal                         uintptr
	dllJsGetGlobal                          uintptr
	dllJsSetGlobal                          uintptr
	dllJsGC                                 uintptr
	dllJsToString                           uintptr
	dllJsToStringW                          uintptr
	dllJsArrayBuffer                        uintptr
	dllWkeSetFileSystem                     uintptr
	dllWkeCreateWebView                     uintptr
	dllWkeGetWebView                        uintptr
	dllWkeGetCaretRect                      uintptr
	dllWkeDestroyWebView                    uintptr
	dllWkeVisitAllCookie                    uintptr
	dllWkePerformCookieCommand              uintptr
	dllWkeRunJsByFrame                      uintptr
	dllWkeSetNpapiPluginsEnabled            uintptr
	dllWkeSetHeadlessEnabled                uintptr
	dllWkeSetTouchEnabled                   uintptr
	dllWkeSetDragEnable                     uintptr
	dllWkeSetMemoryCacheEnable              uintptr
	dllWkeOnMouseOverUrlChanged             uintptr
	dllWkeSetDebugConfig                    uintptr
	dllWkeSetDeviceParameter                uintptr
	dllWkeGetUserAgent                      uintptr
	dllWkeNetCancelRequest                  uintptr
	gonode                                  syscall.Handle
	GOjsArg                                 uintptr
	GOjsToString                            uintptr
	GOwkeRunJS                              uintptr
	GOjsIsTrue                              uintptr
	GOjsToInt                               uintptr
	GOjsToDouble                            uintptr
	GOBindNavigation                        uintptr
	GOBindCreateView                        uintptr
	GOBindLoadUrlBegin                      uintptr
	GOwkeSetNavigationToNewWindowEnable     uintptr
	GOwkeCanGoBack                          uintptr
	GOwkeCanGoForward                       uintptr
)

func loadMBDLL(mbPath, gonodePath string) {

	kr32, _ = syscall.LoadLibrary("kernel32.dll")
	dllGetCurrentThreadId, _ = syscall.GetProcAddress(kr32, "GetCurrentThreadId")

	u32, _ = syscall.LoadLibrary("user32.dll")
	dllPostQuitMessage, _ = syscall.GetProcAddress(u32, "PostQuitMessage")
	dllDispatchMessage, _ = syscall.GetProcAddress(u32, "DispatchMessageW")
	dllTranslateMessage, _ = syscall.GetProcAddress(u32, "TranslateMessage")
	dllPeekMessage, _ = syscall.GetProcAddress(u32, "PeekMessageW")

	mb, _ = syscall.LoadLibrary(mbPath)
	gonode, _ = syscall.LoadLibrary(gonodePath)

	GOjsArg, _ = syscall.GetProcAddress(gonode, "GOjsArg")
	GOjsToString, _ = syscall.GetProcAddress(gonode, "GOjsToString")
	GOwkeRunJS, _ = syscall.GetProcAddress(gonode, "GOwkeRunJS")
	GOjsIsTrue, _ = syscall.GetProcAddress(gonode, "GOjsIsTrue")
	GOjsToInt, _ = syscall.GetProcAddress(gonode, "GOjsToInt")
	GOjsToDouble, _ = syscall.GetProcAddress(gonode, "GOjsToDouble")
	GOBindNavigation, _ = syscall.GetProcAddress(gonode, "GOBindNavigation")
	GOBindCreateView, _ = syscall.GetProcAddress(gonode, "GOBindCreateView")
	GOBindLoadUrlBegin, _ = syscall.GetProcAddress(gonode, "GOBindLoadUrlBegin")
	GOwkeSetNavigationToNewWindowEnable, _ = syscall.GetProcAddress(gonode, "GOwkeSetNavigationToNewWindowEnable")
	GOwkeCanGoBack, _ = syscall.GetProcAddress(gonode, "GOwkeCanGoBack")
	GOwkeCanGoForward, _ = syscall.GetProcAddress(gonode, "GOwkeCanGoForward")

	dllWkeCreateWebWindow, _ = syscall.GetProcAddress(mb, "wkeCreateWebWindow")
	dllWkeDestroyWebWindow, _ = syscall.GetProcAddress(mb, "wkeDestroyWebWindow")
	dllWkeGetWindowHandle, _ = syscall.GetProcAddress(mb, "wkeGetWindowHandle")
	dllWkeShowWindow, _ = syscall.GetProcAddress(mb, "wkeShowWindow")
	dllWkeNetCancelRequest, _ = syscall.GetProcAddress(mb, "wkeNetCancelRequest")

	/*
		dllWkeEnableWindow, _ = syscall.GetProcAddress(mb, "wkeEnableWindow")
		dllWkeMoveWindow, _ = syscall.GetProcAddress(mb, "wkeMoveWindow")*/
	dllWkeMoveToCenter, _ = syscall.GetProcAddress(mb, "wkeMoveToCenter") /*
		dllWkeResizeWindow, _ = syscall.GetProcAddress(mb, "wkeResizeWindow")*/
	dllWkeSetWindowTitleW, _ = syscall.GetProcAddress(mb, "wkeSetWindowTitleW")
	dllWkeInitialize, _ = syscall.GetProcAddress(mb, "wkeInitialize")
	dllWkeIsInitialize, _ = syscall.GetProcAddress(mb, "wkeIsInitialize") /*
		dllWkeInitializeEx, _ = syscall.GetProcAddress(mb, "wkeInitializeEx")*/
	dllWkeFinalize, _ = syscall.GetProcAddress(mb, "wkeFinalize")
	dllWkeGetUserAgent, _ = syscall.GetProcAddress(mb, "wkeGetUserAgent") /*
		dllWkeUpdate, _ = syscall.GetProcAddress(mb, "wkeUpdate")
		dllWkeGetVersion, _ = syscall.GetProcAddress(mb, "wkeGetVersion")
		dllWkeGetVersionString, _ = syscall.GetProcAddress(mb, "wkeGetVersionString")
		dllWkeOnWindowClosing, _ = syscall.GetProcAddress(mb, "wkeOnWindowClosing")*/
	dllWkeOnWindowDestroy, _ = syscall.GetProcAddress(mb, "wkeOnWindowDestroy") /*
		dllWkeSetNavigationToNewWindowEnable, _ = syscall.GetProcAddress(mb, "wkeSetNavigationToNewWindowEnable") /*
			dllWkeLoadW, _ = syscall.GetProcAddress(mb, "wkeLoadW")
			dllWkeSetProxy, _ = syscall.GetProcAddress(mb, "wkeSetProxy")
			dllWkeSetViewProxy, _ = syscall.GetProcAddress(mb, "wkeSetViewProxy")
			dllWkeGetName, _ = syscall.GetProcAddress(mb, "wkeGetName")
			dllWkeSetName, _ = syscall.GetProcAddress(mb, "wkeSetName")*/
	dllWkeSetHandle, _ = syscall.GetProcAddress(mb, "wkeSetHandle") /*
		dllWkeSetHandleOffset, _ = syscall.GetProcAddress(mb, "wkeSetHandleOffset")
		dllWkeSetUserAgentW, _ = syscall.GetProcAddress(mb, "wkeSetUserAgentW")
		dllWkeSetUserAgent, _ = syscall.GetProcAddress(mb, "wkeSetUserAgent")*/
	dllWkeLoadURLW, _ = syscall.GetProcAddress(mb, "wkeLoadURLW")
	dllWkeLoadURL, _ = syscall.GetProcAddress(mb, "wkeLoadURL") /*
		dllWkeLoadHTMLW, _ = syscall.GetProcAddress(mb, "wkeLoadHTMLW")*/
	dllWkeLoadHTML, _ = syscall.GetProcAddress(mb, "wkeLoadHTML") /*
		dllWkeLoadFileW, _ = syscall.GetProcAddress(mb, "wkeLoadFileW")
		dllWkePostURLW, _ = syscall.GetProcAddress(mb, "wkePostURLW")
		dllWkePostURL, _ = syscall.GetProcAddress(mb, "wkePostURL")*/
	dllWkeGetURL, _ = syscall.GetProcAddress(mb, "wkeGetURL") /*
		dllWkeIsLoading, _ = syscall.GetProcAddress(mb, "wkeIsLoading")
		dllWkeIsLoadingSucceeded, _ = syscall.GetProcAddress(mb, "wkeIsLoadingSucceeded")
		dllWkeIsLoadingFailed, _ = syscall.GetProcAddress(mb, "wkeIsLoadingFailed")
		dllWkeIsLoadingCompleted, _ = syscall.GetProcAddress(mb, "wkeIsLoadingCompleted")*/
	dllWkeIsDocumentReady, _ = syscall.GetProcAddress(mb, "wkeIsDocumentReady") /*
		dllWkeStopLoading, _ = syscall.GetProcAddress(mb, "wkeStopLoading")*/
	dllWkeReload, _ = syscall.GetProcAddress(mb, "wkeReload") /*
		dllWkeGetTitleW, _ = syscall.GetProcAddress(mb, "wkeGetTitleW")*/
	dllWkeGetTitle, _ = syscall.GetProcAddress(mb, "wkeGetTitle")
	dllWkeResize, _ = syscall.GetProcAddress(mb, "wkeResize") /*
			dllWkeGetWidth, _ = syscall.GetProcAddress(mb, "wkeGetWidth")
			dllWkeGetHeight, _ = syscall.GetProcAddress(mb, "wkeGetHeight")
			dllWkeGetContentWidth, _ = syscall.GetProcAddress(mb, "wkeGetContentWidth")
			dllWkeGetContentHeight, _ = syscall.GetProcAddress(mb, "wkeGetContentHeight")
		dllWkeRepaintIfNeeded, _ = syscall.GetProcAddress(mb, "wkeRepaintIfNeeded")
			dllWkeSetDirty, _ = syscall.GetProcAddress(mb, "wkeSetDirty")
			dllWkeIsDirty, _ = syscall.GetProcAddress(mb, "wkeIsDirty")
			dllWkeAddDirtyArea, _ = syscall.GetProcAddress(mb, "wkeAddDirtyArea")
			dllWkeLayoutIfNeeded, _ = syscall.GetProcAddress(mb, "wkeLayoutIfNeeded")
			dllWkePaint2, _ = syscall.GetProcAddress(mb, "wkePaint2")
			dllWkePaint, _ = syscall.GetProcAddress(mb, "wkePaint")*/
	dllWkeGetViewDC, _ = syscall.GetProcAddress(mb, "wkeGetViewDC") /*
			dllWkeGC, _ = syscall.GetProcAddress(mb, "wkeGC")
			dllWkeGetHostHWND, _ = syscall.GetProcAddress(mb, "wkeGetHostHWND")
			dllWkeGetSource, _ = syscall.GetProcAddress(mb, "wkeGetSource")
		dllWkeCanGoBack, _ = syscall.GetProcAddress(mb, "wkeCanGoBack")
		dllWkeCanGoForward, _ = syscall.GetProcAddress(mb, "wkeCanGoForward")*/
	dllWkeGoBack, _ = syscall.GetProcAddress(mb, "wkeGoBack")
	dllWkeGoForward, _ = syscall.GetProcAddress(mb, "wkeGoForward") /*
		dllWkeEditorSelectAll, _ = syscall.GetProcAddress(mb, "wkeEditorSelectAll")
		dllWkeEditorUnSelect, _ = syscall.GetProcAddress(mb, "wkeEditorUnSelect")
		dllWkeEditorCopy, _ = syscall.GetProcAddress(mb, "wkeEditorCopy")
		dllWkeEditorCut, _ = syscall.GetProcAddress(mb, "wkeEditorCut")
		dllWkeEditorPaste, _ = syscall.GetProcAddress(mb, "wkeEditorPaste")
		dllWkeEditorDelete, _ = syscall.GetProcAddress(mb, "wkeEditorDelete")
		dllWkeEditorUndo, _ = syscall.GetProcAddress(mb, "wkeEditorUndo")
		dllWkeEditorRedo, _ = syscall.GetProcAddress(mb, "wkeEditorRedo")
		dllWkeGetCookieW, _ = syscall.GetProcAddress(mb, "wkeGetCookieW")*/
	dllWkeGetCookie, _ = syscall.GetProcAddress(mb, "wkeGetCookie") /*
		dllWkeSetCookie, _ = syscall.GetProcAddress(mb, "wkeSetCookie")
		dllWkeSetCookieEnabled, _ = syscall.GetProcAddress(mb, "wkeSetCookieEnabled")
		dllWkeIsCookieEnabled, _ = syscall.GetProcAddress(mb, "wkeIsCookieEnabled")
		dllWkeSetCookieJarPath, _ = syscall.GetProcAddress(mb, "wkeSetCookieJarPath")
		dllWkeSetCookieJarFullPath, _ = syscall.GetProcAddress(mb, "wkeSetCookieJarFullPath")
		dllWkeSetMediaVolume, _ = syscall.GetProcAddress(mb, "wkeSetMediaVolume")
		dllWkeGetMediaVolume, _ = syscall.GetProcAddress(mb, "wkeGetMediaVolume")*/
	dllWkeSetFocus, _ = syscall.GetProcAddress(mb, "wkeSetFocus")
	dllWkeKillFocus, _ = syscall.GetProcAddress(mb, "wkeKillFocus") /*
		dllWkeRunJSW, _ = syscall.GetProcAddress(mb, "wkeRunJSW")*/
	dllWkeRunJS, _ = syscall.GetProcAddress(mb, "wkeRunJS")
	dllWkeGlobalExec, _ = syscall.GetProcAddress(mb, "wkeGlobalExec") /*
		dllWkeSleep, _ = syscall.GetProcAddress(mb, "wkeSleep")
		dllWkeWake, _ = syscall.GetProcAddress(mb, "wkeWake")
		dllWkeIsAwake, _ = syscall.GetProcAddress(mb, "wkeIsAwake")
		dllWkeSetTransparent, _ = syscall.GetProcAddress(mb, "wkeSetTransparent")
		dllWkeSetZoomFactor, _ = syscall.GetProcAddress(mb, "wkeSetZoomFactor")
		dllWkeGetZoomFactor, _ = syscall.GetProcAddress(mb, "wkeGetZoomFactor")
		dllWkeSetEditable, _ = syscall.GetProcAddress(mb, "wkeSetEditable")
		dllWkeGetStringW, _ = syscall.GetProcAddress(mb, "wkeGetStringW")*/
	dllWkeGetString, _ = syscall.GetProcAddress(mb, "wkeGetString") /*
		dllWkeSetStringW, _ = syscall.GetProcAddress(mb, "wkeSetStringW")
		dllWkeGetWebViewForCurrentContext, _ = syscall.GetProcAddress(mb, "wkeGetWebViewForCurrentContext")
		dllWkeSetUserKeyValue, _ = syscall.GetProcAddress(mb, "wkeSetUserKeyValue")
		dllWkeGetUserKeyValue, _ = syscall.GetProcAddress(mb, "wkeGetUserKeyValue")
		dllWkeSetDragFiles, _ = syscall.GetProcAddress(mb, "wkeSetDragFiles")*/
	dllWkeOnTitleChanged, _ = syscall.GetProcAddress(mb, "wkeOnTitleChanged")
	dllWkeOnURLChanged, _ = syscall.GetProcAddress(mb, "wkeOnURLChanged")
	dllWkeOnURLChanged2, _ = syscall.GetProcAddress(mb, "wkeOnURLChanged2")
	dllWkeOnAlertBox, _ = syscall.GetProcAddress(mb, "wkeOnAlertBox") /*
		dllWkeOnConfirmBox, _ = syscall.GetProcAddress(mb, "wkeOnConfirmBox")
		dllWkeOnPromptBox, _ = syscall.GetProcAddress(mb, "wkeOnPromptBox")
		dllWkeOnNavigation, _ = syscall.GetProcAddress(mb, "wkeOnNavigation")*/
	dllWkeOnCreateView, _ = syscall.GetProcAddress(mb, "wkeOnCreateView")
	dllWkeOnDocumentReady, _ = syscall.GetProcAddress(mb, "wkeOnDocumentReady")
	dllWkeOnDocumentReady2, _ = syscall.GetProcAddress(mb, "wkeOnDocumentReady2") /*
			dllWkeOnLoadingFinish, _ = syscall.GetProcAddress(mb, "wkeOnLoadingFinish")
			dllWkeOnDownload, _ = syscall.GetProcAddress(mb, "wkeOnDownload")
			dllWkeOnConsole, _ = syscall.GetProcAddress(mb, "wkeOnConsole")
		dllWkeOnPaintUpdated, _ = syscall.GetProcAddress(mb, "wkeOnPaintUpdated")
			dllWkeOnWillMediaLoad, _ = syscall.GetProcAddress(mb, "wkeOnWillMediaLoad")
		dllWkeFireMouseEvent, _ = syscall.GetProcAddress(mb, "wkeFireMouseEvent")
		dllWkeFireContextMenuEvent, _ = syscall.GetProcAddress(mb, "wkeFireContextMenuEvent")
		dllWkeFireMouseWheelEvent, _ = syscall.GetProcAddress(mb, "wkeFireMouseWheelEvent")
		dllWkeFireKeyUpEvent, _ = syscall.GetProcAddress(mb, "wkeFireKeyUpEvent")
			dllWkeFireKeyDownEvent, _ = syscall.GetProcAddress(mb, "wkeFireKeyDownEvent")
			dllWkeFireKeyPressEvent, _ = syscall.GetProcAddress(mb, "wkeFireKeyPressEvent")
		dllWkeFireWindowsMessage, _ = syscall.GetProcAddress(mb, "wkeFireWindowsMessage")
			dllWkeIsMainFrame, _ = syscall.GetProcAddress(mb, "wkeIsMainFrame")
			dllWkeIsWebRemoteFrame, _ = syscall.GetProcAddress(mb, "wkeIsWebRemoteFrame")
			dllWkeWebFrameGetMainFrame, _ = syscall.GetProcAddress(mb, "wkeWebFrameGetMainFrame")
			dllWkeNetSetMIMEType, _ = syscall.GetProcAddress(mb, "wkeNetSetMIMEType")
			dllWkeNetSetHTTPHeaderField, _ = syscall.GetProcAddress(mb, "wkeNetSetHTTPHeaderField")
			dllWkeNetGetHTTPHeaderField, _ = syscall.GetProcAddress(mb, "wkeNetGetHTTPHeaderField")
			dllWkeNetSetURL, _ = syscall.GetProcAddress(mb, "wkeNetSetURL")
			dllWkeNetSetData, _ = syscall.GetProcAddress(mb, "wkeNetSetData")
			dllWkeGetCursorInfoType, _ = syscall.GetProcAddress(mb, "wkeGetCursorInfoType")
			dllWkeSetCspCheckEnable, _ = syscall.GetProcAddress(mb, "wkeSetCspCheckEnable")
			dllWkeSetViewNetInterface, _ = syscall.GetProcAddress(mb, "wkeSetViewNetInterface")*/
	dllWkeNetHookRequest, _ = syscall.GetProcAddress(mb, "wkeNetHookRequest")
	dllWkeOnLoadUrlBegin, _ = syscall.GetProcAddress(mb, "wkeOnLoadUrlBegin")
	dllWkeOnLoadUrlEnd, _ = syscall.GetProcAddress(mb, "wkeOnLoadUrlEnd") /*
		dllWkeOnDidCreateScriptContext, _ = syscall.GetProcAddress(mb, "wkeOnDidCreateScriptContext")
		dllWkeOnWillReleaseScriptContext, _ = syscall.GetProcAddress(mb, "wkeOnWillReleaseScriptContext")
		dllWkeNetOnResponse, _ = syscall.GetProcAddress(mb, "wkeNetOnResponse")
		dllWkeNetGetMIMEType, _ = syscall.GetProcAddress(mb, "wkeNetGetMIMEType")
		dllWkeWebFrameGetMainWorldScriptContext, _ = syscall.GetProcAddress(mb, "wkeWebFrameGetMainWorldScriptContext")
		dllWkeGetBlinkMainThreadIsolate, _ = syscall.GetProcAddress(mb, "wkeGetBlinkMainThreadIsolate")
		dllWkeCreateStringW, _ = syscall.GetProcAddress(mb, "wkeCreateStringW")
		dllWkeDeleteString, _ = syscall.GetProcAddress(mb, "wkeDeleteString")
		dllWkeSetLocalStorageFullPath, _ = syscall.GetProcAddress(mb, "wkeSetLocalStorageFullPath")*/
	dllWkeJsBindFunction, _ = syscall.GetProcAddress(mb, "wkeJsBindFunction") /*
		dllJsBindFunction, _ = syscall.GetProcAddress(mb, "jsBindFunction")
		dllWkeJsBindGetter, _ = syscall.GetProcAddress(mb, "wkeJsBindGetter")
		dllJsBindGetter, _ = syscall.GetProcAddress(mb, "jsBindGetter")
		dllWkeJsBindSetter, _ = syscall.GetProcAddress(mb, "wkeJsBindSetter")
		dllJsBindSetter, _ = syscall.GetProcAddress(mb, "jsBindSetter")
		dllJsArgCount, _ = syscall.GetProcAddress(mb, "jsArgCount")
		dllJsArgType, _ = syscall.GetProcAddress(mb, "jsArgType")*/
	dllJsArg, _ = syscall.GetProcAddress(mb, "jsArg") /*
		dllJsTypeOf, _ = syscall.GetProcAddress(mb, "jsTypeOf")
		dllJsIsNumber, _ = syscall.GetProcAddress(mb, "jsIsNumber")
		dllJsIsString, _ = syscall.GetProcAddress(mb, "jsIsString")
		dllJsIsBoolean, _ = syscall.GetProcAddress(mb, "jsIsBoolean")
		dllJsIsObject, _ = syscall.GetProcAddress(mb, "jsIsObject")
		dllJsIsFunction, _ = syscall.GetProcAddress(mb, "jsIsFunction")
		dllJsIsUndefined, _ = syscall.GetProcAddress(mb, "jsIsUndefined")
		dllJsIsNull, _ = syscall.GetProcAddress(mb, "jsIsNull")
		dllJsIsArray, _ = syscall.GetProcAddress(mb, "jsIsArray")*/
	dllJsIsTrue, _ = syscall.GetProcAddress(mb, "jsIsTrue")
	dllJsIsFalse, _ = syscall.GetProcAddress(mb, "jsIsFalse")
	dllJsToInt, _ = syscall.GetProcAddress(mb, "jsToInt")
	dllJsToFloat, _ = syscall.GetProcAddress(mb, "jsToFloat")
	dllJsToDouble, _ = syscall.GetProcAddress(mb, "jsToDouble") /*
		dllJsToBoolean, _ = syscall.GetProcAddress(mb, "jsToBoolean")*/
	dllJsToTempString, _ = syscall.GetProcAddress(mb, "jsToTempString") /*
		dllJsInt, _ = syscall.GetProcAddress(mb, "jsInt")
		dllJsFloat, _ = syscall.GetProcAddress(mb, "jsFloat")
		dllJsDouble, _ = syscall.GetProcAddress(mb, "jsDouble")
		dllJsBoolean, _ = syscall.GetProcAddress(mb, "jsBoolean")
		dllJsUndefined, _ = syscall.GetProcAddress(mb, "jsUndefined")
		dllJsNull, _ = syscall.GetProcAddress(mb, "jsNull")
		dllJsTrue, _ = syscall.GetProcAddress(mb, "jsTrue")
		dllJsFalse, _ = syscall.GetProcAddress(mb, "jsFalse")
		dllJsStringW, _ = syscall.GetProcAddress(mb, "jsStringW")
		dllJsEmptyObject, _ = syscall.GetProcAddress(mb, "jsEmptyObject")
		dllJsEmptyArray, _ = syscall.GetProcAddress(mb, "jsEmptyArray")
		dllJsObject, _ = syscall.GetProcAddress(mb, "jsObject")
		dllJsFunction, _ = syscall.GetProcAddress(mb, "jsFunction")
		dllJsGetData, _ = syscall.GetProcAddress(mb, "jsGetData")
		dllJsGet, _ = syscall.GetProcAddress(mb, "jsGet")
		dllJsSet, _ = syscall.GetProcAddress(mb, "jsSet")
		dllJsGetAt, _ = syscall.GetProcAddress(mb, "jsGetAt")
		dllJsSetAt, _ = syscall.GetProcAddress(mb, "jsSetAt")
		dllJsGetLength, _ = syscall.GetProcAddress(mb, "jsGetLength")
		dllJsSetLength, _ = syscall.GetProcAddress(mb, "jsSetLength")
		dllJsGlobalObject, _ = syscall.GetProcAddress(mb, "jsGlobalObject")
		dllJsGetWebView, _ = syscall.GetProcAddress(mb, "jsGetWebView")
		dllJsEvalW, _ = syscall.GetProcAddress(mb, "jsEvalW")
		dllJsEvalExW, _ = syscall.GetProcAddress(mb, "jsEvalExW")
		dllJsCall, _ = syscall.GetProcAddress(mb, "jsCall")
		dllJsCallGlobal, _ = syscall.GetProcAddress(mb, "jsCallGlobal")
		dllJsGetGlobal, _ = syscall.GetProcAddress(mb, "jsGetGlobal")
		dllJsSetGlobal, _ = syscall.GetProcAddress(mb, "jsSetGlobal")
		dllJsGC, _ = syscall.GetProcAddress(mb, "jsGC")*/
	dllJsToString, _ = syscall.GetProcAddress(mb, "jsToString")
	dllJsToStringW, _ = syscall.GetProcAddress(mb, "jsToStringW") /*
		dllJsArrayBuffer, _ = syscall.GetProcAddress(mb, "jsArrayBuffer")
		dllWkeSetFileSystem, _ = syscall.GetProcAddress(mb, "wkeSetFileSystem")*/
	dllWkeCreateWebView, _ = syscall.GetProcAddress(mb, "wkeCreateWebView")
	dllWkeGetWebView, _ = syscall.GetProcAddress(mb, "wkeGetWebView")
	dllWkeGetCaretRect, _ = syscall.GetProcAddress(mb, "wkeGetCaretRect") /*
		dllWkeDestroyWebView, _ = syscall.GetProcAddress(mb, "wkeDestroyWebView")
		dllWkeVisitAllCookie, _ = syscall.GetProcAddress(mb, "wkeVisitAllCookie")
		dllWkePerformCookieCommand, _ = syscall.GetProcAddress(mb, "wkePerformCookieCommand")
		dllWkeRunJsByFrame, _ = syscall.GetProcAddress(mb, "wkeRunJsByFrame")
		dllWkeSetNpapiPluginsEnabled, _ = syscall.GetProcAddress(mb, "wkeSetNpapiPluginsEnabled")
		dllWkeSetHeadlessEnabled, _ = syscall.GetProcAddress(mb, "wkeSetHeadlessEnabled")
		dllWkeSetTouchEnabled, _ = syscall.GetProcAddress(mb, "wkeSetTouchEnabled")
		dllWkeSetDragEnable, _ = syscall.GetProcAddress(mb, "wkeSetDragEnable")
		dllWkeSetMemoryCacheEnable, _ = syscall.GetProcAddress(mb, "wkeSetMemoryCacheEnable")
		dllWkeOnMouseOverUrlChanged, _ = syscall.GetProcAddress(mb, "wkeOnMouseOverUrlChanged")
		dllWkeSetDebugConfig, _ = syscall.GetProcAddress(mb, "wkeSetDebugConfig")
		dllWkeSetDeviceParameter, _ = syscall.GetProcAddress(mb, "wkeSetDeviceParameter")*/

}

func freeMBDLL() {
	syscall.FreeLibrary(mb)
	syscall.FreeLibrary(gonode)
	syscall.FreeLibrary(kr32)
	syscall.FreeLibrary(u32)
}
