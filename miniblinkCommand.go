package gowebui

import (
	"syscall"
	"unicode/utf16"
	"unsafe"
)

type point struct {
	X, Y int32
}
type msg struct {
	HWnd    uintptr
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      point
}

func getCurrentThreadId() uint32 {
	ret, _, _ := syscall.Syscall(dllGetCurrentThreadId, 0,
		0,
		0,
		0)

	return uint32(ret)
}

func postQuitMessage(exitCode int32) {
	syscall.Syscall(dllPostQuitMessage, 1,
		uintptr(exitCode),
		0,
		0)
}

func peekMessage(lpMsg *msg, hWnd uintptr, wMsgFilterMin, wMsgFilterMax, wRemoveMsg uint32) bool {
	ret, _, _ := syscall.Syscall6(dllPeekMessage, 5,
		uintptr(unsafe.Pointer(lpMsg)),
		uintptr(hWnd),
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMax),
		uintptr(wRemoveMsg),
		0)

	return ret != 0
}

func translateMessage(msg *msg) bool {
	ret, _, _ := syscall.Syscall(dllTranslateMessage, 1,
		uintptr(unsafe.Pointer(msg)),
		0,
		0)

	return ret != 0
}

func dispatchMessage(msg *msg) uintptr {
	ret, _, _ := syscall.Syscall(dllDispatchMessage, 1,
		uintptr(unsafe.Pointer(msg)),
		0,
		0)

	return ret
}

func miniblinkCommand(m_WebView WkeWebView, command string, args ...interface{}) interface{} {
	var ret uintptr
	switch command {
	case "wkeIsInitialize":
		ret, _, _ = syscall.Syscall(dllWkeIsInitialize, 0, 0, 0, 0)
		if ptrToInt32(ret) == 1 {
			return true
		} else {
			return false
		}
	case "wkeInitialize":
		ret, _, _ = syscall.Syscall(dllWkeInitialize, 0, 0, 0, 0)
	case "wkeDestroyWebWindow":
		ret, _, _ = syscall.Syscall(dllWkeDestroyWebWindow, 1, wkeWebViewToPtr(m_WebView), 0, 0)
	case "wkeLoadURL":
		ret, _, _ = syscall.Syscall(dllWkeLoadURL, 2, wkeWebViewToPtr(m_WebView), strToPtr(args[0].(string)), 0)
	case "wkeShowWindow":
		ret, _, _ = syscall.Syscall(dllWkeShowWindow, 2, wkeWebViewToPtr(m_WebView), args[0].(uintptr), 0)
	case "wkeMoveToCenter":
		ret, _, _ = syscall.Syscall(dllWkeMoveToCenter, 1, wkeWebViewToPtr(m_WebView), 0, 0)
	case "wkeGetURL":
		ret, _, _ = syscall.Syscall(dllWkeGetURL, 1, wkeWebViewToPtr(m_WebView), 0, 0)
		return PtrToString(ret)
	case "wkeGetTitle":
		ret, _, _ = syscall.Syscall(dllWkeGetTitle, 1, wkeWebViewToPtr(m_WebView), 0, 0)
		return PtrToString(ret)
	case "wkeCreateWebWindow":
		ret, _, _ = syscall.Syscall6(dllWkeCreateWebWindow, 6, int32ToPtr(args[0].(int32)), int32ToPtr(args[1].(int32)), int32ToPtr(args[2].(int32)), int32ToPtr(args[3].(int32)), int32ToPtr(args[4].(int32)), int32ToPtr(args[5].(int32)))
		return WkeWebView(ptrToInt32(ret))
	case "wkeGetWindowHandle":
		ret, _, _ = syscall.Syscall(dllWkeGetWindowHandle, 1, wkeWebViewToPtr(m_WebView), 0, 0)
		return HWnd(ptrToInt32(ret))
	case "wkeSetWindowTitle":
		ret, _, _ = syscall.Syscall(dllWkeSetWindowTitleW, 2, wkeWebViewToPtr(m_WebView), strToUTF16Ptr(args[0].(string)), 0)
	case "wkeLoadHTML":
		ret, _, _ = syscall.Syscall(dllWkeLoadHTML, 2, wkeWebViewToPtr(m_WebView), strToPtr(args[0].(string)), 0)
	case "jsArg":
		ret, _, _ = syscall.Syscall(GOjsArg, 2, jsExecStateToPtr(args[0].(JsExecState)), int32ToPtr(args[1].(int32)), 0)
		return JsValue(PtrToString(ret))
	case "jsToString":
		ret, _, _ = syscall.Syscall(GOjsToString, 2, jsExecStateToPtr(args[0].(JsExecState)), jsValueToPtr(args[1].(JsValue)), 0)
		return PtrToString(ret)
	case "wkeGlobalExec":
		ret, _, _ = syscall.Syscall(dllWkeGlobalExec, 1, wkeWebViewToPtr(m_WebView), 0, 0)
		return JsExecState(ptrToUint32(&ret))
	case "wkeRunJS":
		ret, _, _ = syscall.Syscall(GOwkeRunJS, 2, wkeWebViewToPtr(m_WebView), strToPtr(args[0].(string)), 0)
		return ptrToJsValue(ret)
	case "jsToInt":
		ret, _, _ = syscall.Syscall(GOjsToInt, 2, jsExecStateToPtr(args[0].(JsExecState)), jsValueToPtr(args[1].(JsValue)), 0)
		return ptrToInt32(ret)
	case "jsToDouble":
		ret, _, _ = syscall.Syscall(GOjsToDouble, 2, jsExecStateToPtr(args[0].(JsExecState)), jsValueToPtr(args[1].(JsValue)), 0)
		return ptrToFloat64(ret)
	case "jsIsTrue":
		ret, _, _ = syscall.Syscall(GOjsIsTrue, 1, jsValueToPtr(args[0].(JsValue)), 0, 0)
		if ptrToInt32(ret) == 1 {
			return true
		} else {
			return false
		}
	case "wkeIsDocumentReady":
		ret, _, _ = syscall.Syscall(dllWkeIsDocumentReady, 1, wkeWebViewToPtr(m_WebView), 0, 0)
		if ptrToInt32(ret) == 1 {
			return true
		} else {
			return false
		}
	case "wkeGetCookie":
		ret, _, _ = syscall.Syscall(dllWkeGetCookie, 1, wkeWebViewToPtr(m_WebView), 0, 0)
		return PtrToString(ret)
	case "wkeGetUserAgent":
		ret, _, _ = syscall.Syscall(dllWkeGetUserAgent, 1, wkeWebViewToPtr(m_WebView), 0, 0)
		return PtrToString(ret)
	case "wkeGetString":
		ret, _, _ = syscall.Syscall(dllWkeGetString, 1, wkeStringToPtr(args[0].(WkeString)), 0, 0)
		return PtrToString(ret)
	case "wkeNetHookRequest":
		ret, _, _ = syscall.Syscall(dllWkeNetHookRequest, 1, args[0].(uintptr), 0, 0)
		return 0
	case "wkeNetCancelRequest":
		ret, _, _ = syscall.Syscall(dllWkeNetCancelRequest, 1, args[0].(uintptr), 0, 0)
		return 0
	case "wkeSetNavigationToNewWindowEnable":
		ret, _, _ = syscall.Syscall(GOwkeSetNavigationToNewWindowEnable, 2, wkeWebViewToPtr(m_WebView), args[0].(uintptr), 0)
		return 0
	case "wkeCanGoBack":
		ret, _, _ = syscall.Syscall(GOwkeCanGoBack, 1, wkeWebViewToPtr(m_WebView), 0, 0)
		if ptrToInt32(ret) == 1 {
			return true
		} else {
			return false
		}
	case "wkeGoBack":
		ret, _, _ = syscall.Syscall(dllWkeGoBack, 1, wkeWebViewToPtr(m_WebView), 0, 0)
		return 0
	case "wkeCanGoForward":
		ret, _, _ = syscall.Syscall(GOwkeCanGoForward, 1, wkeWebViewToPtr(m_WebView), 0, 0)
		if ptrToInt32(ret) == 1 {
			return true
		} else {
			return false
		}
	case "wkeGoForward":
		ret, _, _ = syscall.Syscall(dllWkeGoForward, 1, wkeWebViewToPtr(m_WebView), 0, 0)
		return 0
	case "wkeReload":
		ret, _, _ = syscall.Syscall(dllWkeReload, 1, wkeWebViewToPtr(m_WebView), 0, 0)
		return 0
	default:
	}
	return ret
}

