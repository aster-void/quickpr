# QuickPR

Quickly make a PR from CLI!

state: unstable (alpha)

## Dependencies
- gh
- git

## Setup

### Set up gh
```sh
gh auth login
```

### Set up this binary

```sh
mkdir -p $(go env GOPATH)/pkg/src
cd $(go env GOPATH)/pkg/src
git clone git@github.com:aster-void/quickpr.git
cd quickpr
go build . -o ../bin/quickpr
```

and add this to ~/.{bash,zsh}rc

```sh
export PATH=$PATH:/$(go env GOPATH)/pkg/bin
```

### Run

1. find a random repo (better be a personal one) to pr
2. run `quickpr`
3. follow the guide
4. now you have a PR on the remote repo
