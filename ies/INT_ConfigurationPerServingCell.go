package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type INT_ConfigurationPerServingCell struct {
	servingCellId ServCellIndex `madatory`
	positionInDCI int64         `lb:0,ub:maxINT_DCI_PayloadSize_1,madatory`
}

func (ie *INT_ConfigurationPerServingCell) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.servingCellId.Encode(w); err != nil {
		return utils.WrapError("Encode servingCellId", err)
	}
	if err = w.WriteInteger(ie.positionInDCI, &uper.Constraint{Lb: 0, Ub: maxINT_DCI_PayloadSize_1}, false); err != nil {
		return utils.WrapError("WriteInteger positionInDCI", err)
	}
	return nil
}

func (ie *INT_ConfigurationPerServingCell) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.servingCellId.Decode(r); err != nil {
		return utils.WrapError("Decode servingCellId", err)
	}
	var tmp_int_positionInDCI int64
	if tmp_int_positionInDCI, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxINT_DCI_PayloadSize_1}, false); err != nil {
		return utils.WrapError("ReadInteger positionInDCI", err)
	}
	ie.positionInDCI = tmp_int_positionInDCI
	return nil
}
