package configuration

import (
	"dstock-backend/internal/common/composite"
	"dstock-backend/internal/common/httpresponse"
	"dstock-backend/internal/domain/repo"
	"dstock-backend/internal/domain/usecase/stock"
	"dstock-backend/internal/infrastructure/postgresql"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type StockHandler struct {
	routerGroup     *gin.RouterGroup
	appServerConfig *AppConfigServer
	service         *StockService
	logger          *zap.Logger
}

func tySInitStockRouter(logger *zap.Logger, routerGroup *gin.RouterGroup, appServerConfig *AppConfigServer) {
	s := &StockHandler{
		routerGroup:     routerGroup,
		appServerConfig: appServerConfig,
		logger:          logger,
	}
	stockRepoComposite := &composite.StockRepoComposite{
		PersistentRepo: postgresql.NewPostgresqlStockRepo(appServerConfig.gormDB, logger),
	}
	s.service = NewStockService(stockRepoComposite, logger)
	s.initRouter()
}

func (s *StockHandler) initRouter() {
	s.routerGroup.GET("/", s.GetAllStock)
}

func (s *StockHandler) GetAllStock(ginCtx *gin.Context) {
	type QueryOption struct {
		Limit  int `form:"limit"`
		Offset int `form:"offset"`
	}
	var response *httpresponse.Response
	var queryOption QueryOption

	err := ginCtx.ShouldBindQuery(&queryOption)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, &gin.H{
			"error": err.Error(),
		})
		return
	}
	if queryOption.Limit == queryOption.Offset {
		queryOption.Limit = 10
		queryOption.Offset = 0
	}
	repoQuery := &repo.Query{
		Limit:  queryOption.Limit,
		Offset: queryOption.Offset,
	}
	req := &stock.GetAllStockReq{
		Query: repoQuery,
	}
	response = s.service.StockService(ginCtx, req)
	ginCtx.JSON(response.Code, response)
}
