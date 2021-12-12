package usecase

import (
	"time"

	gr "github.com/Tiratom/gin-study/grpc"
)

// 2021-08-21 15:00:01 UTC
var time20210821150001 = time.Date(2021, 8, 21, 15, 0, 1, 0, time.UTC)
var timestamp20210821150001 = 1629558001
var tokyoTimeOfUtc20210821150001 = time.Date(2021, 8, 22, 0, 0, 1, 0, time.FixedZone("Asia/Tokyo", 9*60*60))

// 2021/08/22 15:00:01 UTC
var time20210822150001 = time.Date(2021, 8, 22, 15, 0, 1, 0, time.UTC)
var timestamp20210822150001 = 1629644401
var tokyoTimeOfUtc20210822150001 = time.Date(2021, 8, 23, 0, 0, 1, 0, time.FixedZone("Asia/Tokyo", 9*60*60))

// 2021/08/22 15:00:02 UTC
var time20210822150002 = time.Date(2021, 8, 22, 15, 0, 2, 0, time.UTC)
var timestamp20210822150002 = 1629644402
var tokyoTimeOfUtc20210822150002 = time.Date(2021, 8, 23, 0, 0, 2, 0, time.FixedZone("Asia/Tokyo", 9*60*60))

// 2021/08/22 15:00:03 UTC
var time20210822150003 = time.Date(2021, 8, 22, 15, 0, 3, 0, time.UTC)
var timestamp20210822150003 = 1629644403
var tokyoTimeOfUtc20210822150003 = time.Date(2021, 8, 23, 0, 0, 3, 0, time.FixedZone("Asia/Tokyo", 9*60*60))

// 2021/08/22 15:00:04 UTC
var time20210822150004 = time.Date(2021, 8, 22, 15, 0, 4, 0, time.UTC)
var timestamp20210822150004 = 1629644404
var tokyoTimeOfUtc20210822150004 = time.Date(2021, 8, 23, 0, 0, 4, 0, time.FixedZone("Asia/Tokyo", 9*60*60))

// 2021/08/23 15:00:01 UTC
var time20210823150001 = time.Date(2021, 8, 23, 15, 0, 1, 0, time.UTC)
var timestamp20210823150001 = 1629730801
var tokyoTimeOfUtc20210823150001 = time.Date(2021, 8, 24, 0, 0, 1, 0, time.FixedZone("Asia/Tokyo", 9*60*60))

// 2021/08/24 15:00:01 UTC
var time20210824150001 = time.Date(2021, 8, 24, 15, 0, 1, 0, time.UTC)
var timestamp20210824150001 = 1629817201
var tokyoTimeOfUtc20210824150001 = time.Date(2021, 8, 25, 0, 0, 1, 0, time.FixedZone("Asia/Tokyo", 9*60*60))

// 2021/09/22 15:00:01 UTC
var time20210922150001 = time.Date(2021, 9, 22, 15, 0, 1, 0, time.UTC)
var timestamp20210922150001 = 1632322801
var tokyoTimeOfUtc20210922150001 = time.Date(2021, 9, 23, 0, 0, 1, 0, time.FixedZone("Asia/Tokyo", 9*60*60))

// 2021/09/22 15:00:02 UTC
var time20210922150002 = time.Date(2021, 9, 22, 15, 0, 2, 0, time.UTC)
var timestamp20210922150002 = 1632322802
var tokyoTimeOfUtc20210922150002 = time.Date(2021, 9, 23, 0, 0, 2, 0, time.FixedZone("Asia/Tokyo", 9*60*60))

// 2021/09/22 15:00:03 UTC
var time20210922150003 = time.Date(2021, 9, 22, 15, 0, 3, 0, time.UTC)
var timestamp20210922150003 = 1632322803
var tokyoTimeOfUtc20210922150003 = time.Date(2021, 9, 23, 0, 0, 3, 0, time.FixedZone("Asia/Tokyo", 9*60*60))

// 2021/09/22 15:00:04 UTC
var time20210922150004 = time.Date(2021, 9, 22, 15, 0, 4, 0, time.UTC)
var timestamp20210922150004 = 1632322804
var tokyoTimeOfUtc20210922150004 = time.Date(2021, 9, 23, 0, 0, 4, 0, time.UTC)

var searchTypeForDeadlineAfter = gr.TimestampCompareBy_AFTER
var searchTypeForDeadlineBefore = gr.TimestampCompareBy_BEFORE
var searchTypeForDeadlineSame = gr.TimestampCompareBy_SAME
var searchTypeForDeadlineNone = gr.TimestampCompareBy_NONE
