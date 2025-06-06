// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
type MeetingLiveShareOptions int

const (
    ENABLED_MEETINGLIVESHAREOPTIONS MeetingLiveShareOptions = iota
    DISABLED_MEETINGLIVESHAREOPTIONS
    UNKNOWNFUTUREVALUE_MEETINGLIVESHAREOPTIONS
)

func (i MeetingLiveShareOptions) String() string {
    return []string{"enabled", "disabled", "unknownFutureValue"}[i]
}
func ParseMeetingLiveShareOptions(v string) (any, error) {
    result := ENABLED_MEETINGLIVESHAREOPTIONS
    switch v {
        case "enabled":
            result = ENABLED_MEETINGLIVESHAREOPTIONS
        case "disabled":
            result = DISABLED_MEETINGLIVESHAREOPTIONS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MEETINGLIVESHAREOPTIONS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeMeetingLiveShareOptions(values []MeetingLiveShareOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MeetingLiveShareOptions) isMultiValue() bool {
    return false
}
