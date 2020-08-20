package model

type Config struct {
	EsaAPIKey string `envs:"PIGEON_ESA_KEY"`
	EsaTeam   string `envs:"PIGEON_ESA_TEAM"`

	DiaryRepoURL string `envs:"PIGEON_DIARY_REPO_URL"`
	ArticleDir   string `envs:"PIGEON_PROJECT_ARTICLE_DIR"`
	ImageDir     string `envs:"PIGEON_PROJECT_IMAGE_DIR"`
}
