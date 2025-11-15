package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRM_Config_ue_InactiveTime_Enum_s1            uper.Enumerated = 0
	RRM_Config_ue_InactiveTime_Enum_s2            uper.Enumerated = 1
	RRM_Config_ue_InactiveTime_Enum_s3            uper.Enumerated = 2
	RRM_Config_ue_InactiveTime_Enum_s5            uper.Enumerated = 3
	RRM_Config_ue_InactiveTime_Enum_s7            uper.Enumerated = 4
	RRM_Config_ue_InactiveTime_Enum_s10           uper.Enumerated = 5
	RRM_Config_ue_InactiveTime_Enum_s15           uper.Enumerated = 6
	RRM_Config_ue_InactiveTime_Enum_s20           uper.Enumerated = 7
	RRM_Config_ue_InactiveTime_Enum_s25           uper.Enumerated = 8
	RRM_Config_ue_InactiveTime_Enum_s30           uper.Enumerated = 9
	RRM_Config_ue_InactiveTime_Enum_s40           uper.Enumerated = 10
	RRM_Config_ue_InactiveTime_Enum_s50           uper.Enumerated = 11
	RRM_Config_ue_InactiveTime_Enum_min1          uper.Enumerated = 12
	RRM_Config_ue_InactiveTime_Enum_min1s20       uper.Enumerated = 13
	RRM_Config_ue_InactiveTime_Enum_min1s40       uper.Enumerated = 14
	RRM_Config_ue_InactiveTime_Enum_min2          uper.Enumerated = 15
	RRM_Config_ue_InactiveTime_Enum_min2s30       uper.Enumerated = 16
	RRM_Config_ue_InactiveTime_Enum_min3          uper.Enumerated = 17
	RRM_Config_ue_InactiveTime_Enum_min3s30       uper.Enumerated = 18
	RRM_Config_ue_InactiveTime_Enum_min4          uper.Enumerated = 19
	RRM_Config_ue_InactiveTime_Enum_min5          uper.Enumerated = 20
	RRM_Config_ue_InactiveTime_Enum_min6          uper.Enumerated = 21
	RRM_Config_ue_InactiveTime_Enum_min7          uper.Enumerated = 22
	RRM_Config_ue_InactiveTime_Enum_min8          uper.Enumerated = 23
	RRM_Config_ue_InactiveTime_Enum_min9          uper.Enumerated = 24
	RRM_Config_ue_InactiveTime_Enum_min10         uper.Enumerated = 25
	RRM_Config_ue_InactiveTime_Enum_min12         uper.Enumerated = 26
	RRM_Config_ue_InactiveTime_Enum_min14         uper.Enumerated = 27
	RRM_Config_ue_InactiveTime_Enum_min17         uper.Enumerated = 28
	RRM_Config_ue_InactiveTime_Enum_min20         uper.Enumerated = 29
	RRM_Config_ue_InactiveTime_Enum_min24         uper.Enumerated = 30
	RRM_Config_ue_InactiveTime_Enum_min28         uper.Enumerated = 31
	RRM_Config_ue_InactiveTime_Enum_min33         uper.Enumerated = 32
	RRM_Config_ue_InactiveTime_Enum_min38         uper.Enumerated = 33
	RRM_Config_ue_InactiveTime_Enum_min44         uper.Enumerated = 34
	RRM_Config_ue_InactiveTime_Enum_min50         uper.Enumerated = 35
	RRM_Config_ue_InactiveTime_Enum_hr1           uper.Enumerated = 36
	RRM_Config_ue_InactiveTime_Enum_hr1min30      uper.Enumerated = 37
	RRM_Config_ue_InactiveTime_Enum_hr2           uper.Enumerated = 38
	RRM_Config_ue_InactiveTime_Enum_hr2min30      uper.Enumerated = 39
	RRM_Config_ue_InactiveTime_Enum_hr3           uper.Enumerated = 40
	RRM_Config_ue_InactiveTime_Enum_hr3min30      uper.Enumerated = 41
	RRM_Config_ue_InactiveTime_Enum_hr4           uper.Enumerated = 42
	RRM_Config_ue_InactiveTime_Enum_hr5           uper.Enumerated = 43
	RRM_Config_ue_InactiveTime_Enum_hr6           uper.Enumerated = 44
	RRM_Config_ue_InactiveTime_Enum_hr8           uper.Enumerated = 45
	RRM_Config_ue_InactiveTime_Enum_hr10          uper.Enumerated = 46
	RRM_Config_ue_InactiveTime_Enum_hr13          uper.Enumerated = 47
	RRM_Config_ue_InactiveTime_Enum_hr16          uper.Enumerated = 48
	RRM_Config_ue_InactiveTime_Enum_hr20          uper.Enumerated = 49
	RRM_Config_ue_InactiveTime_Enum_day1          uper.Enumerated = 50
	RRM_Config_ue_InactiveTime_Enum_day1hr12      uper.Enumerated = 51
	RRM_Config_ue_InactiveTime_Enum_day2          uper.Enumerated = 52
	RRM_Config_ue_InactiveTime_Enum_day2hr12      uper.Enumerated = 53
	RRM_Config_ue_InactiveTime_Enum_day3          uper.Enumerated = 54
	RRM_Config_ue_InactiveTime_Enum_day4          uper.Enumerated = 55
	RRM_Config_ue_InactiveTime_Enum_day5          uper.Enumerated = 56
	RRM_Config_ue_InactiveTime_Enum_day7          uper.Enumerated = 57
	RRM_Config_ue_InactiveTime_Enum_day10         uper.Enumerated = 58
	RRM_Config_ue_InactiveTime_Enum_day14         uper.Enumerated = 59
	RRM_Config_ue_InactiveTime_Enum_day19         uper.Enumerated = 60
	RRM_Config_ue_InactiveTime_Enum_day24         uper.Enumerated = 61
	RRM_Config_ue_InactiveTime_Enum_day30         uper.Enumerated = 62
	RRM_Config_ue_InactiveTime_Enum_dayMoreThan30 uper.Enumerated = 63
)

type RRM_Config_ue_InactiveTime struct {
	Value uper.Enumerated `lb:0,ub:63,madatory`
}

func (ie *RRM_Config_ue_InactiveTime) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("Encode RRM_Config_ue_InactiveTime", err)
	}
	return nil
}

func (ie *RRM_Config_ue_InactiveTime) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("Decode RRM_Config_ue_InactiveTime", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
