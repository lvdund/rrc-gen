package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_ReportConfig_groupBasedBeamReporting_disabled_nrofReportedRS_Enum_n1 uper.Enumerated = 0
	CSI_ReportConfig_groupBasedBeamReporting_disabled_nrofReportedRS_Enum_n2 uper.Enumerated = 1
	CSI_ReportConfig_groupBasedBeamReporting_disabled_nrofReportedRS_Enum_n3 uper.Enumerated = 2
	CSI_ReportConfig_groupBasedBeamReporting_disabled_nrofReportedRS_Enum_n4 uper.Enumerated = 3
)

type CSI_ReportConfig_groupBasedBeamReporting_disabled_nrofReportedRS struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *CSI_ReportConfig_groupBasedBeamReporting_disabled_nrofReportedRS) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode CSI_ReportConfig_groupBasedBeamReporting_disabled_nrofReportedRS", err)
	}
	return nil
}

func (ie *CSI_ReportConfig_groupBasedBeamReporting_disabled_nrofReportedRS) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode CSI_ReportConfig_groupBasedBeamReporting_disabled_nrofReportedRS", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
