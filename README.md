# nisekoi

> False love

![build-status](https://circleci.com/gh/Ruenzuo/nisekoi.png?circle-token=f39277cfc2d19ca04b1d5aec1feee4105bc0826e)

Calculate average landing PR times

### Usage

```
NAME:
   nisekoi - Calculate average landing PR times

USAGE:
   nisekoi [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     calc     Calculate average landing PR times
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

### How to run it

#### With `go get`

```
$ go get github.com/Ruenzuo/nisekoi
$ nisekoi calc --username Ruenzuo --access-token 7cd6c98a93e94a0c79407d0e320323a0 CocoaPods/Xcodeproj
Average landing PR time is: 1079.03 hours, for a total of 5 out of 273 landed PRs
```

#### With Docker

```
$ ./docker-run.sh "calc --access-token 7cd6c98a93e94a0c79407d0e320323a0 CocoaPods"
GET https://api.github.com/repos/CocoaPods/Specs/pulls?page=22&per_page=100&state=closed: 403 You have triggered an abuse detection mechanism. Please wait a few minutes before you try again.
Average landing PR time is: 140.92 hours, for a total of 3282 landed PRs
```

### Know issues

If you see something like this in the output log

```
GET https://api.github.com/repos/CocoaPods/CocoaPods/pulls?page=4&per_page=100&state=closed: 403 You have triggered an abuse detection mechanism. Please wait a few minutes before you try again.
```

It means the organization or repository you're trying to calculate is too big. The average returned is calculated with the most recent pull requests.
