# CLI API

## `POST /save`

#### Headers

```json
Content-Type: application/json
```

#### Request

```json
{
  "file_path": "/home/user1/file1.avi",
  "strategy": "copy",  // Optional
  "background": "true" // Optional
}
```

#### Response

```json
{
  "file_path": "/home/user1/file1.avi",
  "file_name": "file1.avi",
  "uuid": "f4c8de96-4e03-4772-b83c-f8dfbe64e998"
}
```

#### CLI Command

`save [--strategy-copy | --strategy-fragments | --strategy-fragment-copy] [-b] path`

* `-b` `--background` - File will be saved asynchronously. User can check status of the operation with `status` command
* `--strategy-copy`
`--strategy-fragments`
`--strategy-fragment-copy` - User can override the default strategy of the server.

---

## `GET /list`

Returns list of existing files and their uuid managed by distributed system.

#### Headers

```json
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

`list`

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
  "strategy": "strategy-name",
  ....................
}
```

### CLI Command

`info uuid `

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
  "uuid": "uuid"
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

`get [-p] [-b] uuid`

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
  "uuid": "uuid"
}
```

#### Response

Status

### CLI Command

`delete [-b] uuid`

* `-b` `--background` - File will be loaded asynchronously. User can check status of the operation with `status` command

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
  "page_size": 10,                              // Optional
  "by_status": ["done", "fail", "run"],         // Optional
  "by_operation": ["delete", "save", "get"]     // Optional
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

`status [-g] [-s] [-d] [-f] [-d] [-r]`

* `-g` `--get` - Table will contain the status of `get` operations
* `-s` `--save` - Table will contain the status of `save` operations
* `-d` `--delete` - Table will contain the status of `delet` operations

* `-f` `--fail` - Table will contain failed processes
* `-d` `--done` - Table will contain succeeded processes
* `-r` `--run` - Table will contain running processes