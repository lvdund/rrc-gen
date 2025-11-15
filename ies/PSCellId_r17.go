package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PSCellId_r17 struct {
	physCellId_r17  PhysCellId    `madatory`
	carrierFreq_r17 ARFCN_ValueNR `madatory`
}

func (ie *PSCellId_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.physCellId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode physCellId_r17", err)
	}
	if err = ie.carrierFreq_r17.Encode(w); err != nil {
		return utils.WrapError("Encode carrierFreq_r17", err)
	}
	return nil
}

func (ie *PSCellId_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.physCellId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode physCellId_r17", err)
	}
	if err = ie.carrierFreq_r17.Decode(r); err != nil {
		return utils.WrapError("Decode carrierFreq_r17", err)
	}
	return nil
}
