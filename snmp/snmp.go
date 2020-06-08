package snmp

import (
	"encoding/binary"
	"fmt"
	"github.com/soniah/gosnmp"
	"time"
)

// parseDateAndTime extracts a UNIX timestamp from an RFC 2579 DateAndTime.
func parseDateAndTime(pdu *gosnmp.SnmpPDU) (float64, error) {
	var (
		v  []byte
		tz *time.Location
		//err error
	)
	// DateAndTime should be a slice of bytes.
	switch pduType := pdu.Value.(type) {
	case []byte:
		v = pdu.Value.([]byte)
	default:
		return 0, fmt.Errorf("invalid DateAndTime type %v", pduType)
	}
	pduLength := len(v)
	// DateAndTime can be 8 or 11 bytes depending if the time zone is included.
	switch pduLength {
	case 8:
		// No time zone included, assume UTC.
		tz = time.UTC
	case 11:
		// Extract the timezone from the last 3 bytes.
		locString := fmt.Sprintf("%s%02d%02d", string(v[8]), v[9], v[10])
		loc, err := time.Parse("-0700", locString)
		if err != nil {
			return 0, fmt.Errorf("error parsing location string: %q, error: %s", locString, err)
		}
		tz = loc.Location()
	default:
		return 0, fmt.Errorf("invalid DateAndTime length %v", pduLength)
	}
	//if err != nil {
	//	return 0, fmt.Errorf("unable to parse DateAndTime %q, error: %s", v, err)
	//}
	// Build the date from the various fields and time zone.
	t := time.Date(
		int(binary.BigEndian.Uint16(v[0:2])),
		time.Month(v[2]),
		int(v[3]),
		int(v[4]),
		int(v[5]),
		int(v[6]),
		int(v[7])*1e+8,
		tz)
	return float64(t.Unix()), nil
}
