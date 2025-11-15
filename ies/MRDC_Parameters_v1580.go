package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MRDC_Parameters_v1580 struct {
	dynamicPowerSharingNEDC *MRDC_Parameters_v1580_dynamicPowerSharingNEDC `optional`
}

func (ie *MRDC_Parameters_v1580) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dynamicPowerSharingNEDC != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dynamicPowerSharingNEDC != nil {
		if err = ie.dynamicPowerSharingNEDC.Encode(w); err != nil {
			return utils.WrapError("Encode dynamicPowerSharingNEDC", err)
		}
	}
	return nil
}

func (ie *MRDC_Parameters_v1580) Decode(r *uper.UperReader) error {
	var err error
	var dynamicPowerSharingNEDCPresent bool
	if dynamicPowerSharingNEDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if dynamicPowerSharingNEDCPresent {
		ie.dynamicPowerSharingNEDC = new(MRDC_Parameters_v1580_dynamicPowerSharingNEDC)
		if err = ie.dynamicPowerSharingNEDC.Decode(r); err != nil {
			return utils.WrapError("Decode dynamicPowerSharingNEDC", err)
		}
	}
	return nil
}
