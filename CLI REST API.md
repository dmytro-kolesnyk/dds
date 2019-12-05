# DDS Client-Daemon Interaction

## Uploading file

### CLI Command

`./ddsctl --upload [--strategy [copy | fragment | fragment-copy] [--background | -b] {PATH}`

* `-b` `--background` - file will be saved asynchronously.
* `--strategy` is optional and can be used together with values:
  * `copy` - stored data is copied to all active nodes without fragmenting.
  * `fragment` - stored data is fragmented and distributed among active nodes.
  * `fragment-copy` - stored data is fragmented and distributed among active nodes. Each fragment has recovery copy on another node.

Strategy aliases (???):
* `-sc` - `copy`
* `-sf` - `fragment`
* `-sfc` - `fragment-copy`

Note: User can override the default strategy in DDS daemon configuration.

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
Possible `strategy` field values:
* `copy`
* `fragment`
* `fragment-copy`
* `default` - if user does not set saving strategy explicitly.

#### Response

```json
{
  "uuid": "f4c8de96-4e03-4772-b83c-f8dfbe64e998",
  "error": "Timeout to DDS daemon"
}
```
Note: Response body can have either `uuid` or `error` field.

---

## Getting list of files

### CLI Command

`./ddsctl --list`

Returns a list of existing (ready to download) files and their UUID managed by distributed storage.

### `GET /files`

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

## Downloading file

### CLI Command

`./ddsctl --download --path [-p] {PATH} [--background | -b] {UUID}`

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

## Deleting file

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

## Getting status of file

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
      "status": "downloading",
      "progress": 31.5
    }
  ]
}
```

Possible `status` field values:
* `ready` - the file is saved in distributed storage and available for downloading.
* `uploading` - the file is currently uploading.
* `downloading` - the file is currently downloading.
* `deleted` - the file is deleted from distributed storage.

The `progress` shows the completion in percents.
For example, for fully uploaded/downloaded file `progress` will be `100.0`.
