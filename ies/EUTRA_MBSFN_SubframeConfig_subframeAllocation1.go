package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	EUTRA_MBSFN_SubframeConfig_subframeAllocation1_Choice_nothing uint64 = iota
	EUTRA_MBSFN_SubframeConfig_subframeAllocation1_Choice_oneFrame
	EUTRA_MBSFN_SubframeConfig_subframeAllocation1_Choice_fourFrames
)

type EUTRA_MBSFN_SubframeConfig_subframeAllocation1 struct {
	Choice     uint64
	oneFrame   uper.BitString `lb:6,ub:6,madatory`
	fourFrames uper.BitString `lb:24,ub:24,madatory`
}

func (ie *EUTRA_MBSFN_SubframeConfig_subframeAllocation1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case EUTRA_MBSFN_SubframeConfig_subframeAllocation1_Choice_oneFrame:
		if err = w.WriteBitString(ie.oneFrame.Bytes, uint(ie.oneFrame.NumBits), &uper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
			err = utils.WrapError("Encode oneFrame", err)
		}
	case EUTRA_MBSFN_SubframeConfig_subframeAllocation1_Choice_fourFrames:
		if err = w.WriteBitString(ie.fourFrames.Bytes, uint(ie.fourFrames.NumBits), &uper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
			err = utils.WrapError("Encode fourFrames", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *EUTRA_MBSFN_SubframeConfig_subframeAllocation1) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case EUTRA_MBSFN_SubframeConfig_subframeAllocation1_Choice_oneFrame:
		var tmp_bs_oneFrame []byte
		var n_oneFrame uint
		if tmp_bs_oneFrame, n_oneFrame, err = r.ReadBitString(&uper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
			return utils.WrapError("Decode oneFrame", err)
		}
		ie.oneFrame = uper.BitString{
			Bytes:   tmp_bs_oneFrame,
			NumBits: uint64(n_oneFrame),
		}
	case EUTRA_MBSFN_SubframeConfig_subframeAllocation1_Choice_fourFrames:
		var tmp_bs_fourFrames []byte
		var n_fourFrames uint
		if tmp_bs_fourFrames, n_fourFrames, err = r.ReadBitString(&uper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
			return utils.WrapError("Decode fourFrames", err)
		}
		ie.fourFrames = uper.BitString{
			Bytes:   tmp_bs_fourFrames,
			NumBits: uint64(n_fourFrames),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
