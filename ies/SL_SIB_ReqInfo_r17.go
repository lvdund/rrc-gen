package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_SIB_ReqInfo_r17_Enum_sib1    uper.Enumerated = 0
	SL_SIB_ReqInfo_r17_Enum_sib2    uper.Enumerated = 1
	SL_SIB_ReqInfo_r17_Enum_sib3    uper.Enumerated = 2
	SL_SIB_ReqInfo_r17_Enum_sib4    uper.Enumerated = 3
	SL_SIB_ReqInfo_r17_Enum_sib5    uper.Enumerated = 4
	SL_SIB_ReqInfo_r17_Enum_sib6    uper.Enumerated = 5
	SL_SIB_ReqInfo_r17_Enum_sib7    uper.Enumerated = 6
	SL_SIB_ReqInfo_r17_Enum_sib8    uper.Enumerated = 7
	SL_SIB_ReqInfo_r17_Enum_sib9    uper.Enumerated = 8
	SL_SIB_ReqInfo_r17_Enum_sib10   uper.Enumerated = 9
	SL_SIB_ReqInfo_r17_Enum_sib11   uper.Enumerated = 10
	SL_SIB_ReqInfo_r17_Enum_sib12   uper.Enumerated = 11
	SL_SIB_ReqInfo_r17_Enum_sib13   uper.Enumerated = 12
	SL_SIB_ReqInfo_r17_Enum_sib14   uper.Enumerated = 13
	SL_SIB_ReqInfo_r17_Enum_sib15   uper.Enumerated = 14
	SL_SIB_ReqInfo_r17_Enum_sib16   uper.Enumerated = 15
	SL_SIB_ReqInfo_r17_Enum_sib17   uper.Enumerated = 16
	SL_SIB_ReqInfo_r17_Enum_sib18   uper.Enumerated = 17
	SL_SIB_ReqInfo_r17_Enum_sib19   uper.Enumerated = 18
	SL_SIB_ReqInfo_r17_Enum_sib20   uper.Enumerated = 19
	SL_SIB_ReqInfo_r17_Enum_sib21   uper.Enumerated = 20
	SL_SIB_ReqInfo_r17_Enum_spare11 uper.Enumerated = 21
	SL_SIB_ReqInfo_r17_Enum_spare10 uper.Enumerated = 22
	SL_SIB_ReqInfo_r17_Enum_spare9  uper.Enumerated = 23
	SL_SIB_ReqInfo_r17_Enum_spare8  uper.Enumerated = 24
	SL_SIB_ReqInfo_r17_Enum_spare7  uper.Enumerated = 25
	SL_SIB_ReqInfo_r17_Enum_spare6  uper.Enumerated = 26
	SL_SIB_ReqInfo_r17_Enum_spare5  uper.Enumerated = 27
	SL_SIB_ReqInfo_r17_Enum_spare4  uper.Enumerated = 28
	SL_SIB_ReqInfo_r17_Enum_spare3  uper.Enumerated = 29
	SL_SIB_ReqInfo_r17_Enum_spare2  uper.Enumerated = 30
	SL_SIB_ReqInfo_r17_Enum_spare1  uper.Enumerated = 31
)

type SL_SIB_ReqInfo_r17 struct {
	Value uper.Enumerated `lb:0,ub:31,madatory`
}

func (ie *SL_SIB_ReqInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Encode SL_SIB_ReqInfo_r17", err)
	}
	return nil
}

func (ie *SL_SIB_ReqInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Decode SL_SIB_ReqInfo_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
