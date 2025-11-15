package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FailureInfoRLC_Bearer_failureType_Enum_rlc_failure uper.Enumerated = 0
	FailureInfoRLC_Bearer_failureType_Enum_spare3      uper.Enumerated = 1
	FailureInfoRLC_Bearer_failureType_Enum_spare2      uper.Enumerated = 2
	FailureInfoRLC_Bearer_failureType_Enum_spare1      uper.Enumerated = 3
)

type FailureInfoRLC_Bearer_failureType struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *FailureInfoRLC_Bearer_failureType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode FailureInfoRLC_Bearer_failureType", err)
	}
	return nil
}

func (ie *FailureInfoRLC_Bearer_failureType) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode FailureInfoRLC_Bearer_failureType", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
