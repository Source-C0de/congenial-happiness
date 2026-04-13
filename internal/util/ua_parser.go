package util

import (
	"strings"
)

// DeviceInfo represents a client's device fingerprint parsed from HTTP headers.
type DeviceInfo struct {
	OS           string
	Browser      string
	Architecture string // "64-bit" or "32-bit"
	IP           string
}

// ParseDeviceInfo parses the User-Agent string and returns a DeviceInfo.
// No external package needed — we parse the raw UA string directly.
func ParseDeviceInfo(userAgent, ip string) DeviceInfo {
	info := DeviceInfo{
		OS:           parseOS(userAgent),
		Browser:      parseBrowser(userAgent),
		Architecture: parseArchitecture(userAgent),
		IP:           ip,
	}
	return info
}

func parseOS(ua string) string {
	ua = strings.ToLower(ua)

	switch {
	case strings.Contains(ua, "windows nt 10"):
		return "Windows 10/11"
	case strings.Contains(ua, "windows nt 6.3"):
		return "Windows 8.1"
	case strings.Contains(ua, "windows nt 6.2"):
		return "Windows 8"
	case strings.Contains(ua, "windows nt 6.1"):
		return "Windows 7"
	case strings.Contains(ua, "windows"):
		return "Windows"
	case strings.Contains(ua, "mac os x 10_15") || strings.Contains(ua, "mac os x 10.15"):
		return "macOS Catalina"
	case strings.Contains(ua, "mac os x 11") || strings.Contains(ua, "mac os x 12") ||
		strings.Contains(ua, "mac os x 13") || strings.Contains(ua, "mac os x 14") ||
		strings.Contains(ua, "mac os x 15"):
		return "macOS (modern)"
	case strings.Contains(ua, "mac os x"):
		return "macOS"
	case strings.Contains(ua, "ubuntu"):
		return "Ubuntu Linux"
	case strings.Contains(ua, "debian"):
		return "Debian Linux"
	case strings.Contains(ua, "fedora"):
		return "Fedora Linux"
	case strings.Contains(ua, "linux"):
		return "Linux"
	case strings.Contains(ua, "android"):
		return "Android"
	case strings.Contains(ua, "iphone"):
		return "iOS (iPhone)"
	case strings.Contains(ua, "ipad"):
		return "iOS (iPad)"
	case strings.Contains(ua, "cros"):
		return "ChromeOS"
	default:
		return "Unknown OS"
	}
}

func parseBrowser(ua string) string {
	ua = strings.ToLower(ua)

	// Order matters — more specific first
	switch {
	case strings.Contains(ua, "edg/") || strings.Contains(ua, "edge/"):
		return "Microsoft Edge"
	case strings.Contains(ua, "opr/") || strings.Contains(ua, "opera"):
		return "Opera"
	case strings.Contains(ua, "brave"):
		return "Brave"
	case strings.Contains(ua, "chromium"):
		return "Chromium"
	case strings.Contains(ua, "chrome/") && !strings.Contains(ua, "chromium"):
		return "Google Chrome"
	case strings.Contains(ua, "firefox/"):
		return "Mozilla Firefox"
	case strings.Contains(ua, "safari/") && !strings.Contains(ua, "chrome"):
		return "Safari"
	case strings.Contains(ua, "trident") || strings.Contains(ua, "msie"):
		return "Internet Explorer"
	case strings.Contains(ua, "curl"):
		return "cURL (API Client)"
	case strings.Contains(ua, "postman"):
		return "Postman"
	case strings.Contains(ua, "insomnia"):
		return "Insomnia"
	default:
		return "Unknown Browser"
	}
}

func parseArchitecture(ua string) string {
	ua = strings.ToLower(ua)

	// 64-bit indicators
	if strings.Contains(ua, "x86_64") ||
		strings.Contains(ua, "win64") ||
		strings.Contains(ua, "wow64") ||
		strings.Contains(ua, "amd64") ||
		strings.Contains(ua, "arm64") ||
		strings.Contains(ua, "aarch64") {
		return "64-bit"
	}

	// Explicit 32-bit
	if strings.Contains(ua, "i686") ||
		strings.Contains(ua, "i386") ||
		strings.Contains(ua, "x86;") {
		return "32-bit"
	}

	// Modern macOS and Windows 10+ are effectively always 64-bit
	if strings.Contains(ua, "mac os x") || strings.Contains(ua, "windows nt 10") {
		return "64-bit"
	}

	return "Unknown"
}
