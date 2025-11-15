package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_ReportConfig_groupBasedBeamReporting_disabled struct {
	nrofReportedRS *CSI_ReportConfig_groupBasedBeamReporting_disabled_nrofReportedRS `optional`
}

func (ie *CSI_ReportConfig_groupBasedBeamReporting_disabled) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.nrofReportedRS != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.nrofReportedRS != nil {
		if err = ie.nrofReportedRS.Encode(w); err != nil {
			return utils.WrapError("Encode nrofReportedRS", err)
		}
	}
	return nil
}

func (ie *CSI_ReportConfig_groupBasedBeamReporting_disabled) Decode(r *uper.UperReader) error {
	var err error
	var nrofReportedRSPresent bool
	if nrofReportedRSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if nrofReportedRSPresent {
		ie.nrofReportedRS = new(CSI_ReportConfig_groupBasedBeamReporting_disabled_nrofReportedRS)
		if err = ie.nrofReportedRS.Decode(r); err != nil {
			return utils.WrapError("Decode nrofReportedRS", err)
		}
	}
	return nil
}
