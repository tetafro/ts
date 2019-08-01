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

$ ts +2h
Timestamp: 1564663857
RFC 3339:  2019-08-01T15:50:57+03:00
RFC 1123:  Thu, 01 Aug 2019 15:50:57 UTC

$ ts -1d
Timestamp: 1564570921
RFC 3339:  2019-07-31T14:02:01+03:00
RFC 1123:  Wed, 31 Jul 2019 14:02:01 MSK
```
