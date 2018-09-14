package swg

// Config config for server
// Fields:
//   Schemes -- 是http还是https []string{swg.SCHEMES_HTTP, swg.SCHEMES_HTTPS}
//   BasePath -- 基础路径例如 "/apis"
//   Version  -- 版本 "v1.0.0"
//   Title    -- API文档标题
//   Description -- 文档描述
//   Host -- 服务器域名
type Config struct {
	Schemes     []Scheme
	BasePath    string
	Version     string
	Title       string
	Description string
	Host        string
}
