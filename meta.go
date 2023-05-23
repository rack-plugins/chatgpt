package chatgpt

import (
	"github.com/fimreal/rack/module"
	"github.com/spf13/cobra"
)

const (
	ID            = "chatgpt"
	Comment       = "[module] chatgpt api"
	RoutePrefix   = "/chatgpt"
	DefaultEnable = false
)

var Module = module.Module{
	ID:      ID,
	Comment: Comment,
	// gin route
	RouteFunc:   AddRoute,
	RoutePrefix: RoutePrefix,
	// cobra flag
	FlagFunc: ServeFlag,
}

func ServeFlag(serveCmd *cobra.Command) {
	serveCmd.Flags().Bool(ID, DefaultEnable, Comment)

	// chatGPT
	serveCmd.Flags().String("chatgpt.api", "https://api.openai.com", "chatgpt API 地址，方便添加个人 api 代理")
	serveCmd.Flags().String("chatgpt.proxyurl", "", "http proxy 地址，方便添加代理")
	serveCmd.Flags().String("chatgpt.token", "", "chatgpt token https://beta.openai.com/account/api-keys")
}
