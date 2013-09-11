package z

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type rtype int

const (
	intRegion rtype = iota
	longRegion
	floatRegion
	float64Region
	dateRegion
	nilValue
)

type Region struct {
	Left      interface{}
	Right     interface{}
	LeftOpen  bool
	RightOpen bool
	_type_    rtype
}

func (r *Region) String() string {
	sb := StringBuilder()
	if r.LeftOpen {
		sb.Append("(")
	} else {
		sb.Append("[")
	}
	if r.HasLeft() {
		sb.Append(r.Left)
	}
	sb.Append(",")
	if r.HasRight() {
		sb.Append(r.Right)
	}
	if r.RightOpen {
		sb.Append(")")
	} else {
		sb.Append("]")
	}
	return sb.String()
}

// 正负整数
const REX_INT string = "^(-?)(\\d+)$"

// 正负浮点数
const REX_FLOAT string = "^(-?)(\\d+)\\.(\\d+)$"

// yyyy-MM-dd
const REX_DATE string = "^([0-9]{3}[1-9]|[0-9]{2}[1-9][0-9]{1}|[0-9]{1}[1-9][0-9]{2}|[1-9][0-9]{3})-(((0[13578]|1[02])-(0[1-9]|[12][0-9]|3[01]))|((0[469]|11)-(0[1-9]|[12][0-9]|30))|(02-(0[1-9]|[1][0-9]|2[0-8])))$"

// yyyy-MM-dd hh:mm:ss
const REX_DATE_TIME string = "([0-9]{3}[1-9]|[0-9]{2}[1-9][0-9]{1}|[0-9]{1}[1-9][0-9]{2}|[1-9][0-9]{3})-(((0[13578]|1[02])-(0[1-9]|[12][0-9]|3[01]))|((0[469]|11)-(0[1-9]|[12][0-9]|30))|(02-(0[1-9]|[1][0-9]|2[0-8])))\\s+(0[0-9]|1[0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9])$"

func MakeRegion(rstr string) *Region {
	r := new(Region)
	leftstr, rightstr, leftOpen, rightOpen := extractLeftAndRight(rstr)
	r.LeftOpen = leftOpen
	r.RightOpen = rightOpen

	var rt rtype
	r.Left, rt = toAppropriateType(leftstr)
	if rt != nilValue {
		r._type_ = rt
	}
	r.Right, rt = toAppropriateType(rightstr)
	if rt != nilValue {
		r._type_ = rt
	}

	return r
}

func toAppropriateType(str string) (interface{}, rtype) {

	// 空
	if IsBlank(str) {
		return nil, nilValue
	} else {
		str = Trim(str)
	}

	regInt := regexp.MustCompile(REX_INT)
	if regInt.MatchString(str) {
		sint, err1 := strconv.Atoi(str)
		if err1 != nil {
			sint64, err2 := strconv.ParseInt(str, 10, 64)
			if err2 != nil {
				panic(err2)
			} else {
				return sint64, longRegion
			}
		} else {
			return sint, intRegion
		}
	}
	regFloat := regexp.MustCompile(REX_FLOAT)
	if regFloat.MatchString(str) {
		sfloat, err1 := strconv.ParseFloat(str, 32)
		if err1 != nil {
			sfloat64, err2 := strconv.ParseFloat(str, 64)
			if err2 != nil {
				panic(err2)
			} else {
				return sfloat64, float64Region
			}
		} else {
			return float32(sfloat), floatRegion
		}
	}
	regDateTime := regexp.MustCompile(REX_DATE_TIME)
	if regDateTime.MatchString(str) {
		sdatetime, err1 := time.ParseInLocation("2006-01-02 15:04:05", TrimExtraSpace(str), time.Local)
		if err1 != nil {
			panic(err1)
		} else {
			return sdatetime, dateRegion
		}
	}
	regDate := regexp.MustCompile(REX_DATE)
	if regDate.MatchString(str) {
		sdate, err1 := time.ParseInLocation("2006-01-02", str, time.Local)
		if err1 != nil {
			panic(err1)
		} else {
			return sdate, dateRegion
		}
	}
	// 没有可以匹配的?
	panic(errors.New("not a region appropriate type, " + str))
}

func extractLeftAndRight(rstr string) (left, right string, lopen, ropen bool) {
	lr := strings.Split(rstr, ",")
	if len(lr) == 2 {
		lopen = strings.HasPrefix(lr[0], "(")
		ropen = strings.HasSuffix(lr[1], ")")
		left = lr[0][1:]
		right = lr[1][:len(lr[1])-1]
		DebugPrintf("region (%s, %s, %v, %v)\n", left, right, lopen, ropen)
		// TODO 左右自动调整位置???

		return
	}
	panic("wrong format for region, " + rstr)
}

func (r *Region) HasLeft() bool {
	return r.Left != nil
}

func (r *Region) HasRight() bool {
	return r.Right != nil
}

func (r *Region) LeftInt() int {
	nleft, ok := r.Left.(int)
	if ok {
		return nleft
	}
	panic("not int type")
}

func (r *Region) RightInt() int {
	nright, ok := r.Right.(int)
	if ok {
		return nright
	}
	panic("not int type")
}

func (r *Region) LeftLong() int64 {
	nleft, ok := r.Left.(int64)
	if ok {
		return nleft
	}
	panic("not long type")
}

func (r *Region) RightLong() int64 {
	nright, ok := r.Right.(int64)
	if ok {
		return nright
	}
	panic("not long type")
}

func (r *Region) LeftFloat() float32 {
	nleft, ok := r.Left.(float32)
	if ok {
		return nleft
	}
	panic("not float32 type")
}

func (r *Region) RightFloat() float32 {
	nright, ok := r.Right.(float32)
	if ok {
		return nright
	}
	panic("not float32 type")
}

func (r *Region) LeftFloat64() float64 {
	nleft, ok := r.Left.(float64)
	if ok {
		return nleft
	}
	panic("not float64 type")
}

func (r *Region) RightFloat64() float64 {
	nright, ok := r.Right.(float64)
	if ok {
		return nright
	}
	panic("not float64 type")
}

func (r *Region) LeftDate() time.Time {
	nleft, ok := r.Left.(time.Time)
	if ok {
		return nleft
	}
	panic("not date type")
}

func (r *Region) RightDate() time.Time {
	nright, ok := r.Right.(time.Time)
	if ok {
		return nright
	}
	panic("not date type")
}
