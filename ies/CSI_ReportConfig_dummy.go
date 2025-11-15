package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_ReportConfig_dummy_Enum_n1 uper.Enumerated = 0
	CSI_ReportConfig_dummy_Enum_n2 uper.Enumerated = 1
)

type CSI_ReportConfig_dummy struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *CSI_ReportConfig_dummy) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode CSI_ReportConfig_dummy", err)
	}
	return nil
}

func (ie *CSI_ReportConfig_dummy) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode CSI_ReportConfig_dummy", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
