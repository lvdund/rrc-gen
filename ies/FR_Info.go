package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FR_Info struct {
	servCellIndex ServCellIndex   `madatory`
	fr_Type       FR_Info_fr_Type `madatory`
}

func (ie *FR_Info) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.servCellIndex.Encode(w); err != nil {
		return utils.WrapError("Encode servCellIndex", err)
	}
	if err = ie.fr_Type.Encode(w); err != nil {
		return utils.WrapError("Encode fr_Type", err)
	}
	return nil
}

func (ie *FR_Info) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.servCellIndex.Decode(r); err != nil {
		return utils.WrapError("Decode servCellIndex", err)
	}
	if err = ie.fr_Type.Decode(r); err != nil {
		return utils.WrapError("Decode fr_Type", err)
	}
	return nil
}
