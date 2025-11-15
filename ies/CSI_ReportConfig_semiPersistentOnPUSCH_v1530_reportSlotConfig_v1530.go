package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_ReportConfig_semiPersistentOnPUSCH_v1530_reportSlotConfig_v1530_Enum_sl4  uper.Enumerated = 0
	CSI_ReportConfig_semiPersistentOnPUSCH_v1530_reportSlotConfig_v1530_Enum_sl8  uper.Enumerated = 1
	CSI_ReportConfig_semiPersistentOnPUSCH_v1530_reportSlotConfig_v1530_Enum_sl16 uper.Enumerated = 2
)

type CSI_ReportConfig_semiPersistentOnPUSCH_v1530_reportSlotConfig_v1530 struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *CSI_ReportConfig_semiPersistentOnPUSCH_v1530_reportSlotConfig_v1530) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode CSI_ReportConfig_semiPersistentOnPUSCH_v1530_reportSlotConfig_v1530", err)
	}
	return nil
}

func (ie *CSI_ReportConfig_semiPersistentOnPUSCH_v1530_reportSlotConfig_v1530) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode CSI_ReportConfig_semiPersistentOnPUSCH_v1530_reportSlotConfig_v1530", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
