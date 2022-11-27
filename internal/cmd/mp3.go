package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"music/utility/mp3"
	"path/filepath"
)

func mp3Download(ctx context.Context, in cMainMp3Input) (err error) {
	if in.Play {
		mp3Play(ctx, in.Path)
		return
	}
	if in.Id != "" {
		return mp3.DownById(in.Id, in.Path)
	}
	if in.Link != "" {
		return mp3.MusicGet(in.Link, in.Path)
	}
	return nil
}
func mp3Play(ctx context.Context, mp3Path string) (err error) {
	err = mp3.Play(ctx, mp3Path)
	if err != nil {
		logger.Errorf(ctx, "mp3Play mp3Path:%v err:%v\n", mp3Path, err)
		return err
	}
	return nil
}

func mp3InitTodo(ctx context.Context) (err error) {
	v := g.Cfg().MustGet(ctx, "app")
	conf := v.MapStrVar()
	musicTask := conf["musicTask"].MapStrVar()
	dirPath := musicTask["dirPath"].String()
	todoMap := musicTask["todo"].MapStrVar()
	for k, v := range todoMap {
		logger.Debug(ctx, k)
		musicInfo := v.MapStrVar()
		name := musicInfo["name"].String() + ".mp3"
		user := musicInfo["user"].String()
		link := musicInfo["link"].String()
		id := musicInfo["id"].String()
		path := filepath.Join(dirPath, user, name)
		in := cMainMp3Input{
			Link: link,
			Id:   id,
			Path: path,
		}
		logger.Info(ctx, in)
		err = mp3Download(ctx, in)
		if err != nil {
			logger.Fatalf(ctx, "mp3InitTodo err:%v\n", err)
			return err
		}
	}
	return nil
}
