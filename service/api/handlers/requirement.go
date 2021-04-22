package handlers

import (
	internalAction "github.com/Confialink/wallet-kyc/internal/action"
	"github.com/Confialink/wallet-kyc/internal/action/request"
	"github.com/Confialink/wallet-kyc/internal/adapter"
	serviceRequirement "github.com/Confialink/wallet-kyc/internal/service/requirement"
	"github.com/Confialink/wallet-kyc/service/api/errcodes"
	"github.com/Confialink/wallet-kyc/service/api/forms"
	"github.com/Confialink/wallet-kyc/service/api/params"
	"github.com/gin-gonic/gin"
	"strconv"
)

type RequirementController struct {
	actionUserRequirementUpdate   *internalAction.UserRequirementUpdate
	actionAdminRequirementUpdate  *internalAction.AdminRequirementUpdate
	actionRequirementUpdateStatus *internalAction.AdminRequirementUpdateStatus
	adapterUser                   *adapter.User
}

func NewRequirementController() *RequirementController {
	adapterUserRequirement := adapter.NewUserRequirement()

	return &RequirementController{
		actionUserRequirementUpdate: internalAction.NewUserRequirementUpdate(
			adapter.NewTierRequirement(),
			adapter.NewUserRequirementValue(),
			adapterUserRequirement,
			serviceRequirement.NewService(
				adapter.NewUserRequirementValue(),
				adapter.NewUserRequirement(),
				adapter.NewUser(),
				adapter.NewTier(),
				adapter.NewFile(),
			),
			request.NewGetByTierUserId(
				adapter.NewUserRequest(),
				adapter.NewTier(),
			),
		),

		actionAdminRequirementUpdate: internalAction.NewAdminRequirementUpdate(
			adapter.NewTierRequirement(),
			adapter.NewUserRequirementValue(),
			adapterUserRequirement,
			adapter.NewLog(),
			adapter.NewTier(),
			adapter.NewUserRequest(),
			serviceRequirement.NewService(
				adapter.NewUserRequirementValue(),
				adapter.NewUserRequirement(),
				adapter.NewUser(),
				adapter.NewTier(),
				adapter.NewFile(),
			),
			request.NewGetByTierUserId(
				adapter.NewUserRequest(),
				adapter.NewTier(),
			),
		),

		actionRequirementUpdateStatus: internalAction.NewAdminRequirementUpdateStatus(
			adapterUserRequirement,
			adapter.NewTierRequirement(),
			adapter.NewUser(),
			adapter.NewNotification(),
			adapter.NewLog(),
		),

		adapterUser: adapter.NewUser(),
	}
}

func (s *RequirementController) Update(ctx *gin.Context) {
	user := params.Helper.GetAuthUser(ctx)
	requirementIdSt := ctx.Params.ByName("requirementId")
	requirementIdInt, err := strconv.ParseInt(requirementIdSt, 10, 64)
	if err != nil {
		errcodes.AddError(ctx, err)
		return
	}

	form := forms.UpdateRequirement{}

	if !forms.Bind(ctx, &form) {
		return
	}

	err = s.actionUserRequirementUpdate.Do(ctx, user, uint64(requirementIdInt), form.Values)
	if err != nil {
		errcodes.AddError(ctx, err)
		return
	}
}

func (s *RequirementController) AdminUpdate(ctx *gin.Context) {
	requirementIdSt := ctx.Params.ByName("requirementId")
	requirementIdInt, err := strconv.ParseInt(requirementIdSt, 10, 64)
	if err != nil {
		errcodes.AddError(ctx, err)
		return
	}

	userIdSt := ctx.Params.ByName("userId")
	user, err := s.adapterUser.GetUserById(ctx, userIdSt)
	if err != nil {
		errcodes.AddError(ctx, err)
		return
	}

	form := forms.UpdateRequirement{}

	if !forms.Bind(ctx, &form) {
		return
	}

	authUser := params.Helper.GetAuthUser(ctx)

	err = s.actionAdminRequirementUpdate.Do(ctx, authUser, *user, uint64(requirementIdInt), form.Values)
	if err != nil {
		errcodes.AddError(ctx, err)
		return
	}
}

func (r *RequirementController) UpdateStatus(ctx *gin.Context) {

	requirementIdSt := ctx.Params.ByName("requirementId")
	requirementIdInt, err := strconv.ParseInt(requirementIdSt, 10, 64)
	if err != nil {
		errcodes.AddError(ctx, err)
		return
	}
	userIdSt := ctx.Params.ByName("userId")

	form := forms.UpdateStatusRequirement{}

	if !forms.Bind(ctx, &form) {
		return
	}

	authUser := params.Helper.GetAuthUser(ctx)

	err = r.actionRequirementUpdateStatus.Do(ctx, authUser, uint64(requirementIdInt), userIdSt, form.Status)
	if err != nil {
		errcodes.AddError(ctx, err)
		return
	}
}
