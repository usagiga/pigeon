package model

import "testing"

type gitRepoDirArgs struct {
	ProjectRootDir string
	ArticleDir     string
	ImageStoreDir  string
	ImageViewDir   string
}

type gitRepoDirResults struct {
	ArticleDir string
	ImageDir   string
}

var testCases = []struct {
	IsExpectedError  bool
	TestingArgs      gitRepoDirArgs
	ExpectingResults gitRepoDirResults
}{
	// Test for root dir
	{
		IsExpectedError: false,
		TestingArgs: gitRepoDirArgs{
			ProjectRootDir: "~/test/",
			ArticleDir:     "article/",
			ImageStoreDir:  "image/",
		},
		ExpectingResults: gitRepoDirResults{
			ArticleDir: "~/test/article",
			ImageDir:   "~/test/image",
		},
	},
	{
		IsExpectedError: false,
		TestingArgs: gitRepoDirArgs{
			ProjectRootDir: "./test/",
			ArticleDir:     "article/",
			ImageStoreDir:  "image/",
		},
		ExpectingResults: gitRepoDirResults{
			ArticleDir: "test/article",
			ImageDir:   "test/image",
		},
	},
	{
		IsExpectedError: false,
		TestingArgs: gitRepoDirArgs{
			ProjectRootDir: "test/",
			ArticleDir:     "article/",
			ImageStoreDir:  "image/",
		},
		ExpectingResults: gitRepoDirResults{
			ArticleDir: "test/article",
			ImageDir:   "test/image",
		},
	},
	{
		IsExpectedError: false,
		TestingArgs: gitRepoDirArgs{
			ProjectRootDir: "/tmp/",
			ArticleDir:     "article/",
			ImageStoreDir:  "image/",
		},
		ExpectingResults: gitRepoDirResults{
			ArticleDir: "/tmp/article",
			ImageDir:   "/tmp/image",
		},
	},
	{
		IsExpectedError: false,
		TestingArgs: gitRepoDirArgs{
			ProjectRootDir: "test",
			ArticleDir:     "article/",
			ImageStoreDir:  "image/",
		},
		ExpectingResults: gitRepoDirResults{
			ArticleDir: "test/article",
			ImageDir:   "test/image",
		},
	},
	// Test for article dir
	{
		IsExpectedError: true,
		TestingArgs: gitRepoDirArgs{
			ProjectRootDir: "test",
			ArticleDir:     "~/article/",
			ImageStoreDir:  "image/",
		},
	},
	{
		IsExpectedError: false,
		TestingArgs: gitRepoDirArgs{
			ProjectRootDir: "test",
			ArticleDir:     "./article/",
			ImageStoreDir:  "image/",
		},
		ExpectingResults: gitRepoDirResults{
			ArticleDir: "test/article",
			ImageDir:   "test/image",
		},
	},
	{
		IsExpectedError: true,
		TestingArgs: gitRepoDirArgs{
			ProjectRootDir: "test",
			ArticleDir:     "/tmp/",
			ImageStoreDir:  "image/",
		},
	},
	// Test for image dir
	{
		IsExpectedError: true,
		TestingArgs: gitRepoDirArgs{
			ProjectRootDir: "test",
			ArticleDir:     "article/",
			ImageStoreDir:  "~/image/",
		},
	},
	{
		IsExpectedError: false,
		TestingArgs: gitRepoDirArgs{
			ProjectRootDir: "test",
			ArticleDir:     "article/",
			ImageStoreDir:  "./image/",
		},
		ExpectingResults: gitRepoDirResults{
			ArticleDir: "test/article",
			ImageDir:   "test/image",
		},
	},
	{
		IsExpectedError: true,
		TestingArgs: gitRepoDirArgs{
			ProjectRootDir: "test",
			ArticleDir:     "article/",
			ImageStoreDir:  "/tmp/",
		},
	},
}

func TestNewGitRepoDir(t *testing.T) {
	for i, testCase := range testCases {
		caseNum := i + 1
		isExpectedError := testCase.IsExpectedError
		testingValue := testCase.TestingArgs

		_, err := NewGitRepoDir(
			testingValue.ProjectRootDir,
			testingValue.ArticleDir,
			testingValue.ImageStoreDir,
			testingValue.ImageViewDir,
		)

		// When raising NOT expected error
		if err != nil && !isExpectedError {
			t.Errorf("Case %d: This case is not expected to raise error, but error raised; %v", caseNum, err)
		}

		// When NOT raising expected error
		if err == nil && isExpectedError {
			t.Errorf("Case %d: This case is expected to raise error, but error didn't raised", caseNum)
		}
	}
}

func TestGitRepoDir_ArticleDir(t *testing.T) {
	for i, testCase := range testCases {
		caseNum := i + 1
		testingValue := testCase.TestingArgs

		gitRepoDir, err := NewGitRepoDir(
			testingValue.ProjectRootDir,
			testingValue.ArticleDir,
			testingValue.ImageStoreDir,
			testingValue.ImageViewDir,
		)

		// All errors are NOT for this method
		if err != nil {
			continue
		}

		expected := testCase.ExpectingResults.ArticleDir
		actual := gitRepoDir.ArticleDir()

		// When actual value isn't equal expected value
		if expected != actual {
			t.Errorf("Case %d: Actual value isn't equal expected value.\nExpected:\t%v,\nActual:\t%v", caseNum, expected, actual)
		}
	}
}

func TestGitRepoDir_RelativeArticleDir(t *testing.T) {
	t.Skip("There's no test. This method is trivial.")
}

func TestGitRepoDir_ImageDir(t *testing.T) {
	for i, testCase := range testCases {
		caseNum := i + 1
		testingValue := testCase.TestingArgs

		gitRepoDir, err := NewGitRepoDir(
			testingValue.ProjectRootDir,
			testingValue.ArticleDir,
			testingValue.ImageStoreDir,
			testingValue.ImageViewDir,
		)

		// All errors are NOT for this method
		if err != nil {
			continue
		}

		expected := testCase.ExpectingResults.ImageDir
		actual := gitRepoDir.ImageDir()

		// When actual value isn't equal expected value
		if expected != actual {
			t.Errorf("Case %d: Actual value isn't equal expected value.\nExpected:\t%v,\nActual:\t%v", caseNum, expected, actual)
		}
	}
}

func TestGitRepoDir_RelativeImageDir(t *testing.T) {
	t.Skip("There's no test. This method is trivial.")
}
