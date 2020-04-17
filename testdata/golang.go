package testdata

import "github.com/urfave/cli/v2"

var (
	app     *cli.App
	comment string          // 单行注释符号
	path    string // 格式化路径
	replace bool        // 是否替换
	ext     string   // 扩展名
)

