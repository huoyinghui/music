package cmd

import "github.com/gogf/gf/v2/frame/g"

type CMain struct {
	g.Meta `name:"main" brief:"start main"`
}

type cMainOutput struct{}

type cMainHttpInput struct {
	g.Meta `name:"http" brief:"start http server"`
}

type cMainInitInput struct {
	g.Meta `name:"init" brief:"init data"`
}

type cMainMp3Input struct {
	g.Meta `name:"mp3" brief:"download mp3."`
	Link   string `name:"link" short:"link" brief:"详情页面/分享链接. eg:https://music.163.com/#/song?id=298317 "`
	Id     string `name:"id" short:"id" brief:"详情页面/分享链接. eg:id=298317"`
	Path   string `v:"required" name:"path" short:"path" brief:"下载地址.eg:~/Music/网易云音乐/屋顶.mp3, ./tmp/abc.mp3"`
	Play   bool   `name:"play" short:"p" brief:"播放mp3.eg:~/Music/网易云音乐/屋顶.mp3"`
}
