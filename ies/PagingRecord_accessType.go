package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PagingRecord_accessType_Enum_non3GPP uper.Enumerated = 0
)

type PagingRecord_accessType struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *PagingRecord_accessType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode PagingRecord_accessType", err)
	}
	return nil
}

func (ie *PagingRecord_accessType) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode PagingRecord_accessType", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
