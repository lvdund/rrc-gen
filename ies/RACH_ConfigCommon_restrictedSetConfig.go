package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigCommon_restrictedSetConfig_Enum_unrestrictedSet    uper.Enumerated = 0
	RACH_ConfigCommon_restrictedSetConfig_Enum_restrictedSetTypeA uper.Enumerated = 1
	RACH_ConfigCommon_restrictedSetConfig_Enum_restrictedSetTypeB uper.Enumerated = 2
)

type RACH_ConfigCommon_restrictedSetConfig struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *RACH_ConfigCommon_restrictedSetConfig) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode RACH_ConfigCommon_restrictedSetConfig", err)
	}
	return nil
}

func (ie *RACH_ConfigCommon_restrictedSetConfig) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode RACH_ConfigCommon_restrictedSetConfig", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
