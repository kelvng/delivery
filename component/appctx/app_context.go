package appctx

import (
	"awesomeProject1/component/uploadprovider"
	"awesomeProject1/pubsub"
	"awesomeProject1/skio"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMaiDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	GetPubSub() pubsub.Pubsub
	GetRealtimeEngine() skio.RealtimeEngine
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
	secretKey      string
	ps             pubsub.Pubsub
	RtEngine       skio.RealtimeEngine
}

func NewAppContext(
	db *gorm.DB, uploadProvider uploadprovider.UploadProvider,
	secretKey string, ps pubsub.Pubsub,
) *appCtx {
	return &appCtx{db: db, uploadProvider: uploadProvider, secretKey: secretKey, ps: ps}
}

func (ctx *appCtx) GetMaiDBConnection() *gorm.DB                  { return ctx.db }
func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider { return ctx.uploadProvider }
func (ctx *appCtx) SecretKey() string                             { return ctx.secretKey }
func (ctx *appCtx) GetPubSub() pubsub.Pubsub                      { return ctx.ps }
func (ctx *appCtx) GetRealtimeEngine() skio.RealtimeEngine        { return ctx.RtEngine }
func (ctx *appCtx) SetRealtimeEngine(rt skio.RealtimeEngine)      { ctx.RtEngine = rt }
