package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNRDC_v1650 struct {
	supportedCellGrouping_r16 *uper.BitString `lb:1,ub:maxCellGroupings_r16,optional`
}

func (ie *CA_ParametersNRDC_v1650) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.supportedCellGrouping_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.supportedCellGrouping_r16 != nil {
		if err = w.WriteBitString(ie.supportedCellGrouping_r16.Bytes, uint(ie.supportedCellGrouping_r16.NumBits), &uper.Constraint{Lb: 1, Ub: maxCellGroupings_r16}, false); err != nil {
			return utils.WrapError("Encode supportedCellGrouping_r16", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNRDC_v1650) Decode(r *uper.UperReader) error {
	var err error
	var supportedCellGrouping_r16Present bool
	if supportedCellGrouping_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if supportedCellGrouping_r16Present {
		var tmp_bs_supportedCellGrouping_r16 []byte
		var n_supportedCellGrouping_r16 uint
		if tmp_bs_supportedCellGrouping_r16, n_supportedCellGrouping_r16, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: maxCellGroupings_r16}, false); err != nil {
			return utils.WrapError("Decode supportedCellGrouping_r16", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_supportedCellGrouping_r16,
			NumBits: uint64(n_supportedCellGrouping_r16),
		}
		ie.supportedCellGrouping_r16 = &tmp_bitstring
	}
	return nil
}
