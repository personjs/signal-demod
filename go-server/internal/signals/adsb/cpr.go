package adsb

import "math"

const (
	NL_MAX = 59
)

type cprFrame struct {
	Timestamp int64
	LatCpr    int
	LonCpr    int
	IsOdd     bool
}

var cprCache = map[string][2]cprFrame{} // key = ICAO, value = [even, odd]

// Global CPR decoding adapted from ICAO Annex 10, Appendix B
func DecodeCPR(evenLat, evenLon, oddLat, oddLon int, tEven, tOdd int64, isSurface bool) (float64, float64, bool) {
	// CPR resolution
	var dLatEven, dLatOdd float64
	if isSurface {
		dLatEven = 90.0 / 60
		dLatOdd = 90.0 / 59
	} else {
		dLatEven = 360.0 / 60
		dLatOdd = 360.0 / 59
	}

	// Compute j
	j := math.Floor(((59.0*float64(evenLat))-(60.0*float64(oddLat)))/131072.0 + 0.5)

	// Latitude
	rlatEven := dLatEven * (math.Mod(j, 60.0) + float64(evenLat)/131072.0)
	rlatOdd := dLatOdd * (math.Mod(j, 59.0) + float64(oddLat)/131072.0)

	// Normalize latitude
	if rlatEven >= 270.0 {
		rlatEven -= 360.0
	}
	if rlatOdd >= 270.0 {
		rlatOdd -= 360.0
	}

	// Check latitude zone match
	nl := NL(rlatEven)
	if nl == NL(rlatOdd) {
		var ni float64
		var m float64
		var rlat, rlon float64

		if tEven > tOdd {
			rlat = rlatEven
			nl := float64(NL(rlat))
			ni = math.Max(1, nl)
			m = math.Floor(((float64(evenLon)*(nl-1) - float64(oddLon)*nl) / 131072.0) + 0.5)
			rlon = (360.0 / ni) * (math.Mod(m, ni) + float64(evenLon)/131072.0)
		} else {
			rlat = rlatOdd
			nl := float64(NL(rlat))
			ni = math.Max(1, nl-1)
			m = math.Floor(((float64(evenLon)*(nl-1) - float64(oddLon)*nl) / 131072.0) + 0.5)
			rlon = (360.0 / ni) * (math.Mod(m, ni) + float64(oddLon)/131072.0)
		}

		if rlon > 180 {
			rlon -= 360
		}

		return rlat, rlon, true
	}

	return 0, 0, false
}

func NL(lat float64) int {
	if lat < 0 {
		lat = -lat
	}
	if lat < 10.47047130 {
		return 59
	} else if lat < 14.82817437 {
		return 58
	} else if lat < 18.18626357 {
		return 57
	} else if lat < 21.02939493 {
		return 56
	} else if lat < 23.54504487 {
		return 55
	} else if lat < 25.82924707 {
		return 54
	} else if lat < 27.93898710 {
		return 53
	} else if lat < 29.91135686 {
		return 52
	} else if lat < 31.77209708 {
		return 51
	} else if lat < 33.53993436 {
		return 50
	} else if lat < 35.22899598 {
		return 49
	} else if lat < 36.85025108 {
		return 48
	} else if lat < 38.41241892 {
		return 47
	} else if lat < 39.92256684 {
		return 46
	} else if lat < 41.38651832 {
		return 45
	} else if lat < 42.80914012 {
		return 44
	} else if lat < 44.19454951 {
		return 43
	} else if lat < 45.54626723 {
		return 42
	} else if lat < 46.86733252 {
		return 41
	} else if lat < 48.16039128 {
		return 40
	} else if lat < 49.42776439 {
		return 39
	} else if lat < 50.67150166 {
		return 38
	} else if lat < 51.89342469 {
		return 37
	} else if lat < 53.09516153 {
		return 36
	} else if lat < 54.27817472 {
		return 35
	} else if lat < 55.44378444 {
		return 34
	} else if lat < 56.59318756 {
		return 33
	} else if lat < 57.72747354 {
		return 32
	} else if lat < 58.84763776 {
		return 31
	} else if lat < 59.95459277 {
		return 30
	} else if lat < 61.04917774 {
		return 29
	} else if lat < 62.13216659 {
		return 28
	} else if lat < 63.20427479 {
		return 27
	} else if lat < 64.26616523 {
		return 26
	} else if lat < 65.31845310 {
		return 25
	} else if lat < 66.36171008 {
		return 24
	} else if lat < 67.39646774 {
		return 23
	} else if lat < 68.42322022 {
		return 22
	} else if lat < 69.44242631 {
		return 21
	} else if lat < 70.45451075 {
		return 20
	} else if lat < 71.45986473 {
		return 19
	} else if lat < 72.45884545 {
		return 18
	} else if lat < 73.45177442 {
		return 17
	} else if lat < 74.43893416 {
		return 16
	} else if lat < 75.42056257 {
		return 15
	} else if lat < 76.39684391 {
		return 14
	} else if lat < 77.36789461 {
		return 13
	} else if lat < 78.33374083 {
		return 12
	} else if lat < 79.29428225 {
		return 11
	} else if lat < 80.24923213 {
		return 10
	} else if lat < 81.19801349 {
		return 9
	} else if lat < 82.13956981 {
		return 8
	} else if lat < 83.07199445 {
		return 7
	} else if lat < 83.99173563 {
		return 6
	} else if lat < 84.89166191 {
		return 5
	} else if lat < 85.75541621 {
		return 4
	} else if lat < 86.53536998 {
		return 3
	} else if lat < 87.00000000 {
		return 2
	} else {
		return 1
	}
}
