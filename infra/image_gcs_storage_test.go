package infra

import (
	"cloud.google.com/go/storage"
	"context"
	"crypto/tls"
	"github.com/usagiga/pigeon/model"
	"google.golang.org/api/option"
	"net/http"
	"testing"
	"time"
)

func initTestGCSClient(t *testing.T) (client *storage.Client, err error) {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	httpClient := &http.Client{Transport: transCfg}

	clientChan := make(chan *storage.Client)
	go func() {
		for i := 0; i <= 5; i++ {
			c, err := storage.NewClient(context.TODO(), option.WithEndpoint("https://localhost:4443/storage/v1/"), option.WithHTTPClient(httpClient))
			if err != nil {
				time.Sleep(2 * time.Duration(i) * time.Second)
				t.Log("Can't initialize GCS client: ", err)
				continue
			}

			clientChan <- c
			return
		}

		clientChan <- nil
	}()

	client = <-clientChan
	if client == nil {
		t.Fatal("Can't initialize GCS client. Aborting...")
	}

	return client, nil
}

func TestImageGCSStorageInfraImpl_Exists(t *testing.T) {
	// Initialize
	gcsClient, err := initTestGCSClient(t)
	if err != nil {
		t.Fatalf("Can't init test GCS client: %v", err)
	}

	//noinspection GoUnhandledErrorResult
	defer gcsClient.Close()

	imageInfra := NewImageGCSStorageInfra("pigeon-assets", gcsClient)

	// Declare test cases
	type Arg struct {
		repoDir *model.GitRepoDir
		fileName string
	}

	type Result struct {
		exists bool
	}

	testCases := []struct {
		isExpectedError bool
		arg             Arg
		result          Result
	}{
		// Nominal scenario (existing)
		{
			isExpectedError: false,
			arg: Arg{
				fileName: "landscape.jpg",
			},
			result: Result{
				exists: true,
			},
		},
		// Nominal scenario (not existing)
		{
			isExpectedError: false,
			arg: Arg{
				fileName: "not-existing.png",
			},
			result: Result{
				exists: false,
			},
		},
	}

	// Run test
	for i, v := range testCases {
		caseNum := i + 1
		exists, err := imageInfra.Exists(v.arg.repoDir, v.arg.fileName)

		// When raising NOT expected error
		if err != nil && !v.isExpectedError {
			t.Errorf("Case %d: This case is not expected to raise error, but error raised; %v", caseNum, err)
		}

		// When NOT raising expected error
		if err == nil && v.isExpectedError {
			t.Errorf("Case %d: This case is expected to raise error, but error didn't raised", caseNum)
		}

		// When returns NOT expected result
		if exists != v.result.exists {
			t.Errorf(
				"Case %d: Returned 'storedUrl' is not expected.\nExpected:\t%v\nActual:\t%v",
				caseNum,
				v.result.exists,
				exists,
			)
		}
	}
}

func TestImageGCSStorageInfraImpl_Fetch(t *testing.T) {
	// Initialize
	gcsClient, err := initTestGCSClient(t)
	if err != nil {
		t.Fatalf("Can't init test GCS client: %v", err)
	}

	//noinspection GoUnhandledErrorResult
	defer gcsClient.Close()

	imageInfra := NewImageGCSStorageInfra("pigeon-assets", gcsClient)

	// Declare test cases
	type Arg struct {
		repoDir *model.GitRepoDir
		srcUrl   string
	}
	type Result struct {
		skipped bool
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
				srcUrl: "https://example.com/index.html",
			},
			result: Result{
				skipped: false,
			},
		},
		// Nominal scenario (Skipped)
		{
			isExpectedError: false,
			arg: Arg{
				srcUrl: "https://example.com/index.html",
			},
			result: Result{
				skipped: true,
			},
		},
		// On error on fetch image
		{
			isExpectedError: true,
			arg: Arg{
				srcUrl: "error",
			},
			result: Result{
				skipped: false,
			},
		},
	}

	// Run test
	for i, v := range testCases {
		caseNum := i + 1
		skipped, err := imageInfra.Fetch(v.arg.repoDir, v.arg.srcUrl)

		// When raising NOT expected error
		if err != nil && !v.isExpectedError {
			t.Errorf("Case %d: This case is not expected to raise error, but error raised; %v", caseNum, err)
		}

		// When NOT raising expected error
		if err == nil && v.isExpectedError {
			t.Errorf("Case %d: This case is expected to raise error, but error didn't raised", caseNum)
		}

		// When returns NOT expected result
		if skipped != v.result.skipped {
			t.Errorf(
				"Case %d: Returned 'skipped' is not expected.\nExpected:\t%v\nActual:\t%v",
				caseNum,
				v.result.skipped,
				skipped,
			)
		}
	}
}
