package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_ResourceExt_v1610_format_v1610_occ_v1610 struct {
	occ_Length_v1610 *PUCCH_ResourceExt_v1610_format_v1610_occ_v1610_occ_Length_v1610 `optional`
	occ_Index_v1610  *PUCCH_ResourceExt_v1610_format_v1610_occ_v1610_occ_Index_v1610  `optional`
}

func (ie *PUCCH_ResourceExt_v1610_format_v1610_occ_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.occ_Length_v1610 != nil, ie.occ_Index_v1610 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.occ_Length_v1610 != nil {
		if err = ie.occ_Length_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode occ_Length_v1610", err)
		}
	}
	if ie.occ_Index_v1610 != nil {
		if err = ie.occ_Index_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode occ_Index_v1610", err)
		}
	}
	return nil
}

func (ie *PUCCH_ResourceExt_v1610_format_v1610_occ_v1610) Decode(r *uper.UperReader) error {
	var err error
	var occ_Length_v1610Present bool
	if occ_Length_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var occ_Index_v1610Present bool
	if occ_Index_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	if occ_Length_v1610Present {
		ie.occ_Length_v1610 = new(PUCCH_ResourceExt_v1610_format_v1610_occ_v1610_occ_Length_v1610)
		if err = ie.occ_Length_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode occ_Length_v1610", err)
		}
	}
	if occ_Index_v1610Present {
		ie.occ_Index_v1610 = new(PUCCH_ResourceExt_v1610_format_v1610_occ_v1610_occ_Index_v1610)
		if err = ie.occ_Index_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode occ_Index_v1610", err)
		}
	}
	return nil
}
