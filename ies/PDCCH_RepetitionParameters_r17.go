package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCCH_RepetitionParameters_r17 struct {
	supportedMode_r17   PDCCH_RepetitionParameters_r17_supportedMode_r17    `madatory`
	limitX_PerCC_r17    *PDCCH_RepetitionParameters_r17_limitX_PerCC_r17    `optional`
	limitX_AcrossCC_r17 *PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17 `optional`
}

func (ie *PDCCH_RepetitionParameters_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.limitX_PerCC_r17 != nil, ie.limitX_AcrossCC_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.supportedMode_r17.Encode(w); err != nil {
		return utils.WrapError("Encode supportedMode_r17", err)
	}
	if ie.limitX_PerCC_r17 != nil {
		if err = ie.limitX_PerCC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode limitX_PerCC_r17", err)
		}
	}
	if ie.limitX_AcrossCC_r17 != nil {
		if err = ie.limitX_AcrossCC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode limitX_AcrossCC_r17", err)
		}
	}
	return nil
}

func (ie *PDCCH_RepetitionParameters_r17) Decode(r *uper.UperReader) error {
	var err error
	var limitX_PerCC_r17Present bool
	if limitX_PerCC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var limitX_AcrossCC_r17Present bool
	if limitX_AcrossCC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.supportedMode_r17.Decode(r); err != nil {
		return utils.WrapError("Decode supportedMode_r17", err)
	}
	if limitX_PerCC_r17Present {
		ie.limitX_PerCC_r17 = new(PDCCH_RepetitionParameters_r17_limitX_PerCC_r17)
		if err = ie.limitX_PerCC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode limitX_PerCC_r17", err)
		}
	}
	if limitX_AcrossCC_r17Present {
		ie.limitX_AcrossCC_r17 = new(PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17)
		if err = ie.limitX_AcrossCC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode limitX_AcrossCC_r17", err)
		}
	}
	return nil
}
