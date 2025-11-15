package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RLC_BearerConfig_servedRadioBearer_Choice_nothing uint64 = iota
	RLC_BearerConfig_servedRadioBearer_Choice_srb_Identity
	RLC_BearerConfig_servedRadioBearer_Choice_drb_Identity
)

type RLC_BearerConfig_servedRadioBearer struct {
	Choice       uint64
	srb_Identity *SRB_Identity
	drb_Identity *DRB_Identity
}

func (ie *RLC_BearerConfig_servedRadioBearer) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RLC_BearerConfig_servedRadioBearer_Choice_srb_Identity:
		if err = ie.srb_Identity.Encode(w); err != nil {
			err = utils.WrapError("Encode srb_Identity", err)
		}
	case RLC_BearerConfig_servedRadioBearer_Choice_drb_Identity:
		if err = ie.drb_Identity.Encode(w); err != nil {
			err = utils.WrapError("Encode drb_Identity", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RLC_BearerConfig_servedRadioBearer) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RLC_BearerConfig_servedRadioBearer_Choice_srb_Identity:
		ie.srb_Identity = new(SRB_Identity)
		if err = ie.srb_Identity.Decode(r); err != nil {
			return utils.WrapError("Decode srb_Identity", err)
		}
	case RLC_BearerConfig_servedRadioBearer_Choice_drb_Identity:
		ie.drb_Identity = new(DRB_Identity)
		if err = ie.drb_Identity.Decode(r); err != nil {
			return utils.WrapError("Decode drb_Identity", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
