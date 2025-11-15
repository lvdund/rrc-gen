package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ConfiguredGrantConfig_repK_RV_Enum_s1_0231 uper.Enumerated = 0
	ConfiguredGrantConfig_repK_RV_Enum_s2_0303 uper.Enumerated = 1
	ConfiguredGrantConfig_repK_RV_Enum_s3_0000 uper.Enumerated = 2
)

type ConfiguredGrantConfig_repK_RV struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *ConfiguredGrantConfig_repK_RV) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode ConfiguredGrantConfig_repK_RV", err)
	}
	return nil
}

func (ie *ConfiguredGrantConfig_repK_RV) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode ConfiguredGrantConfig_repK_RV", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
