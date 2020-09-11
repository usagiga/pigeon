# pigeon

Push an article on esa.io into github repository through `git pull` and `git push`.

Before pushing, All of images will be uploaded GCS and its URLs replace with GCS one.


## Installation

```shell script
$ go get github.com/usagiga/pigeon
```

And then, [configure as docs.](./docs/configure.md)


## Usage

Run it,

```shell script
$ pigeon -id 1234
```

- `id` is article ID on esa.io
    - you can check it on URL of article page
    - ex: In https://EXAMPLE.esa.io/posts/1234, `1234` is article ID
- Pigeon treats article title as file title, and article body as file contents.
    - ex: `diary/2020-01-01-test.md` will be stored on `${PIGEON_PROJECT_ARTICLE_DIR}/2020-01-01-test.md`

NOTE : If you want to use it through Web API, I recommend [adnanh/webhook](https://github.com/adnanh/webhook)


## Dependencies

- Go (1.15 or higher)
    - [github.com/hiroakis/esa-go](https://github.com/hiroakis/esa-go/)
- [esa.io](https://esa.io/)
- Git
- Google Cloud Storage
- Git Repository on GitHub
    - Perhaps another services can use...?


## License

MIT
