// Copyright 2018 The gooid Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.


package config

import (
	"github.com/gooid/gooid/internal/ndk"
)

const (
	ORIENTATION_ANY    = app.CONFIGURATION_ORIENTATION_ANY
	ORIENTATION_PORT   = app.CONFIGURATION_ORIENTATION_PORT
	ORIENTATION_LAND   = app.CONFIGURATION_ORIENTATION_LAND
	ORIENTATION_SQUARE = app.CONFIGURATION_ORIENTATION_SQUARE

	TOUCHSCREEN_ANY     = app.CONFIGURATION_TOUCHSCREEN_ANY
	TOUCHSCREEN_NOTOUCH = app.CONFIGURATION_TOUCHSCREEN_NOTOUCH
	TOUCHSCREEN_STYLUS  = app.CONFIGURATION_TOUCHSCREEN_STYLUS
	TOUCHSCREEN_FINGER  = app.CONFIGURATION_TOUCHSCREEN_FINGER

	DENSITY_DEFAULT = app.CONFIGURATION_DENSITY_DEFAULT
	DENSITY_LOW     = app.CONFIGURATION_DENSITY_LOW
	DENSITY_MEDIUM  = app.CONFIGURATION_DENSITY_MEDIUM
	DENSITY_TV      = app.CONFIGURATION_DENSITY_TV
	DENSITY_HIGH    = app.CONFIGURATION_DENSITY_HIGH
	DENSITY_XHIGH   = app.CONFIGURATION_DENSITY_XHIGH
	DENSITY_XXHIGH  = app.CONFIGURATION_DENSITY_XXHIGH
	DENSITY_XXXHIGH = app.CONFIGURATION_DENSITY_XXXHIGH
	DENSITY_NONE    = app.CONFIGURATION_DENSITY_NONE

	KEYBOARD_ANY    = app.CONFIGURATION_KEYBOARD_ANY
	KEYBOARD_NOKEYS = app.CONFIGURATION_KEYBOARD_NOKEYS
	KEYBOARD_QWERTY = app.CONFIGURATION_KEYBOARD_QWERTY
	KEYBOARD_12KEY  = app.CONFIGURATION_KEYBOARD_12KEY

	NAVIGATION_ANY       = app.CONFIGURATION_NAVIGATION_ANY
	NAVIGATION_NONAV     = app.CONFIGURATION_NAVIGATION_NONAV
	NAVIGATION_DPAD      = app.CONFIGURATION_NAVIGATION_DPAD
	NAVIGATION_TRACKBALL = app.CONFIGURATION_NAVIGATION_TRACKBALL
	NAVIGATION_WHEEL     = app.CONFIGURATION_NAVIGATION_WHEEL

	KEYSHIDDEN_ANY  = app.CONFIGURATION_KEYSHIDDEN_ANY
	KEYSHIDDEN_NO   = app.CONFIGURATION_KEYSHIDDEN_NO
	KEYSHIDDEN_YES  = app.CONFIGURATION_KEYSHIDDEN_YES
	KEYSHIDDEN_SOFT = app.CONFIGURATION_KEYSHIDDEN_SOFT

	NAVHIDDEN_ANY = app.CONFIGURATION_NAVHIDDEN_ANY
	NAVHIDDEN_NO  = app.CONFIGURATION_NAVHIDDEN_NO
	NAVHIDDEN_YES = app.CONFIGURATION_NAVHIDDEN_YES

	SCREENSIZE_ANY    = app.CONFIGURATION_SCREENSIZE_ANY
	SCREENSIZE_SMALL  = app.CONFIGURATION_SCREENSIZE_SMALL
	SCREENSIZE_NORMAL = app.CONFIGURATION_SCREENSIZE_NORMAL
	SCREENSIZE_LARGE  = app.CONFIGURATION_SCREENSIZE_LARGE
	SCREENSIZE_XLARGE = app.CONFIGURATION_SCREENSIZE_XLARGE

	SCREENLONG_ANY = app.CONFIGURATION_SCREENLONG_ANY
	SCREENLONG_NO  = app.CONFIGURATION_SCREENLONG_NO
	SCREENLONG_YES = app.CONFIGURATION_SCREENLONG_YES

	UI_MODE_TYPE_ANY        = app.CONFIGURATION_UI_MODE_TYPE_ANY
	UI_MODE_TYPE_NORMAL     = app.CONFIGURATION_UI_MODE_TYPE_NORMAL
	UI_MODE_TYPE_DESK       = app.CONFIGURATION_UI_MODE_TYPE_DESK
	UI_MODE_TYPE_CAR        = app.CONFIGURATION_UI_MODE_TYPE_CAR
	UI_MODE_TYPE_TELEVISION = app.CONFIGURATION_UI_MODE_TYPE_TELEVISION
	UI_MODE_TYPE_APPLIANCE  = app.CONFIGURATION_UI_MODE_TYPE_APPLIANCE

	UI_MODE_NIGHT_ANY = app.CONFIGURATION_UI_MODE_NIGHT_ANY
	UI_MODE_NIGHT_NO  = app.CONFIGURATION_UI_MODE_NIGHT_NO
	UI_MODE_NIGHT_YES = app.CONFIGURATION_UI_MODE_NIGHT_YES

	SCREEN_WIDTH_DP_ANY = app.CONFIGURATION_SCREEN_WIDTH_DP_ANY

	SCREEN_HEIGHT_DP_ANY = app.CONFIGURATION_SCREEN_HEIGHT_DP_ANY

	SMALLEST_SCREEN_WIDTH_DP_ANY = app.CONFIGURATION_SMALLEST_SCREEN_WIDTH_DP_ANY

	LAYOUTDIR_ANY = app.CONFIGURATION_LAYOUTDIR_ANY
	LAYOUTDIR_LTR = app.CONFIGURATION_LAYOUTDIR_LTR
	LAYOUTDIR_RTL = app.CONFIGURATION_LAYOUTDIR_RTL

	MCC                  = app.CONFIGURATION_MCC
	MNC                  = app.CONFIGURATION_MNC
	LOCALE               = app.CONFIGURATION_LOCALE
	TOUCHSCREEN          = app.CONFIGURATION_TOUCHSCREEN
	KEYBOARD             = app.CONFIGURATION_KEYBOARD
	KEYBOARD_HIDDEN      = app.CONFIGURATION_KEYBOARD_HIDDEN
	NAVIGATION           = app.CONFIGURATION_NAVIGATION
	ORIENTATION          = app.CONFIGURATION_ORIENTATION
	DENSITY              = app.CONFIGURATION_DENSITY
	SCREEN_SIZE          = app.CONFIGURATION_SCREEN_SIZE
	VERSION              = app.CONFIGURATION_VERSION
	SCREEN_LAYOUT        = app.CONFIGURATION_SCREEN_LAYOUT
	UI_MODE              = app.CONFIGURATION_UI_MODE
	SMALLEST_SCREEN_SIZE = app.CONFIGURATION_SMALLEST_SCREEN_SIZE
	LAYOUTDIR            = app.CONFIGURATION_LAYOUTDIR
)

type Configuration = app.Configuration
