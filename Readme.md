# Shell Completer for Camus2

## TL;DR

``` shell
./c2completer -i
source ~/.bashrc # or zsh or fish
c2 <tab><tab>
```

## Install

```
wget -O c2completer.gz -q https://github.com/gmhta/c2completer/releases/download/v1.0.0/c2completer_1.0.0_linux_amd64.gz
gzip -d c2completer.gz 
chmod +x c2completer
```

## Build

```
git clone https://github.com/gmhta/c2completer.git
cd c2completer
go install
```

## Release

Releases are done using `goreleaser`. Details not provided.

## Todo

- Context sensitive completion for endpoints, deployments etc.
- Move releases to a github action.
