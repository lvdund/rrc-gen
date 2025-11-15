package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_ReportConfig_cqi_BitsPerSubband_r17_Enum_bits4 uper.Enumerated = 0
)

type CSI_ReportConfig_cqi_BitsPerSubband_r17 struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *CSI_ReportConfig_cqi_BitsPerSubband_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode CSI_ReportConfig_cqi_BitsPerSubband_r17", err)
	}
	return nil
}

func (ie *CSI_ReportConfig_cqi_BitsPerSubband_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode CSI_ReportConfig_cqi_BitsPerSubband_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
