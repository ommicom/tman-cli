package main

import (
	"crypto/md5"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type KeyRoutingT struct {
	Src    string
	Dst    string
	SrcRes string
	DstRes string
	Stage  int
	Who    string
	Action string
	OldKey string
}

func main() {
	var md = md5.New()
	k := func(src string, dst string, who string, stage int, action string) string {
		var ret string
		ret = fmt.Sprintf("%x", md.Sum([]byte(src+":"+dst+":"+who+":"+strconv.Itoa(stage)+":"+action)))
		return ret
	}
	f, err := os.Create("insert.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	ff, err := os.Create("update.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer ff.Close()

	ins := func(src string, dst string, who string, stage int, action string, src_res string, dst_res string) string {
		var ret string
		m := k(src, dst, who, stage, action)

		ret = "INSERT INTO tman.routing_rule(key_routing, src, dst,who, stage, src_result, dst_result, \"desc\", act) VALUES('" + m + "','" + src + "', '" + dst + "', '" + who + "', " + strconv.Itoa(stage) + ", '" + src_res + "', '" + dst_res + "', '', '" + action + "');"
		_, err = f.WriteString(ret + "\n")
		if err != nil {
			log.Fatal(err)
		}
		return ret
	}
	upd := func(src string, dst string, who string, stage int, action string) string {
		var ret string
		m := k(src, dst, who, stage, action)
		mOld := k(src, dst, who, stage, "")
		ret = "UPDATE tman.routing_rule SET key_routing= '" + m + "', act='regular' WHERE key_routing = '" + mOld + "';"
		_, err = ff.WriteString(ret + "\n")
		if err != nil {
			log.Fatal(err)
		}
		return ret
	}

	csvFile, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	reader.Comma = ';'
	recs, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	//var routList []KeyRoutingT
	for _, rec := range recs {
		rec := KeyRoutingT{
			Src:    rec[0],
			Dst:    rec[1],
			SrcRes: rec[2],
			DstRes: rec[3],
			Stage:  0,
			Who:    rec[5],
			Action: rec[6],
			OldKey: rec[7],
		}
		//stage, err := strconv.Atoi(rec[4])
		//if err == nil {
		//	log.Println("Ошибка прит парсинге stage")
		//}
		//rec.Stage = 0
		//routList = append(routList, rec)
		fmt.Println(ins(rec.Src, rec.Dst, rec.Who, rec.Stage, rec.Action, rec.SrcRes, rec.DstRes))
		fmt.Println(upd(rec.Src, rec.Dst, rec.Who, rec.Stage, rec.Action))
	}

	//fmt.Println(ins("mart", "tman", "application", 1, "primary", "tman", "flc"))
	//fmt.Println(ins("mart", "tman", "application", 1, "regular", "tman", "flc"))

	fmt.Println("\n")

	//fmt.Println(upd("mart", "tman", "application", 1, "primary"))
}
