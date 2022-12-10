package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/zeek0x/covid19-ogp-lambda/env"
)

const url = "https://raw.githubusercontent.com/tokyo-metropolitan-gov/covid19/development/data/daily_positive_detail.json"
const dateForm = "2006-01-02"
const htmlFrom = `<!DOCTYPE HTML>
<html xmlns="http://www.w3.org/1999/xhtml" lang="ja" xml:lang="ja" xmlns:og="http://ogp.me/ns#" xmlns:fb="http://www.facebook.com/2008/fbml">
<head>
	<meta property="og:title" content="東京の感染者数"/>
	<meta property="og:type" content="website"/>
	<meta property="og:description" content="%s"/>
	<meta property="og:url" content="%s"/>
	<title>t東京の感染者数</title>
</head>
<body>
	%s
</body>
</html>
`

type dailyPositiveDetail struct {
	Date string `json:"name"`
	Data []struct {
		DiagnosedDate      string  `json:"diagnosed_date"`
		Count              int     `json:"count"`
		MissingCount       int     `json:"missing_count"`
		ReportedCount      int     `json:"reported_count"`
		WeeklyGainRatio    float32 `json:"weekly_gain_ratio"`
		UntrackedPercent   float32 `json:"untracked_percent"`
		WeeklyAverageCount float32 `json:"weekly_average_count"`
	} `json:"data"`
}

func weekDay(index time.Weekday) string {
	return []string{"日", "月", "火", "水", "木", "金", "土"}[index]
}

func Handle() (string, error) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	var detail *dailyPositiveDetail
	decoder := json.NewDecoder(resp.Body)
	_ = decoder.Decode(&detail)

	length := len(detail.Data)
	n := 7
	thisWeek := detail.Data[length-n:]
	lastWeek := detail.Data[length-n*2 : length-n]

	var text string
	for i := 0; i < n; i++ {
		thisDay, lastDay := thisWeek[i], lastWeek[i]
		t, _ := time.Parse(dateForm, thisDay.DiagnosedDate)
		text += fmt.Sprintf(
			", %d(%s): %d人(%d%%)",
			t.Day(), weekDay(t.Weekday()), thisDay.Count, int(thisDay.Count*100/lastDay.Count),
		)
	}

	return fmt.Sprintf(htmlFrom, text[2:], url, text[2:]), nil
}

func main() {
	env.Main(Handle)
}
