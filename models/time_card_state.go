// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
type TimeCardState int

const (
    CLOCKEDIN_TIMECARDSTATE TimeCardState = iota
    ONBREAK_TIMECARDSTATE
    CLOCKEDOUT_TIMECARDSTATE
    UNKNOWNFUTUREVALUE_TIMECARDSTATE
)

func (i TimeCardState) String() string {
    return []string{"clockedIn", "onBreak", "clockedOut", "unknownFutureValue"}[i]
}
func ParseTimeCardState(v string) (any, error) {
    result := CLOCKEDIN_TIMECARDSTATE
    switch v {
        case "clockedIn":
            result = CLOCKEDIN_TIMECARDSTATE
        case "onBreak":
            result = ONBREAK_TIMECARDSTATE
        case "clockedOut":
            result = CLOCKEDOUT_TIMECARDSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TIMECARDSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTimeCardState(values []TimeCardState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TimeCardState) isMultiValue() bool {
    return false
}
