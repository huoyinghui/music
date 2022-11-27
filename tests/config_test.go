package tests

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

var ctx context.Context

func init() {
	ctx = gctx.New()
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("manifest/config/config.yaml")
}

// 测试针对: config_tpl.toml 进行测试.
// 项目中真实使用.config.toml.
// config_tpl.toml 必须实时同步.
// 更新config_tpl.toml 需要更新本测试文件. git action会进行测试
// 本地测试: go test -v ./config

func TestLogger(t *testing.T) {
	//g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("manifest/config/config.yaml")
	// 多个日志实例
	// 1.默认日志
	g.Log().Info(ctx, "i am in logger default")
	g.Log("debug").Info(ctx, "i am in logger debug")
	g.Log("test").Info(ctx, "i am in logger test")
}

func TestAppConf(t *testing.T) {
	v := g.Cfg().MustGet(ctx, "app")
	conf := v.MapStrVar()
	t.Log(conf)
	casbin := conf["casbin"].Map()
	t.Log("casbin:", casbin)
	model := casbin["model"]
	policy := casbin["policy"]
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(model, "./manifest/config/rbac_model.conf")
		t.AssertEQ(policy, "./manifest/config/rbac_policy.csv")
		//t.AssertEQ(conf["shareLink"].String(), "http://192.168.0.100:9000/preview/#/index")
		//t.AssertEQ(conf["machineId"].Uint16(), uint16(1))
		//// tpl
		//t.AssertEQ(conf["jwtSalt"].String(), "testsalt")
		//t.AssertEQ(conf["jwtExp"].Int64(), int64(604800))
		//// shp
		//t.AssertEQ(conf["shpOutMinLevel"].Int(), -3)
		//t.AssertEQ(conf["shpOutMaxLevel"].Int(), 0)
		//t.AssertEQ(conf["shpInDoorMinLevel"].Int(), 1)
		//t.AssertEQ(conf["shpInDoorMaxLevel"].Int(), 10)
		//// theme
		//t.AssertEQ(conf["themeIconSysDir"].String(), "/static/mapimgs")
		//// op log size
		//t.AssertEQ(conf["opLogPoolSize"].Int(), 5)
		//// tmp dir delete
		//t.AssertEQ(conf["tmpDirDelete"].Bool(), true)
	})

}
