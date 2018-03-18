使用百度音乐的API

url = **http://sug.music.baidu.com/info/suggestion**

这个地址获得songid 

mess = song_name + singer

Payload = {'word': mess, 'version': '2', 'from':'0'}

url = **http://music.baidu.com/data/music/fmlink**

这一地址提供音乐下载

playload = {'songIds': song_id, 'type':'flac'}





Golang通过http.NewRequest实现模拟请求，添加请求头和请求参数：

https://my.oschina.net/idufei/blog/668649



原项目地址：

https://github.com/lifei6671/NeteaseCloudMusicFlac