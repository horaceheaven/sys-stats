##Purpose
This tool was created to provide system stats on a Raspberry PI via a Rest API.

## API Doc
#### /stats
The following will be returned from the `/stats` endpoint
	
```golang
	TotalMemory         string
	AvailableMemory     string
	UsedMemPercentage   float64
	CPULoad             string
	UpTimeDays          string
	DiskUsagePercentage float64
	DiskTotal           string
	DiskUsed            string
```

## Example of How to run
`$ ./sys-stats-arm-build -drive "/media"`

## How to compile for Raspberry PI 2
`$ GOOS=linux GOARCH=arm GOARM=7 go build .`

[Additional build flags from the go docs](https://golang.org/doc/install/source#environment)
