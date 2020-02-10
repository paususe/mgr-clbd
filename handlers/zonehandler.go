package hdl

import (
	"github.com/gin-gonic/gin"
	"github.com/isbm/mgr-clbd/backend"
	"github.com/isbm/mgr-clbd/dbx"
	"net/http"
)

type ZoneHandler struct {
	BaseHandler
	db  *dbx.Dbx
	bnd *backend.Zones
}

func NewZoneHandler(root string) *ZoneHandler {
	zh := new(ZoneHandler)
	zh.bnd = backend.NewZonesBackend()
	zh.PrepareRoot(root)
	return zh
}

// Backend returns an underlying backend interface
func (zh *ZoneHandler) Backend() backend.Backend {
	return zh.bnd
}

// SetDbx sets the Dbx instance pointer
func (sh *ZoneHandler) SetDbx(db *dbx.Dbx) {
	sh.db = db
	sh.bnd.SetDbx(sh.db)
}

// Handlers returns a map of supported handlers and their configuration
func (zh *ZoneHandler) Handlers() []*HandlerMeta {
	return []*HandlerMeta{
		&HandlerMeta{
			Route:   zh.ToRoute("list"),
			Handle:  zh.ListZones,
			Methods: []string{GET},
		},
		&HandlerMeta{
			Route:   zh.ToRoute("add"),
			Handle:  zh.AddZone,
			Methods: []string{POST},
		},
		&HandlerMeta{
			Route:   zh.ToRoute("remove"),
			Handle:  zh.RemoveZone,
			Methods: []string{POST}, // XXX: Probably should be DELETE instead
		},
		&HandlerMeta{
			Route:   zh.ToRoute("update"),
			Handle:  zh.UpdateZone,
			Methods: []string{POST},
		},
		&HandlerMeta{
			Route:   zh.ToRoute("stats"),
			Handle:  zh.ZoneStats,
			Methods: []string{GET},
		},
	}
}

// ZoneStats godoc
// @Summary Return Zone stats.
// @Description ZoneStats returns data about zone.
// @ID zone-stats
// @Accept json
// @Produce json
// @Param name query string true "Name of the Zone"
// @Header 200 {string} Token "0"
// @Router /api/v1/zones/stats [get]
func (zh *ZoneHandler) ZoneStats(ctx *gin.Context) {
	zh.GetLogger().Errorln("Zone stats not yet implemented")
	ctx.JSON(http.StatusOK, gin.H{"error": "Not implemented yet"})
}

// AddZone godoc
// @Summary Define a cluster Zone.
// @Description AddZone creates a new empty zone in the cluster.
// @ID add-zone
// @Accept json
// @Produce json
// @Param name query string true "Name of the Zone"
// @Param description query string true "Zone description"
// @Header 200 {string} Token "0"
// @Router /api/v1/zones/add [post]
func (zh *ZoneHandler) AddZone(ctx *gin.Context) {
	err := ctx.Request.ParseForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	errcode, msg := zh.GetValidators().VerifyRequired(ctx.Request, "name", "description")
	if errcode != http.StatusOK {
		ctx.JSON(errcode, gin.H{"error": msg})
		return
	}

	name := ctx.Request.Form.Get("name")
	descr := ctx.Request.Form.Get("description")
	zh.bnd.CreateZone(name, descr)

	ctx.JSON(200, gin.H{"foo": "bar"})
}

// UpdateZone godoc
// @Summary Update a cluster Zone
// @Description UpdateZone updates a zone data,
// @ID update-zone
// @Accept json
// @Produce json
// @Param description query string true "Zone description"
// @Header 200 {string} Token "0"
// @Router /api/v1/zones/update [post]
func (zh *ZoneHandler) UpdateZone(ctx *gin.Context) {
}

// RemoveZone godoc
// @Summary Remove an empty cluster Zone
// @Description RemoveZone removes a zone from the cluster, but only if it is empty (no nodes assigned to it).
// @ID remove-zone
// @Accept json
// @Produce json
// @Param name query string true "Name of the Zone"
// @Header 200 {string} Token "0"
// @Router /api/v1/zones/remove [delete]
func (zh *ZoneHandler) RemoveZone(ctx *gin.Context) {
}

// ListZones godoc
// @Summary List cluster zones
// @Description List all zones in the Cluster.
// @ID list-zones
// @Accept  json
// @Produce  json
// @Header 200 {string} Token "0"
// @Router /api/v1/zones/list [get]
func (nh *ZoneHandler) ListZones(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"zones": gin.H{
			"name": "foo",
		},
	})
}
