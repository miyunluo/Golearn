**原项目地址**：

https://github.com/lifei6671/NeteaseCloudMusicFlac



Baidu Music API

url = **http://sug.music.baidu.com/info/suggestion**

获得songid 

mess = song_name + singer

Payload = {'word': mess, 'version': '2', 'from':'0'}



url = **http://music.baidu.com/data/music/fmlink**

获得下载地址

playload = {'songIds': song_id, 'type':'flac'}





Golang通过http.NewRequest实现模拟请求，添加请求头和请求参数：

https://my.oschina.net/idufei/blog/668649


