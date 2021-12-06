package analyse

import (
	"strings"
)


type Speed struct {
	Language *string
	DurationMilliSeconds *string
}


func ExtractSpeedTest(line *string) *Speed{
	v := strings.Fields(*line)  
	var duration string
	outside:
	for i, cell := range v {
		if (i >0) && (cell == "Duration:") && (string(v[i-1]) != "Billed")  {
			duration = string(v[i+1])
			break outside
		}
	}
	return &Speed{DurationMilliSeconds: &duration}
}