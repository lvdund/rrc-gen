package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ServingCellAndBWP_Id_r17 struct {
	servingcell_r17 ServCellIndex `madatory`
	bwp_r17         BWP_Id        `madatory`
}

func (ie *ServingCellAndBWP_Id_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.servingcell_r17.Encode(w); err != nil {
		return utils.WrapError("Encode servingcell_r17", err)
	}
	if err = ie.bwp_r17.Encode(w); err != nil {
		return utils.WrapError("Encode bwp_r17", err)
	}
	return nil
}

func (ie *ServingCellAndBWP_Id_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.servingcell_r17.Decode(r); err != nil {
		return utils.WrapError("Decode servingcell_r17", err)
	}
	if err = ie.bwp_r17.Decode(r); err != nil {
		return utils.WrapError("Decode bwp_r17", err)
	}
	return nil
}
