package fetchreport

import "fmt"

// take report struct and print it as csv with delimiter ;
// Name;Memory;Init;Cold;Billed
func (r *Report) Print() {
	fmt.Printf("%s;%v;%v;%v;%v\n", r.Name, r.MemorySize, r.InitDuration, r.Duration, r.BilledDuration)
}
