package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UCI_OnPUSCH_DCI_0_2_r16 struct {
	betaOffsetsDCI_0_2_r16 *UCI_OnPUSCH_DCI_0_2_r16_betaOffsetsDCI_0_2_r16 `lb:2,ub:2,optional`
	scalingDCI_0_2_r16     UCI_OnPUSCH_DCI_0_2_r16_scalingDCI_0_2_r16      `madatory`
}

func (ie *UCI_OnPUSCH_DCI_0_2_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.betaOffsetsDCI_0_2_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.betaOffsetsDCI_0_2_r16 != nil {
		if err = ie.betaOffsetsDCI_0_2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode betaOffsetsDCI_0_2_r16", err)
		}
	}
	if err = ie.scalingDCI_0_2_r16.Encode(w); err != nil {
		return utils.WrapError("Encode scalingDCI_0_2_r16", err)
	}
	return nil
}

func (ie *UCI_OnPUSCH_DCI_0_2_r16) Decode(r *uper.UperReader) error {
	var err error
	var betaOffsetsDCI_0_2_r16Present bool
	if betaOffsetsDCI_0_2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if betaOffsetsDCI_0_2_r16Present {
		ie.betaOffsetsDCI_0_2_r16 = new(UCI_OnPUSCH_DCI_0_2_r16_betaOffsetsDCI_0_2_r16)
		if err = ie.betaOffsetsDCI_0_2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode betaOffsetsDCI_0_2_r16", err)
		}
	}
	if err = ie.scalingDCI_0_2_r16.Decode(r); err != nil {
		return utils.WrapError("Decode scalingDCI_0_2_r16", err)
	}
	return nil
}
