package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_UplinkTxSwitch_v1720 struct {
	bandCombination_v1720                   *BandCombination_v1720                                                        `optional`
	uplinkTxSwitching_OptionSupport2T2T_r17 *BandCombination_UplinkTxSwitch_v1720_uplinkTxSwitching_OptionSupport2T2T_r17 `optional`
}

func (ie *BandCombination_UplinkTxSwitch_v1720) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.bandCombination_v1720 != nil, ie.uplinkTxSwitching_OptionSupport2T2T_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.bandCombination_v1720 != nil {
		if err = ie.bandCombination_v1720.Encode(w); err != nil {
			return utils.WrapError("Encode bandCombination_v1720", err)
		}
	}
	if ie.uplinkTxSwitching_OptionSupport2T2T_r17 != nil {
		if err = ie.uplinkTxSwitching_OptionSupport2T2T_r17.Encode(w); err != nil {
			return utils.WrapError("Encode uplinkTxSwitching_OptionSupport2T2T_r17", err)
		}
	}
	return nil
}

func (ie *BandCombination_UplinkTxSwitch_v1720) Decode(r *uper.UperReader) error {
	var err error
	var bandCombination_v1720Present bool
	if bandCombination_v1720Present, err = r.ReadBool(); err != nil {
		return err
	}
	var uplinkTxSwitching_OptionSupport2T2T_r17Present bool
	if uplinkTxSwitching_OptionSupport2T2T_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if bandCombination_v1720Present {
		ie.bandCombination_v1720 = new(BandCombination_v1720)
		if err = ie.bandCombination_v1720.Decode(r); err != nil {
			return utils.WrapError("Decode bandCombination_v1720", err)
		}
	}
	if uplinkTxSwitching_OptionSupport2T2T_r17Present {
		ie.uplinkTxSwitching_OptionSupport2T2T_r17 = new(BandCombination_UplinkTxSwitch_v1720_uplinkTxSwitching_OptionSupport2T2T_r17)
		if err = ie.uplinkTxSwitching_OptionSupport2T2T_r17.Decode(r); err != nil {
			return utils.WrapError("Decode uplinkTxSwitching_OptionSupport2T2T_r17", err)
		}
	}
	return nil
}
