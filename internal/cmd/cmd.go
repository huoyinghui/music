package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *CMain) Http(ctx context.Context, in cMainHttpInput) (out *cMainOutput, err error) {
	s := g.Server()
	s = RegisterRouter(s, ctx, in)
	s.SetOpenApiPath("/api.json")
	s.Run()
	out = &cMainOutput{}
	return out, nil
}

func (c *CMain) Mp3(ctx context.Context, in cMainMp3Input) (out *cMainOutput, err error) {
	mp3Download(ctx, in)
	out = &cMainOutput{}
	return out, nil
}

func (c *CMain) Init(ctx context.Context, in cMainInitInput) (out *cMainOutput, err error) {
	mp3InitTodo(ctx)
	out = &cMainOutput{}
	return out, nil
}
