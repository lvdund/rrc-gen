package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_ReselectionConfig_r17 struct {
	sl_RSRP_Thresh_r17           *SL_RSRP_Range_r16 `optional`
	sl_FilterCoefficientRSRP_r17 *FilterCoefficient `optional`
	sl_HystMin_r17               *Hysteresis        `optional`
}

func (ie *SL_ReselectionConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_RSRP_Thresh_r17 != nil, ie.sl_FilterCoefficientRSRP_r17 != nil, ie.sl_HystMin_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_RSRP_Thresh_r17 != nil {
		if err = ie.sl_RSRP_Thresh_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_RSRP_Thresh_r17", err)
		}
	}
	if ie.sl_FilterCoefficientRSRP_r17 != nil {
		if err = ie.sl_FilterCoefficientRSRP_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_FilterCoefficientRSRP_r17", err)
		}
	}
	if ie.sl_HystMin_r17 != nil {
		if err = ie.sl_HystMin_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_HystMin_r17", err)
		}
	}
	return nil
}

func (ie *SL_ReselectionConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	var sl_RSRP_Thresh_r17Present bool
	if sl_RSRP_Thresh_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_FilterCoefficientRSRP_r17Present bool
	if sl_FilterCoefficientRSRP_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_HystMin_r17Present bool
	if sl_HystMin_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_RSRP_Thresh_r17Present {
		ie.sl_RSRP_Thresh_r17 = new(SL_RSRP_Range_r16)
		if err = ie.sl_RSRP_Thresh_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_RSRP_Thresh_r17", err)
		}
	}
	if sl_FilterCoefficientRSRP_r17Present {
		ie.sl_FilterCoefficientRSRP_r17 = new(FilterCoefficient)
		if err = ie.sl_FilterCoefficientRSRP_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_FilterCoefficientRSRP_r17", err)
		}
	}
	if sl_HystMin_r17Present {
		ie.sl_HystMin_r17 = new(Hysteresis)
		if err = ie.sl_HystMin_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_HystMin_r17", err)
		}
	}
	return nil
}
