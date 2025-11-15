package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	Resourcetype_r16_SRS_PosResource_r16_Choice_nothing uint64 = iota
	Resourcetype_r16_SRS_PosResource_r16_Choice_aperiodic_r16
	Resourcetype_r16_SRS_PosResource_r16_Choice_semi_persistent_r16
	Resourcetype_r16_SRS_PosResource_r16_Choice_periodic_r16
)

type Resourcetype_r16_SRS_PosResource_r16 struct {
	Choice              uint64
	aperiodic_r16       *Resourcetype_r16_SRS_PosResource_r16_aperiodic_r16
	semi_persistent_r16 *Resourcetype_r16_SRS_PosResource_r16_semi_persistent_r16
	periodic_r16        *Resourcetype_r16_SRS_PosResource_r16_periodic_r16
}

func (ie *Resourcetype_r16_SRS_PosResource_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case Resourcetype_r16_SRS_PosResource_r16_Choice_aperiodic_r16:
		if err = ie.aperiodic_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode aperiodic_r16", err)
		}
	case Resourcetype_r16_SRS_PosResource_r16_Choice_semi_persistent_r16:
		if err = ie.semi_persistent_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode semi_persistent_r16", err)
		}
	case Resourcetype_r16_SRS_PosResource_r16_Choice_periodic_r16:
		if err = ie.periodic_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode periodic_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *Resourcetype_r16_SRS_PosResource_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case Resourcetype_r16_SRS_PosResource_r16_Choice_aperiodic_r16:
		ie.aperiodic_r16 = new(Resourcetype_r16_SRS_PosResource_r16_aperiodic_r16)
		if err = ie.aperiodic_r16.Decode(r); err != nil {
			return utils.WrapError("Decode aperiodic_r16", err)
		}
	case Resourcetype_r16_SRS_PosResource_r16_Choice_semi_persistent_r16:
		ie.semi_persistent_r16 = new(Resourcetype_r16_SRS_PosResource_r16_semi_persistent_r16)
		if err = ie.semi_persistent_r16.Decode(r); err != nil {
			return utils.WrapError("Decode semi_persistent_r16", err)
		}
	case Resourcetype_r16_SRS_PosResource_r16_Choice_periodic_r16:
		ie.periodic_r16 = new(Resourcetype_r16_SRS_PosResource_r16_periodic_r16)
		if err = ie.periodic_r16.Decode(r); err != nil {
			return utils.WrapError("Decode periodic_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
