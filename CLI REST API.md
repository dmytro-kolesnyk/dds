# CLI API

## Upload file

### CLI Command

`./ddsctl --upload [--strategy [copy | fragment | fragment-copy] [--background | -b] path`

* `-b` `--background` - file will be saved asynchronously. User can check status of the operation with `status` command
* --strategy:
    - `copy` - stored data is copied to all active nodes without fragmenting.
    - `fragment` - stored data is fragmented and distributed among active nodes.
    - `fragment-copy` - stored data is fragmented and distributed among active nodes. Each fragment has recovery copy on another node.

Aliases (???):
* `-sc` - copy
* `-sf` `- fragment
* `-sfc` `- fragment-copy

User can override the default strategy in DDS daemon configuration.

### `POST /files/{UUID}`

#### Headers

```text
Content-Type: application/json
```

#### Request

```json
{
  "file_path": "/home/user1/file1.avi",
  "strategy": "copy",
  "store_locally": true
}
```
The field "strategy":
- copy
- fragment
- fragment-copy
- default (if user doesn't set strategy explicitly)

#### Response

```json
{
  "uuid": "f4c8de96-4e03-4772-b83c-f8dfbe64e998",
  "error": "Timeout to DDS daemon"
}
```

---

## List files

### CLI Command

`./ddsctl --list`

### `GET /files`

Returns list of existing files and their UUID managed by distributed system.

#### Headers

Empty.

#### Response

```json
{
   "files": [{
      "uuid": "f4c8de96-4e03-4772-b83c-f8dfbe64e998",
      "file_name": "file1.avi"
    }]
}
```

---

## Download file

### CLI Command

`./ddsctl --download --path [-p] "/save/dir/path" [--background | -b] {UUID}`

* `--path` `-p` - Path to the directory where file will be stored.
* `--background` `-b` - File will be downloaded asynchronously.

### `GET /files/{UUID}?dirpath=%2Fsave%2Fdir%2Fpath`

#### Headers

Empty.

#### Request

No body.

#### Response

```json
{
  "uuid": "f4c8de96-4e03-4772-b83c-f8dfbe64e998",
  "error": "Timeout to DDS daemon"
}
```

---

## Delete file

### CLI Command

`./ddsctl --delete [-b] {UUID}`

* `-b` `--background` - File will be loaded asynchronously. User can check status of the operation with `status` command.

### `DELETE /files/{UUID}`

#### Headers

Empty.

#### Request

No body.

#### Response

```json
{
  "uuid": "f4c8de96-4e03-4772-b83c-f8dfbe64e998",
  "error": "Timeout to DDS daemon"
}
```

---

## Status of file

### CLI Command

`./ddsctl --status {UUID}`

### `GET /files/{UUID}/status`

#### Headers

Empty.

#### Request

No body.

#### Response

```json
{
  "statuses": [
    {
      "uuid": "f4c8de96-4e03-4772-b83c-f8dfbe64e998",
      "status": "download",
      "progress": 31.5
    }
  ]
}
```

Status:
* `upload`
* `download`
* `delete`

Note: for fully uploaded/downloaded file `progress` will be 100%.
