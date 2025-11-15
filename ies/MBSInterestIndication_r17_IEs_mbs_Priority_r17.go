package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MBSInterestIndication_r17_IEs_mbs_Priority_r17_Enum_true uper.Enumerated = 0
)

type MBSInterestIndication_r17_IEs_mbs_Priority_r17 struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *MBSInterestIndication_r17_IEs_mbs_Priority_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode MBSInterestIndication_r17_IEs_mbs_Priority_r17", err)
	}
	return nil
}

func (ie *MBSInterestIndication_r17_IEs_mbs_Priority_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode MBSInterestIndication_r17_IEs_mbs_Priority_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
