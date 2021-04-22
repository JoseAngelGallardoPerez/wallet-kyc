package handlers

import (
	"github.com/Confialink/wallet-kyc/internal/action"
	"github.com/Confialink/wallet-kyc/internal/action/tier"
	"github.com/Confialink/wallet-kyc/internal/adapter"
	"github.com/Confialink/wallet-kyc/internal/response"
	"github.com/Confialink/wallet-kyc/service/api/errcodes"
	"github.com/Confialink/wallet-kyc/service/api/forms"
	"github.com/Confialink/wallet-kyc/service/api/serializers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TierSettingsController struct {
	actionAdminCountries          *action.AdminCountries
	actionAdminTiersByCountryCode *action.AdminTiersByCountryCode
	actionAdminTier               *action.AdminTier
	actionTierUpdateLimitation    *tier.UpdateLimitation
}

func NewTierSettingsController() *TierSettingsController {
	return &TierSettingsController{
		actionAdminCountries: action.NewAdminCountries(
			adapter.NewCountry(),
		),
		actionAdminTiersByCountryCode: action.NewAdminTiersByCountryCode(
			adapter.NewTier(),
		),
		actionAdminTier: action.NewAdminTier(
			adapter.NewTier(),
			tier.NewBindLimitations(
				adapter.NewTierLimitation(),
			),
		),
		actionTierUpdateLimitation: tier.NewTierUpdateLimitation(
			adapter.NewTierLimitation(),
		),
	}
}

func (s *TierSettingsController) ListCountries(ctx *gin.Context) {
	objects, err := s.actionAdminCountries.Do(ctx)
	if err != nil {
		errcodes.AddError(ctx, err)
		return
	}
	serialized := serializers.AdminCountries.SerializeList(ctx, objects)
	ctx.JSON(http.StatusOK, response.New().SetData(serialized))
}

func (s *TierSettingsController) ListTiers(ctx *gin.Context) {
	codeSt := ctx.Params.ByName("country_code")
	objects, err := s.actionAdminTiersByCountryCode.Do(ctx, codeSt)
	if err != nil {
		errcodes.AddError(ctx, err)
		return
	}
	serialized := serializers.AdminTiers.SerializeList(ctx, objects)
	ctx.JSON(http.StatusOK, response.New().SetData(serialized))
}

func (s *TierSettingsController) GetTier(ctx *gin.Context) {
	tierIdSt := ctx.Params.ByName("tierId")
	tierIdInt, err := strconv.ParseInt(tierIdSt, 10, 64)
	if err != nil {
		errcodes.AddError(ctx, err)
		return
	}

	object, err := s.actionAdminTier.Do(ctx, uint64(tierIdInt))
	if err != nil {
		errcodes.AddError(ctx, err)
		return
	}

	serialized := serializers.AdminTiers.Serialize(ctx, &object)
	ctx.JSON(http.StatusOK, response.New().SetData(serialized))
}

func (s *TierSettingsController) UpdateTier(ctx *gin.Context) {
	form := forms.AdminUpdateTier{}

	if !forms.Bind(ctx, &form) {
		return
	}

	object, err := form.Unserialize(ctx, form)
	if err != nil {
		errcodes.AddErrorLog(ctx, errcodes.SerializeError, err)
		return
	}

	tierIdSt := ctx.Params.ByName("tierId")
	tierIdInt, err := strconv.ParseInt(tierIdSt, 10, 64)
	if err != nil {
		errcodes.AddError(ctx, err)
		return
	}

	err = s.actionTierUpdateLimitation.Do(ctx, uint64(tierIdInt), object.Limitations)
	if err != nil {
		errcodes.AddError(ctx, err)
		return
	}
}
