# QuickPR

Quickly make a PR from CLI!

state: alpha

issues: Welcome

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
mkdir -p $(go env GOPATH)/src
cd $(go env GOPATH)/src
git clone git@github.com:aster-void/quickpr.git
cd quickpr
./build.sh
```

and add this to ~/.{bash,zsh}rc

```sh
export PATH=$PATH:/$(go env GOPATH)/bin
```

### Run

1. find a random repo (better be a personal one, because it's not stable yet) to pr
2. run `quickpr`
3. follow the guide
4. now you have a PR on the remote repo

### Stop

Enter `q` or ctrl+C to stop it. It will not change Git state until the last step is done.

### Different base branch

you can also make a PR to a different branch than main.

```sh
quickpr master
```

should make a PR to branch `master` (unchecked)


