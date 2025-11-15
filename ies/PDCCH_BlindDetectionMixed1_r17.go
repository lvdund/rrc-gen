package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCCH_BlindDetectionMixed1_r17 struct {
	pdcch_BlindDetectionCA_Mixed1_r17    *PDCCH_BlindDetectionCA_Mixed1_r17                                   `optional`
	pdcch_BlindDetectionCG_UE_Mixed1_r17 *PDCCH_BlindDetectionMixed1_r17_pdcch_BlindDetectionCG_UE_Mixed1_r17 `optional`
}

func (ie *PDCCH_BlindDetectionMixed1_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.pdcch_BlindDetectionCA_Mixed1_r17 != nil, ie.pdcch_BlindDetectionCG_UE_Mixed1_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.pdcch_BlindDetectionCA_Mixed1_r17 != nil {
		if err = ie.pdcch_BlindDetectionCA_Mixed1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pdcch_BlindDetectionCA_Mixed1_r17", err)
		}
	}
	if ie.pdcch_BlindDetectionCG_UE_Mixed1_r17 != nil {
		if err = ie.pdcch_BlindDetectionCG_UE_Mixed1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pdcch_BlindDetectionCG_UE_Mixed1_r17", err)
		}
	}
	return nil
}

func (ie *PDCCH_BlindDetectionMixed1_r17) Decode(r *uper.UperReader) error {
	var err error
	var pdcch_BlindDetectionCA_Mixed1_r17Present bool
	if pdcch_BlindDetectionCA_Mixed1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcch_BlindDetectionCG_UE_Mixed1_r17Present bool
	if pdcch_BlindDetectionCG_UE_Mixed1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if pdcch_BlindDetectionCA_Mixed1_r17Present {
		ie.pdcch_BlindDetectionCA_Mixed1_r17 = new(PDCCH_BlindDetectionCA_Mixed1_r17)
		if err = ie.pdcch_BlindDetectionCA_Mixed1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pdcch_BlindDetectionCA_Mixed1_r17", err)
		}
	}
	if pdcch_BlindDetectionCG_UE_Mixed1_r17Present {
		ie.pdcch_BlindDetectionCG_UE_Mixed1_r17 = new(PDCCH_BlindDetectionMixed1_r17_pdcch_BlindDetectionCG_UE_Mixed1_r17)
		if err = ie.pdcch_BlindDetectionCG_UE_Mixed1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pdcch_BlindDetectionCG_UE_Mixed1_r17", err)
		}
	}
	return nil
}
