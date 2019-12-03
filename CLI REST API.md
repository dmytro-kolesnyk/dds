# CLI API

## `POST /save`

#### Headers

```text
Content-Type: application/json
```

#### Request

```json
{
  "file_path": "/home/user1/file1.avi",
  "strategy": "copy",
  "background": "true"
}
```
where `strategy` and `background` fields are optional.

#### Response

```json
{
  "file_path": "/home/user1/file1.avi",
  "file_name": "file1.avi",
  "uuid": "f4c8de96-4e03-4772-b83c-f8dfbe64e998"
}
```

#### CLI Command

`./ddsctl --save [--strategy-copy | --strategy-fragment | --strategy-fragment-copy] [--background | -b] path`

* `-b` `--background` - file will be saved asynchronously. User can check status of the operation with `status` command
* `-sc` `--strategy-copy` - stored data is copied to all active nodes without fragmenting.
* `-sf` `--strategy-fragment` - stored data is fragmented and distributed among active nodes.
* `-sfc` `--strategy-fragment-copy` - stored data is fragmented and distributed among active nodes. Each fragment has recovery copy on another node.

User can override the default strategy in DDS daemon configuration.

---

## `GET /list`

Returns list of existing files and their UUID managed by distributed system.

#### Headers

```text
Content-Type: application/json
```

#### Response

```json
{
   "files": [{
      "uuid": "f4c8de96-4e03-4772-b83c-f8dfbe64e998",
      "file_name": "file1.avi"
    }]
}
```

### CLI Command

`./ddsctl --list`

---

## `GET /info`

#### Headers

```
Content-Type: application/json
```

#### Request

```json
{
  "uuid": "f4c8de96-4e03-4772-b83c-f8dfbe64e998"
}
```

#### Response

```json
{
  "uuid": "f4c8de96-4e03-4772-b83c-f8dfbe64e998",
  "file_name": "file1.avi",
  "size": "213421412",
  "strategy": "strategy-copy"
}
```
Response fields are TBD.

### CLI Command

`./ddsctl --info {UUID}`

> We can support info by `fileName` and in case of duplication we show all available uuids

---

## `POST /get`

#### Headers

```
Content-Type: application/json
```

#### Request

```json
{
  "uuid": "f4c8de96-4e03-4772-b83c-f8dfbe64e998"
}
```

#### Response

```json
{
  "local_file_path": "/home/new/path/file1.avi",
  "file_name": "file1.avi"
}
```

### CLI Command

`./ddsctl --get [-p] [-b] {UUID}`

* `-b` `--background` - File will be loaded asynchronously. User can check status of the operation with `status` command
* `-p` `--path` - Path to the directory where file will be stored.

> We can support get by `fileName` and in case of duplication we show all available uuids

---

## `DELETE /delete`

#### Headers

```
Content-Type: application/json
```

#### Request

```json
{
  "uuid": "f4c8de96-4e03-4772-b83c-f8dfbe64e998"
}
```

#### Response

Status

### CLI Command

`./ddsctl --delete [-b] {UUID}`

* `-b` `--background` - File will be loaded asynchronously. User can check status of the operation with `status` command.

> We can support get by `fileName` and in case of duplication we show all available uuids

---

## `POST /status`

#### Headers

```
Content-Type: application/json
```

#### Request

```json
{
  "page_size": 10,                              
  "by_status": ["done", "fail", "run"],        
  "by_operation": ["delete", "save", "get"]     
}
```
where `page_size`, `by_status`, `by_operation` fields are optional.

#### Response

```json
{
  "statuses": [
    {
      "file_name": "file1.avi",
      "operation": "get",
      "status": "run",
      "progress": 31.5
    },
    {
      "file_name": "file2.avi",
      "operation": "delete",
      "status": "run",
      "progress": 20
    },
    {
      "file_name": "file3.avi",
      "operation": "save",
      "status": "run",
      "progress": 90.2
    },
    {
      "file_name": "file4.avi",
      "operation": "save",
      "status": "fail"
    },
    {
      "file_name": "file5.avi",
      "operation": "delete",
      "status": "done"
    }
  ]
}
```

### CLI Command

`./ddsctl --status [-g] [-s] [-d] [-f] [-d] [-r]`

* `-g` `--get` - Table will contain the status of `get` operations
* `-s` `--save` - Table will contain the status of `save` operations
* `-d` `--delete` - Table will contain the status of `delete` operations

* `-f` `--fail` - Table will contain failed processes
* `-d` `--done` - Table will contain succeeded processes
* `-r` `--run` - Table will contain running processes
