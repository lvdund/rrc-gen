package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PCI_ARFCN_NR_r16 struct {
	physCellId_r16  PhysCellId    `madatory`
	carrierFreq_r16 ARFCN_ValueNR `madatory`
}

func (ie *PCI_ARFCN_NR_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.physCellId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode physCellId_r16", err)
	}
	if err = ie.carrierFreq_r16.Encode(w); err != nil {
		return utils.WrapError("Encode carrierFreq_r16", err)
	}
	return nil
}

func (ie *PCI_ARFCN_NR_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.physCellId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode physCellId_r16", err)
	}
	if err = ie.carrierFreq_r16.Decode(r); err != nil {
		return utils.WrapError("Decode carrierFreq_r16", err)
	}
	return nil
}
