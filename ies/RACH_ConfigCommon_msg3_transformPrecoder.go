package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigCommon_msg3_transformPrecoder_Enum_enabled uper.Enumerated = 0
)

type RACH_ConfigCommon_msg3_transformPrecoder struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *RACH_ConfigCommon_msg3_transformPrecoder) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode RACH_ConfigCommon_msg3_transformPrecoder", err)
	}
	return nil
}

func (ie *RACH_ConfigCommon_msg3_transformPrecoder) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode RACH_ConfigCommon_msg3_transformPrecoder", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
