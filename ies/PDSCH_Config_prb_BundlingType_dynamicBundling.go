package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_Config_prb_BundlingType_dynamicBundling struct {
	bundleSizeSet1 *PDSCH_Config_prb_BundlingType_dynamicBundling_bundleSizeSet1 `optional`
	bundleSizeSet2 *PDSCH_Config_prb_BundlingType_dynamicBundling_bundleSizeSet2 `optional`
}

func (ie *PDSCH_Config_prb_BundlingType_dynamicBundling) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.bundleSizeSet1 != nil, ie.bundleSizeSet2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.bundleSizeSet1 != nil {
		if err = ie.bundleSizeSet1.Encode(w); err != nil {
			return utils.WrapError("Encode bundleSizeSet1", err)
		}
	}
	if ie.bundleSizeSet2 != nil {
		if err = ie.bundleSizeSet2.Encode(w); err != nil {
			return utils.WrapError("Encode bundleSizeSet2", err)
		}
	}
	return nil
}

func (ie *PDSCH_Config_prb_BundlingType_dynamicBundling) Decode(r *uper.UperReader) error {
	var err error
	var bundleSizeSet1Present bool
	if bundleSizeSet1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var bundleSizeSet2Present bool
	if bundleSizeSet2Present, err = r.ReadBool(); err != nil {
		return err
	}
	if bundleSizeSet1Present {
		ie.bundleSizeSet1 = new(PDSCH_Config_prb_BundlingType_dynamicBundling_bundleSizeSet1)
		if err = ie.bundleSizeSet1.Decode(r); err != nil {
			return utils.WrapError("Decode bundleSizeSet1", err)
		}
	}
	if bundleSizeSet2Present {
		ie.bundleSizeSet2 = new(PDSCH_Config_prb_BundlingType_dynamicBundling_bundleSizeSet2)
		if err = ie.bundleSizeSet2.Decode(r); err != nil {
			return utils.WrapError("Decode bundleSizeSet2", err)
		}
	}
	return nil
}
