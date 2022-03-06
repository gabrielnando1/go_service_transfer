package dto_common

import (
	"encoding/json"
	"time"
)

var _ json.Unmarshaler = &JSONDATE{}

const timeFormatYMDhms = "2006-01-02" // Time format used by json
type JSONDATE time.Time

// JSONDATE deserialization
func (self *JSONDATE) UnmarshalJSON(data []byte) (err error) {
	var s string
	err = json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation(timeFormatYMDhms, s, time.UTC)
	if err != nil {
		return err
	}
	*self = JSONDATE(t)
	return nil
}

// JSONDATE serialization
func (self *JSONDATE) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(*self).Format(timeFormatYMDhms))
}

//Output string
func (self JSONDATE) String() string {
	return time.Time(self).Format(timeFormatYMDhms)
}

func (self *JSONDATE) ToTime() time.Time {
	//v, _ := time.ParseInLocation("\""+timeFormatYMDhms+"\"", self.String(), time.Local)
	return time.Time(*self)
}
