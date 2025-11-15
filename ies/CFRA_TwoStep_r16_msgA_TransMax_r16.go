package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CFRA_TwoStep_r16_msgA_TransMax_r16_Enum_n1   uper.Enumerated = 0
	CFRA_TwoStep_r16_msgA_TransMax_r16_Enum_n2   uper.Enumerated = 1
	CFRA_TwoStep_r16_msgA_TransMax_r16_Enum_n4   uper.Enumerated = 2
	CFRA_TwoStep_r16_msgA_TransMax_r16_Enum_n6   uper.Enumerated = 3
	CFRA_TwoStep_r16_msgA_TransMax_r16_Enum_n8   uper.Enumerated = 4
	CFRA_TwoStep_r16_msgA_TransMax_r16_Enum_n10  uper.Enumerated = 5
	CFRA_TwoStep_r16_msgA_TransMax_r16_Enum_n20  uper.Enumerated = 6
	CFRA_TwoStep_r16_msgA_TransMax_r16_Enum_n50  uper.Enumerated = 7
	CFRA_TwoStep_r16_msgA_TransMax_r16_Enum_n100 uper.Enumerated = 8
	CFRA_TwoStep_r16_msgA_TransMax_r16_Enum_n200 uper.Enumerated = 9
)

type CFRA_TwoStep_r16_msgA_TransMax_r16 struct {
	Value uper.Enumerated `lb:0,ub:9,madatory`
}

func (ie *CFRA_TwoStep_r16_msgA_TransMax_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return utils.WrapError("Encode CFRA_TwoStep_r16_msgA_TransMax_r16", err)
	}
	return nil
}

func (ie *CFRA_TwoStep_r16_msgA_TransMax_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return utils.WrapError("Decode CFRA_TwoStep_r16_msgA_TransMax_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
