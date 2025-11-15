package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_PosResource_r16_transmissionComb_r16_Choice_nothing uint64 = iota
	SRS_PosResource_r16_transmissionComb_r16_Choice_n2_r16
	SRS_PosResource_r16_transmissionComb_r16_Choice_n4_r16
	SRS_PosResource_r16_transmissionComb_r16_Choice_n8_r16
)

type SRS_PosResource_r16_transmissionComb_r16 struct {
	Choice uint64
	n2_r16 *SRS_PosResource_r16_transmissionComb_r16_n2_r16
	n4_r16 *SRS_PosResource_r16_transmissionComb_r16_n4_r16
	n8_r16 *SRS_PosResource_r16_transmissionComb_r16_n8_r16
}

func (ie *SRS_PosResource_r16_transmissionComb_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_PosResource_r16_transmissionComb_r16_Choice_n2_r16:
		if err = ie.n2_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode n2_r16", err)
		}
	case SRS_PosResource_r16_transmissionComb_r16_Choice_n4_r16:
		if err = ie.n4_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode n4_r16", err)
		}
	case SRS_PosResource_r16_transmissionComb_r16_Choice_n8_r16:
		if err = ie.n8_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode n8_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SRS_PosResource_r16_transmissionComb_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_PosResource_r16_transmissionComb_r16_Choice_n2_r16:
		ie.n2_r16 = new(SRS_PosResource_r16_transmissionComb_r16_n2_r16)
		if err = ie.n2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode n2_r16", err)
		}
	case SRS_PosResource_r16_transmissionComb_r16_Choice_n4_r16:
		ie.n4_r16 = new(SRS_PosResource_r16_transmissionComb_r16_n4_r16)
		if err = ie.n4_r16.Decode(r); err != nil {
			return utils.WrapError("Decode n4_r16", err)
		}
	case SRS_PosResource_r16_transmissionComb_r16_Choice_n8_r16:
		ie.n8_r16 = new(SRS_PosResource_r16_transmissionComb_r16_n8_r16)
		if err = ie.n8_r16.Decode(r); err != nil {
			return utils.WrapError("Decode n8_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
