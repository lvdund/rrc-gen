package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCCH_ConfigSIB1 struct {
	controlResourceSetZero ControlResourceSetZero `madatory`
	searchSpaceZero        SearchSpaceZero        `madatory`
}

func (ie *PDCCH_ConfigSIB1) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.controlResourceSetZero.Encode(w); err != nil {
		return utils.WrapError("Encode controlResourceSetZero", err)
	}
	if err = ie.searchSpaceZero.Encode(w); err != nil {
		return utils.WrapError("Encode searchSpaceZero", err)
	}
	return nil
}

func (ie *PDCCH_ConfigSIB1) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.controlResourceSetZero.Decode(r); err != nil {
		return utils.WrapError("Decode controlResourceSetZero", err)
	}
	if err = ie.searchSpaceZero.Decode(r); err != nil {
		return utils.WrapError("Decode searchSpaceZero", err)
	}
	return nil
}
