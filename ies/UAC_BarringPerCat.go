package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UAC_BarringPerCat struct {
	accessCategory          int64                   `lb:1,ub:maxAccessCat_1,madatory`
	uac_barringInfoSetIndex UAC_BarringInfoSetIndex `madatory`
}

func (ie *UAC_BarringPerCat) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.accessCategory, &uper.Constraint{Lb: 1, Ub: maxAccessCat_1}, false); err != nil {
		return utils.WrapError("WriteInteger accessCategory", err)
	}
	if err = ie.uac_barringInfoSetIndex.Encode(w); err != nil {
		return utils.WrapError("Encode uac_barringInfoSetIndex", err)
	}
	return nil
}

func (ie *UAC_BarringPerCat) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_accessCategory int64
	if tmp_int_accessCategory, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxAccessCat_1}, false); err != nil {
		return utils.WrapError("ReadInteger accessCategory", err)
	}
	ie.accessCategory = tmp_int_accessCategory
	if err = ie.uac_barringInfoSetIndex.Decode(r); err != nil {
		return utils.WrapError("Decode uac_barringInfoSetIndex", err)
	}
	return nil
}
