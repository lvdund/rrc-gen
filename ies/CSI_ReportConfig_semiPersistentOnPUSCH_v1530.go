package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_ReportConfig_semiPersistentOnPUSCH_v1530 struct {
	reportSlotConfig_v1530 CSI_ReportConfig_semiPersistentOnPUSCH_v1530_reportSlotConfig_v1530 `madatory`
}

func (ie *CSI_ReportConfig_semiPersistentOnPUSCH_v1530) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.reportSlotConfig_v1530.Encode(w); err != nil {
		return utils.WrapError("Encode reportSlotConfig_v1530", err)
	}
	return nil
}

func (ie *CSI_ReportConfig_semiPersistentOnPUSCH_v1530) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.reportSlotConfig_v1530.Decode(r); err != nil {
		return utils.WrapError("Decode reportSlotConfig_v1530", err)
	}
	return nil
}
