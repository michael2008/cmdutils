package cmdutils

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/exp/constraints"
	"os"
)

func DisableFlagSort(cmd *cobra.Command) {
	cmd.Flags().SortFlags = false
	cmd.PersistentFlags().SortFlags = false
}

func flagFail(cmd *cobra.Command, name string) {
	fmt.Printf("ERROR: flag '%s' is required\n", name)
	_ = cmd.Usage()
	os.Exit(1)
}

func MustGetFlagString(cmd *cobra.Command, name string) string {
	flag, err := cmd.Flags().GetString(name)
	if err != nil || flag == "" {
		flagFail(cmd, name)
	}
	return flag
}

func MustGetFlagBool(cmd *cobra.Command, name string) bool {
	flag, err := cmd.Flags().GetBool(name)
	if err != nil {
		flagFail(cmd, name)
	}
	return flag
}

func MustGetFlagInt(cmd *cobra.Command, name string) int {
	flag, err := cmd.Flags().GetInt(name)
	if err != nil {
		flagFail(cmd, name)
	}
	return flag
}

func MustGetFlagUint(cmd *cobra.Command, name string) uint {
	flag, err := cmd.Flags().GetUint(name)
	if err != nil {
		flagFail(cmd, name)
	}
	return flag
}

func MustGetFlagUint64(cmd *cobra.Command, name string) uint64 {
	flag, err := cmd.Flags().GetUint64(name)
	if err != nil {
		flagFail(cmd, name)
	}
	return flag
}

func MustGetFlagFloat64(cmd *cobra.Command, name string) float64 {
	flag, err := cmd.Flags().GetFloat64(name)
	if err != nil {
		flagFail(cmd, name)
	}
	return flag
}

func GetFlagStringWithDefault(cmd *cobra.Command, name, defaultVal string) string {
	flag, err := cmd.Flags().GetString(name)
	if err != nil || flag == "" {
		return defaultVal
	}
	return flag
}

func GetFlagStringEmpty(cmd *cobra.Command, name string) string {
	return GetFlagStringWithDefault(cmd, name, "")
}

type Number interface {
	constraints.Integer | constraints.Float
}

func FitInRange[T Number](v, low, high T) T {
	if v < low {
		return low
	}
	if v > high {
		return high
	}
	return v
}

func GetFlagIntWithDefault(cmd *cobra.Command, name string, defaultVal int) int {
	val, err := cmd.Flags().GetInt(name)
	if err != nil {
		return defaultVal
	}
	return val
}

func GetFlagUintWithDefault(cmd *cobra.Command, name string, defaultVal uint) uint {
	val, err := cmd.Flags().GetUint(name)
	if err != nil {
		return defaultVal
	}
	return val
}

func GetFlagUint64WithDefault(cmd *cobra.Command, name string, defaultVal uint64) uint64 {
	val, err := cmd.Flags().GetUint64(name)
	if err != nil {
		return defaultVal
	}
	return val
}

func GetFlagBoolWithDefault(cmd *cobra.Command, name string, defaultVal bool) bool {
	val, err := cmd.Flags().GetBool(name)
	if err != nil {
		return defaultVal
	}
	return val
}
