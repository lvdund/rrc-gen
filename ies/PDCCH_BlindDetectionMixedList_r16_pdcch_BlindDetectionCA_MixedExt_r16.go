package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDCCH_BlindDetectionMixedList_r16_pdcch_BlindDetectionCA_MixedExt_r16_Choice_nothing uint64 = iota
	PDCCH_BlindDetectionMixedList_r16_pdcch_BlindDetectionCA_MixedExt_r16_Choice_pdcch_BlindDetectionCA_Mixed_v16a0
	PDCCH_BlindDetectionMixedList_r16_pdcch_BlindDetectionCA_MixedExt_r16_Choice_pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_v16a0
)

type PDCCH_BlindDetectionMixedList_r16_pdcch_BlindDetectionCA_MixedExt_r16 struct {
	Choice                                            uint64
	pdcch_BlindDetectionCA_Mixed_v16a0                *PDCCH_BlindDetectionCA_MixedExt_r16
	pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_v16a0 *PDCCH_BlindDetectionCA_MixedExt_r16
}

func (ie *PDCCH_BlindDetectionMixedList_r16_pdcch_BlindDetectionCA_MixedExt_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PDCCH_BlindDetectionMixedList_r16_pdcch_BlindDetectionCA_MixedExt_r16_Choice_pdcch_BlindDetectionCA_Mixed_v16a0:
		if err = ie.pdcch_BlindDetectionCA_Mixed_v16a0.Encode(w); err != nil {
			err = utils.WrapError("Encode pdcch_BlindDetectionCA_Mixed_v16a0", err)
		}
	case PDCCH_BlindDetectionMixedList_r16_pdcch_BlindDetectionCA_MixedExt_r16_Choice_pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_v16a0:
		if err = ie.pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_v16a0.Encode(w); err != nil {
			err = utils.WrapError("Encode pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_v16a0", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PDCCH_BlindDetectionMixedList_r16_pdcch_BlindDetectionCA_MixedExt_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PDCCH_BlindDetectionMixedList_r16_pdcch_BlindDetectionCA_MixedExt_r16_Choice_pdcch_BlindDetectionCA_Mixed_v16a0:
		ie.pdcch_BlindDetectionCA_Mixed_v16a0 = new(PDCCH_BlindDetectionCA_MixedExt_r16)
		if err = ie.pdcch_BlindDetectionCA_Mixed_v16a0.Decode(r); err != nil {
			return utils.WrapError("Decode pdcch_BlindDetectionCA_Mixed_v16a0", err)
		}
	case PDCCH_BlindDetectionMixedList_r16_pdcch_BlindDetectionCA_MixedExt_r16_Choice_pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_v16a0:
		ie.pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_v16a0 = new(PDCCH_BlindDetectionCA_MixedExt_r16)
		if err = ie.pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_v16a0.Decode(r); err != nil {
			return utils.WrapError("Decode pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_v16a0", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
