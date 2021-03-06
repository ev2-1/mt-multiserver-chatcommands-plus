/*
mt-multiserver-chatcommands-plus contains more useful chatcommands
*/
package main

import (
	"fmt"
	"io"

	proxy "github.com/HimbeerserverDE/mt-multiserver-proxy"
)

func init() {
	proxy.RegisterChatCmd(proxy.ChatCmd{
		Name:  "add_server",
		Perm:  "add_server",
		Help:  "Add a minetest server dynamicaly",
		Usage: "add_server <name> <addr> <mediaPool>",
		Handler: func(cc *proxy.ClientConn, w io.Writer, args ...string) string {
			if len(args) != 3 {
				return "Usage: <name> <addr> <mediaPool>"
			}

			err := proxy.AddServer(args[0], proxy.Server{
				Addr:      args[1],
				MediaPool: args[2],
			})
			if err { // is boolean
				return "error"
			}
			return "ok"
		},
	})

	proxy.RegisterChatCmd(proxy.ChatCmd{
		Name:  "get_server",
		Perm:  "get_server",
		Help:  "Get Parameters of server",
		Usage: "get_server <name>",
		Handler: func(cc *proxy.ClientConn, w io.Writer, args ...string) string {
			if len(args) != 1 {
				return "Usage: <name>"
			}

			srv, found := proxy.Conf().Servers[args[0]]
			if !found {
				return "server dosn't exist"
			}

			return fmt.Sprintf("--- %s ---\nAddr: %s\n MediaPool: %s\n Fallbacks: %s", args[0], srv.Addr, srv.MediaPool, srv.Fallbacks)
		},
	})
}
