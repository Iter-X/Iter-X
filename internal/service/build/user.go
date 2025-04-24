package build

import (
	userV1 "github.com/iter-x/iter-x/internal/api/user/v1"
)

// GetTimeFormatString converts TimeFormat enum to string representation
func GetTimeFormatString(format userV1.TimeFormat) string {
	switch format {
	case userV1.TimeFormat_H12:
		return "12h"
	case userV1.TimeFormat_H24:
		return "24h"
	default:
		return "24h"
	}
}

// GetDistanceUnitString converts DistanceUnit enum to string representation
func GetDistanceUnitString(unit userV1.DistanceUnit) string {
	switch unit {
	case userV1.DistanceUnit_Kilometer:
		return "km"
	case userV1.DistanceUnit_Mile:
		return "mile"
	default:
		return "km"
	}
}

// GetDarkModeString converts DarkMode enum to string representation
func GetDarkModeString(darkMode userV1.DarkMode) string {
	switch darkMode {
	case userV1.DarkMode_On:
		return "on"
	case userV1.DarkMode_Off:
		return "off"
	case userV1.DarkMode_System:
		return "system"
	default:
		return "system"
	}
}

// GetTimeFormatProto converts string to TimeFormat enum
func GetTimeFormatProto(format string) userV1.TimeFormat {
	switch format {
	case "12h":
		return userV1.TimeFormat_H12
	case "24h":
		return userV1.TimeFormat_H24
	default:
		return userV1.TimeFormat_H24
	}
}

// GetDistanceUnitProto converts string to DistanceUnit enum
func GetDistanceUnitProto(unit string) userV1.DistanceUnit {
	switch unit {
	case "km":
		return userV1.DistanceUnit_Kilometer
	case "mile":
		return userV1.DistanceUnit_Mile
	default:
		return userV1.DistanceUnit_Kilometer
	}
}

// GetDarkModeProto converts string to DarkMode enum
func GetDarkModeProto(darkMode string) userV1.DarkMode {
	switch darkMode {
	case "on":
		return userV1.DarkMode_On
	case "off":
		return userV1.DarkMode_Off
	case "system":
		return userV1.DarkMode_System
	default:
		return userV1.DarkMode_System
	}
}
