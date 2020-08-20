package domain

import (
	"fmt"
	"github.com/usagiga/pigeon/infra"
	"github.com/usagiga/pigeon/model"
	"github.com/usagiga/pigeon/util/urlnode"
	"golang.org/x/xerrors"
	"io/ioutil"
	"path"
	"regexp"
	"strings"
)

type ArticleBuilderUseCaseImpl struct {
	imageMatchRule          *regexp.Regexp
	imageStoreKeeperUseCase ImageStoreKeeperUseCase
	esaInfra                infra.EsaInfra
}

func NewArticleBuilderUseCase(
	imageStoreKeeperUseCase ImageStoreKeeperUseCase,
	esaInfra infra.EsaInfra,
) (domain ArticleBuilderUseCase) {
	//noinspection RegExpRedundantEscape
	return &ArticleBuilderUseCaseImpl{
		imageMatchRule: regexp.MustCompile(
			`(?:<img.*src="(https?://.+/.+\\.(?:png|jpeg|jpg|jfif|gif|webp)).*">|!\[.*\]\((https?://.+/.+\\.(?:png|jpeg|jpg|jfif|gif|webp)).*\))`,
		),
		imageStoreKeeperUseCase: imageStoreKeeperUseCase,
		esaInfra:                esaInfra,
	}
}

func (d *ArticleBuilderUseCaseImpl) GetRawArticle(postId int) (rawArticle *model.Article, err error) {
	rawArticle, err = d.esaInfra.GetArticle(postId)
	if err != nil {
		return nil, xerrors.Errorf("can't get article: %w", err)
	}

	return rawArticle, nil
}

func (d *ArticleBuilderUseCaseImpl) FormatArticle(repoDir *model.GitRepoDir, rawArticle *model.Article) (formattedArticle *model.Article, err error) {
	contents := rawArticle.Contents

	// Replace image tag
	matches := d.imageMatchRule.FindAllStringSubmatch(contents, -1)
	for _, m := range matches {
		// imageMatchRule has 2 groups only
		if len(m) != 3 {
			continue
		}

		imageTag := m[0]                    // m[0] has original image tag
		imageUrl := strings.Join(m[1:], "") // m[1] or [2] have image urls and the others are blank.

		// Store and get new url
		newUrl, err := d.imageStoreKeeperUseCase.Store(repoDir, imageUrl)
		if err != nil {
			return nil, xerrors.Errorf("can't store image: %w", err)
		}
		newFileName, err := urlnode.GetLastNodeFromString(newUrl)
		newTag := fmt.Sprintf("![%s](%s)", newFileName, newUrl)

		// Replace original image tag to new one
		contents = strings.Replace(contents, imageTag, newTag, -1)
	}

	return &model.Article{
		Title:    rawArticle.Title,
		Contents: contents,
	}, nil
}

func (d *ArticleBuilderUseCaseImpl) Store(repoDir *model.GitRepoDir, article *model.Article) (err error) {
	destPath := path.Join(repoDir.ProjectRootDir, article.Title)

	textBytes := []byte(article.Contents) // textBytes are UTF-8(according to Go implementation)
	err = ioutil.WriteFile(destPath, textBytes, 0666)
	if err != nil {
		return xerrors.Errorf("can't store article: %w", err)
	}

	return nil
}