//以下为类型转换
func strToPtr(s string) uintptr {
	url := []byte(s)
	url = append(url, byte(0))
	return uintptr(unsafe.Pointer(&url[0]))
}
func strToUTF16Ptr(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
}

func PtrToString(ptr uintptr) string {
	p := (*byte)(unsafe.Pointer(ptr))
	data := make([]byte, 0)
	for *p != 0 {
		data = append(data, *p)
		ptr += unsafe.Sizeof(byte(0))
		p = (*byte)(unsafe.Pointer(ptr))
	}
	return string(data)
}

func utf16ptrToString(po *uintptr) string {
	p := (*uint16)(unsafe.Pointer(po))
	ptr := *po
	data := make([]uint16, 0)
	for *p != 0 {
		data = append(data, *p)
		ptr += unsafe.Sizeof(uint16(0))
		p = (*uint16)(unsafe.Pointer(ptr))
	}
	return string(utf16.Decode(data))
}

func ptrToUint32(p *uintptr) uint32 {
	return *(*uint32)(unsafe.Pointer(p))
}

func ptrToInt32(p uintptr) int32 {
	return int32(p)
}

func ptrToFloat64(p uintptr) float64 {
	return *(*float64)(unsafe.Pointer(p))
}

func int32ToPtr(i int32) uintptr {
	return uintptr(i)
}

func wkeStringToPtr(s WkeString) uintptr {
	return uintptr(s)
}

func int64ToPtr(i int64) uintptr {
	return uintptr(unsafe.Pointer(&i))
}

func ptrToInt64(p *uintptr) int64 {
	return int64(*(*uint64)(unsafe.Pointer(p)))
}

func ptrToInt641(p uintptr) int64 {
	return int64(*(*uint64)(unsafe.Pointer(&p)))
}

func jsExecStateToPtr(es JsExecState) uintptr {
	return uintptr(es)
}

func wkeWebViewToPtr(w WkeWebView) uintptr {
	return uintptr(w)
}

func jsValueToPtr(v JsValue) uintptr {
	url := []byte(v)
	url = append(url, byte(0))
	return uintptr(unsafe.Pointer(&url[0]))
}

func ptrToJsValue(ptr uintptr) JsValue {
	p := (*byte)(unsafe.Pointer(ptr))
	data := make([]byte, 0)
	for *p != 0 {
		data = append(data, *p)
		ptr += unsafe.Sizeof(byte(0))
		p = (*byte)(unsafe.Pointer(ptr))
	}
	return JsValue(string(data))
}

func hWndToPtr(h HWnd) uintptr {
	return uintptr(h)
}
