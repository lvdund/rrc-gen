package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CG_ConfigInfo_v1610_IEs_alignedDRX_Indication_Enum_true uper.Enumerated = 0
)

type CG_ConfigInfo_v1610_IEs_alignedDRX_Indication struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *CG_ConfigInfo_v1610_IEs_alignedDRX_Indication) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode CG_ConfigInfo_v1610_IEs_alignedDRX_Indication", err)
	}
	return nil
}

func (ie *CG_ConfigInfo_v1610_IEs_alignedDRX_Indication) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode CG_ConfigInfo_v1610_IEs_alignedDRX_Indication", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
