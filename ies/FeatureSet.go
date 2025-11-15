package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSet_Choice_nothing uint64 = iota
	FeatureSet_Choice_eutra
	FeatureSet_Choice_nr
)

type FeatureSet struct {
	Choice uint64
	eutra  *FeatureSet_eutra
	nr     *FeatureSet_nr
}

func (ie *FeatureSet) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case FeatureSet_Choice_eutra:
		if err = ie.eutra.Encode(w); err != nil {
			err = utils.WrapError("Encode eutra", err)
		}
	case FeatureSet_Choice_nr:
		if err = ie.nr.Encode(w); err != nil {
			err = utils.WrapError("Encode nr", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *FeatureSet) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case FeatureSet_Choice_eutra:
		ie.eutra = new(FeatureSet_eutra)
		if err = ie.eutra.Decode(r); err != nil {
			return utils.WrapError("Decode eutra", err)
		}
	case FeatureSet_Choice_nr:
		ie.nr = new(FeatureSet_nr)
		if err = ie.nr.Decode(r); err != nil {
			return utils.WrapError("Decode nr", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
