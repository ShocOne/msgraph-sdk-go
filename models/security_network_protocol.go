package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the security singleton.
type SecurityNetworkProtocol int

const (
    UNKNOWN_SECURITYNETWORKPROTOCOL SecurityNetworkProtocol = iota
    IP_SECURITYNETWORKPROTOCOL
    ICMP_SECURITYNETWORKPROTOCOL
    IGMP_SECURITYNETWORKPROTOCOL
    GGP_SECURITYNETWORKPROTOCOL
    IPV4_SECURITYNETWORKPROTOCOL
    TCP_SECURITYNETWORKPROTOCOL
    PUP_SECURITYNETWORKPROTOCOL
    UDP_SECURITYNETWORKPROTOCOL
    IDP_SECURITYNETWORKPROTOCOL
    IPV6_SECURITYNETWORKPROTOCOL
    IPV6ROUTINGHEADER_SECURITYNETWORKPROTOCOL
    IPV6FRAGMENTHEADER_SECURITYNETWORKPROTOCOL
    IPSECENCAPSULATINGSECURITYPAYLOAD_SECURITYNETWORKPROTOCOL
    IPSECAUTHENTICATIONHEADER_SECURITYNETWORKPROTOCOL
    ICMPV6_SECURITYNETWORKPROTOCOL
    IPV6NONEXTHEADER_SECURITYNETWORKPROTOCOL
    IPV6DESTINATIONOPTIONS_SECURITYNETWORKPROTOCOL
    ND_SECURITYNETWORKPROTOCOL
    RAW_SECURITYNETWORKPROTOCOL
    IPX_SECURITYNETWORKPROTOCOL
    SPX_SECURITYNETWORKPROTOCOL
    SPXII_SECURITYNETWORKPROTOCOL
    UNKNOWNFUTUREVALUE_SECURITYNETWORKPROTOCOL
)

func (i SecurityNetworkProtocol) String() string {
    return []string{"UNKNOWN", "IP", "ICMP", "IGMP", "GGP", "IPV4", "TCP", "PUP", "UDP", "IDP", "IPV6", "IPV6ROUTINGHEADER", "IPV6FRAGMENTHEADER", "IPSECENCAPSULATINGSECURITYPAYLOAD", "IPSECAUTHENTICATIONHEADER", "ICMPV6", "IPV6NONEXTHEADER", "IPV6DESTINATIONOPTIONS", "ND", "RAW", "IPX", "SPX", "SPXII", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseSecurityNetworkProtocol(v string) (interface{}, error) {
    result := UNKNOWN_SECURITYNETWORKPROTOCOL
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_SECURITYNETWORKPROTOCOL
        case "IP":
            result = IP_SECURITYNETWORKPROTOCOL
        case "ICMP":
            result = ICMP_SECURITYNETWORKPROTOCOL
        case "IGMP":
            result = IGMP_SECURITYNETWORKPROTOCOL
        case "GGP":
            result = GGP_SECURITYNETWORKPROTOCOL
        case "IPV4":
            result = IPV4_SECURITYNETWORKPROTOCOL
        case "TCP":
            result = TCP_SECURITYNETWORKPROTOCOL
        case "PUP":
            result = PUP_SECURITYNETWORKPROTOCOL
        case "UDP":
            result = UDP_SECURITYNETWORKPROTOCOL
        case "IDP":
            result = IDP_SECURITYNETWORKPROTOCOL
        case "IPV6":
            result = IPV6_SECURITYNETWORKPROTOCOL
        case "IPV6ROUTINGHEADER":
            result = IPV6ROUTINGHEADER_SECURITYNETWORKPROTOCOL
        case "IPV6FRAGMENTHEADER":
            result = IPV6FRAGMENTHEADER_SECURITYNETWORKPROTOCOL
        case "IPSECENCAPSULATINGSECURITYPAYLOAD":
            result = IPSECENCAPSULATINGSECURITYPAYLOAD_SECURITYNETWORKPROTOCOL
        case "IPSECAUTHENTICATIONHEADER":
            result = IPSECAUTHENTICATIONHEADER_SECURITYNETWORKPROTOCOL
        case "ICMPV6":
            result = ICMPV6_SECURITYNETWORKPROTOCOL
        case "IPV6NONEXTHEADER":
            result = IPV6NONEXTHEADER_SECURITYNETWORKPROTOCOL
        case "IPV6DESTINATIONOPTIONS":
            result = IPV6DESTINATIONOPTIONS_SECURITYNETWORKPROTOCOL
        case "ND":
            result = ND_SECURITYNETWORKPROTOCOL
        case "RAW":
            result = RAW_SECURITYNETWORKPROTOCOL
        case "IPX":
            result = IPX_SECURITYNETWORKPROTOCOL
        case "SPX":
            result = SPX_SECURITYNETWORKPROTOCOL
        case "SPXII":
            result = SPXII_SECURITYNETWORKPROTOCOL
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SECURITYNETWORKPROTOCOL
        default:
            return 0, errors.New("Unknown SecurityNetworkProtocol value: " + v)
    }
    return &result, nil
}
func SerializeSecurityNetworkProtocol(values []SecurityNetworkProtocol) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}