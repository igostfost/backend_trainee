package types

type BannerRequest struct {
	BannerId  int     `json:"banner_id" db:"banner_id"`
	TagIds    []int   `json:"tag_ids" db:"tags_id"`
	FeatureId int     `json:"feature_id" db:"feature_id"`
	Content   Content `json:"content"`
	IsActive  bool    `json:"is_active" db:"is_active"`
}

type BannerResponse struct {
	BannerId  int     `json:"banner_id" db:"banner_id"`
	TagIds    []int   `json:"tag_ids" db:"tags_id"`
	FeatureId int     `json:"feature_id" db:"feature_id"`
	Content   Content `json:"content"`
	IsActive  bool    `json:"is_active" db:"is_active"`
}

type Content struct {
	Title string `json:"title" db:"title"`
	Text  string `json:"text" db:"text"`
	Url   string `json:"url" db:"url"`
}

type GetInputBanners struct {
	FeatureId int `json:"feature_id" db:"feature_id"`
	TagIds    int `json:"tag_id" db:"tags_id"`
	Limit     int `json:"limit"`
	Offset    int `json:"offset"`
}

type GetInputUserBanners struct {
	FeatureId       int  `json:"feature_id" db:"feature_id"`
	TagIds          int  `json:"tag_id" db:"tags_id"`
	UseLastRevision bool `json:"use_last_revision"`
}
