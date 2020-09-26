package domain

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/usagiga/pigeon/model"
	"github.com/usagiga/pigeon/util/mock/mock_domain"
	"github.com/usagiga/pigeon/util/mock/mock_infra"
	"reflect"
	"testing"
)

func TestArticleBuilderUseCaseImpl_FormatArticle(t *testing.T) {
	// Initialize go mock
	ctrl := gomock.NewController(t)

	imageStoreKeeper := mock_domain.NewMockImageStoreKeeperUseCase(ctrl)
	imageStoreKeeper.EXPECT().Store(gomock.Any(), "https://example.com/test.png").Return("https://example.com/formatted/test.png", nil).MinTimes(1)
	imageStoreKeeper.EXPECT().Store(gomock.Any(), "https://example.com/error.png").Return("", errors.New("error")).MinTimes(1)

	esaInfra := mock_infra.NewMockEsaInfra(ctrl)

	articleBuilder := NewArticleBuilderUseCase(imageStoreKeeper, esaInfra)

	// Declare test cases
	type Arg struct {
		repoDir    *model.GitRepoDir
		rawArticle *model.Article
	}

	type Result struct {
		formattedArticle *model.Article
	}

	testCases := []struct {
		isExpectedError bool
		arg             Arg
		result          Result
	}{
		// Nominal scenario
		{
			isExpectedError: false,
			arg: Arg{
				repoDir: nil,
				rawArticle: &model.Article{
					Title:    "test_title",
					Contents: "![test.png](https://example.com/test.png)",
				},
			},
			result: Result{
				formattedArticle: &model.Article{
					Title:    "test_title",
					Contents: "![test.png](https://example.com/formatted/test.png)",
				},
			},
		},
		// On error on ImageStoreKeeper.Store()
		{
			isExpectedError: true,
			arg: Arg{
				repoDir: nil,
				rawArticle: &model.Article{
					Title:    "test_title",
					Contents: "![error.png](https://example.com/error.png)",
				},
			},
			result: Result{
				formattedArticle: nil,
			},
		},
	}

	// Run test
	for i, v := range testCases {
		caseNum := i + 1
		formattedArticle, err := articleBuilder.FormatArticle(v.arg.repoDir, v.arg.rawArticle)

		// When raising NOT expected error
		if err != nil && !v.isExpectedError {
			t.Errorf("Case %d: This case is not expected to raise error, but error raised; %v", caseNum, err)
		}

		// When NOT raising expected error
		if err == nil && v.isExpectedError {
			t.Errorf("Case %d: This case is expected to raise error, but error didn't raised", caseNum)
		}

		// When returns NOT expected result
		if !reflect.DeepEqual(formattedArticle, v.result.formattedArticle) {
			t.Errorf(
				"Case %d: Returned 'formattedArticle' is not expected.\nExpected:\t%v\nActual:\t%v",
				caseNum,
				v.result.formattedArticle,
				formattedArticle,
			)
		}
	}
}

func TestArticleBuilderUseCaseImpl_GetRawArticle(t *testing.T) {
	t.Skip("TestArticleBuilderUseCaseImpl_GetRawArticle is not implemented")
}

func TestArticleBuilderUseCaseImpl_Store(t *testing.T) {
	t.Skip("TestArticleBuilderUseCaseImpl_Store is not implemented")
}
