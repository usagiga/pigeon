package domain

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/usagiga/pigeon/model"
	"github.com/usagiga/pigeon/util/mock/mock_infra"
	"reflect"
	"testing"
)

func TestImageStoreKeeperUseCaseImpl_Store(t *testing.T) {
	// Initialize go mock
	ctrl := gomock.NewController(t)

	imageInfra := mock_infra.NewMockImageStorageInfra(ctrl)
	imageInfra.EXPECT().Fetch("https://example.com/test.png").Return(false, nil).Times(2)
	imageInfra.EXPECT().Fetch("https://example.com/error.png").Return(false, errors.New("error")).Times(1)

	imageStoreKeeper := NewImageStoreKeeperUseCase(imageInfra)

	// Declare test cases
	type Arg struct {
		repoDir *model.GitRepoDir
		srcUrl  string
	}

	type Result struct {
		storedUrl string
	}

	// Initialize GitRepoDir for test cases
	urlRepoDir, _ := model.NewGitRepoDir("test", "test", "test", "https://example.com/formatted")
	fileRepoDir, _ := model.NewGitRepoDir("test", "test", "test", "test")

	testCases := []struct {
		isExpectedError bool
		arg             Arg
		result          Result
	}{
		// Nominal scenario (Store into image server)
		{
			isExpectedError: false,
			arg: Arg{
				repoDir: urlRepoDir,
				srcUrl:  "https://example.com/test.png",
			},
			result: Result{
				storedUrl: "https://example.com/formatted/test.png",
			},
		},
		// Nominal scenario (Store into file)
		{
			isExpectedError: false,
			arg: Arg{
				repoDir: fileRepoDir,
				srcUrl:  "https://example.com/test.png",
			},
			result: Result{
				storedUrl: "test/test.png",
			},
		},
		// On error on ImageInfra.Fetch()
		{
			isExpectedError: true,
			arg: Arg{
				repoDir: urlRepoDir,
				srcUrl:  "https://example.com/error.png",
			},
			result: Result{
				storedUrl: "",
			},
		},
	}

	// Run test
	for i, v := range testCases {
		caseNum := i + 1
		storedUrl, err := imageStoreKeeper.Store(v.arg.repoDir, v.arg.srcUrl)

		// When raising NOT expected error
		if err != nil && !v.isExpectedError {
			t.Errorf("Case %d: This case is not expected to raise error, but error raised; %v", caseNum, err)
		}

		// When NOT raising expected error
		if err == nil && v.isExpectedError {
			t.Errorf("Case %d: This case is expected to raise error, but error didn't raised", caseNum)
		}

		// When returns NOT expected result
		if !reflect.DeepEqual(storedUrl, v.result.storedUrl) {
			t.Errorf(
				"Case %d: Returned 'storedUrl' is not expected.\nExpected:\t%v\nActual:\t%v",
				caseNum,
				v.result.storedUrl,
				storedUrl,
			)
		}
	}
}
