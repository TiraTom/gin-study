package domain_obj

import "time"

var dummyNowTime = time.Date(2021, 8, 26, 14, 16, 18, 0, time.UTC)

// 2021-08-23 00:00:01 UTC
var time20210823000001 = time.Date(2021, 8, 23, 0, 0, 1, 0, time.UTC)
var timestamp20210823000001 = int64(1629676801)

// 2021-08-23 00:00:02 UTC
var time20210823000002 = time.Date(2021, 8, 23, 0, 0, 2, 0, time.UTC)
var timestamp20210823000002 = int64(1629676802)

// 2021-08-23 00:00:03 UTC
var time20210823000003 = time.Date(2021, 8, 23, 0, 0, 3, 0, time.UTC)
var timestamp20210823000003 = int64(1629676803)
