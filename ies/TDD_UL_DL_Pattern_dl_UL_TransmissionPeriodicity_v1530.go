package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_v1530_Enum_ms3 uper.Enumerated = 0
	TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_v1530_Enum_ms4 uper.Enumerated = 1
)

type TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_v1530 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_v1530) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_v1530", err)
	}
	return nil
}

func (ie *TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_v1530) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_v1530", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
