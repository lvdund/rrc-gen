package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	VictimSystemType_wlan_Enum_true uper.Enumerated = 0
)

type VictimSystemType_wlan struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *VictimSystemType_wlan) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode VictimSystemType_wlan", err)
	}
	return nil
}

func (ie *VictimSystemType_wlan) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode VictimSystemType_wlan", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
