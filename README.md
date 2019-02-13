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

Many of our existing microservices deploy their artifacts using a Node script I
write in about an hour near the end of 2017. This script was initially meant as
a quick and dirty solution to get artifacts up to our S3 bucket.

My original script has been copy/pasted into nearly every one of our
microservices and in many cases has not been updated since. Dependencies the
script relies on are way out of date and in some cases completely deprecated
(eg., [`tar.gz`](https://www.npmjs.com/package/tar.gz)).

The script has the following additional drawbacks:

- bloats `package.json` with dependencies that only pertain to getting artifacts
  on S3
- increases CI build times due to needing to install said packages during `npm install`
- microservices have different versions of the script, sometimes even using
  different environment variables
- can only be used in NodeJS projects (or on systems with NodeJS installed)

`archivist` aims to solve all of these issues.

- standalone binary means no `package.json` bloat nor pervasive NPM-related
  security issues
- decreases CI build times
- if installed globally on CI, then all projects use the same version
- can be used in any language
- if CI defines certain environment variables, then it has no external
  dependencies

## Usage

### Commands

#### `archivist compress [dir to compress] [destination dir]`

Takes a directory to compress and creates a `.tar.gz` archive in the destination
directory. This archive is named for the current git SHA and will de-compress
into a similarly named folder.

NOTE: `archivist` includes a Go implemtation of `tar`, so it is not necessary to
have `tar` installed.

Example:

```shell
$ git rev-parse HEAD
12345abcde

$ tree build
build
├── a-file.txt
└── b-file.txt

$ archivist compress build dist

$ tree dist
dist
├── 12345abcde
│   ├── a-file.txt
│   └── b-file.txt
└── 12345abcde.tar.gz

$ tar zxf dist/12345abcde.tar.gz

$ tree 12345abcde
12345abcde
├── a-file.txt
└── b-file.txt

```

#### `archivist upload [name of project] [path to file to upload] [flags]`

Takes a name and a file to upload and places the file into the configured S3
bucket at the following location:

```
s3://your-bucket/name-of-project/branch-name/git-sha.tar.gz
```

See the section below about [configuring access to
S3](https://github.com/kofile/archivist#s3-configcredentials).

#### `archivist archive [name of project] [directory to compress] [flags]`

Combines the `compress` and `upload` commands into one. See the above sections
on what each command does.

You can pass all of the `--s3-*` flags to this command as well as to `upload`.

### External Dependencies

If you do not have the environment variables `GIT_BRANCH` or `GIT_COMMIT` set,
then `git` needs to be present; usually these variables are set in most CI
environments, like Jenkins or Travis.

Otherwise there are no external dependencies!

### S3 Config/Credentials

You can provide S3 credentials using configuration files or environment
variables [per the official Amazon AWS SDK
docs](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html).

Additionally, you can use the following flags:

- `--s3-bucket` (default: artifacts-kofile-systems)
- `--s3-region`
- `--s3-access-key`
- `--s3-secret-access-key`

## Developing

`make help`

**NOTE**: Binaries produced with `make cross-build` are compiled with
`-ldflags="-s -w"`. The `linux/amd64` binary is packed with
[UPX](https://upx.github.io/).

## Author

[neezer](https://github.com/neezer)

---

<p align="center">
  <img width="480" src="https://user-images.githubusercontent.com/29997/52675386-8e340480-2edb-11e9-8d8c-cf212218e3f8.jpg">
</p>
