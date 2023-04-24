package internal

import (
	"fmt"
	"unicode"
)

var (
	bar Bar
)

// Bar
type Bar struct {
	percent int64 // progress percentage
	cur     int64 // current progress
	total   int64
	rate    string // the actual progress bar to be printed
	graph   string // the fill value for progress bar
}

func NewBar() {
	bar.cur = 0
	bar.total = 1
	if bar.graph == "" {
		bar.graph = "#"
	}
	bar.percent = bar.getPercent()
	for i := 0; i < int(bar.percent); i += 2 {
		bar.rate += bar.graph // initial progress position
	}
}

func (bar *Bar) getPercent() int64 {
	return int64((float32(bar.cur) / float32(bar.total)) * 100)
}

func addToBar(n int64) {
	bar.total = bar.total + n
}

func printBar(cur int64, str string) {
	bar.cur = bar.cur + cur
	last := bar.percent
	bar.percent = bar.getPercent()
	if bar.percent != last && bar.percent%2 == 0 {
		bar.rate += bar.graph
	}
	// [%-50s]
	// bar.rate,
	erase := "\r                                                                                                              "
	max := len(erase)
	out := fmt.Sprintf("\r\u001b[1m %3d%% %s/%s \033[0m writing %s.. ", bar.percent, fmtByteCount(bar.cur), fmtByteCount(bar.total), str)
	fmt.Print(erase)
	fmt.Print(trunc(out, max))

}

func FinishBar() {
	fmt.Println()
	fmt.Println("finished")
}

func trunc(text string, maxLen int) string {
	lastSpaceIx := maxLen
	len := 0
	for i, r := range text {
		if unicode.IsSpace(r) {
			lastSpaceIx = i
		}
		len++
		if len > maxLen {
			return text[:lastSpaceIx] + ".."
		}
	}
	// If here, string is shorter or equal to maxLen
	return text
}

func fmtByteCount(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}
