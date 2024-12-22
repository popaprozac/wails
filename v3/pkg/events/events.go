package events

type ApplicationEventType uint
type WindowEventType uint

var Common = newCommonEvents()

type commonEvents struct {
	ApplicationStarted        ApplicationEventType
	WindowMaximise            WindowEventType
	WindowUnMaximise          WindowEventType
	WindowFullscreen          WindowEventType
	WindowUnFullscreen        WindowEventType
	WindowRestore             WindowEventType
	WindowMinimise            WindowEventType
	WindowUnMinimise          WindowEventType
	WindowClosing             WindowEventType
	WindowZoom                WindowEventType
	WindowZoomIn              WindowEventType
	WindowZoomOut             WindowEventType
	WindowZoomReset           WindowEventType
	WindowFocus               WindowEventType
	WindowLostFocus           WindowEventType
	WindowShow                WindowEventType
	WindowHide                WindowEventType
	WindowDPIChanged          WindowEventType
	WindowFilesDropped        WindowEventType
	WindowRuntimeReady        WindowEventType
	ThemeChanged              ApplicationEventType
	WindowDidMove             WindowEventType
	WindowDidResize           WindowEventType
	ApplicationOpenedWithFile ApplicationEventType
}

func newCommonEvents() commonEvents {
	return commonEvents{
		ApplicationStarted:        1203,
		WindowMaximise:            1204,
		WindowUnMaximise:          1205,
		WindowFullscreen:          1206,
		WindowUnFullscreen:        1207,
		WindowRestore:             1208,
		WindowMinimise:            1209,
		WindowUnMinimise:          1210,
		WindowClosing:             1211,
		WindowZoom:                1212,
		WindowZoomIn:              1213,
		WindowZoomOut:             1214,
		WindowZoomReset:           1215,
		WindowFocus:               1216,
		WindowLostFocus:           1217,
		WindowShow:                1218,
		WindowHide:                1219,
		WindowDPIChanged:          1220,
		WindowFilesDropped:        1221,
		WindowRuntimeReady:        1222,
		ThemeChanged:              1223,
		WindowDidMove:             1224,
		WindowDidResize:           1225,
		ApplicationOpenedWithFile: 1226,
	}
}

var Linux = newLinuxEvents()

type linuxEvents struct {
	SystemThemeChanged ApplicationEventType
	WindowLoadChanged  WindowEventType
	WindowDeleteEvent  WindowEventType
	WindowDidMove      WindowEventType
	WindowDidResize    WindowEventType
	WindowFocusIn      WindowEventType
	WindowFocusOut     WindowEventType
	ApplicationStartup ApplicationEventType
}

func newLinuxEvents() linuxEvents {
	return linuxEvents{
		SystemThemeChanged: 1024,
		WindowLoadChanged:  1025,
		WindowDeleteEvent:  1026,
		WindowDidMove:      1027,
		WindowDidResize:    1028,
		WindowFocusIn:      1029,
		WindowFocusOut:     1030,
		ApplicationStartup: 1031,
	}
}

var Mac = newMacEvents()

