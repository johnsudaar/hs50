package config

const (
	// Global communication vars
	START_TRANSMISSION = 0x02
	END_TRANSMISSION   = 0x03

	// Bus Commands
	SET_BUS                = "SBUS"
	CROSS_POINT_QUERY_BUS  = "QBUS" // Return the crosspoint status (XPT_X can be mapped to inputs via interface)"
	CROSS_POINT_STATUS_BUS = "ABUS"
	RAW_QUERY_BUS          = "QBSC" // Show raw seletion (CAM1, CAM2, ..., PGM, PVW, KEYOUT, MV, etc.)
	RAW_STATUS_BUS         = "ABSC"

	// Misc Commands
	SEND_ACTION_FADE = "SAUT"
	SEND_ACTION_CUR  = "SCUT"

	// Parameters
	BUS_A     = ":00"
	BUS_B     = ":01"
	BUS_PGM   = ":02"
	BUS_PVW   = ":03"
	BUS_KEY_F = ":04"
	BUS_KEY_S = ":05"
	BUS_PINP  = ":10"
	BUS_AUX   = ":12"

	CROSS_POINT_1  = ":00"
	CROSS_POINT_2  = ":01"
	CROSS_POINT_3  = ":02"
	CROSS_POINT_4  = ":03"
	CROSS_POINT_5  = ":04"
	CROSS_POINT_6  = ":05"
	CROSS_POINT_7  = ":06"
	CROSS_POINT_8  = ":07"
	CROSS_POINT_9  = ":08"
	CROSS_POINT_10 = ":09"

	INPUT_1 = ":50"
	INPUT_2 = ":51"
	INPUT_3 = ":52"
	INPUT_4 = ":53"
	INPUT_5 = ":54"

	COLOR_BARS       = ":70"
	COLOR_BACKGROUND = ":71"
	BLACK            = ":72"
	FRAME_1          = ":73"
	FRAME_2          = ":74"
	PGM              = ":77"
	PVW              = ":78"
	KEYOUT           = ":79"
	CLN              = ":80"
	MULTI_VIEW       = ":81"

	NO_SELECTION = ":99"

	TALLY_ON  = ":01"
	TALLY_OFF = ":00"

	// For SCUT and ACUT

	BACKGROUND = ":00"
	KEY        = ":01"
	PINP       = ":04"
	FTB        = ":06"

	// Trigger
	TRIGGER = ":00"

	DEFAULT_IP   = "192.168.0.8"
	DEFAULT_PORT = 60040
)
