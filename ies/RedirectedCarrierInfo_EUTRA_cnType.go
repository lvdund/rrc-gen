package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RedirectedCarrierInfo_EUTRA_cnType_Enum_epc    uper.Enumerated = 0
	RedirectedCarrierInfo_EUTRA_cnType_Enum_fiveGC uper.Enumerated = 1
)

type RedirectedCarrierInfo_EUTRA_cnType struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *RedirectedCarrierInfo_EUTRA_cnType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode RedirectedCarrierInfo_EUTRA_cnType", err)
	}
	return nil
}

func (ie *RedirectedCarrierInfo_EUTRA_cnType) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode RedirectedCarrierInfo_EUTRA_cnType", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
