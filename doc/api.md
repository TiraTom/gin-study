# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [gin-study.proto](#gin-study.proto)
    - [CreateTaskRequestParam](#.CreateTaskRequestParam)
    - [DeleteTaskRequestParam](#.DeleteTaskRequestParam)
    - [GetMyCatMessage](#.GetMyCatMessage)
    - [GetTaskByConditionRequestParam](#.GetTaskByConditionRequestParam)
    - [GetTaskByIdRequestParam](#.GetTaskByIdRequestParam)
    - [Importance](#.Importance)
    - [MyCatResponse](#.MyCatResponse)
    - [Task](#.Task)
    - [Tasks](#.Tasks)
    - [UpdateTaskRequestParam](#.UpdateTaskRequestParam)
  
    - [TimestampCompareBy](#.TimestampCompareBy)
  
    - [CatService](#.CatService)
    - [TaskService](#.TaskService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="gin-study.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gin-study.proto



<a name=".CreateTaskRequestParam"></a>

### CreateTaskRequestParam



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| details | [string](#string) |  |  |
| importanceName | [string](#string) |  |  |
| deadline | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | 期限日時（タイムスタンプ） |






<a name=".DeleteTaskRequestParam"></a>

### DeleteTaskRequestParam



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name=".GetMyCatMessage"></a>

### GetMyCatMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| target_cat | [string](#string) |  |  |






<a name=".GetTaskByConditionRequestParam"></a>

### GetTaskByConditionRequestParam



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| details | [string](#string) |  |  |
| importanceName | [string](#string) |  | 重要度 |
| deadline | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | 期限日時（タイムスタンプ） |
| searchTypeForDeadline | [TimestampCompareBy](#TimestampCompareBy) |  |  |






<a name=".GetTaskByIdRequestParam"></a>

### GetTaskByIdRequestParam



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name=".Importance"></a>

### Importance



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| level | [uint32](#uint32) |  |  |






<a name=".MyCatResponse"></a>

### MyCatResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| kind | [string](#string) |  |  |






<a name=".Task"></a>

### Task



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| details | [string](#string) |  |  |
| importanceName | [string](#string) |  | 重要度 |
| registered_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | 登録日時（タイムスタンプ） |
| deadline | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | 期限日時（タイムスタンプ） |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | 更新日時（タイムスタンプ） |






<a name=".Tasks"></a>

### Tasks



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tasks | [Task](#Task) | repeated |  |






<a name=".UpdateTaskRequestParam"></a>

### UpdateTaskRequestParam



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| details | [string](#string) |  |  |
| importanceName | [string](#string) |  |  |
| deadline | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | 期限日時（タイムスタンプ） |





 


<a name=".TimestampCompareBy"></a>

### TimestampCompareBy


| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 |  |
| SAME | 1 |  |
| BEFORE | 2 |  |
| AFTER | 3 |  |


 

 


<a name=".CatService"></a>

### CatService
お試し用

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetMyCat | [.GetMyCatMessage](#GetMyCatMessage) | [.MyCatResponse](#MyCatResponse) |  |


<a name=".TaskService"></a>

### TaskService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetAllTasks | [.google.protobuf.Empty](#google.protobuf.Empty) | [.Tasks](#Tasks) | 全タスクの取得 |
| GetTasks | [.GetTaskByConditionRequestParam](#GetTaskByConditionRequestParam) | [.Tasks](#Tasks) | 条件付きタスクの取得 |
| GetTask | [.GetTaskByIdRequestParam](#GetTaskByIdRequestParam) | [.Task](#Task) | id指定でタスクの取得 ・エラー返却について - NotFound：idに対応するタスクが存在しない - InvalidArgument：引数の内容が不適切 |
| CreateTask | [.CreateTaskRequestParam](#CreateTaskRequestParam) | [.Task](#Task) | タスクの新規作成 ・エラー返却について - InvalidArgument：引数の内容が不適切 |
| UpdateTask | [.UpdateTaskRequestParam](#UpdateTaskRequestParam) | [.Task](#Task) | タスクの更新 ・エラー返却について - NotFound：idに対応するタスクが存在しない - InvalidArgument：引数の内容が不適切 |
| DeleteTask | [.DeleteTaskRequestParam](#DeleteTaskRequestParam) | [.google.protobuf.Empty](#google.protobuf.Empty) | タスクの一件削除 ・エラー返却について - NotFound：idに対応するタスクが存在しない - InvalidArgument：引数の内容が不適切 |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

