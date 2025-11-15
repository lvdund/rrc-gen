package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17_Enum_n4      uper.Enumerated = 0
	PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17_Enum_n8      uper.Enumerated = 1
	PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17_Enum_n16     uper.Enumerated = 2
	PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17_Enum_n32     uper.Enumerated = 3
	PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17_Enum_n44     uper.Enumerated = 4
	PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17_Enum_n64     uper.Enumerated = 5
	PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17_Enum_n128    uper.Enumerated = 6
	PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17_Enum_n256    uper.Enumerated = 7
	PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17_Enum_n512    uper.Enumerated = 8
	PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17_Enum_nolimit uper.Enumerated = 9
)

type PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17 struct {
	Value uper.Enumerated `lb:0,ub:9,madatory`
}

func (ie *PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return utils.WrapError("Encode PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17", err)
	}
	return nil
}

func (ie *PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return utils.WrapError("Decode PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
