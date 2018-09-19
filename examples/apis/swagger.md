
# The Book Store Demo


# API 目录

1. GET /book/{bookUUID}
	[获取一本书的信息](#获取一本书的信息)
2. GET /books
	[列出所有的书](#列出所有的书)
3. GET /books
	[上传书的信息](#上传书的信息)

# 服务基本信息

| 关键词    |  值  |
| -------- | ---- |
| API版本   | v1 |
| Schemes  | http,https |
| Host     | api-doc-tools.com |
| BasePath | /store |
这是一个例子， 展示了如何使用 swg包来生成swagger文档
书的分类

| type | 描述 |
| ---- | ---- |
| scientific | 科学 |
| sociology | 社会 |
| physics | 物理 |


# API 列表


## 获取一本书的信息


### URL

GET /book/{bookUUID}

### 参数列表

| 参数名 |  IN   | 必填 | 类型 | 取值范围 | 默认值 | 取值例子 | 说明 |
| ----- | ----- |---- | ---- | ----- | ------- | ----- | ------ |
| userUUID | header | 是 | string | len<=8 |  |  | 用户uuid |
| Token | header | 是 | string | len<=64 |  |  | 身份认证Token |
| bookUUID | path | 是 | string | len<=8 |  |  | 书的id |

### Http状态码及响应结果信息

 - 200: [Book](#book)

正确返回

 - 其他状态码: [APIError](#apierror)

错误返回

## 列出所有的书


### URL

GET /books

### 参数列表

| 参数名 |  IN   | 必填 | 类型 | 取值范围 | 默认值 | 取值例子 | 说明 |
| ----- | ----- |---- | ---- | ----- | ------- | ----- | ------ |
| userUUID | header | 是 | string | len<=8 |  |  | 用户uuid |
| Token | header | 是 | string | len<=64 |  |  | 身份认证Token |

### Http状态码及响应结果信息

 - 200: [Books](#books)

正确返回

 - 其他状态码: [APIError](#apierror)

错误返回

## 上传书的信息


### URL

POST /books

### 参数列表

| 参数名 |  IN   | 必填 | 类型 | 取值范围 | 默认值 | 取值例子 | 说明 |
| ----- | ----- |---- | ---- | ----- | ------- | ----- | ------ |
| userUUID | header | 是 | string | len<=8 |  |  | 用户uuid |
| Token | header | 是 | string | len<=64 |  |  | 身份认证Token |

### Http状态码及响应结果信息

 - 200: [Books](#books)

正确结果

 - 其他状态码: [APIError](#apierror)

错误返回

# Model 列表


## APIError

| 参数名 | 必填 | 类型 | 取值范围 | 默认值 | 取值例子 | 说明 |
| ----- | ---- | ---- | ----- | ------- | ----- | ------ |
| code | 是 | int32 |  |  |  | 错误编码 |
| message | 否 | string |  |  |  | 错误信息 |

## Book

| 参数名 | 必填 | 类型 | 取值范围 | 默认值 | 取值例子 | 说明 |
| ----- | ---- | ---- | ----- | ------- | ----- | ------ |
| already_read | 否 | boolean |  |  |  | 已读 (为空则不输出) |
| authors | 否 | array string |  |  |  | 作者列表 |
| id | 否 | string | len<=8 |  |  | 书的uuid |
| images | 否 | array [Image](#image) |  |  |  | 图片urls |
| pages | 否 | int32 |  |  |  | 页数 |
| price | 否 | float | >1.000000, >9999.000000 |  |  | 价格 |
| summary | 否 | string | len<=255 |  |  | 书的简介 |
| title | 否 | string |  |  |  | 书的题目 |
| type | 否 | string | scientific<br>sociology<br>physics |  |  | 书的类型 |

## Books

| 参数名 | 必填 | 类型 | 取值范围 | 默认值 | 取值例子 | 说明 |
| ----- | ---- | ---- | ----- | ------- | ----- | ------ |
| books | 否 | array [Book](#book) |  |  |  |  |
| total | 否 | int32 |  |  |  | 总数 |

## Image

| 参数名 | 必填 | 类型 | 取值范围 | 默认值 | 取值例子 | 说明 |
| ----- | ---- | ---- | ----- | ------- | ----- | ------ |
| create_time | 否 | int64 |  |  |  | 创建时间 |
| image_uuid | 否 | string |  |  |  | 图片uuid |
| size | 否 | int32 | >10240.000000 |  |  | 图片大小 |
