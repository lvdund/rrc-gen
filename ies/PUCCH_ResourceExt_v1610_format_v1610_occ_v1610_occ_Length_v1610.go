package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_ResourceExt_v1610_format_v1610_occ_v1610_occ_Length_v1610_Enum_n2 uper.Enumerated = 0
	PUCCH_ResourceExt_v1610_format_v1610_occ_v1610_occ_Length_v1610_Enum_n4 uper.Enumerated = 1
)

type PUCCH_ResourceExt_v1610_format_v1610_occ_v1610_occ_Length_v1610 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *PUCCH_ResourceExt_v1610_format_v1610_occ_v1610_occ_Length_v1610) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode PUCCH_ResourceExt_v1610_format_v1610_occ_v1610_occ_Length_v1610", err)
	}
	return nil
}

func (ie *PUCCH_ResourceExt_v1610_format_v1610_occ_v1610_occ_Length_v1610) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode PUCCH_ResourceExt_v1610_format_v1610_occ_v1610_occ_Length_v1610", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
