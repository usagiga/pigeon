# Example config for adnanh/webhook

## webhook.json (disabled image uploading)

```json
[
{
    "id": "pigeon",
        "execute-command": "PUT_YOUR_PIGEON_BIN_PATH",
        "incoming-payload-content-type": "application/json",
        "pass-environment-to-command": [
        {
            "source": "string",
            "envname": "PIGEON_ESA_KEY",
            "name": "PUT_YOUR_ESA_ACCESS_TOKEN"
        },
        {
            "source": "string",
            "envname": "PIGEON_ESA_TEAM",
            "name": "PUT_YOUR_ESA_TEAM"
        },
        {
            "source": "string",
            "envname": "PIGEON_DIARY_REPO_URL",
            "name": "PUT_YOUR_REPO_URL"
        },
        {
            "source": "string",
            "envname": "PIGEON_PROJECT_ARTICLE_DIR",
            "name": "content/posts"
        },
        {
            "source": "string",
            "envname": "PIGEON_IMAGE_MODE",
            "name": "none"
        },
        {
            "source": "string",
            "envname": "PIGEON_GIT_BIN",
            "name": "/usr/bin/git"
        }
    ],
    "pass-arguments-to-command": [
    {
        "source": "string",
        "name": "-id"
    },
    {
        "source": "payload",
        "name": "post.number"
    }
    ],
    "http-methods": ["POST"],
    "trigger-rule": {
        "match":
        {
            "type": "payload-hash-sha256",
            "secret": "PUT_YOUR_SECRET",
            "parameter":
            {
                "source": "header",
                "name": "X-Esa-Signature"
            }
        }
    }
}
]
```
