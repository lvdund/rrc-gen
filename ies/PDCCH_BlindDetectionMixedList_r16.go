package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCCH_BlindDetectionMixedList_r16 struct {
	pdcch_BlindDetectionCA_MixedExt_r16    *PDCCH_BlindDetectionMixedList_r16_pdcch_BlindDetectionCA_MixedExt_r16    `optional`
	pdcch_BlindDetectionCG_UE_MixedExt_r16 *PDCCH_BlindDetectionMixedList_r16_pdcch_BlindDetectionCG_UE_MixedExt_r16 `optional`
}

func (ie *PDCCH_BlindDetectionMixedList_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.pdcch_BlindDetectionCA_MixedExt_r16 != nil, ie.pdcch_BlindDetectionCG_UE_MixedExt_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.pdcch_BlindDetectionCA_MixedExt_r16 != nil {
		if err = ie.pdcch_BlindDetectionCA_MixedExt_r16.Encode(w); err != nil {
			return utils.WrapError("Encode pdcch_BlindDetectionCA_MixedExt_r16", err)
		}
	}
	if ie.pdcch_BlindDetectionCG_UE_MixedExt_r16 != nil {
		if err = ie.pdcch_BlindDetectionCG_UE_MixedExt_r16.Encode(w); err != nil {
			return utils.WrapError("Encode pdcch_BlindDetectionCG_UE_MixedExt_r16", err)
		}
	}
	return nil
}

func (ie *PDCCH_BlindDetectionMixedList_r16) Decode(r *uper.UperReader) error {
	var err error
	var pdcch_BlindDetectionCA_MixedExt_r16Present bool
	if pdcch_BlindDetectionCA_MixedExt_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcch_BlindDetectionCG_UE_MixedExt_r16Present bool
	if pdcch_BlindDetectionCG_UE_MixedExt_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if pdcch_BlindDetectionCA_MixedExt_r16Present {
		ie.pdcch_BlindDetectionCA_MixedExt_r16 = new(PDCCH_BlindDetectionMixedList_r16_pdcch_BlindDetectionCA_MixedExt_r16)
		if err = ie.pdcch_BlindDetectionCA_MixedExt_r16.Decode(r); err != nil {
			return utils.WrapError("Decode pdcch_BlindDetectionCA_MixedExt_r16", err)
		}
	}
	if pdcch_BlindDetectionCG_UE_MixedExt_r16Present {
		ie.pdcch_BlindDetectionCG_UE_MixedExt_r16 = new(PDCCH_BlindDetectionMixedList_r16_pdcch_BlindDetectionCG_UE_MixedExt_r16)
		if err = ie.pdcch_BlindDetectionCG_UE_MixedExt_r16.Decode(r); err != nil {
			return utils.WrapError("Decode pdcch_BlindDetectionCG_UE_MixedExt_r16", err)
		}
	}
	return nil
}
