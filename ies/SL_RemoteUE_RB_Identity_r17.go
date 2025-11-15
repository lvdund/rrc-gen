package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_RemoteUE_RB_Identity_r17_Choice_nothing uint64 = iota
	SL_RemoteUE_RB_Identity_r17_Choice_srb_Identity_r17
	SL_RemoteUE_RB_Identity_r17_Choice_drb_Identity_r17
)

type SL_RemoteUE_RB_Identity_r17 struct {
	Choice           uint64
	srb_Identity_r17 int64 `lb:0,ub:3,madatory`
	drb_Identity_r17 *DRB_Identity
}

func (ie *SL_RemoteUE_RB_Identity_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_RemoteUE_RB_Identity_r17_Choice_srb_Identity_r17:
		if err = w.WriteInteger(int64(ie.srb_Identity_r17), &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			err = utils.WrapError("Encode srb_Identity_r17", err)
		}
	case SL_RemoteUE_RB_Identity_r17_Choice_drb_Identity_r17:
		if err = ie.drb_Identity_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode drb_Identity_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SL_RemoteUE_RB_Identity_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_RemoteUE_RB_Identity_r17_Choice_srb_Identity_r17:
		var tmp_int_srb_Identity_r17 int64
		if tmp_int_srb_Identity_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode srb_Identity_r17", err)
		}
		ie.srb_Identity_r17 = tmp_int_srb_Identity_r17
	case SL_RemoteUE_RB_Identity_r17_Choice_drb_Identity_r17:
		ie.drb_Identity_r17 = new(DRB_Identity)
		if err = ie.drb_Identity_r17.Decode(r); err != nil {
			return utils.WrapError("Decode drb_Identity_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