type macEvents struct {
	ApplicationDidBecomeActive                              ApplicationEventType
	ApplicationDidChangeBackingProperties                   ApplicationEventType
	ApplicationDidChangeEffectiveAppearance                 ApplicationEventType
	ApplicationDidChangeIcon                                ApplicationEventType
	ApplicationDidChangeOcclusionState                      ApplicationEventType
	ApplicationDidChangeScreenParameters                    ApplicationEventType
	ApplicationDidChangeStatusBarFrame                      ApplicationEventType
	ApplicationDidChangeStatusBarOrientation                ApplicationEventType
	ApplicationDidFinishLaunching                           ApplicationEventType
	ApplicationDidHide                                      ApplicationEventType
	ApplicationDidResignActiveNotification                  ApplicationEventType
	ApplicationDidUnhide                                    ApplicationEventType
	ApplicationDidUpdate                                    ApplicationEventType
	ApplicationWillBecomeActive                             ApplicationEventType
	ApplicationWillFinishLaunching                          ApplicationEventType
	ApplicationWillHide                                     ApplicationEventType
	ApplicationWillResignActive                             ApplicationEventType
	ApplicationWillTerminate                                ApplicationEventType
	ApplicationWillUnhide                                   ApplicationEventType
	ApplicationWillUpdate                                   ApplicationEventType
	ApplicationDidChangeTheme                               ApplicationEventType
	ApplicationShouldHandleReopen                           ApplicationEventType
	WindowDidBecomeKey                                      WindowEventType
	WindowDidBecomeMain                                     WindowEventType
	WindowDidBeginSheet                                     WindowEventType
	WindowDidChangeAlpha                                    WindowEventType
	WindowDidChangeBackingLocation                          WindowEventType
	WindowDidChangeBackingProperties                        WindowEventType
	WindowDidChangeCollectionBehavior                       WindowEventType
	WindowDidChangeEffectiveAppearance                      WindowEventType
	WindowDidChangeOcclusionState                           WindowEventType
	WindowDidChangeOrderingMode                             WindowEventType
	WindowDidChangeScreen                                   WindowEventType
	WindowDidChangeScreenParameters                         WindowEventType
	WindowDidChangeScreenProfile                            WindowEventType
	WindowDidChangeScreenSpace                              WindowEventType
	WindowDidChangeScreenSpaceProperties                    WindowEventType
	WindowDidChangeSharingType                              WindowEventType
	WindowDidChangeSpace                                    WindowEventType
	WindowDidChangeSpaceOrderingMode                        WindowEventType
	WindowDidChangeTitle                                    WindowEventType
	WindowDidChangeToolbar                                  WindowEventType
	WindowDidDeminiaturize                                  WindowEventType
	WindowDidEndSheet                                       WindowEventType
	WindowDidEnterFullScreen                                WindowEventType
	WindowMaximise                                          WindowEventType
	WindowUnMaximise                                        WindowEventType
	WindowDidZoom                                           WindowEventType
	WindowZoomIn                                            WindowEventType
	WindowZoomOut                                           WindowEventType
	WindowZoomReset                                         WindowEventType
	WindowDidEnterVersionBrowser                            WindowEventType
	WindowDidExitFullScreen                                 WindowEventType
	WindowDidExitVersionBrowser                             WindowEventType
	WindowDidExpose                                         WindowEventType
	WindowDidFocus                                          WindowEventType
	WindowDidMiniaturize                                    WindowEventType
	WindowDidMove                                           WindowEventType
	WindowDidOrderOffScreen                                 WindowEventType
	WindowDidOrderOnScreen                                  WindowEventType
	WindowDidResignKey                                      WindowEventType
	WindowDidResignMain                                     WindowEventType
	WindowDidResize                                         WindowEventType
	WindowDidUpdate                                         WindowEventType
	WindowDidUpdateAlpha                                    WindowEventType
	WindowDidUpdateCollectionBehavior                       WindowEventType
	WindowDidUpdateCollectionProperties                     WindowEventType
	WindowDidUpdateShadow                                   WindowEventType
	WindowDidUpdateTitle                                    WindowEventType
	WindowDidUpdateToolbar                                  WindowEventType
	WindowShouldClose                                       WindowEventType
	WindowWillBecomeKey                                     WindowEventType
	WindowWillBecomeMain                                    WindowEventType
	WindowWillBeginSheet                                    WindowEventType
	WindowWillChangeOrderingMode                            WindowEventType
	WindowWillClose                                         WindowEventType
	WindowWillDeminiaturize                                 WindowEventType
	WindowWillEnterFullScreen                               WindowEventType
	WindowWillEnterVersionBrowser                           WindowEventType
	WindowWillExitFullScreen                                WindowEventType
	WindowWillExitVersionBrowser                            WindowEventType
	WindowWillFocus                                         WindowEventType
	WindowWillMiniaturize                                   WindowEventType
	WindowWillMove                                          WindowEventType
	WindowWillOrderOffScreen                                WindowEventType
	WindowWillOrderOnScreen                                 WindowEventType
	WindowWillResignMain                                    WindowEventType
	WindowWillResize                                        WindowEventType
	WindowWillUnfocus                                       WindowEventType
	WindowWillUpdate                                        WindowEventType
	WindowWillUpdateAlpha                                   WindowEventType
	WindowWillUpdateCollectionBehavior                      WindowEventType
	WindowWillUpdateCollectionProperties                    WindowEventType
	WindowWillUpdateShadow                                  WindowEventType
	WindowWillUpdateTitle                                   WindowEventType
	WindowWillUpdateToolbar                                 WindowEventType
	WindowWillUpdateVisibility                              WindowEventType
	WindowWillUseStandardFrame                              WindowEventType
	MenuWillOpen                                            ApplicationEventType
	MenuDidOpen                                             ApplicationEventType
	MenuDidClose                                            ApplicationEventType
	MenuWillSendAction                                      ApplicationEventType
	MenuDidSendAction                                       ApplicationEventType
	MenuWillHighlightItem                                   ApplicationEventType
	MenuDidHighlightItem                                    ApplicationEventType
	MenuWillDisplayItem                                     ApplicationEventType
	MenuDidDisplayItem                                      ApplicationEventType
	MenuWillAddItem                                         ApplicationEventType
	MenuDidAddItem                                          ApplicationEventType
	MenuWillRemoveItem                                      ApplicationEventType
	MenuDidRemoveItem                                       ApplicationEventType
	MenuWillBeginTracking                                   ApplicationEventType
	MenuDidBeginTracking                                    ApplicationEventType
	MenuWillEndTracking                                     ApplicationEventType
	MenuDidEndTracking                                      ApplicationEventType
	MenuWillUpdate                                          ApplicationEventType
	MenuDidUpdate                                           ApplicationEventType
	MenuWillPopUp                                           ApplicationEventType
	MenuDidPopUp                                            ApplicationEventType
	MenuWillSendActionToItem                                ApplicationEventType
	MenuDidSendActionToItem                                 ApplicationEventType
	WebViewDidStartProvisionalNavigation                    WindowEventType
	WebViewDidReceiveServerRedirectForProvisionalNavigation WindowEventType
	WebViewDidFinishNavigation                              WindowEventType
	WebViewDidCommitNavigation                              WindowEventType
	WindowFileDraggingEntered                               WindowEventType
	WindowFileDraggingPerformed                             WindowEventType
	WindowFileDraggingExited                                WindowEventType
}

