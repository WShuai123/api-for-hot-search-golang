package app

import (
	"api/utils"
	"io"
	"net/http"
)

func Renminwang() map[string]interface{} {
	url := "http://www.people.com.cn/GB/59476/index.html"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")

	pattern := `<li><a href="(.*?)" target="_blank">(.*?)</a></li>`
	matched := utils.ExtractMatches(string(pageBytes), pattern)

	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "人民网"

	var obj []map[string]interface{}

	for index, item := range matched {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item[2]
		result["url"] = item[1]
		obj = append(obj, result)
	}
	api["obj"] = obj
	return api
}
