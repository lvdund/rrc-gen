package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_ResourceExt_v1610_format_v1610_occ_v1610_occ_Index_v1610_Enum_n0 uper.Enumerated = 0
	PUCCH_ResourceExt_v1610_format_v1610_occ_v1610_occ_Index_v1610_Enum_n1 uper.Enumerated = 1
	PUCCH_ResourceExt_v1610_format_v1610_occ_v1610_occ_Index_v1610_Enum_n2 uper.Enumerated = 2
	PUCCH_ResourceExt_v1610_format_v1610_occ_v1610_occ_Index_v1610_Enum_n3 uper.Enumerated = 3
)

type PUCCH_ResourceExt_v1610_format_v1610_occ_v1610_occ_Index_v1610 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *PUCCH_ResourceExt_v1610_format_v1610_occ_v1610_occ_Index_v1610) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode PUCCH_ResourceExt_v1610_format_v1610_occ_v1610_occ_Index_v1610", err)
	}
	return nil
}

func (ie *PUCCH_ResourceExt_v1610_format_v1610_occ_v1610_occ_Index_v1610) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode PUCCH_ResourceExt_v1610_format_v1610_occ_v1610_occ_Index_v1610", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
