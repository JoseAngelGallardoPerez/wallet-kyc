package serializers

import (
	"github.com/Confialink/wallet-kyc/internal/model"
	"github.com/Confialink/wallet-pkg-model_serializer"
	"context"
)

type adminCountries struct {
	fields []interface{}
}

func (t adminCountries) Serialize(ctx context.Context, model interface{}) map[string]interface{} {
	return model_serializer.Serialize(model, t.fields)
}

func (t adminCountries) SerializeList(ctx context.Context, models []model.Country) []map[string]interface{} {
	serialized := make([]map[string]interface{}, len(models))

	for i, v := range models {
		serialized[i] = model_serializer.Serialize(&v, t.fields)
	}
	return serialized
}

var AdminCountries = adminCountries{
	fields: []interface{}{
		"Code",
		"Name",
	},
}
