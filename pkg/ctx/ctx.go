package ctx

//Ctx 程序核心变量
var Ctx *ctx

type ctx struct {
	Version string
}

func init() {
	Ctx = &ctx{Version: "1.0.0"}
}
