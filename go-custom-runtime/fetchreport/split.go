package fetchreport

import (
	"regexp"
	"strconv"
)

func Split(input string) *Report {
	//  input := "2023-07-29 18:30:11.606 REPORT RequestId: 2b8f9fc5-cef3-49d3-bc19-58c1f6171b2c Duration: 8.33 ms Billed Duration: 9 ms Memory Size: 128 MB Max Memory Used: 66 MB Init Duration: 166.90 ms"

	line := Report{}

	// Regular expression patterns to match the required values
	memorySizePattern := regexp.MustCompile(`Memory Size: (\d+) MB`)
	durationPattern := regexp.MustCompile(`Duration: ([\d.]+) ms`)
	billedDurationPattern := regexp.MustCompile(`Billed Duration: (\d+) ms`)
	initDurationPattern := regexp.MustCompile(`Init Duration: ([\d.]+) ms`)

	// Find the matches for each pattern
	memorySizeMatch := memorySizePattern.FindStringSubmatch(input)
	durationMatch := durationPattern.FindStringSubmatch(input)
	billedDurationMatch := billedDurationPattern.FindStringSubmatch(input)
	initDurationMatch := initDurationPattern.FindStringSubmatch(input)

	// Extract the values from the matches
	var memorySize int
	var duration, billedDuration, initDuration float64

	if len(memorySizeMatch) > 1 {
		memorySize, _ = strconv.Atoi(memorySizeMatch[1])
		line.MemorySize = memorySize
	}

	if len(durationMatch) > 1 {
		duration, _ = strconv.ParseFloat(durationMatch[1], 64)
		line.Duration = duration
	}

	if len(billedDurationMatch) > 1 {
		billedDuration, _ = strconv.ParseFloat(billedDurationMatch[1], 64)
		line.BilledDuration = billedDuration
	}

	if len(initDurationMatch) > 1 {
		initDuration, _ = strconv.ParseFloat(initDurationMatch[1], 64)
		line.InitDuration = initDuration
	}

	return &line
}
