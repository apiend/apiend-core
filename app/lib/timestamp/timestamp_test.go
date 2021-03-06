/*
    fileName: timestamp
    author: diogoxiang@qq.com
    date: 2019/6/26
*/
package timestamp

import (
	"github.com/globalsign/mgo/bson"
	"github.com/gogf/gf/g/os/gtime"
	"strconv"
	"testing"
	"time"
)

type SampleItem struct {
	ObjId      bson.ObjectId `bson:"_id,omitempty" json:"-"`
	SampleDate Timestamp      `bson:"sampleDate" json:"sampleDate"`
}
func TestTimestamp_Time(t *testing.T) {
	var item SampleItem

	item.SampleDate = Timestamp(time.Now().UTC())

	// t.Log(atime)
	// fmt.Printf(item.SampleDate)
	t.Log(item.SampleDate.String())

	t.Log(item.SampleDate.Time())

	t.Log(item.SampleDate.GetBSON())
}


func TestMarshalJSON(t *testing.T) {
	tm := time.Unix(3000, 0)
	ts := Timestamp(tm)

	b, err := ts.MarshalJSON()
	if err != nil {
		t.Error(err)
	}

	temp, err := strconv.Atoi(string(b))
	if err != nil {
		t.Error(err)
	}

	if temp != 3000 {
		t.Fail()
	}
}

func TestUnmarshalJSON(t *testing.T) {
	tm := time.Unix(3000, 0).UTC()
	ts := Timestamp(tm)

	b, err := ts.MarshalJSON()
	if err != nil {
		t.Error(err)
	}

	var temp Timestamp

	if err := temp.UnmarshalJSON(b); err != nil {
		t.Error(err)
	}

	if temp != ts {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	tm := time.Unix(3000, 0)
	ts := Timestamp(tm)

	if tm.String() != ts.String() {
		t.Fail()
	}
	t.Log(ts.String())
}

func TestGetBSON(t *testing.T) {
	tm := time.Unix(3000, 0)
	ts := Timestamp(tm)

	result, err := ts.GetBSON()
	if err != nil {
		t.Error(err)
	}

	if result != tm {
		t.Fail()
	}
}

func TestTime(t *testing.T) {
	tm := time.Unix(3000, 0)
	ts := Timestamp(tm)

	if ts.Time() != tm {
		t.Fail()
	}
	t.Log(ts.Time())
}

func TestNow(t *testing.T) {
	tm := time.Now()
	ts := Now()

	if tm.Unix() != ts.Time().Unix() {
		t.Fail()
	}

	t.Log(ts.Time().Unix())
}

func TestUnix(t *testing.T) {
	tm := time.Unix(3000, 0)
	ts := Timestamp(tm)

	if ts.Unix() != tm.Unix() {
		t.Fail()
	}
}

func TestFromTime(t *testing.T) {
	tm := time.Unix(3000, 0)
	ts := Time(tm)

	if ts.Unix() != tm.Unix() {
		t.Fail()
	}
}

func TestFromUnix(t *testing.T) {
	tm := time.Unix(3000, 0).UTC()
	ts := Unix(3000, 0)

	if ts.Time() != tm {
		t.Fail()
	}
}

func TestToMili(t *testing.T) {
	numSeconds := int64(3000)
	tm := time.Unix(numSeconds, 0)
	ts := Timestamp(tm)

	result := ts.ToMili()
	if result != numSeconds*1000 {
		t.Fail()
	}
}


func Test_gtime(t *testing.T)  {
	var temp gtime.Time

	tm := temp.String()

	t.Log(tm)
}