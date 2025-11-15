package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UCI_OnPUSCH struct {
	betaOffsets *UCI_OnPUSCH_betaOffsets `lb:4,ub:4,optional`
	scaling     UCI_OnPUSCH_scaling      `madatory`
}

func (ie *UCI_OnPUSCH) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.betaOffsets != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.betaOffsets != nil {
		if err = ie.betaOffsets.Encode(w); err != nil {
			return utils.WrapError("Encode betaOffsets", err)
		}
	}
	if err = ie.scaling.Encode(w); err != nil {
		return utils.WrapError("Encode scaling", err)
	}
	return nil
}

func (ie *UCI_OnPUSCH) Decode(r *uper.UperReader) error {
	var err error
	var betaOffsetsPresent bool
	if betaOffsetsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if betaOffsetsPresent {
		ie.betaOffsets = new(UCI_OnPUSCH_betaOffsets)
		if err = ie.betaOffsets.Decode(r); err != nil {
			return utils.WrapError("Decode betaOffsets", err)
		}
	}
	if err = ie.scaling.Decode(r); err != nil {
		return utils.WrapError("Decode scaling", err)
	}
	return nil
}
