package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ConfiguredGrantConfig_repK_v1710_Enum_n12 uper.Enumerated = 0
	ConfiguredGrantConfig_repK_v1710_Enum_n16 uper.Enumerated = 1
	ConfiguredGrantConfig_repK_v1710_Enum_n24 uper.Enumerated = 2
	ConfiguredGrantConfig_repK_v1710_Enum_n32 uper.Enumerated = 3
)

type ConfiguredGrantConfig_repK_v1710 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *ConfiguredGrantConfig_repK_v1710) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode ConfiguredGrantConfig_repK_v1710", err)
	}
	return nil
}

func (ie *ConfiguredGrantConfig_repK_v1710) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode ConfiguredGrantConfig_repK_v1710", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
