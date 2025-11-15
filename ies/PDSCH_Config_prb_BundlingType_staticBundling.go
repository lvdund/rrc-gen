package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_Config_prb_BundlingType_staticBundling struct {
	bundleSize *PDSCH_Config_prb_BundlingType_staticBundling_bundleSize `optional`
}

func (ie *PDSCH_Config_prb_BundlingType_staticBundling) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.bundleSize != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.bundleSize != nil {
		if err = ie.bundleSize.Encode(w); err != nil {
			return utils.WrapError("Encode bundleSize", err)
		}
	}
	return nil
}

func (ie *PDSCH_Config_prb_BundlingType_staticBundling) Decode(r *uper.UperReader) error {
	var err error
	var bundleSizePresent bool
	if bundleSizePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if bundleSizePresent {
		ie.bundleSize = new(PDSCH_Config_prb_BundlingType_staticBundling_bundleSize)
		if err = ie.bundleSize.Decode(r); err != nil {
			return utils.WrapError("Decode bundleSize", err)
		}
	}
	return nil
}
