package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_ReportConfig_groupBasedBeamReporting_v1710 struct {
	nrofReportedGroups_r17 CSI_ReportConfig_groupBasedBeamReporting_v1710_nrofReportedGroups_r17 `madatory`
}

func (ie *CSI_ReportConfig_groupBasedBeamReporting_v1710) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.nrofReportedGroups_r17.Encode(w); err != nil {
		return utils.WrapError("Encode nrofReportedGroups_r17", err)
	}
	return nil
}

func (ie *CSI_ReportConfig_groupBasedBeamReporting_v1710) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.nrofReportedGroups_r17.Decode(r); err != nil {
		return utils.WrapError("Decode nrofReportedGroups_r17", err)
	}
	return nil
}
