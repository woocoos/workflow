package identity

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/web"
	"net/http"
	"strconv"
)

var tenantContextKey = "github.com_woocoos_entco_tenant_id"

// RegistryTenantIDMiddleware register a middleware to get tenant id from request header
func RegistryTenantIDMiddleware() web.Option {
	return web.RegisterMiddlewareByFunc("tenant", func(cfg *conf.Configuration) gin.HandlerFunc {
		typ := cfg.String("type")
		return func(c *gin.Context) {
			tid := c.GetHeader("X-Tenant-ID")
			if tid != "" {
				if typ == "int" {
					if _, err := strconv.Atoi(tid); err != nil {
						c.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid tenant id %s:%v", tid, err))
						return
					}
				} else {
					c.Set(tenantContextKey, tid)
				}
			}
		}
	})
}

// TenantIDFromWebContext returns the tenant id from web context.tenant id is int format
func TenantIDFromWebContext(ctx context.Context) (val int, err error) {
	gctx := ctx.Value(gin.ContextKey).(*gin.Context)
	tid := gctx.Value(tenantContextKey)
	if tid == nil {
		return
	}
	switch v := tid.(type) {
	case int:
		val = v
	case string:
		val, err = strconv.Atoi(v)
	}
	return
}
