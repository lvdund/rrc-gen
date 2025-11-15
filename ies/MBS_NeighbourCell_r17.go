package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MBS_NeighbourCell_r17 struct {
	physCellId_r17  PhysCellId     `madatory`
	carrierFreq_r17 *ARFCN_ValueNR `optional`
}

func (ie *MBS_NeighbourCell_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.carrierFreq_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.physCellId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode physCellId_r17", err)
	}
	if ie.carrierFreq_r17 != nil {
		if err = ie.carrierFreq_r17.Encode(w); err != nil {
			return utils.WrapError("Encode carrierFreq_r17", err)
		}
	}
	return nil
}

func (ie *MBS_NeighbourCell_r17) Decode(r *uper.UperReader) error {
	var err error
	var carrierFreq_r17Present bool
	if carrierFreq_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.physCellId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode physCellId_r17", err)
	}
	if carrierFreq_r17Present {
		ie.carrierFreq_r17 = new(ARFCN_ValueNR)
		if err = ie.carrierFreq_r17.Decode(r); err != nil {
			return utils.WrapError("Decode carrierFreq_r17", err)
		}
	}
	return nil
}
