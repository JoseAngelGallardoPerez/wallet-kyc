package model

type TierRequirementElement struct {
	Name    string                          `json:"name"`
	Index   string                          `json:"index"`
	Type    string                          `json:"type"`
	Values  []*UserRequirementValue         `json:"values"`
	Options []*TierRequirementElementOption `json:"options"`
}

type TierRequirementElementOption struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
