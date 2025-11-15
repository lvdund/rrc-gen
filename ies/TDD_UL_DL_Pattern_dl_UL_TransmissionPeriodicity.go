package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_Enum_ms0p5   uper.Enumerated = 0
	TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_Enum_ms0p625 uper.Enumerated = 1
	TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_Enum_ms1     uper.Enumerated = 2
	TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_Enum_ms1p25  uper.Enumerated = 3
	TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_Enum_ms2     uper.Enumerated = 4
	TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_Enum_ms2p5   uper.Enumerated = 5
	TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_Enum_ms5     uper.Enumerated = 6
	TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_Enum_ms10    uper.Enumerated = 7
)

type TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity", err)
	}
	return nil
}

func (ie *TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
