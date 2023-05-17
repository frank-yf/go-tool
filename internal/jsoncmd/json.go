package jsoncmd

import (
	"strings"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/frank-yf/go-tool/internal/pkg/cjson"
	"github.com/frank-yf/go-tool/internal/pkg/tool"
)

func runJson(cmd *cobra.Command, args []string) error {
	if data != "" {
		var v interface{}
		if err := parseData(data, &v); err != nil {
			return err
		}
		return printValue(v)
	}

	if dataFile != "" {
		// todo 待实现
	}

	logger.Infoln("什么也没做")
	return nil
}

var (
	logger = zap.S()

	data      string
	dataFile  string
	pretty    bool
	useNumber bool
)

func Init(rootCmd *cobra.Command) {
	jsonCmd := &cobra.Command{
		Use:   "json",
		Short: "json工具",
		RunE:  runJson,
	}

	rootCmd.AddCommand(jsonCmd)

	jsonCmd.Flags().
		StringVarP(&data, "data", "d", "", "序列化数据")
	jsonCmd.Flags().
		StringVarP(&dataFile, "file", "f", "", "序列化数据文件")
	jsonCmd.Flags().
		BoolVar(&pretty, "pretty", true, "json美化")
	jsonCmd.Flags().
		BoolVar(&useNumber, "use-number", true, "序列化方式定义：使用 UseNumber")

}

func parseData(str string, v any) error {
	sr := strings.NewReader(str)
	var dec cjson.Decoder
	if useNumber {
		dec = cjson.NewDecoder(sr, cjson.UseNumber())
	} else {
		dec = cjson.NewDecoder(sr)
	}
	return dec.Decode(v)
}

func printValue(v any) error {
	var data string
	if pretty {
		bs, err := cjson.MarshalIndent(v, "", "  ")
		if err != nil {
			return err
		}
		data = tool.BytesToString(bs)
	} else {
		marshal, err := cjson.MarshalToString(v)
		if err != nil {
			return err
		}
		data = marshal
	}
	logger.Infoln(data)
	return nil
}
