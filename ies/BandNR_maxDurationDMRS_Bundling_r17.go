package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandNR_maxDurationDMRS_Bundling_r17 struct {
	fdd_r17 *BandNR_maxDurationDMRS_Bundling_r17_fdd_r17 `optional`
	tdd_r17 *BandNR_maxDurationDMRS_Bundling_r17_tdd_r17 `optional`
}

func (ie *BandNR_maxDurationDMRS_Bundling_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.fdd_r17 != nil, ie.tdd_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.fdd_r17 != nil {
		if err = ie.fdd_r17.Encode(w); err != nil {
			return utils.WrapError("Encode fdd_r17", err)
		}
	}
	if ie.tdd_r17 != nil {
		if err = ie.tdd_r17.Encode(w); err != nil {
			return utils.WrapError("Encode tdd_r17", err)
		}
	}
	return nil
}

func (ie *BandNR_maxDurationDMRS_Bundling_r17) Decode(r *uper.UperReader) error {
	var err error
	var fdd_r17Present bool
	if fdd_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tdd_r17Present bool
	if tdd_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if fdd_r17Present {
		ie.fdd_r17 = new(BandNR_maxDurationDMRS_Bundling_r17_fdd_r17)
		if err = ie.fdd_r17.Decode(r); err != nil {
			return utils.WrapError("Decode fdd_r17", err)
		}
	}
	if tdd_r17Present {
		ie.tdd_r17 = new(BandNR_maxDurationDMRS_Bundling_r17_tdd_r17)
		if err = ie.tdd_r17.Decode(r); err != nil {
			return utils.WrapError("Decode tdd_r17", err)
		}
	}
	return nil
}
