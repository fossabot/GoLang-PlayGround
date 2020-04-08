package timestamp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseTimestamp(t *testing.T) {
	timestamp, err := ParseTimestamp("2006-01-02 15:04:05", "Asia/Shanghai", "2018-07-11 15:07:51")
	if err != nil {
		t.Errorf("ParseTimestamp err %v", err)
	}
	t.Logf("timestamp %v", timestamp)
	assert.NotZero(t, timestamp)
}

func TestParseTimestampSecond(t *testing.T) {
	timestampSecond, err := ParseTimestampSecond("2018-07-11 15:07:51")
	if err != nil {
		t.Errorf("ParseTimestampSecond err %v", err)
	}
	t.Logf("timestampSecond %v", timestampSecond)
	assert.NotZero(t, timestampSecond)

	errorTimeStr := "2018:07:11 15:07:51"
	str2TimestampSecondErr, err := ParseTimestampSecond(errorTimeStr)
	if err == nil {
		t.Errorf("ParseTimestampSecond not found error at %v", errorTimeStr)
	}
	assert.Zero(t, str2TimestampSecondErr)
}

func TestParesTimestampMicro(t *testing.T) {
	timestampMicro, err := ParesTimestampMicro("2018-07-11 15:07:51.456123")
	if err != nil {
		t.Errorf("ParesTimestampMicro err %v", err)
	}
	t.Logf("timestampMicro -> %v", timestampMicro)
	assert.NotZero(t, timestampMicro)

	errorTimeStr := "2018:07:11 15:07:51.456123"
	str2TimestampMicroErr, err := ParesTimestampMicro(errorTimeStr)
	if err == nil {
		t.Errorf("ParesTimestampMicro not found error at %v", errorTimeStr)
	}
	assert.Zero(t, str2TimestampMicroErr)
}

func TestParseLocation(t *testing.T) {
	testTime := "2020-02-25T16:52:18Z"
	parseLocation, err := ParseLocation("2006-01-02T15:04:05Z", "2006-01-02 15:04:05", testTime, "UTC", "Asia/Shanghai")
	if err != nil {
		t.Errorf("ParseLocation err: %v", err)
	}
	t.Logf("parseLocation -> %v", parseLocation)
	assert.NotEmpty(t, parseLocation)
}

func TestParseLocationISO8601TimeSecond(t *testing.T) {
	testTime := "2020-02-25T16:52:18Z"
	locationISO8601TimeSecond, err := ParseLocationISO8601TimeSecond(testTime, "UTC", "Asia/Shanghai")
	if err != nil {
		t.Errorf("ParseLocationISO8601TimeSecond err: %v", err)
	}
	t.Logf("ParseLocationISO8601TimeSecond -> %v", locationISO8601TimeSecond)
	assert.NotEmpty(t, locationISO8601TimeSecond)
}

func TestParseLocationISO8601TimeSecondUTC(t *testing.T) {
	testTime := "2020-02-25T16:52:18Z"
	locationISO8601TimeSecondUTC, err := ParseLocationISO8601TimeSecondUTC(testTime)
	if err != nil {
		t.Errorf("ParseLocationISO8601TimeSecondUTC err: %v", err)
	}
	t.Logf("ParseLocationISO8601TimeSecondUTC -> %v", locationISO8601TimeSecondUTC)
	assert.NotEmpty(t, locationISO8601TimeSecondUTC)
}

func TestParseLocationSecond(t *testing.T) {
	parseLocationSecond, err := ParseLocationSecond("2020-02-26 10:08:58", "UTC", "Asia/Shanghai")
	if err != nil {
		t.Errorf("ParseLocationSecond err: %v", err)
	}
	t.Logf("parseLocationSecond -> %v", parseLocationSecond)
	assert.NotEmpty(t, parseLocationSecond)

	errorTimeString := "2020:02#26 10:08:58"
	locationSecondErr, err := ParseLocationSecond(errorTimeString, "UTC", "Asia/Shanghai")
	if err == nil {
		t.Errorf("ParseLocationSecond not return error")
	}
	assert.Empty(t, locationSecondErr)
}

func TestParseLocationMicro(t *testing.T) {
	parseLocationMicro, err := ParseLocationMicro("2020-02-26 10:08:58", "UTC", "Asia/Shanghai")
	if err != nil {
		t.Errorf("ParseLocationMicro err: %v", parseLocationMicro)
	}
	t.Logf("parseLocationMicro -> %v", parseLocationMicro)
	assert.NotEmpty(t, parseLocationMicro)

	errorTimeString := "2020:02#26 10:08:58"
	locationMicroErr, err := ParseLocationMicro(errorTimeString, "UTC", "Asia/Shanghai")
	if err == nil {
		t.Errorf("ParseLocationMicro not return error")
	}
	assert.Empty(t, locationMicroErr)
}
