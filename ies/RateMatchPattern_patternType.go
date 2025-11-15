package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RateMatchPattern_patternType_Choice_nothing uint64 = iota
	RateMatchPattern_patternType_Choice_bitmaps
	RateMatchPattern_patternType_Choice_controlResourceSet
)

type RateMatchPattern_patternType struct {
	Choice             uint64
	bitmaps            *RateMatchPattern_patternType_bitmaps
	controlResourceSet *ControlResourceSetId
}

func (ie *RateMatchPattern_patternType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RateMatchPattern_patternType_Choice_bitmaps:
		if err = ie.bitmaps.Encode(w); err != nil {
			err = utils.WrapError("Encode bitmaps", err)
		}
	case RateMatchPattern_patternType_Choice_controlResourceSet:
		if err = ie.controlResourceSet.Encode(w); err != nil {
			err = utils.WrapError("Encode controlResourceSet", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RateMatchPattern_patternType) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RateMatchPattern_patternType_Choice_bitmaps:
		ie.bitmaps = new(RateMatchPattern_patternType_bitmaps)
		if err = ie.bitmaps.Decode(r); err != nil {
			return utils.WrapError("Decode bitmaps", err)
		}
	case RateMatchPattern_patternType_Choice_controlResourceSet:
		ie.controlResourceSet = new(ControlResourceSetId)
		if err = ie.controlResourceSet.Decode(r); err != nil {
			return utils.WrapError("Decode controlResourceSet", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