func newMacEvents() macEvents {
	return macEvents{
		ApplicationDidBecomeActive:               1032,
		ApplicationDidChangeBackingProperties:    1033,
		ApplicationDidChangeEffectiveAppearance:  1034,
		ApplicationDidChangeIcon:                 1035,
		ApplicationDidChangeOcclusionState:       1036,
		ApplicationDidChangeScreenParameters:     1037,
		ApplicationDidChangeStatusBarFrame:       1038,
		ApplicationDidChangeStatusBarOrientation: 1039,
		ApplicationDidFinishLaunching:            1040,
		ApplicationDidHide:                       1041,
		ApplicationDidResignActiveNotification:   1042,
		ApplicationDidUnhide:                     1043,
		ApplicationDidUpdate:                     1044,
		ApplicationWillBecomeActive:              1045,
		ApplicationWillFinishLaunching:           1046,
		ApplicationWillHide:                      1047,
		ApplicationWillResignActive:              1048,
		ApplicationWillTerminate:                 1049,
		ApplicationWillUnhide:                    1050,
		ApplicationWillUpdate:                    1051,
		ApplicationDidChangeTheme:                1052,
		ApplicationShouldHandleReopen:            1053,
		WindowDidBecomeKey:                       1054,
		WindowDidBecomeMain:                      1055,
		WindowDidBeginSheet:                      1056,
		WindowDidChangeAlpha:                     1057,
		WindowDidChangeBackingLocation:           1058,
		WindowDidChangeBackingProperties:         1059,
		WindowDidChangeCollectionBehavior:        1060,
		WindowDidChangeEffectiveAppearance:       1061,
		WindowDidChangeOcclusionState:            1062,
		WindowDidChangeOrderingMode:              1063,
		WindowDidChangeScreen:                    1064,
		WindowDidChangeScreenParameters:          1065,
		WindowDidChangeScreenProfile:             1066,
		WindowDidChangeScreenSpace:               1067,
		WindowDidChangeScreenSpaceProperties:     1068,
		WindowDidChangeSharingType:               1069,
		WindowDidChangeSpace:                     1070,
		WindowDidChangeSpaceOrderingMode:         1071,
		WindowDidChangeTitle:                     1072,
		WindowDidChangeToolbar:                   1073,
		WindowDidDeminiaturize:                   1074,
		WindowDidEndSheet:                        1075,
		WindowDidEnterFullScreen:                 1076,
		WindowMaximise:                           1077,
		WindowUnMaximise:                         1078,
		WindowDidZoom:                            1079,
		WindowZoomIn:                             1080,
		WindowZoomOut:                            1081,
		WindowZoomReset:                          1082,
		WindowDidEnterVersionBrowser:             1083,
		WindowDidExitFullScreen:                  1084,
		WindowDidExitVersionBrowser:              1085,
		WindowDidExpose:                          1086,
		WindowDidFocus:                           1087,
		WindowDidMiniaturize:                     1088,
		WindowDidMove:                            1089,
		WindowDidOrderOffScreen:                  1090,
		WindowDidOrderOnScreen:                   1091,
		WindowDidResignKey:                       1092,
		WindowDidResignMain:                      1093,
		WindowDidResize:                          1094,
		WindowDidUpdate:                          1095,
		WindowDidUpdateAlpha:                     1096,
		WindowDidUpdateCollectionBehavior:        1097,
		WindowDidUpdateCollectionProperties:      1098,
		WindowDidUpdateShadow:                    1099,
		WindowDidUpdateTitle:                     1100,
		WindowDidUpdateToolbar:                   1101,
		WindowShouldClose:                        1102,
		WindowWillBecomeKey:                      1103,
		WindowWillBecomeMain:                     1104,
		WindowWillBeginSheet:                     1105,
		WindowWillChangeOrderingMode:             1106,
		WindowWillClose:                          1107,
		WindowWillDeminiaturize:                  1108,
		WindowWillEnterFullScreen:                1109,
		WindowWillEnterVersionBrowser:            1110,
		WindowWillExitFullScreen:                 1111,
		WindowWillExitVersionBrowser:             1112,
		WindowWillFocus:                          1113,
		WindowWillMiniaturize:                    1114,
		WindowWillMove:                           1115,
		WindowWillOrderOffScreen:                 1116,
		WindowWillOrderOnScreen:                  1117,
		WindowWillResignMain:                     1118,
		WindowWillResize:                         1119,
		WindowWillUnfocus:                        1120,
		WindowWillUpdate:                         1121,
		WindowWillUpdateAlpha:                    1122,
		WindowWillUpdateCollectionBehavior:       1123,
		WindowWillUpdateCollectionProperties:     1124,
		WindowWillUpdateShadow:                   1125,
		WindowWillUpdateTitle:                    1126,
		WindowWillUpdateToolbar:                  1127,
		WindowWillUpdateVisibility:               1128,
		WindowWillUseStandardFrame:               1129,
		MenuWillOpen:                             1130,
		MenuDidOpen:                              1131,
		MenuDidClose:                             1132,
		MenuWillSendAction:                       1133,
		MenuDidSendAction:                        1134,
		MenuWillHighlightItem:                    1135,
		MenuDidHighlightItem:                     1136,
		MenuWillDisplayItem:                      1137,
		MenuDidDisplayItem:                       1138,
		MenuWillAddItem:                          1139,
		MenuDidAddItem:                           1140,
		MenuWillRemoveItem:                       1141,
		MenuDidRemoveItem:                        1142,
		MenuWillBeginTracking:                    1143,
		MenuDidBeginTracking:                     1144,
		MenuWillEndTracking:                      1145,
		MenuDidEndTracking:                       1146,
		MenuWillUpdate:                           1147,
		MenuDidUpdate:                            1148,
		MenuWillPopUp:                            1149,
		MenuDidPopUp:                             1150,
		MenuWillSendActionToItem:                 1151,
		MenuDidSendActionToItem:                  1152,
		WebViewDidStartProvisionalNavigation:     1153,
		WebViewDidReceiveServerRedirectForProvisionalNavigation: 1154,
		WebViewDidFinishNavigation:                              1155,
		WebViewDidCommitNavigation:                              1156,
		WindowFileDraggingEntered:                               1157,
		WindowFileDraggingPerformed:                             1158,
		WindowFileDraggingExited:                                1159,
	}
}

