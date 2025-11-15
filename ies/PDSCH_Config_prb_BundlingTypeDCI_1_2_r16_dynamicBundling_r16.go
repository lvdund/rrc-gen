package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_dynamicBundling_r16 struct {
	bundleSizeSet1_r16 *PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_dynamicBundling_r16_bundleSizeSet1_r16 `optional`
	bundleSizeSet2_r16 *PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_dynamicBundling_r16_bundleSizeSet2_r16 `optional`
}

func (ie *PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_dynamicBundling_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.bundleSizeSet1_r16 != nil, ie.bundleSizeSet2_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.bundleSizeSet1_r16 != nil {
		if err = ie.bundleSizeSet1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode bundleSizeSet1_r16", err)
		}
	}
	if ie.bundleSizeSet2_r16 != nil {
		if err = ie.bundleSizeSet2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode bundleSizeSet2_r16", err)
		}
	}
	return nil
}

func (ie *PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_dynamicBundling_r16) Decode(r *uper.UperReader) error {
	var err error
	var bundleSizeSet1_r16Present bool
	if bundleSizeSet1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var bundleSizeSet2_r16Present bool
	if bundleSizeSet2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if bundleSizeSet1_r16Present {
		ie.bundleSizeSet1_r16 = new(PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_dynamicBundling_r16_bundleSizeSet1_r16)
		if err = ie.bundleSizeSet1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode bundleSizeSet1_r16", err)
		}
	}
	if bundleSizeSet2_r16Present {
		ie.bundleSizeSet2_r16 = new(PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_dynamicBundling_r16_bundleSizeSet2_r16)
		if err = ie.bundleSizeSet2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode bundleSizeSet2_r16", err)
		}
	}
	return nil
}
