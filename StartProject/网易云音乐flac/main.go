package main

import (
	"fmt"
	"os"
	"io"
	"strings"
	"net/http"
	"net/url"
	"io/ioutil"
	"compress/gzip"
	"encoding/json"
	"regexp"
	"sync"
)

const(
	SuggestionUrl = "http://sug.music.baidu.com/info/suggestion"
	Fmlink = "http://music.baidu.com/data/music/fmlink"
)

func DownloadString(remoteUrl string, queryValues url.Values) (body []byte, err error) {
	client := &http.Client{}
	body = nil
	uri, err := url.Parse(remoteUrl)

	//fmt.Println("uri: ", uri)

	if(err != nil){
		return
	}
	if(queryValues != nil){
		values := uri.Query()
		if(values != nil){
			for k,v := range values {
				queryValues[k] = v
			}
		}
		uri.RawQuery = queryValues.Encode()
	}
	request, err := http.NewRequest("GET", uri.String(), nil)
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	request.Header.Add("Accept-Encoding", "gzip, deflate");
	request.Header.Add("Accept-Language", "zh-cn,zh;q=0.8,en-us;q=0.5,en;q=0.3");
	request.Header.Add("Connection", "keep-alive");
	request.Header.Add("Host", uri.Host);
	request.Header.Add("Referer", uri.String());
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0");

	response, err := client.Do(request)
	defer response.Body.Close();
	if(err != nil){
		return ;
	}

	//fmt.Println("response encoding: ", response.Header.Get("Content-Encoding"))

	if response.StatusCode == 200 {
		switch response.Header.Get("Content-Encoding") {
		case "gzip":
			reader, _ := gzip.NewReader(response.Body)
			for {
				buf := make([]byte, 1024)
				n, err := reader.Read(buf)

				if err != nil && err != io.EOF {
					panic(err)
				}

				if n == 0 {
					break
				}
				body = append(body,buf...);
			}
		default:
			body, _ = ioutil.ReadAll(response.Body)

		}
	}

	return
}


func WriteStringToFile(filepath, s string) error {
	fo, err := os.Create(filepath)
	if (err != nil) {
		return err
	}
	defer fo.Close()

	_,err = io.Copy(fo, strings.NewReader(s))
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if(len(os.Args) <= 1) {
		fmt.Println("输入网易云音乐歌单链接.")
		return
	}

	fmt.Println("正在访问 ", os.Args[1])

	nurl := strings.Replace(os.Args[1], "#/", "", -1)

	//fmt.Println("New url: ", nurl)

	response, err := DownloadString(nurl, nil)
	if(err != nil) {
		fmt.Println("访问url出错: ", err)
		return
	}

	WriteStringToFile("info.txt", string(response))


	//generate path to save the song
	var separator string
	if os.IsPathSeparator('\\') {
		separator = "\\"
	} else {
		separator = "/"
	}

	dir, _ := os.Getwd()
	dir = dir + separator + "songs_dir"

	if _, err := os.Stat(dir); err != nil {
		err = os.Mkdir(dir, os.ModePerm)
		if(err != nil){
			fmt.Println("创建目录失败: ", err)
			return
		}
	}

	reg := regexp.MustCompile(`<ul class="f-hide">(.*?)</ul>`)
	mm := reg.FindAllString(string(response), -1)
	//WriteStringToFile("reginfo.txt", mm[0])

	// reg = regexp.MustCompile(`<li><a .*?>(.*?)</a></li>`)
	// mm = reg.FindAllString(string(response), -1)
	// for _, v := range mm{
	// 	fmt.Println(v)
	// }
	 
	waitGroup := sync.WaitGroup{}
	
	if(len(mm) > 0) {
	 	reg = regexp.MustCompile(`<li><a .*?>(.*?)</a></li>`)

	 	contents := mm[0]
	 	urlli := reg.FindAllSubmatch([]byte(contents), -1)

	 	for _,item := range urlli {
	 		query := url.Values{}
	 		query.Set("word", string(item[1]))
	 		query.Set("version", "2")
	 		query.Set("from","0")

	 		res,err := DownloadString(SuggestionUrl,query);
			if(err != nil){
				fmt.Println("获取音乐列表时出错：",err);
				continue;
			}

			//fmt.Println(string(res))

			var dat map[string]interface{};

			err = json.Unmarshal([]byte(res), &dat);

			if err != nil {
				fmt.Println("反序列化JSON时出错:",err);
				continue;
			}

			if _,ok := dat["data"]; ok == false{
				fmt.Println("没有找到音乐资源:",string(item[1]));
				continue;
			}

			songid := dat["data"].(map[string]interface{})["song"].([]interface{})[0].(map[string]interface{})["songid"].(string);

			query = url.Values{};
			query.Set("songIds",songid);
			query.Set("type","flac");


			res ,err = DownloadString(Fmlink,query);

			if(err != nil){
				fmt.Println("获取音乐文件时出错：",err);
				continue;
			}

			var data map[string]interface{};

			err = json.Unmarshal(res,&data);

			if code,ok:= data["errorCode"]; ((ok && code.(float64) == 22005) || err != nil){
				fmt.Println("解析音乐文件时出错：",err);
				continue;
			}

			songlink := data["data"].(map[string]interface{})["songList"].([]interface{})[0].(map[string]interface{})["songLink"].(string);

			r := []rune(songlink)
			if(len(r) < 10){
				fmt.Println("没有无损音乐地址:",string(item[1]));
				continue;
			}

			songname :=  data["data"].(map[string]interface{})["songList"].([]interface{})[0].(map[string]interface{})["songName"].(string);

			artistName :=  data["data"].(map[string]interface{})["songList"].([]interface{})[0].(map[string]interface{})["artistName"].(string);

			filename := dir + separator + songname+"-"+artistName+".flac"

			waitGroup.Add(1);
			go func() {
				fmt.Println("正在下载 ", songname," ......");
				defer waitGroup.Done();

				songRes ,err:= http.Get(songlink);
				if(err != nil){
					fmt.Println("下载文件时出错：",songlink);
					return ;
				}

				songFile,err := os.Create(filename);
				written,err := io.Copy(songFile,songRes.Body);
				if(err != nil){
					fmt.Println("保存音乐文件时出错：",err);
					return ;
				}
				fmt.Println(songname,"下载完成,文件大小：",fmt.Sprintf("%.2f", (float64(written)/(1024*1024))),"MB");
			}();


	 	}
	}

	waitGroup.Wait()

}