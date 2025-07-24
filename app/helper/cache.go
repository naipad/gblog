package helper

import (
	"encoding/json"

	"github.com/VictoriaMetrics/fastcache"
)

func ObjCachedSet(mc *fastcache.Cache, k []byte, v interface{}) {
	switch v2 := v.(type) {
	case string:
		mc.Set(k, []byte(v2))
	case []byte:
		mc.Set(k, v2)
	default:
		if jb, err := json.Marshal(v); err == nil {
			mc.Set(k, jb)
		}
	}
}

func ObjCachedSetBig(mc *fastcache.Cache, k []byte, v interface{}) {
	switch v2 := v.(type) {
	case string:
		mc.SetBig(k, []byte(v2))
	case []byte:
		mc.SetBig(k, v2)
	default:
		if jb, err := json.Marshal(v); err == nil {
			mc.SetBig(k, jb)
		}
	}
}

func ObjCachedGet(mc *fastcache.Cache, k []byte, v interface{}, getByte bool) (mcValue []byte, exist bool) {
	if mcValue = mc.Get(nil, k); len(mcValue) > 0 {
		if getByte {
			exist = true
			return
		} else {
			err := json.Unmarshal(mcValue, v)
			if err == nil {
				exist = true
				return
			}
		}
	}
	return
}

func ObjCachedGetBig(mc *fastcache.Cache, k []byte, v interface{}, getByte bool) (mcValue []byte, exist bool) {
	if mcValue = mc.GetBig(nil, k); len(mcValue) > 0 {
		if getByte {
			exist = true
			return
		} else {
			err := json.Unmarshal(mcValue, v)
			if err == nil {
				exist = true
				return
			}
		}
	}
	return
}
