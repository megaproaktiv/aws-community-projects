package analyse_test

import (
	"analyse"
	"testing"

	"gotest.tools/assert"
)


func TestExtractSpeedCold(t *testing.T){
	cold := "REPORT RequestId: a5c93795-5c22-4035-8e53-19f05230aee9	Duration: 40.55 ms	Billed Duration: 41 ms	Memory Size: 1024 MB	Max Memory Used: 81 MB	Init Duration: 410.71 ms"
	
	bench := analyse.ExtractSpeedTest(&cold)
	assert.Equal(t, "40.55", *bench.DurationMilliSeconds)
}
func TestExtractSpeedWarm(t *testing.T){
	warm := "REPORT RequestId: faeec3c3-8313-4f74-9b4e-5191b8e2cdc1	Duration: 3.59 ms	Billed Duration: 4 ms	Memory Size: 1024 MB	Max Memory Used: 81 MB"
	
	bench := analyse.ExtractSpeedTest(&warm)
	assert.Equal(t, "3.59", *bench.DurationMilliSeconds)

}