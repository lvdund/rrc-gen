package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_ms0dot5 uper.Enumerated = 0
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_ms1     uper.Enumerated = 1
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_ms2     uper.Enumerated = 2
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_ms3     uper.Enumerated = 3
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_ms4     uper.Enumerated = 4
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_ms5     uper.Enumerated = 5
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_ms6     uper.Enumerated = 6
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_ms7     uper.Enumerated = 7
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_ms8     uper.Enumerated = 8
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_ms9     uper.Enumerated = 9
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_ms10    uper.Enumerated = 10
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_spare5  uper.Enumerated = 11
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_spare4  uper.Enumerated = 12
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_spare3  uper.Enumerated = 13
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_spare2  uper.Enumerated = 14
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_spare1  uper.Enumerated = 15
)

type PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17 struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17", err)
	}
	return nil
}

func (ie *PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