var Windows = newWindowsEvents()

type windowsEvents struct {
	SystemThemeChanged         ApplicationEventType
	APMPowerStatusChange       ApplicationEventType
	APMSuspend                 ApplicationEventType
	APMResumeAutomatic         ApplicationEventType
	APMResumeSuspend           ApplicationEventType
	APMPowerSettingChange      ApplicationEventType
	ApplicationStarted         ApplicationEventType
	WebViewNavigationCompleted WindowEventType
	WindowInactive             WindowEventType
	WindowActive               WindowEventType
	WindowClickActive          WindowEventType
	WindowMaximise             WindowEventType
	WindowUnMaximise           WindowEventType
	WindowFullscreen           WindowEventType
	WindowUnFullscreen         WindowEventType
	WindowRestore              WindowEventType
	WindowMinimise             WindowEventType
	WindowUnMinimise           WindowEventType
	WindowClosing              WindowEventType
	WindowSetFocus             WindowEventType
	WindowKillFocus            WindowEventType
	WindowDragDrop             WindowEventType
	WindowDragEnter            WindowEventType
	WindowDragLeave            WindowEventType
	WindowDragOver             WindowEventType
	WindowDidMove              WindowEventType
	WindowDidResize            WindowEventType
	WindowShow                 WindowEventType
	WindowHide                 WindowEventType
	WindowStartMove            WindowEventType
	WindowEndMove              WindowEventType
	WindowStartResize          WindowEventType
	WindowEndResize            WindowEventType
	WindowKeyDown              WindowEventType
	WindowKeyUp                WindowEventType
	WindowZOrderChanged        WindowEventType
	WindowPaint                WindowEventType
	WindowBackgroundErase      WindowEventType
	WindowNonClientHit         WindowEventType
	WindowNonClientMouseDown   WindowEventType
	WindowNonClientMouseUp     WindowEventType
	WindowNonClientMouseMove   WindowEventType
	WindowNonClientMouseLeave  WindowEventType
}

