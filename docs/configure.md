# Configure

## Set your git client

Pigeon do `git clone`, `git clone` and `git clone`.
So, to use pigeon, git client is needed to set up as below.

- Pigeon needs git client it has cred to push into `PIGEON_DIARY_REPO_URL`


## Create bucket on GCS

Pigeon uploads all images on a target article into GCS.

- Create bucket
- Get cred JSON to use GCS


## Set env vars

Please set env vars as below.

- `PIGEON_ESA_TEAM` : Your esa team
- `PIGEON_ESA_KEY` : Your esa Personal access tokens( you can found it in `/user/applications` )
- `PIGEON_DIARY_REPO_URL` : Your diary repository on GitHub
- `PIGEON_PROJECT_ARTICLE_DIR` : Directory where article store. It needs relative path(its origin is project root).
- `PIGEON_IMAGE_MODE` : Mode to change behavior treating images. `file` or `gcs` or `none`.
    - If you choose `file`, all images stored on `PIGEON_PROJECT_IMAGE_DIR`.
    - If you choose `gcs`, all images stored on GCS.
    - If you choose `none`, all images are left as it is.
- `PIGEON_PROJECT_IMAGE_DIR` : (Optional) Directory where images store. It needs relative path.
- `PIGEON_PROJECT_IMAGE_VIEW_DIR` : Base URL where images store.
- `PIGEON_PROJECT_ID` : (Optional) Project ID on GCP
- `PIGEON_BUCKET_ID` : (Optional) GCS Bucket ID on `PIGEON_PROJECT_ID`
- `GOOGLE_APPLICATION_CREDENTIALS` : (Optional) Path to JSON of GCP credential to use GCS


## (Optional) Create article template

See [Example esa.io Template](./example_template.md).


## (Optional) Set up GitHub Pages

I recommend to enable target repository's GitHub Pages.
To know how to enable GitHub Pages, see [official docs](https://docs.github.com/en/github/working-with-github-pages/getting-started-with-github-pages).
