package handlers

import (
	internalAction "github.com/Confialink/wallet-kyc/internal/action"
	"github.com/Confialink/wallet-kyc/internal/action/requirement"
	"github.com/Confialink/wallet-kyc/internal/action/tier"
	"github.com/Confialink/wallet-kyc/internal/adapter"
	"github.com/Confialink/wallet-kyc/internal/model"
	"github.com/Confialink/wallet-kyc/internal/response"
	ServiceRequirement "github.com/Confialink/wallet-kyc/internal/service/requirement"
	"github.com/Confialink/wallet-kyc/service/api/errcodes"
	"github.com/Confialink/wallet-kyc/service/api/params"
	"github.com/Confialink/wallet-kyc/service/api/serializers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TierController struct {
	actionUserTiers         *internalAction.UserTiers
	actionUserCurrentTier   *internalAction.UserCurrentTier
	actionUserTier          *internalAction.UserTier
	actionAdminTiersForUser *internalAction.AdminTiersForUser
}

func NewTierController() *TierController {
	adapterTier := adapter.NewTier()
	adapterUserRequest := adapter.NewUserRequest()
	adapterUserRequirement := adapter.NewUserRequirement()
	adapterUserRequirementValue := adapter.NewUserRequirementValue()
	adapterTierRequirement := adapter.NewTierRequirement()
	adapterUser := adapter.NewUser()

	actionTierGetByUser := tier.NewGetByUser(
		adapterTier,
		adapterTierRequirement,
		adapterUserRequest,
		adapterUserRequirement,
	)

	actionTierTierBindLimitations := tier.NewBindLimitations(
		adapter.NewTierLimitation(),
	)

	serviceRequirement := ServiceRequirement.NewService(
		adapterUserRequirementValue,
		adapter.NewUserRequirement(),
		adapterUser,
		adapter.NewTier(),
		adapter.NewFile(),
	)

	actionRequirementBindRequirementsAndElements := requirement.NewBindRequirementsAndElements(
		adapterUserRequirement,
		serviceRequirement,
	)

	return &TierController{
		actionUserTiers: internalAction.NewUserTiers(
			adapterTier,
			adapterUserRequest,
			actionRequirementBindRequirementsAndElements,
			actionTierTierBindLimitations,
			actionTierGetByUser,
		),
		actionUserCurrentTier: internalAction.NewUserCurrentTier(
			actionTierGetByUser,
			actionTierTierBindLimitations,
		),
		actionUserTier: internalAction.NewUserTier(
			adapterTier,
			actionRequirementBindRequirementsAndElements,
			actionTierGetByUser,
		),
		actionAdminTiersForUser: internalAction.NewAdminTiersForUser(
			actionTierGetByUser,
			actionTierTierBindLimitations,
			actionRequirementBindRequirementsAndElements,
			adapterTier,
			adapterUserRequest,
			adapterUser,
		),
	}
}

func (t *TierController) ListForUser(ctx *gin.Context) {
	user := params.Helper.GetAuthUser(ctx)

	tiers, err := t.actionUserTiers.Do(ctx, user)

	if err != nil {
		errcodes.AddError(ctx, err)
		return
	}

	serialized := serializers.TiersForUser.SerializeList(ctx, tiers)

	ctx.JSON(http.StatusOK, response.New().SetData(serialized))
}

func (t *TierController) ListForAdmin(ctx *gin.Context) {
	userIdSt := ctx.Params.ByName("userId")

	tiers, err := t.actionAdminTiersForUser.Do(ctx, userIdSt)

	if err != nil {
		errcodes.AddError(ctx, err)
		return
	}

	serialized := serializers.TiersForAdmin.SerializeList(ctx, *tiers)

	ctx.JSON(http.StatusOK, response.New().SetData(serialized))
}

func (t *TierController) Get(ctx *gin.Context) {
	tierIdSt := ctx.Params.ByName("tierId")
	user := params.Helper.GetAuthUser(ctx)

	tierIdInt, err := strconv.ParseInt(tierIdSt, 10, 64)
	if err != nil {
		errcodes.AddError(ctx, err)
		return
	}

	var tierObject *model.Tier

	tierObject, err = t.actionUserTier.Do(ctx, user, uint64(tierIdInt))
	if err != nil {
		errcodes.AddError(ctx, err)
		return
	}
	serialized := serializers.TierForUser.Serialize(ctx, tierObject)

	ctx.JSON(http.StatusOK, response.New().SetData(serialized))
}

func (t *TierController) Current(ctx *gin.Context) {
	user := params.Helper.GetAuthUser(ctx)

	tierObject, err := t.actionUserCurrentTier.Do(ctx, user)
	if err != nil {
		errcodes.AddError(ctx, err)
		return
	}

	serialized := serializers.CurrentTier.Serialize(ctx, tierObject)

	ctx.JSON(http.StatusOK, response.New().SetData(serialized))
}
