package snippet

import (
	"reflect"
	"strings"
	"strconv"
)

func (m MapStringInterface) AssignTo(ptr interface{}, tagName string) bool {
	v := reflect.ValueOf(ptr)
	if v.IsValid() == false {
		return false
	}

	//找到最后指向的值，或空指针，空指针是需要进行初使化的
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}

	tv := v
	if tv.Kind() == reflect.Ptr && tv.CanSet() {
		tv.Set(reflect.New(tv.Type().Elem()))
		tv = tv.Elem()
	}

	if tv.Kind() != reflect.Struct {
		return false
	}

	if assign(tv, m, tagName) { //赋值成功，将临时变量赋给原值
		if v.Kind() == reflect.Ptr {
			v.Set(tv.Addr())
		} else {
			v.Set(tv)
		}
		return true
	} else {
		return false
	}

	return true
}

func assign(dstVal reflect.Value, src interface{}, tagName string) bool {
	sv := reflect.ValueOf(src)
	if !dstVal.IsValid() || !sv.IsValid() {
		return false
	}

	if dstVal.Kind() == reflect.Ptr {
		if dstVal.IsNil() && dstVal.CanSet() {
			dstVal.Set(reflect.New(dstVal.Type().Elem()))
		}
		dstVal = dstVal.Elem()
	}

	if !dstVal.CanSet() {
		return false
	}

	switch dstVal.Kind() {
	case reflect.String:
		dstVal.Set(sv)
	case reflect.Int:
		vInt ,err := strconv.Atoi(sv.Interface().(string))
		if err != nil {
			return false
		}
		dstVal.SetInt(int64(vInt))
	case reflect.Struct:
		if sv.Kind() != reflect.Map || sv.Type().Key().Kind() != reflect.String {
			return false
		}
		success := false
		for i := 0; i < dstVal.NumField(); i++ {
			fv := dstVal.Field(i)
			if fv.IsValid() == false || fv.CanSet() == false {
				continue
			}

			ft := dstVal.Type().Field(i)
			name := ft.Name
			strs := strings.Split(ft.Tag.Get(tagName), ",")
			if strs[0] == "-" {
				continue
			}

			if len(strs[0]) > 0 {
				name = strs[0]
			}

			fsv := sv.MapIndex(reflect.ValueOf(name))
			if fsv.IsValid() {
				if fv.Kind() == reflect.Ptr && fv.IsNil() {
					pv := reflect.New(fv.Type().Elem())
					if assign(pv, fsv.Interface(), tagName) {
						fv.Set(pv)
						success = true
					}
				} else {
					if assign(fv, fsv.Interface(), tagName) {
						success = true
					}
				}
			} else if ft.Anonymous {
				if fv.Kind() == reflect.Ptr && fv.IsNil() {
					pv := reflect.New(fv.Type().Elem())
					if assign(pv, src, tagName) {
						fv.Set(pv)
						success = true
					}
				} else {
					if assign(fv, src, tagName) {
						success = true
					}
				}
			}
		}
		return success
	default:
		return false
	}
	return true
}
