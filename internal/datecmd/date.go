package datecmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func runDate(cmd *cobra.Command, args []string) error {
	loc, err := parseLocation(location)
	if err != nil {
		return err
	}

	// 解析时间字符串
	if timeStr != "" {
		if layout == "" {
			return errors.New("解析时间字符串时 layout 参数必须填写")
		}
		t, err := time.ParseInLocation(layout, timeStr, loc)
		if err != nil {
			return err
		}
		return printTime(t)
	}

	// 解析时间戳
	if timestamp > 0 {
		t, err := parseTimestamp(timestamp, duration)
		if err != nil {
			return err
		}
		return printTime(t)
	}

	// 默认：输出当前时间
	return printTime(time.Now())
}

var (
	timestamp int64
	timeStr   string
	duration  time.Duration
	output    string
	layout    string
	location  string

	logger = zap.S()
)

func Init(rootCmd *cobra.Command) {
	dateCmd := &cobra.Command{
		Use:   "date",
		Short: "时间工具",
		RunE:  runDate,
	}
	rootCmd.AddCommand(dateCmd)

	dateCmd.Flags().
		StringVarP(&timeStr, "time-string", "t", "",
			"输入时间字符串，解析使用的时间格式可以通过 --layout 指定")

	dateCmd.Flags().
		Int64Var(&timestamp, "ts", 0,
			"输入时间戳，解析使用的时间单位可以通过 --duration 指定")

	dateCmd.Flags().
		DurationVarP(&duration, "duration", "d", time.Millisecond, "时间单位")
	dateCmd.Flags().
		StringVar(&output, "output", timeOutputTypeStringer, "时间输出格式")
	dateCmd.Flags().
		StringVarP(&layout, "layout", "l", time.RFC3339Nano, "时间格式化")
	dateCmd.Flags().
		StringVar(&location, "loc", time.Local.String(), "时区")

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
		logger.Infoln(t.Format(layout))
	case timeOutputTypeTimestamp:
		ts, err := formatTimestamp(t, duration)
		if err != nil {
			return err
		}
		logger.Infoln(ts)
	case timeOutputTypeGoString:
		logger.Infoln(t.GoString())
	case timeOutputTypeStringer:
		logger.Infoln(t)
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
		// 获取超过秒为单位的时间场景太少见了
		err = fmt.Errorf("暂不支持时间单位 '%s'", unit.String())
	}
	return
}

func parseLocation(name string) (loc *time.Location, err error) {
	switch name {
	case "":
		// 默认使用本地时区，而不是 UTC
		loc = time.Local
	case "Local":
		loc = time.Local
	case "UTC", "utc":
		loc = time.UTC
	default:
		loc, err = time.LoadLocation(name)
	}
	return
}
