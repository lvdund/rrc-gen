package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FR2_2_AccessParamsPerBand_r17_enhancedPDCCH_monitoringSCS_960kHz_r17 struct {
	pdcch_monitoring4_1_r17 *FR2_2_AccessParamsPerBand_r17_enhancedPDCCH_monitoringSCS_960kHz_r17_pdcch_monitoring4_1_r17 `optional`
	pdcch_monitoring4_2_r17 *FR2_2_AccessParamsPerBand_r17_enhancedPDCCH_monitoringSCS_960kHz_r17_pdcch_monitoring4_2_r17 `optional`
	pdcch_monitoring8_4_r17 *FR2_2_AccessParamsPerBand_r17_enhancedPDCCH_monitoringSCS_960kHz_r17_pdcch_monitoring8_4_r17 `optional`
}

func (ie *FR2_2_AccessParamsPerBand_r17_enhancedPDCCH_monitoringSCS_960kHz_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.pdcch_monitoring4_1_r17 != nil, ie.pdcch_monitoring4_2_r17 != nil, ie.pdcch_monitoring8_4_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.pdcch_monitoring4_1_r17 != nil {
		if err = ie.pdcch_monitoring4_1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pdcch_monitoring4_1_r17", err)
		}
	}
	if ie.pdcch_monitoring4_2_r17 != nil {
		if err = ie.pdcch_monitoring4_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pdcch_monitoring4_2_r17", err)
		}
	}
	if ie.pdcch_monitoring8_4_r17 != nil {
		if err = ie.pdcch_monitoring8_4_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pdcch_monitoring8_4_r17", err)
		}
	}
	return nil
}

func (ie *FR2_2_AccessParamsPerBand_r17_enhancedPDCCH_monitoringSCS_960kHz_r17) Decode(r *uper.UperReader) error {
	var err error
	var pdcch_monitoring4_1_r17Present bool
	if pdcch_monitoring4_1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcch_monitoring4_2_r17Present bool
	if pdcch_monitoring4_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcch_monitoring8_4_r17Present bool
	if pdcch_monitoring8_4_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if pdcch_monitoring4_1_r17Present {
		ie.pdcch_monitoring4_1_r17 = new(FR2_2_AccessParamsPerBand_r17_enhancedPDCCH_monitoringSCS_960kHz_r17_pdcch_monitoring4_1_r17)
		if err = ie.pdcch_monitoring4_1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pdcch_monitoring4_1_r17", err)
		}
	}
	if pdcch_monitoring4_2_r17Present {
		ie.pdcch_monitoring4_2_r17 = new(FR2_2_AccessParamsPerBand_r17_enhancedPDCCH_monitoringSCS_960kHz_r17_pdcch_monitoring4_2_r17)
		if err = ie.pdcch_monitoring4_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pdcch_monitoring4_2_r17", err)
		}
	}
	if pdcch_monitoring8_4_r17Present {
		ie.pdcch_monitoring8_4_r17 = new(FR2_2_AccessParamsPerBand_r17_enhancedPDCCH_monitoringSCS_960kHz_r17_pdcch_monitoring8_4_r17)
		if err = ie.pdcch_monitoring8_4_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pdcch_monitoring8_4_r17", err)
		}
	}
	return nil
}
