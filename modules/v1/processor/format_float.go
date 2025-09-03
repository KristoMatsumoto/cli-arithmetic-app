package processor

import (
	"fmt"
	"math"
)

func FormatFloat(f float64) string {
	if math.IsNaN(f) {
		return "NaN"
	}
	if math.Mod(f, 1) == 0 {
		return fmt.Sprintf("%.0f", f)
	}
	return fmt.Sprintf("%.2f", f)
}
