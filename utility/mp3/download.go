package mp3

import (
	"fmt"
	"github.com/gogf/gf/v2/os/glog"
	"net/url"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
)

const ApiPreFmt = "http://music.163.com/song/media/outer/url?id=%v.mp3"

var logger *glog.Logger

func init() {
	logger = g.Log()
}

//网易云音乐有版权的歌必须要开会员才能下载、但即使是开会员你所下载的歌也只能用网易云音乐播放器播放、因为格式做了版权保护。必须要下载MP3格式的才能通用。
//
//所以🙃 教你一个网易云所有歌曲免费下载的方法！
//
//网易云音乐在线网站：https://music.163.com
//
//浏览器打开网易云音乐官网、点开自己喜欢听的歌直接把浏览器地址栏的链接复制下来、或者找到自己喜欢听的歌、选择分享、在选择复制地址
//
//例如：尚东峰的《无骨无花、无我无他》
//
//网易在线播放链接是：
//
//https://music.163.com/#/song?id=863706084
//
//
//
//复制红线标出来的ID粘贴到下面这个链接ID=和.mp3的中间
//
//http://music.163.com/song/media/outer/url?id=863706084.mp3

// ParseSrcLink
// https://music.163.com/song?id=1901371647&userid=429237423
func ParseSrcLink(src string) (id string, err error) {
	// '#'
	src = strings.Replace(src, "#/", "", 1)
	data, err := url.Parse(src)
	if err != nil {
		return "", err
	}
	query := data.Query()
	id = query.Get("id")
	return id, err
}

func DownById(id string, path string) (err error) {
	ctx := gctx.New()
	url := fmt.Sprintf(ApiPreFmt, id)
	result, err := g.Client().Get(ctx, url)
	if err != nil {
		return err
	}
	defer result.Close()
	data := result.ReadAll()
	logger.Debugf(ctx, "path:%v len:%v\n", path, len(data))
	err = gfile.PutBytes(path, data)
	if err != nil {
		return err
	}
	return nil
}

func MusicGet(src string, path string) (err error) {
	id, err := ParseSrcLink(src)
	if err != nil {
		return err
	}
	return DownById(id, path)
}
