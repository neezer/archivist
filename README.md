> Easily compress and upload application artifacts.

<p align="center">
  <img width="840" src="https://cdn.jsdelivr.net/gh/kofile/archivist/example.svg">
</p>

# Archivist

Worry about building your app: let Archivist handle compressing/uploading it.

- Standalone binary
- Works with AWS S3

**_NOTE_**: Currently has some hard-coded configuration suited to
https://github.com/kofile projects.

## Why?

[write up about previous script]

## Usage

**TL;DR** `archivist archive [name of project] [build directory]`

```
archivist assists in creating artifact archives suitable for deployment
to an asset server. You can use it to generate compressed archives, upload an
archive to an asset server, or both in one command.

Usage:
  archivist [command]

Available Commands:
  archive     Create compressed artifact and upload to server
  compress    Compress an artifact directory
  help        Help about any command
  upload      upload an artifact to s3 artifacts bucket

Flags:
  -h, --help   help for archivist

Use "archivist [command] --help" for more information about a command.
```

- Default bucket: `artifacts-kofile-systems`
- [`~/.aws/credentials` and `~/.aws/config`
  files](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html)
- `git` required if `GIT_BRANCH` and `GIT_COMMIT` environment variables are not
  set (these are often already set in CI environments like Jenkins or TravisCI)

## Developing

[TODO]

---

Author: [neezer](https://github.com/neezer)

![20dc3d3e1ef3251a52f92dded2b1c785](https://user-images.githubusercontent.com/29997/52675386-8e340480-2edb-11e9-8d8c-cf212218e3f8.jpg)
