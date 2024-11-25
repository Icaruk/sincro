# sincro

CLI tool written in Go to sync your files between a source and multiple destinations.

![sincro circle](https://i.imgur.com/NCAcwYQ.png)


[![built with Codeium](https://codeium.com/badges/main)](https://codeium.com)



# Use cases

- sync files with **JSDoc** from backend and frontend
- sync **types** from main repository to multiple destinations



# Getting started

```bash
sincro init
```



# Config

sincro.json
```json
{
	"$schema": "https://raw.githubusercontent.com/Icaruk/sincro/main/json-schema.json",
	"version": 1,
	"id": "project_m3xdD55Apsq0Uv83",
	"type": "source",
	"sync": [
		{
			"source": "./relative/path/to/source",
			"destinations": [
				"C:/absolute/path/to/destination/1",
				"C:/absolute/path/to/destination/2"
			]
		}
	]
}
```

- `sync.*.source` must be relative path from root.
- `sync.*.destinations` must be absolute path.



# Usage

```bash
sincro [command] [flags]

sincro init
sincro push
sincro watch
```
