package router

type CtxParams struct {
	Val map[string]string
}

func NewCtxParam(basepath, allowDir, BaseUrl string) CtxParams {

	return CtxParams{map[string]string{
		"BasePath": basepath,
		"AllowDir": allowDir,
		"BaseUrl":  BaseUrl,
	}}
}

func (v CtxParams) Get(key string) string {
	return v.Val[key]
}
