package algolia

type Settings struct {
	AttributesToIndex        []string `json:"attributesToIndex,omitempty"`
	AttributesForFaceting    []string `json:"attributesForFaceting,omitempty"`
	NumericAttributesToIndex []string `json:"numericAttributesToIndex,omitempty"`
	AttributeForDistinct     string   `json:"attributeForDistinct,omitempty"`
	Ranking                  []string `json:"ranking,omitempty"`
	CustomRanking            []string `json:"custom_ranking,omitempty"`
}

type SettingsBuilder struct {
	S     *Settings
	Index Index
}

func NewSettingsBuilder(index Index) *SettingsBuilder {
	return &SettingsBuilder{
		S:     &Settings{},
		Index: index,
	}
}

func (b *SettingsBuilder) AttributesToIndex(values ...string) *SettingsBuilder {
	b.S.AttributesToIndex = values
	return b
}

func (b *SettingsBuilder) AttributesForFaceting(values ...string) *SettingsBuilder {
	b.S.AttributesForFaceting = values
	return b
}

func (b *SettingsBuilder) NumericAttributesToIndex(values ...string) *SettingsBuilder {
	b.S.NumericAttributesToIndex = values
	return b
}

func (b *SettingsBuilder) AttributeForDistinct(value string) *SettingsBuilder {
	b.S.AttributeForDistinct = value
	return b
}

func (b *SettingsBuilder) Ranking(values ...string) *SettingsBuilder {
	b.S.Ranking = values
	return b
}

func (b *SettingsBuilder) CustomRanking(values ...string) *SettingsBuilder {
	b.S.CustomRanking = values
	return b
}

func (b *SettingsBuilder) Save() error {
	_, err := b.Index.SetSettings(b.S)
	return err
}
