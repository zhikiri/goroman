# goroman [![Build Status](https://travis-ci.org/zhikiri/goroman.svg?branch=master)](https://travis-ci.org/zhikiri/goroman)

Numbers convertion library.
Convert arabic to roman numbers and backwards.

## Description

Package contain 2 public types: `Arabic` and `Roman`.

- `Arabic` based on int type and have `toRoman` method

- `Roman` based on string type and have `toArabic` method

## Example

Here is the usage example:

```go
var rom Roman
rom = "CCVII"
fmt.Println(rom.toArabic())
// 207
```

The same logic applying backwards.

## Summary

Project is finished.
