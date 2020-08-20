package model

type Config struct {
	EsaAPIKey string `env:"PIGEON_ESA_KEY"`
	EsaTeam   string `env:"PIGEON_ESA_TEAM"`

	DiaryRepoURL string `env:"PIGEON_DIARY_REPO_URL"`
	ArticleDir   string `env:"PIGEON_PROJECT_ARTICLE_DIR"`
	ImageDir     string `env:"PIGEON_PROJECT_IMAGE_DIR"`
}
