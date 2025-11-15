package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_staticBundling_r16 struct {
	bundleSize_r16 *PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_staticBundling_r16_bundleSize_r16 `optional`
}

func (ie *PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_staticBundling_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.bundleSize_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.bundleSize_r16 != nil {
		if err = ie.bundleSize_r16.Encode(w); err != nil {
			return utils.WrapError("Encode bundleSize_r16", err)
		}
	}
	return nil
}

func (ie *PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_staticBundling_r16) Decode(r *uper.UperReader) error {
	var err error
	var bundleSize_r16Present bool
	if bundleSize_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if bundleSize_r16Present {
		ie.bundleSize_r16 = new(PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_staticBundling_r16_bundleSize_r16)
		if err = ie.bundleSize_r16.Decode(r); err != nil {
			return utils.WrapError("Decode bundleSize_r16", err)
		}
	}
	return nil
}