func newWindowsEvents() windowsEvents {
	return windowsEvents{
		SystemThemeChanged:         1160,
		APMPowerStatusChange:       1161,
		APMSuspend:                 1162,
		APMResumeAutomatic:         1163,
		APMResumeSuspend:           1164,
		APMPowerSettingChange:      1165,
		ApplicationStarted:         1166,
		WebViewNavigationCompleted: 1167,
		WindowInactive:             1168,
		WindowActive:               1169,
		WindowClickActive:          1170,
		WindowMaximise:             1171,
		WindowUnMaximise:           1172,
		WindowFullscreen:           1173,
		WindowUnFullscreen:         1174,
		WindowRestore:              1175,
		WindowMinimise:             1176,
		WindowUnMinimise:           1177,
		WindowClosing:              1178,
		WindowSetFocus:             1179,
		WindowKillFocus:            1180,
		WindowDragDrop:             1181,
		WindowDragEnter:            1182,
		WindowDragLeave:            1183,
		WindowDragOver:             1184,
		WindowDidMove:              1185,
		WindowDidResize:            1186,
		WindowShow:                 1187,
		WindowHide:                 1188,
		WindowStartMove:            1189,
		WindowEndMove:              1190,
		WindowStartResize:          1191,
		WindowEndResize:            1192,
		WindowKeyDown:              1193,
		WindowKeyUp:                1194,
		WindowZOrderChanged:        1195,
		WindowPaint:                1196,
		WindowBackgroundErase:      1197,
		WindowNonClientHit:         1198,
		WindowNonClientMouseDown:   1199,
		WindowNonClientMouseUp:     1200,
		WindowNonClientMouseMove:   1201,
		WindowNonClientMouseLeave:  1202,
	}
}

