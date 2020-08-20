package model

type Config struct {
	EsaAPIKey string `envs:"PIGEON_ESA_KEY"`
	EsaTeam   string `envs:"PIGEON_ESA_TEAM"`

	GitBinPath string `envs:"PIGEON_GIT_BIN"`

	DiaryRepoURL  string `envs:"PIGEON_DIARY_REPO_URL"`
	ArticleDir    string `envs:"PIGEON_PROJECT_ARTICLE_DIR"`
	ImageStoreDir string `envs:"PIGEON_PROJECT_IMAGE_DIR"`
	ImageViewDir  string `envs:"PIGEON_PROJECT_IMAGE_VIEW_DIR"`
}
