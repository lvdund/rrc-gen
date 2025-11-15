package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PHR_Config_phr_ModeOtherCG_Enum_real    uper.Enumerated = 0
	PHR_Config_phr_ModeOtherCG_Enum_virtual uper.Enumerated = 1
)

type PHR_Config_phr_ModeOtherCG struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *PHR_Config_phr_ModeOtherCG) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode PHR_Config_phr_ModeOtherCG", err)
	}
	return nil
}

func (ie *PHR_Config_phr_ModeOtherCG) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode PHR_Config_phr_ModeOtherCG", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
