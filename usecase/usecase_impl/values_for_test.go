package usecase_impl_test

import (
	"time"

	gr "github.com/Tiratom/gin-study/grpc"
)

// 2021/08/22 15:00:01 UTC
var time20210822150001 = time.Date(2021, 8, 22, 15, 0, 1, 0, time.UTC)
var timestamp20210822150001 = int64(1629644401)

// 2021/08/22 15:00:02 UTC
var time20210822150002 = time.Date(2021, 8, 22, 15, 0, 2, 0, time.UTC)

// 2021/08/22 15:00:03 UTC
var time20210822150003 = time.Date(2021, 8, 22, 15, 0, 3, 0, time.UTC)

var searchTypeForDeadlineAfter = gr.TimestampCompareBy_AFTER
var searchTypeForDeadlineNone = gr.TimestampCompareBy_NONE
