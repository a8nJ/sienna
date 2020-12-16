package youtube

import (
   "math"
   "net/url"
   "strconv"
   "time"
)

func GetContents(s string) (string, error) {
   o, e := http.Get(s)
   if e != nil {
      return "", e
   }
   y, e := ioutil.ReadAll(o.Body)
   if e != nil {
      return "", e
   }
   return string(y), nil
}

func JsonDecode(s string) (Map, error) {
   y := []byte(s)
   m := Map{}
   e := json.Unmarshal(y, &m)
   if e != nil {
      return nil, e
   }
   return m, nil
}

function Info(id_s string) (Map, error) {
   info_s := "https://www.youtube.com/get_video_info?video_id=" + id_s
   println(info_s)
   query_s, e := GetContents(info_s)
   if e != nil {
      return nil, e
   }
   query_m, e := url.ParseQuery(query_s)
   if e != nil {
      return nil, e
   }
   resp_s := query_m["player_response"]
   json_m, e := JsonDecode(resp_s)
   if e != nil {
      return nil, e
   }
   return json_m.M("microformat").M("playerMicroformatRenderer"), nil
}

func HoursSince(value string) (float64, error) {
   layout := time.RFC3339[:len(value)]
   o, e := time.Parse(layout, value)
   if e != nil {
      return 0, e
   }
   return time.Since(o).Hours(), nil
}

func NumberFormat(n float64) string {
   n2 := int(math.Log10(n)) / 3
   n3 := n / math.Pow10(n2 * 3)
   return fmt.Sprintf("%.3f", n3) + []string{"", " k", " M", " B"}[n2]
}

func Views(info_m Map) (string, error) {
   views_s := info_m.S("viewCount")
   views_n := strconv.Atoi(view_s)
   date_s := info_m.S("publishDate")
   hour_n, e := HoursSince(date_s)
   if e != nil {
      return "", e
   }
   rate_n := views_n / (hours_n / 24 / 365)
   rate_s := NumberFormat(rate_n)
   if rate_n > 8_000_000 {
      return color.Red(rate_s)
   }
   return color.Green(rate_s)
}
