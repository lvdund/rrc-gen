package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CarrierTypePair_r16 struct {
	carrierForCSI_Measurement_r16 PUCCH_Grp_CarrierTypes_r16 `madatory`
	carrierForCSI_Reporting_r16   PUCCH_Grp_CarrierTypes_r16 `madatory`
}

func (ie *CarrierTypePair_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.carrierForCSI_Measurement_r16.Encode(w); err != nil {
		return utils.WrapError("Encode carrierForCSI_Measurement_r16", err)
	}
	if err = ie.carrierForCSI_Reporting_r16.Encode(w); err != nil {
		return utils.WrapError("Encode carrierForCSI_Reporting_r16", err)
	}
	return nil
}

func (ie *CarrierTypePair_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.carrierForCSI_Measurement_r16.Decode(r); err != nil {
		return utils.WrapError("Decode carrierForCSI_Measurement_r16", err)
	}
	if err = ie.carrierForCSI_Reporting_r16.Decode(r); err != nil {
		return utils.WrapError("Decode carrierForCSI_Reporting_r16", err)
	}
	return nil
}
