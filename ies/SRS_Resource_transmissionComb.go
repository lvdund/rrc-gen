package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_Resource_transmissionComb_Choice_nothing uint64 = iota
	SRS_Resource_transmissionComb_Choice_n2
	SRS_Resource_transmissionComb_Choice_n4
)

type SRS_Resource_transmissionComb struct {
	Choice uint64
	n2     *SRS_Resource_transmissionComb_n2
	n4     *SRS_Resource_transmissionComb_n4
}

func (ie *SRS_Resource_transmissionComb) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_Resource_transmissionComb_Choice_n2:
		if err = ie.n2.Encode(w); err != nil {
			err = utils.WrapError("Encode n2", err)
		}
	case SRS_Resource_transmissionComb_Choice_n4:
		if err = ie.n4.Encode(w); err != nil {
			err = utils.WrapError("Encode n4", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SRS_Resource_transmissionComb) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_Resource_transmissionComb_Choice_n2:
		ie.n2 = new(SRS_Resource_transmissionComb_n2)
		if err = ie.n2.Decode(r); err != nil {
			return utils.WrapError("Decode n2", err)
		}
	case SRS_Resource_transmissionComb_Choice_n4:
		ie.n4 = new(SRS_Resource_transmissionComb_n4)
		if err = ie.n4.Decode(r); err != nil {
			return utils.WrapError("Decode n4", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
