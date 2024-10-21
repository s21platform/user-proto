# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [user.proto](#user-proto)
    - [Birthday](#-Birthday)
    - [GetLoginByUUIDIn](#-GetLoginByUUIDIn)
    - [GetLoginByUUIDOut](#-GetLoginByUUIDOut)
    - [GetUserByLoginIn](#-GetUserByLoginIn)
    - [GetUserByLoginOut](#-GetUserByLoginOut)
    - [GetUserInfoByUUIDIn](#-GetUserInfoByUUIDIn)
    - [GetUserInfoByUUIDOut](#-GetUserInfoByUUIDOut)
    - [IsUserExistByUUIDIn](#-IsUserExistByUUIDIn)
    - [IsUserExistByUUIDOut](#-IsUserExistByUUIDOut)
  
    - [UserService](#-UserService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="user-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## user.proto



<a name="-Birthday"></a>

### Birthday



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| day | [int64](#int64) |  |  |
| month | [int64](#int64) |  |  |
| year | [int64](#int64) |  |  |






<a name="-GetLoginByUUIDIn"></a>

### GetLoginByUUIDIn



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  |  |






<a name="-GetLoginByUUIDOut"></a>

### GetLoginByUUIDOut



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| login | [string](#string) |  |  |






<a name="-GetUserByLoginIn"></a>

### GetUserByLoginIn
Data in request or getting uuid by login. If User doesnt exist - user will be creating


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| login | [string](#string) |  | Email of target user |






<a name="-GetUserByLoginOut"></a>

### GetUserByLoginOut
Message for response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID of user |
| isNewUser | [bool](#bool) |  | Flag for indicate of new user |






<a name="-GetUserInfoByUUIDIn"></a>

### GetUserInfoByUUIDIn
Request data fo getting user info (for initiator page)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID for target user |






<a name="-GetUserInfoByUUIDOut"></a>

### GetUserInfoByUUIDOut
Response data for initiator page


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nickname | [string](#string) |  |  |
| avatar | [string](#string) |  |  |
| name | [string](#string) | optional |  |
| surname | [string](#string) | optional |  |
| birthdate | [Birthday](#Birthday) | optional |  |
| phone | [string](#string) | optional |  |
| city | [string](#string) | optional |  |
| telegram | [string](#string) | optional |  |
| git | [string](#string) | optional |  |
| os | [string](#string) | optional |  |
| work | [string](#string) | optional |  |
| university | [string](#string) | optional |  |
| skills | [string](#string) | repeated |  |
| hobbies | [string](#string) | repeated |  |






<a name="-IsUserExistByUUIDIn"></a>

### IsUserExistByUUIDIn
Message for request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | UUID for target user |






<a name="-IsUserExistByUUIDOut"></a>

### IsUserExistByUUIDOut
Message for response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| isExist | [bool](#bool) |  | Flag of indicate user exist |





 

 

 


<a name="-UserService"></a>

### UserService
Service for friends

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetUserByLogin | [.GetUserByLoginIn](#GetUserByLoginIn) | [.GetUserByLoginOut](#GetUserByLoginOut) | Add friends method |
| IsUserExistByUUID | [.IsUserExistByUUIDIn](#IsUserExistByUUIDIn) | [.IsUserExistByUUIDOut](#IsUserExistByUUIDOut) |  |
| GetUserInfoByUUID | [.GetUserInfoByUUIDIn](#GetUserInfoByUUIDIn) | [.GetUserInfoByUUIDOut](#GetUserInfoByUUIDOut) |  |
| GetLoginByUUID | [.GetLoginByUUIDIn](#GetLoginByUUIDIn) | [.GetLoginByUUIDOut](#GetLoginByUUIDOut) |  |

 



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

