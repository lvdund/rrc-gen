package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BH_LogicalChannelIdentity_r16_Choice_nothing uint64 = iota
	BH_LogicalChannelIdentity_r16_Choice_bh_LogicalChannelIdentity_r16
	BH_LogicalChannelIdentity_r16_Choice_bh_LogicalChannelIdentityExt_r16
)

type BH_LogicalChannelIdentity_r16 struct {
	Choice                           uint64
	bh_LogicalChannelIdentity_r16    *LogicalChannelIdentity
	bh_LogicalChannelIdentityExt_r16 *BH_LogicalChannelIdentity_Ext_r16
}

func (ie *BH_LogicalChannelIdentity_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BH_LogicalChannelIdentity_r16_Choice_bh_LogicalChannelIdentity_r16:
		if err = ie.bh_LogicalChannelIdentity_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode bh_LogicalChannelIdentity_r16", err)
		}
	case BH_LogicalChannelIdentity_r16_Choice_bh_LogicalChannelIdentityExt_r16:
		if err = ie.bh_LogicalChannelIdentityExt_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode bh_LogicalChannelIdentityExt_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *BH_LogicalChannelIdentity_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BH_LogicalChannelIdentity_r16_Choice_bh_LogicalChannelIdentity_r16:
		ie.bh_LogicalChannelIdentity_r16 = new(LogicalChannelIdentity)
		if err = ie.bh_LogicalChannelIdentity_r16.Decode(r); err != nil {
			return utils.WrapError("Decode bh_LogicalChannelIdentity_r16", err)
		}
	case BH_LogicalChannelIdentity_r16_Choice_bh_LogicalChannelIdentityExt_r16:
		ie.bh_LogicalChannelIdentityExt_r16 = new(BH_LogicalChannelIdentity_Ext_r16)
		if err = ie.bh_LogicalChannelIdentityExt_r16.Decode(r); err != nil {
			return utils.WrapError("Decode bh_LogicalChannelIdentityExt_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
