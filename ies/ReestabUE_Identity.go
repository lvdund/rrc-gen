package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReestabUE_Identity struct {
	c_RNTI     RNTI_Value `madatory`
	physCellId PhysCellId `madatory`
	shortMAC_I ShortMAC_I `madatory`
}

func (ie *ReestabUE_Identity) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.c_RNTI.Encode(w); err != nil {
		return utils.WrapError("Encode c_RNTI", err)
	}
	if err = ie.physCellId.Encode(w); err != nil {
		return utils.WrapError("Encode physCellId", err)
	}
	if err = ie.shortMAC_I.Encode(w); err != nil {
		return utils.WrapError("Encode shortMAC_I", err)
	}
	return nil
}

func (ie *ReestabUE_Identity) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.c_RNTI.Decode(r); err != nil {
		return utils.WrapError("Decode c_RNTI", err)
	}
	if err = ie.physCellId.Decode(r); err != nil {
		return utils.WrapError("Decode physCellId", err)
	}
	if err = ie.shortMAC_I.Decode(r); err != nil {
		return utils.WrapError("Decode shortMAC_I", err)
	}
	return nil
}
