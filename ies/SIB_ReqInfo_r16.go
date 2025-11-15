package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SIB_ReqInfo_r16_Enum_sib12       uper.Enumerated = 0
	SIB_ReqInfo_r16_Enum_sib13       uper.Enumerated = 1
	SIB_ReqInfo_r16_Enum_sib14       uper.Enumerated = 2
	SIB_ReqInfo_r16_Enum_sib20_v1700 uper.Enumerated = 3
	SIB_ReqInfo_r16_Enum_sib21_v1700 uper.Enumerated = 4
	SIB_ReqInfo_r16_Enum_spare3      uper.Enumerated = 5
	SIB_ReqInfo_r16_Enum_spare2      uper.Enumerated = 6
	SIB_ReqInfo_r16_Enum_spare1      uper.Enumerated = 7
)

type SIB_ReqInfo_r16 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SIB_ReqInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SIB_ReqInfo_r16", err)
	}
	return nil
}

func (ie *SIB_ReqInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SIB_ReqInfo_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
