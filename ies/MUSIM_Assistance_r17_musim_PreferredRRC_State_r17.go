package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MUSIM_Assistance_r17_musim_PreferredRRC_State_r17_Enum_idle           uper.Enumerated = 0
	MUSIM_Assistance_r17_musim_PreferredRRC_State_r17_Enum_inactive       uper.Enumerated = 1
	MUSIM_Assistance_r17_musim_PreferredRRC_State_r17_Enum_outOfConnected uper.Enumerated = 2
)

type MUSIM_Assistance_r17_musim_PreferredRRC_State_r17 struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *MUSIM_Assistance_r17_musim_PreferredRRC_State_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode MUSIM_Assistance_r17_musim_PreferredRRC_State_r17", err)
	}
	return nil
}

func (ie *MUSIM_Assistance_r17_musim_PreferredRRC_State_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode MUSIM_Assistance_r17_musim_PreferredRRC_State_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
