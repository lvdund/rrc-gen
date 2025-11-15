package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlink_v1610_intraFreqDAPS_r16 struct {
	intraFreqDiffSCS_DAPS_r16 *FeatureSetDownlink_v1610_intraFreqDAPS_r16_intraFreqDiffSCS_DAPS_r16 `optional`
	intraFreqAsyncDAPS_r16    *FeatureSetDownlink_v1610_intraFreqDAPS_r16_intraFreqAsyncDAPS_r16    `optional`
}

func (ie *FeatureSetDownlink_v1610_intraFreqDAPS_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.intraFreqDiffSCS_DAPS_r16 != nil, ie.intraFreqAsyncDAPS_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.intraFreqDiffSCS_DAPS_r16 != nil {
		if err = ie.intraFreqDiffSCS_DAPS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode intraFreqDiffSCS_DAPS_r16", err)
		}
	}
	if ie.intraFreqAsyncDAPS_r16 != nil {
		if err = ie.intraFreqAsyncDAPS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode intraFreqAsyncDAPS_r16", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlink_v1610_intraFreqDAPS_r16) Decode(r *uper.UperReader) error {
	var err error
	var intraFreqDiffSCS_DAPS_r16Present bool
	if intraFreqDiffSCS_DAPS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var intraFreqAsyncDAPS_r16Present bool
	if intraFreqAsyncDAPS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if intraFreqDiffSCS_DAPS_r16Present {
		ie.intraFreqDiffSCS_DAPS_r16 = new(FeatureSetDownlink_v1610_intraFreqDAPS_r16_intraFreqDiffSCS_DAPS_r16)
		if err = ie.intraFreqDiffSCS_DAPS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode intraFreqDiffSCS_DAPS_r16", err)
		}
	}
	if intraFreqAsyncDAPS_r16Present {
		ie.intraFreqAsyncDAPS_r16 = new(FeatureSetDownlink_v1610_intraFreqDAPS_r16_intraFreqAsyncDAPS_r16)
		if err = ie.intraFreqAsyncDAPS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode intraFreqAsyncDAPS_r16", err)
		}
	}
	return nil
}
