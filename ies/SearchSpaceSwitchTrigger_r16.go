package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpaceSwitchTrigger_r16 struct {
	servingCellId_r16 ServCellIndex `madatory`
	positionInDCI_r16 int64         `lb:0,ub:maxSFI_DCI_PayloadSize_1,madatory`
}

func (ie *SearchSpaceSwitchTrigger_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.servingCellId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode servingCellId_r16", err)
	}
	if err = w.WriteInteger(ie.positionInDCI_r16, &uper.Constraint{Lb: 0, Ub: maxSFI_DCI_PayloadSize_1}, false); err != nil {
		return utils.WrapError("WriteInteger positionInDCI_r16", err)
	}
	return nil
}

func (ie *SearchSpaceSwitchTrigger_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.servingCellId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode servingCellId_r16", err)
	}
	var tmp_int_positionInDCI_r16 int64
	if tmp_int_positionInDCI_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxSFI_DCI_PayloadSize_1}, false); err != nil {
		return utils.WrapError("ReadInteger positionInDCI_r16", err)
	}
	ie.positionInDCI_r16 = tmp_int_positionInDCI_r16
	return nil
}
