##Purpose
This tool was created to provide system stats on a Raspberry PI via a Rest API.

## API Doc
#### ***HTTP GET***
#### /stats
The following will be returned from the `/stats` endpoint

```json
	{
		"TotalMemory": "733 MB",
		"AvailableMemory": "108 MB",
		"UsedMemPercentage": 85.17118162985999,
		"CPULoad": "{load1:0.08,load5:0.08,load15:0.12}",
		"UpTimeDays": "13 Days",
		"DiskUsagePercentage": 28.302848394298085,
		"DiskTotal": "7 GB",
		"DiskUsed": "2 GB"
	}
```

## Example of How to run
`$ ./sys-stats-arm-build -drive "/media"`

## How to compile for Raspberry PI 2
`$ GOOS=linux GOARCH=arm GOARM=7 go build .`

[Additional build flags from the go docs](https://golang.org/doc/install/source#environment)
