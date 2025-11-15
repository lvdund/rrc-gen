package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDCCH_RepetitionParameters_r17_supportedMode_r17_Enum_intra_span uper.Enumerated = 0
	PDCCH_RepetitionParameters_r17_supportedMode_r17_Enum_inter_span uper.Enumerated = 1
	PDCCH_RepetitionParameters_r17_supportedMode_r17_Enum_both       uper.Enumerated = 2
)

type PDCCH_RepetitionParameters_r17_supportedMode_r17 struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *PDCCH_RepetitionParameters_r17_supportedMode_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode PDCCH_RepetitionParameters_r17_supportedMode_r17", err)
	}
	return nil
}

func (ie *PDCCH_RepetitionParameters_r17_supportedMode_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode PDCCH_RepetitionParameters_r17_supportedMode_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
