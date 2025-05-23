// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
// Contains value for auto-update superseded apps.
type Win32LobAutoUpdateSupersededAppsState int

const (
    // Indicates that the auto-update superseded apps state is not configured and the app will not auto-update the superseded apps.
    NOTCONFIGURED_WIN32LOBAUTOUPDATESUPERSEDEDAPPSSTATE Win32LobAutoUpdateSupersededAppsState = iota
    // Indicates that the auto-update superseded apps state is enabled and the app will auto-update the superseded apps if the superseded apps are installed on the device.
    ENABLED_WIN32LOBAUTOUPDATESUPERSEDEDAPPSSTATE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_WIN32LOBAUTOUPDATESUPERSEDEDAPPSSTATE
)

func (i Win32LobAutoUpdateSupersededAppsState) String() string {
    return []string{"notConfigured", "enabled", "unknownFutureValue"}[i]
}
func ParseWin32LobAutoUpdateSupersededAppsState(v string) (any, error) {
    result := NOTCONFIGURED_WIN32LOBAUTOUPDATESUPERSEDEDAPPSSTATE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_WIN32LOBAUTOUPDATESUPERSEDEDAPPSSTATE
        case "enabled":
            result = ENABLED_WIN32LOBAUTOUPDATESUPERSEDEDAPPSSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WIN32LOBAUTOUPDATESUPERSEDEDAPPSSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWin32LobAutoUpdateSupersededAppsState(values []Win32LobAutoUpdateSupersededAppsState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Win32LobAutoUpdateSupersededAppsState) isMultiValue() bool {
    return false
}
