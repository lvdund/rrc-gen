package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RateMatchPatternGroupItem_Choice_nothing uint64 = iota
	RateMatchPatternGroupItem_Choice_cellLevel
	RateMatchPatternGroupItem_Choice_bwpLevel
)

type RateMatchPatternGroupItem struct {
	Choice    uint64
	cellLevel *RateMatchPatternId
	bwpLevel  *RateMatchPatternId
}

func (ie *RateMatchPatternGroupItem) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RateMatchPatternGroupItem_Choice_cellLevel:
		if err = ie.cellLevel.Encode(w); err != nil {
			err = utils.WrapError("Encode cellLevel", err)
		}
	case RateMatchPatternGroupItem_Choice_bwpLevel:
		if err = ie.bwpLevel.Encode(w); err != nil {
			err = utils.WrapError("Encode bwpLevel", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RateMatchPatternGroupItem) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RateMatchPatternGroupItem_Choice_cellLevel:
		ie.cellLevel = new(RateMatchPatternId)
		if err = ie.cellLevel.Decode(r); err != nil {
			return utils.WrapError("Decode cellLevel", err)
		}
	case RateMatchPatternGroupItem_Choice_bwpLevel:
		ie.bwpLevel = new(RateMatchPatternId)
		if err = ie.bwpLevel.Decode(r); err != nil {
			return utils.WrapError("Decode bwpLevel", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
