package tman_cli

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

}
