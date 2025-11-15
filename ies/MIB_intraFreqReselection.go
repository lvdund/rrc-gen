package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MIB_intraFreqReselection_Enum_allowed    uper.Enumerated = 0
	MIB_intraFreqReselection_Enum_notAllowed uper.Enumerated = 1
)

type MIB_intraFreqReselection struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *MIB_intraFreqReselection) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode MIB_intraFreqReselection", err)
	}
	return nil
}

func (ie *MIB_intraFreqReselection) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode MIB_intraFreqReselection", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
