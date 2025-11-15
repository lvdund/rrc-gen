package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_MaxCodeRate_Enum_zeroDot08 uper.Enumerated = 0
	PUCCH_MaxCodeRate_Enum_zeroDot15 uper.Enumerated = 1
	PUCCH_MaxCodeRate_Enum_zeroDot25 uper.Enumerated = 2
	PUCCH_MaxCodeRate_Enum_zeroDot35 uper.Enumerated = 3
	PUCCH_MaxCodeRate_Enum_zeroDot45 uper.Enumerated = 4
	PUCCH_MaxCodeRate_Enum_zeroDot60 uper.Enumerated = 5
	PUCCH_MaxCodeRate_Enum_zeroDot80 uper.Enumerated = 6
)

type PUCCH_MaxCodeRate struct {
	Value uper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *PUCCH_MaxCodeRate) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode PUCCH_MaxCodeRate", err)
	}
	return nil
}

func (ie *PUCCH_MaxCodeRate) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode PUCCH_MaxCodeRate", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
