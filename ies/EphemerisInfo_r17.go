package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	EphemerisInfo_r17_Choice_nothing uint64 = iota
	EphemerisInfo_r17_Choice_positionVelocity_r17
	EphemerisInfo_r17_Choice_orbital_r17
)

type EphemerisInfo_r17 struct {
	Choice               uint64
	positionVelocity_r17 *PositionVelocity_r17
	orbital_r17          *Orbital_r17
}

func (ie *EphemerisInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case EphemerisInfo_r17_Choice_positionVelocity_r17:
		if err = ie.positionVelocity_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode positionVelocity_r17", err)
		}
	case EphemerisInfo_r17_Choice_orbital_r17:
		if err = ie.orbital_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode orbital_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *EphemerisInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case EphemerisInfo_r17_Choice_positionVelocity_r17:
		ie.positionVelocity_r17 = new(PositionVelocity_r17)
		if err = ie.positionVelocity_r17.Decode(r); err != nil {
			return utils.WrapError("Decode positionVelocity_r17", err)
		}
	case EphemerisInfo_r17_Choice_orbital_r17:
		ie.orbital_r17 = new(Orbital_r17)
		if err = ie.orbital_r17.Decode(r); err != nil {
			return utils.WrapError("Decode orbital_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
