// nolint: err113
package options

import (
	"github.com/go-basic-frames/internal/apiserver"
	"github.com/go-basic-frames/pkg/options"
)

type ServerOptions struct {
	MySQLOptions *options.MySQLOptions `json:"mysql" mapstructure:"mysql"`
}

// NewServerOptions 创建带有默认值的 ServerOptions 实例.
func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		MySQLOptions: options.NewMySQLOptions(),
	}
}

// Validate 校验 ServerOptions 中的选项是否合法.
// 提示：Validate 方法中的具体校验逻辑可以由 Claude、DeepSeek、GPT 等 LLM 自动生成。
func (o *ServerOptions) Validate() error {
	if err := o.MySQLOptions.Validate(); err != nil {
		return err
	}

	return nil
}

// Config 基于 ServerOptions 构建 apiserver.Config
func (o *ServerOptions) Config() (*apiserver.Config, error) {
	return &apiserver.Config{
		MySQLOptions: o.MySQLOptions,
	}, nil
}
