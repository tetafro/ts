# Time parser and formatter

Run `ts` to parse input time of different formats and get timestamp and
RFC representations.

## Install

```sh
go get -u github.com/tetafro/ts
```

## Usage

```sh
$ ts 1562913751
Timestamp: 1562913751
RFC 3339:  2019-07-12T06:42:31Z
RFC 1123:  Fri, 12 Jul 2019 06:42:31 UTC

$ ts 2019-07-12T09:42:31+03:00
Timestamp: 1562913751
RFC 3339:  2019-07-12T06:42:31Z
RFC 1123:  Fri, 12 Jul 2019 06:42:31 UTC

$ ts 2019/05/23 12:54
Timestamp: 1558616040
RFC 3339:  2019-05-23T12:54:00Z
RFC 1123:  Thu, 23 May 2019 12:54:00 UTC
```
