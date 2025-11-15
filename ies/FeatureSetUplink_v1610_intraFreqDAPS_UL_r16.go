package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplink_v1610_intraFreqDAPS_UL_r16 struct {
	dummy                     *FeatureSetUplink_v1610_intraFreqDAPS_UL_r16_dummy                     `optional`
	intraFreqTwoTAGs_DAPS_r16 *FeatureSetUplink_v1610_intraFreqDAPS_UL_r16_intraFreqTwoTAGs_DAPS_r16 `optional`
	dummy1                    *FeatureSetUplink_v1610_intraFreqDAPS_UL_r16_dummy1                    `optional`
	dummy2                    *FeatureSetUplink_v1610_intraFreqDAPS_UL_r16_dummy2                    `optional`
	dummy3                    *FeatureSetUplink_v1610_intraFreqDAPS_UL_r16_dummy3                    `optional`
}

func (ie *FeatureSetUplink_v1610_intraFreqDAPS_UL_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dummy != nil, ie.intraFreqTwoTAGs_DAPS_r16 != nil, ie.dummy1 != nil, ie.dummy2 != nil, ie.dummy3 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dummy != nil {
		if err = ie.dummy.Encode(w); err != nil {
			return utils.WrapError("Encode dummy", err)
		}
	}
	if ie.intraFreqTwoTAGs_DAPS_r16 != nil {
		if err = ie.intraFreqTwoTAGs_DAPS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode intraFreqTwoTAGs_DAPS_r16", err)
		}
	}
	if ie.dummy1 != nil {
		if err = ie.dummy1.Encode(w); err != nil {
			return utils.WrapError("Encode dummy1", err)
		}
	}
	if ie.dummy2 != nil {
		if err = ie.dummy2.Encode(w); err != nil {
			return utils.WrapError("Encode dummy2", err)
		}
	}
	if ie.dummy3 != nil {
		if err = ie.dummy3.Encode(w); err != nil {
			return utils.WrapError("Encode dummy3", err)
		}
	}
	return nil
}

func (ie *FeatureSetUplink_v1610_intraFreqDAPS_UL_r16) Decode(r *uper.UperReader) error {
	var err error
	var dummyPresent bool
	if dummyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var intraFreqTwoTAGs_DAPS_r16Present bool
	if intraFreqTwoTAGs_DAPS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dummy1Present bool
	if dummy1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dummy2Present bool
	if dummy2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dummy3Present bool
	if dummy3Present, err = r.ReadBool(); err != nil {
		return err
	}
	if dummyPresent {
		ie.dummy = new(FeatureSetUplink_v1610_intraFreqDAPS_UL_r16_dummy)
		if err = ie.dummy.Decode(r); err != nil {
			return utils.WrapError("Decode dummy", err)
		}
	}
	if intraFreqTwoTAGs_DAPS_r16Present {
		ie.intraFreqTwoTAGs_DAPS_r16 = new(FeatureSetUplink_v1610_intraFreqDAPS_UL_r16_intraFreqTwoTAGs_DAPS_r16)
		if err = ie.intraFreqTwoTAGs_DAPS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode intraFreqTwoTAGs_DAPS_r16", err)
		}
	}
	if dummy1Present {
		ie.dummy1 = new(FeatureSetUplink_v1610_intraFreqDAPS_UL_r16_dummy1)
		if err = ie.dummy1.Decode(r); err != nil {
			return utils.WrapError("Decode dummy1", err)
		}
	}
	if dummy2Present {
		ie.dummy2 = new(FeatureSetUplink_v1610_intraFreqDAPS_UL_r16_dummy2)
		if err = ie.dummy2.Decode(r); err != nil {
			return utils.WrapError("Decode dummy2", err)
		}
	}
	if dummy3Present {
		ie.dummy3 = new(FeatureSetUplink_v1610_intraFreqDAPS_UL_r16_dummy3)
		if err = ie.dummy3.Decode(r); err != nil {
			return utils.WrapError("Decode dummy3", err)
		}
	}
	return nil
}