func JSEvent(event uint) string {
	return eventToJS[event]
}

var eventToJS = map[uint]string{
	1024: "linux:SystemThemeChanged",
	1025: "linux:WindowLoadChanged",
	1026: "linux:WindowDeleteEvent",
	1027: "linux:WindowDidMove",
	1028: "linux:WindowDidResize",
	1029: "linux:WindowFocusIn",
	1030: "linux:WindowFocusOut",
	1031: "linux:ApplicationStartup",
	1032: "mac:ApplicationDidBecomeActive",
	1033: "mac:ApplicationDidChangeBackingProperties",
	1034: "mac:ApplicationDidChangeEffectiveAppearance",
	1035: "mac:ApplicationDidChangeIcon",
	1036: "mac:ApplicationDidChangeOcclusionState",
	1037: "mac:ApplicationDidChangeScreenParameters",
	1038: "mac:ApplicationDidChangeStatusBarFrame",
	1039: "mac:ApplicationDidChangeStatusBarOrientation",
	1040: "mac:ApplicationDidFinishLaunching",
	1041: "mac:ApplicationDidHide",
	1042: "mac:ApplicationDidResignActiveNotification",
	1043: "mac:ApplicationDidUnhide",
	1044: "mac:ApplicationDidUpdate",
	1045: "mac:ApplicationWillBecomeActive",
	1046: "mac:ApplicationWillFinishLaunching",
	1047: "mac:ApplicationWillHide",
	1048: "mac:ApplicationWillResignActive",
	1049: "mac:ApplicationWillTerminate",
	1050: "mac:ApplicationWillUnhide",
	1051: "mac:ApplicationWillUpdate",
	1052: "mac:ApplicationDidChangeTheme!",
	1053: "mac:ApplicationShouldHandleReopen!",
	1054: "mac:WindowDidBecomeKey",
	1055: "mac:WindowDidBecomeMain",
	1056: "mac:WindowDidBeginSheet",
	1057: "mac:WindowDidChangeAlpha",
	1058: "mac:WindowDidChangeBackingLocation",
	1059: "mac:WindowDidChangeBackingProperties",
	1060: "mac:WindowDidChangeCollectionBehavior",
	1061: "mac:WindowDidChangeEffectiveAppearance",
	1062: "mac:WindowDidChangeOcclusionState",
	1063: "mac:WindowDidChangeOrderingMode",
	1064: "mac:WindowDidChangeScreen",
	1065: "mac:WindowDidChangeScreenParameters",
	1066: "mac:WindowDidChangeScreenProfile",
	1067: "mac:WindowDidChangeScreenSpace",
	1068: "mac:WindowDidChangeScreenSpaceProperties",
	1069: "mac:WindowDidChangeSharingType",
	1070: "mac:WindowDidChangeSpace",
	1071: "mac:WindowDidChangeSpaceOrderingMode",
	1072: "mac:WindowDidChangeTitle",
	1073: "mac:WindowDidChangeToolbar",
	1074: "mac:WindowDidDeminiaturize",
	1075: "mac:WindowDidEndSheet",
	1076: "mac:WindowDidEnterFullScreen",
	1077: "mac:WindowMaximise",
	1078: "mac:WindowUnMaximise",
	1079: "mac:WindowDidZoom!",
	1080: "mac:WindowZoomIn!",
	1081: "mac:WindowZoomOut!",
	1082: "mac:WindowZoomReset!",
	1083: "mac:WindowDidEnterVersionBrowser",
	1084: "mac:WindowDidExitFullScreen",
	1085: "mac:WindowDidExitVersionBrowser",
	1086: "mac:WindowDidExpose",
	1087: "mac:WindowDidFocus",
	1088: "mac:WindowDidMiniaturize",
	1089: "mac:WindowDidMove",
	1090: "mac:WindowDidOrderOffScreen",
	1091: "mac:WindowDidOrderOnScreen",
	1092: "mac:WindowDidResignKey",
	1093: "mac:WindowDidResignMain",
	1094: "mac:WindowDidResize",
	1095: "mac:WindowDidUpdate",
	1096: "mac:WindowDidUpdateAlpha",
	1097: "mac:WindowDidUpdateCollectionBehavior",
	1098: "mac:WindowDidUpdateCollectionProperties",
	1099: "mac:WindowDidUpdateShadow",
	1100: "mac:WindowDidUpdateTitle",
	1101: "mac:WindowDidUpdateToolbar",
	1102: "mac:WindowShouldClose!",
	1103: "mac:WindowWillBecomeKey",
	1104: "mac:WindowWillBecomeMain",
	1105: "mac:WindowWillBeginSheet",
	1106: "mac:WindowWillChangeOrderingMode",
	1107: "mac:WindowWillClose",
	1108: "mac:WindowWillDeminiaturize",
	1109: "mac:WindowWillEnterFullScreen",
	1110: "mac:WindowWillEnterVersionBrowser",
	1111: "mac:WindowWillExitFullScreen",
	1112: "mac:WindowWillExitVersionBrowser",
	1113: "mac:WindowWillFocus",
	1114: "mac:WindowWillMiniaturize",
	1115: "mac:WindowWillMove",
	1116: "mac:WindowWillOrderOffScreen",
	1117: "mac:WindowWillOrderOnScreen",
	1118: "mac:WindowWillResignMain",
	1119: "mac:WindowWillResize",
	1120: "mac:WindowWillUnfocus",
	1121: "mac:WindowWillUpdate",
	1122: "mac:WindowWillUpdateAlpha",
	1123: "mac:WindowWillUpdateCollectionBehavior",
	1124: "mac:WindowWillUpdateCollectionProperties",
	1125: "mac:WindowWillUpdateShadow",
	1126: "mac:WindowWillUpdateTitle",
	1127: "mac:WindowWillUpdateToolbar",
	1128: "mac:WindowWillUpdateVisibility",
	1129: "mac:WindowWillUseStandardFrame",
	1130: "mac:MenuWillOpen",
	1131: "mac:MenuDidOpen",
	1132: "mac:MenuDidClose",
	1133: "mac:MenuWillSendAction",
	1134: "mac:MenuDidSendAction",
	1135: "mac:MenuWillHighlightItem",
	1136: "mac:MenuDidHighlightItem",
	1137: "mac:MenuWillDisplayItem",
	1138: "mac:MenuDidDisplayItem",
	1139: "mac:MenuWillAddItem",
	1140: "mac:MenuDidAddItem",
	1141: "mac:MenuWillRemoveItem",
	1142: "mac:MenuDidRemoveItem",
	1143: "mac:MenuWillBeginTracking",
	1144: "mac:MenuDidBeginTracking",
	1145: "mac:MenuWillEndTracking",
	1146: "mac:MenuDidEndTracking",
	1147: "mac:MenuWillUpdate",
	1148: "mac:MenuDidUpdate",
	1149: "mac:MenuWillPopUp",
	1150: "mac:MenuDidPopUp",
	1151: "mac:MenuWillSendActionToItem",
	1152: "mac:MenuDidSendActionToItem",
	1153: "mac:WebViewDidStartProvisionalNavigation",
	1154: "mac:WebViewDidReceiveServerRedirectForProvisionalNavigation",
	1155: "mac:WebViewDidFinishNavigation",
	1156: "mac:WebViewDidCommitNavigation",
	1157: "mac:WindowFileDraggingEntered",
	1158: "mac:WindowFileDraggingPerformed",
	1159: "mac:WindowFileDraggingExited",
	1160: "windows:SystemThemeChanged",
	1161: "windows:APMPowerStatusChange",
	1162: "windows:APMSuspend",
	1163: "windows:APMResumeAutomatic",
	1164: "windows:APMResumeSuspend",
	1165: "windows:APMPowerSettingChange",
	1166: "windows:ApplicationStarted",
	1167: "windows:WebViewNavigationCompleted",
	1168: "windows:WindowInactive",
	1169: "windows:WindowActive",
	1170: "windows:WindowClickActive",
	1171: "windows:WindowMaximise",
	1172: "windows:WindowUnMaximise",
	1173: "windows:WindowFullscreen",
	1174: "windows:WindowUnFullscreen",
	1175: "windows:WindowRestore",
	1176: "windows:WindowMinimise",
	1177: "windows:WindowUnMinimise",
	1178: "windows:WindowClosing",
	1179: "windows:WindowSetFocus",
	1180: "windows:WindowKillFocus",
	1181: "windows:WindowDragDrop",
	1182: "windows:WindowDragEnter",
	1183: "windows:WindowDragLeave",
	1184: "windows:WindowDragOver",
	1185: "windows:WindowDidMove",
	1186: "windows:WindowDidResize",
	1187: "windows:WindowShow",
	1188: "windows:WindowHide",
	1189: "windows:WindowStartMove",
	1190: "windows:WindowEndMove",
	1191: "windows:WindowStartResize",
	1192: "windows:WindowEndResize",
	1193: "windows:WindowKeyDown",
	1194: "windows:WindowKeyUp",
	1195: "windows:WindowZOrderChanged",
	1196: "windows:WindowPaint",
	1197: "windows:WindowBackgroundErase",
	1198: "windows:WindowNonClientHit",
	1199: "windows:WindowNonClientMouseDown",
	1200: "windows:WindowNonClientMouseUp",
	1201: "windows:WindowNonClientMouseMove",
	1202: "windows:WindowNonClientMouseLeave",
	1203: "common:ApplicationStarted",
	1204: "common:WindowMaximise",
	1205: "common:WindowUnMaximise",
	1206: "common:WindowFullscreen",
	1207: "common:WindowUnFullscreen",
	1208: "common:WindowRestore",
	1209: "common:WindowMinimise",
	1210: "common:WindowUnMinimise",
	1211: "common:WindowClosing",
	1212: "common:WindowZoom",
	1213: "common:WindowZoomIn",
	1214: "common:WindowZoomOut",
	1215: "common:WindowZoomReset",
	1216: "common:WindowFocus",
	1217: "common:WindowLostFocus",
	1218: "common:WindowShow",
	1219: "common:WindowHide",
	1220: "common:WindowDPIChanged",
	1221: "common:WindowFilesDropped",
	1222: "common:WindowRuntimeReady",
	1223: "common:ThemeChanged",
	1224: "common:WindowDidMove",
	1225: "common:WindowDidResize",
	1226: "common:ApplicationOpenedWithFile",
}
