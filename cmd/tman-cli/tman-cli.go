package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {
	var md = md5.New()
	k := func(src string, dst string, who string, stage int, action string) string {
		var ret string
		ret = fmt.Sprintf("%x", md.Sum([]byte(src+":"+dst+":"+who+":"+strconv.Itoa(stage)+":"+action)))
		return ret
	}

	ins := func(src string, dst string, who string, stage int, action string, src_res string, dst_res string) string {
		var ret string
		m := k(src, dst, who, stage, action)

		ret = "INSERT INTO tman.routing_rule(key_routing, src, dst,who, stage, src_result, dst_result, \"desc\", act) VALUES('" + m + "','" + src + "', '" + dst + "', '" + who + "', " + strconv.Itoa(stage) + ", '" + src_res + "', '" + dst_res + "', '', '" + action + "');"
		return ret
	}
	fmt.Println(ins("mart", "tman", "application", 1, "primary", "tman", "flc"))
	fmt.Println(ins("mart", "tman", "application", 1, "regular", "tman", "flc"))

	upd := func(src string, dst string, who string, stage int, action string) string {
		var ret string
		m := k(src, dst, who, stage, action)
		mOld := k(src, dst, who, stage, "")
		ret = `UPDATE tman.routing_rule
				SET key_routing= "` + m + `", act='primary'
				WHERE key_routing = "` + mOld + `";
			  `

		return ret
	}
	fmt.Println(upd("mart", "tman", "application", 1, "primary"))
}
