# ignoreit

[![Go Report Card](https://goreportcard.com/badge/github.com/danielkvist/ignoreit)](https://goreportcard.com/report/github.com/danielkvist/ignoreit)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)

ignoreit is a CLI writen in Go to create `.gitignore` files using [gitignore.io](https://gitignore.io).

## Example

```bash
ignoreit linux golang
```

The command above should generate the following `.gitignore` file:

```txt

# Created by https://www.gitignore.io/api/linux,golang
# Edit at https://www.gitignore.io/?templates=linux,golang

#!! ERROR: golang is undefined. Use list command to see defined gitignore types !!#

### Linux ###
*~

# temporary files which can be created if a process still has a handle open of a deleted file
.fuse_hidden*

# KDE directory preferences
.directory

# Linux trash folder which might appear on any partition or disk
.Trash-*

# .nfs files are created when an open file is removed but is still being accessed
.nfs*

# End of https://www.gitignore.io/api/linux,golang

```

> If a `.gitignore` already exists in the directory the command will fail.

## Installation

```bash
go get -u github.com/danielkvist/ignoreit

# And

go install github.com/danielkvist/ignoreit
```
