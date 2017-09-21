# hba

## Depencies

This projects uses the packages bellow to build:

```bash
go get github.com/goreleaser/goreleaser
```

And (for now) have a ruby gem dependency:

```bash
gem install fpm
```

### osx related dependecies:

```bash
brew install rpm
brew install dpkg
```
> More details here: http://timperrett.com/2014/03/23/enabling-rpmbuild-on-mac-osx/

## How to built it

```bash
goreleaser --rm-dist
```