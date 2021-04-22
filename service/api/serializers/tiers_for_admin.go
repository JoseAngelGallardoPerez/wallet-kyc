package serializers

import (
	"github.com/Confialink/wallet-kyc/internal/model"
	"github.com/Confialink/wallet-pkg-model_serializer"
	"context"
)

type tiersForAdmin struct {
	fields []interface{}
}

func (t tiersForAdmin) Serialize(ctx context.Context, model interface{}) map[string]interface{} {
	return model_serializer.Serialize(model, t.fields)
}

func (t tiersForAdmin) SerializeList(ctx context.Context, models []model.Tier) []map[string]interface{} {
	serialized := make([]map[string]interface{}, len(models))

	for i, v := range models {
		serialized[i] = model_serializer.Serialize(&v, t.fields)
	}
	return serialized
}

var TiersForAdmin = tiersForAdmin{
	fields: []interface{}{
		"ID",
		"Name",
		"Level",
		map[string][]interface{}{
			"Requirements": {
				"ID", "Name",
				model_serializer.FieldSerializer(func(modelSerializer interface{}) (fieldName string, value interface{}) {

					requirement := modelSerializer.(*model.TierRequirement)
					userRequirements := requirement.UserRequirements[0]

					if userRequirements.Status != model.RequirementStatus.NotFilled {
						return "updatedAt", userRequirements.UpdatedAt
					} else {
						return "updatedAt", nil
					}
				}),

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

				model_serializer.FieldSerializer(func(modelSerializer interface{}) (fieldName string, value interface{}) {
					var elements []struct {
						Index   string `json:"index"`
						Name    string `json:"name"`
						Type    string `json:"type"`
						Options []*struct {
							Name  string `json:"name"`
							Value string `json:"value"`
						} `json:"options"`
						Value string `json:"value"`
					}
					requirement := modelSerializer.(*model.TierRequirement)
					for a := 0; a < len(requirement.Elements); a++ {
						el := requirement.Elements[a]

						value := ""
						if len(requirement.UserRequirements) > 0 {
							for v := 0; v < len(requirement.UserRequirements[0].Values); v++ {
								if requirement.UserRequirements[0].Values[v].Index == requirement.Elements[a].Index {

									value = requirement.UserRequirements[0].Values[v].Value
								}
							}
						}

						var options []*struct {
							Name  string `json:"name"`
							Value string `json:"value"`
						}

						for o := 0; o < len(el.Options); o++ {
							options = append(options, &struct {
								Name  string `json:"name"`
								Value string `json:"value"`
							}{
								Name:  el.Options[o].Name,
								Value: el.Options[o].Value,
							})
						}

						elements = append(elements, struct {
							Index   string `json:"index"`
							Name    string `json:"name"`
							Type    string `json:"type"`
							Options []*struct {
								Name  string `json:"name"`
								Value string `json:"value"`
							} `json:"options"`
							Value string `json:"value"`
						}{
							Index:   el.Index,
							Name:    el.Name,
							Type:    el.Type,
							Options: options,
							Value:   value,
						})
					}

					return "elements", elements
				}),
			},
		},
		model_serializer.FieldSerializer(func(modelSerializer interface{}) (fieldName string, value interface{}) {

			tier := modelSerializer.(*model.Tier)

			if len(tier.Requests) > 0 {
				return "requestId", tier.Requests[0].ID
			} else {
				return "requestId", nil
			}

		}),

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
