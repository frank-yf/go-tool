package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func runDate(cmd *cobra.Command, args []string) error {

	// 时间戳转化
	if timestamp > 0 {
		t, err := parseTimestamp(timestamp, duration)
		if err != nil {
			return err
		}
		printTime(t)
		return nil
	}

	// 默认：输出当前时间
	printTime(time.Now())
	return nil
}

var dateCmd = &cobra.Command{
	Use:   "date",
	Short: "时间工具",
	RunE:  runDate,
}

var (
	logger = log.New(os.Stdout, "", 0)

	timestamp int64
	duration  time.Duration
	output    string
	layout    string
)

func init() {
	rootCmd.AddCommand(dateCmd)

	dateCmd.Flags().
		Int64Var(&timestamp, "ts", 0, "时间戳转换，单位可以通过 --duration 指定")
	dateCmd.Flags().
		DurationVarP(&duration, "duration", "d", time.Millisecond, "时间单位")
	dateCmd.Flags().
		StringVar(&output, "output", timeOutputTypeStringer, "时间输出格式")
	dateCmd.Flags().
		StringVarP(&layout, "layout", "l", time.RFC3339Nano, "时间格式化")

}

const (
	timeOutputTypeStringer  = "stringer"
	timeOutputTypeFormat    = "format"
	timeOutputTypeGoString  = "goString"
	timeOutputTypeTimestamp = "ts"
)

func printTime(t time.Time) error {
	switch output {
	case timeOutputTypeFormat:
		if layout == "" {
			return errors.New("output='format' 必须指定 layout 参数")
		}
		logger.Println(t.Format(layout))
	case timeOutputTypeTimestamp:
		ts, err := formatTimestamp(t, duration)
		if err != nil {
			return err
		}
		logger.Println(ts)
	case timeOutputTypeGoString:
		logger.Println(t.GoString())
	case timeOutputTypeStringer:
		logger.Println(t)
	default:
		return fmt.Errorf("不支持的 output '%s'", output)
	}
	return nil
}

func parseTimestamp(ts int64, unit time.Duration) (t time.Time, err error) {
	switch unit {
	case time.Millisecond:
		t = time.UnixMilli(ts)
	case time.Second:
		t = time.Unix(ts, 0)
	default:
		err = fmt.Errorf("暂不支持时间单位 '%s'", unit.String())
	}
	return
}

func formatTimestamp(t time.Time, unit time.Duration) (ts int64, err error) {
	switch unit {
	case time.Nanosecond:
		ts = t.UnixNano()
	case time.Microsecond:
		ts = t.UnixMicro()
	case time.Millisecond:
		ts = t.UnixMilli()
	case time.Second:
		ts = t.Unix()
	default:
		err = fmt.Errorf("暂不支持时间单位 '%s'", unit.String())
	}
	return
}
