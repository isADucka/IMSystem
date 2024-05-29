/*
 * @Author: cyy
 * @Description: --
 */
package domain

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

type CLientContext struct {
	Ip string `ip:"ip"`
}

type IpConfContext struct {
	Ctx       *context.Context
	AppCtx    *app.RequestContext
	ClientCxt *CLientContext
}

func NewIPConfContext(c *context.Context, ctx *app.RequestContext) *IpConfContext {
	return &IpConfContext{
		Ctx:       c,
		AppCtx:    ctx,
		ClientCxt: &CLientContext{},
	}
}
