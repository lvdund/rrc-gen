package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RA_Prioritization struct {
	powerRampingStepHighPriority RA_Prioritization_powerRampingStepHighPriority `madatory`
	scalingFactorBI              *RA_Prioritization_scalingFactorBI             `optional`
}

func (ie *RA_Prioritization) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.scalingFactorBI != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.powerRampingStepHighPriority.Encode(w); err != nil {
		return utils.WrapError("Encode powerRampingStepHighPriority", err)
	}
	if ie.scalingFactorBI != nil {
		if err = ie.scalingFactorBI.Encode(w); err != nil {
			return utils.WrapError("Encode scalingFactorBI", err)
		}
	}
	return nil
}

func (ie *RA_Prioritization) Decode(r *uper.UperReader) error {
	var err error
	var scalingFactorBIPresent bool
	if scalingFactorBIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.powerRampingStepHighPriority.Decode(r); err != nil {
		return utils.WrapError("Decode powerRampingStepHighPriority", err)
	}
	if scalingFactorBIPresent {
		ie.scalingFactorBI = new(RA_Prioritization_scalingFactorBI)
		if err = ie.scalingFactorBI.Decode(r); err != nil {
			return utils.WrapError("Decode scalingFactorBI", err)
		}
	}
	return nil
}
