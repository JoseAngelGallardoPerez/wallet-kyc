package serializers

import (
	"github.com/Confialink/wallet-kyc/internal/model"
	"github.com/Confialink/wallet-pkg-model_serializer"
	"context"
)

type tiersForUser struct {
	fields []interface{}
}

func (t tiersForUser) Serialize(ctx context.Context, model interface{}) map[string]interface{} {
	return model_serializer.Serialize(model, t.fields)
}

func (t tiersForUser) SerializeList(ctx context.Context, models []model.Tier) []map[string]interface{} {
	serialized := make([]map[string]interface{}, len(models))

	for i, v := range models {
		serialized[i] = model_serializer.Serialize(&v, t.fields)
	}
	return serialized
}

var TiersForUser = tiersForUser{
	fields: []interface{}{
		"ID",
		"Name",
		"Level",
		map[string][]interface{}{
			"Requirements": {
				"ID", "Name",
				model_serializer.FieldSerializer(func(modelSerializer interface{}) (fieldName string, value interface{}) {

					requirement := modelSerializer.(*model.TierRequirement)

					var status string

					if len(requirement.UserRequirements) > 0 {
						status = requirement.UserRequirements[0].Status
					} else {
						status = model.RequirementStatus.NotFilled
					}
					return "status", status
				}),
			},
		},
		model_serializer.FieldSerializer(func(modelSerializer interface{}) (fieldName string, value interface{}) {

			tier := modelSerializer.(*model.Tier)

			var status string

			if len(tier.Requests) > 0 {
				status = tier.Requests[0].Status
			} else {
				status = model.RequestStatus.NotAvailable
			}
			return "status", status
		}),
	},
}
