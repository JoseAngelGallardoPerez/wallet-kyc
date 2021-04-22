package handlers

import (
	internalAction "github.com/Confialink/wallet-kyc/internal/action"
	"github.com/Confialink/wallet-kyc/internal/action/request"
	"github.com/Confialink/wallet-kyc/internal/adapter"
	"github.com/Confialink/wallet-kyc/internal/connection"
	"github.com/Confialink/wallet-kyc/internal/model"
	"github.com/Confialink/wallet-kyc/internal/service/limit"
	"github.com/Confialink/wallet-kyc/service/api/errcodes"
	"github.com/Confialink/wallet-kyc/service/api/forms"
	"github.com/Confialink/wallet-kyc/service/api/params"
	"github.com/Confialink/wallet-kyc/service/api/serializers"
	"github.com/Confialink/wallet-pkg-list_params"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RequestController struct {
	actionCreateRequest       *internalAction.UserCreateRequest
	actionRequestUpdateStatus *internalAction.AdminRequestUpdateStatus
	actionRequestBindUsers    *request.BindUsers
	adapterTier               *adapter.Tier
	adapterUserRequest        *adapter.UserRequest
}

func NewRequestController() *RequestController {

	adapterRequest := adapter.NewUserRequest()

	return &RequestController{
		actionCreateRequest: internalAction.NewUserCreateRequest(
			adapterRequest,
			adapter.NewUserRequirement(),
			adapter.NewTierRequirement(),
			adapter.NewNotification(),
		),
		actionRequestUpdateStatus: internalAction.NewAdminRequestUpdateStatus(
			adapterRequest,
			adapter.NewNotification(),
			adapter.NewUser(),
			adapter.NewTier(),
			adapter.NewLog(),
			adapter.NewTierRequirement(),
			adapter.NewUserRequirement(),
			limit.NewService(connection.GetRpcLimit()),
		),
		adapterTier:        adapter.NewTier(),
		adapterUserRequest: adapter.NewUserRequest(),
		actionRequestBindUsers: request.NewBindUsers(
			adapter.NewUser(),
		),
	}
}

func (r *RequestController) Create(ctx *gin.Context) {
	user, _ := ctx.Get("auth_user")

	fCreate := forms.CreateRequest{}

	if !forms.Bind(ctx, &fCreate) {
		return
	}

	tier, err := r.adapterTier.FindById(ctx, fCreate.TierId)
	if err != nil {
		errcodes.AddError(ctx, err)
		return
	}

	_, err = r.actionCreateRequest.Do(ctx, user.(model.User), tier)
	if err != nil {
		errcodes.AddError(ctx, err)
		return
	}
}

func (r *RequestController) List(ctx *gin.Context) {

	listParams := params.CreateParams(ctx, model.UserRequest{})
	listParams.Includes.AddIncludes("Tier")
	listParams.Sortings = append(listParams.Sortings, list_params.SortingListParameter{Field: "updatedAt", Direction: list_params.DescDirection})

	query := params.CreateQuery(ctx, listParams)

	list, err := r.adapterUserRequest.FindByQuery(ctx, query)

	if err != nil {
		errcodes.AddError(ctx, err)
		return
	}

	list, err = r.actionRequestBindUsers.Do(ctx, list)
	if err != nil {
		errcodes.AddError(ctx, err)
		return
	}

	queryCount := params.CreateQueryCount(ctx, listParams)
	count, err := r.adapterUserRequest.FindByQueryCount(ctx, queryCount)
	if err != nil {
		errcodes.AddError(ctx, err)
		return
	}

	serialized := serializers.Request.SerializeList(ctx, list)

	ctx.JSON(http.StatusOK, params.NewWithListAndLinksAndPagination(serialized, count, listParams))
}

func (r *RequestController) UpdateStatus(ctx *gin.Context) {
	requestIdSt := ctx.Params.ByName("requestId")
	requestIdInt, err := strconv.ParseInt(requestIdSt, 10, 64)
	if err != nil {
		errcodes.AddError(ctx, err)
		return
	}

	form := forms.UpdateStatusRequest{}
	if !forms.Bind(ctx, &form) {
		return
	}

	authUser := params.Helper.GetAuthUser(ctx)

	_, err = r.actionRequestUpdateStatus.Do(ctx, authUser, uint64(requestIdInt), form.Status)

	if err != nil {
		errcodes.AddError(ctx, err)
		return
	}
}
