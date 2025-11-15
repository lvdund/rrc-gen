package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_MRDC_Capability_v15g0 struct {
	rf_ParametersMRDC_v15g0 *RF_ParametersMRDC_v15g0 `optional`
	nonCriticalExtension    interface{}              `optional`
}

func (ie *UE_MRDC_Capability_v15g0) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.rf_ParametersMRDC_v15g0 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.rf_ParametersMRDC_v15g0 != nil {
		if err = ie.rf_ParametersMRDC_v15g0.Encode(w); err != nil {
			return utils.WrapError("Encode rf_ParametersMRDC_v15g0", err)
		}
	}
	return nil
}

func (ie *UE_MRDC_Capability_v15g0) Decode(r *uper.UperReader) error {
	var err error
	var rf_ParametersMRDC_v15g0Present bool
	if rf_ParametersMRDC_v15g0Present, err = r.ReadBool(); err != nil {
		return err
	}
	if rf_ParametersMRDC_v15g0Present {
		ie.rf_ParametersMRDC_v15g0 = new(RF_ParametersMRDC_v15g0)
		if err = ie.rf_ParametersMRDC_v15g0.Decode(r); err != nil {
			return utils.WrapError("Decode rf_ParametersMRDC_v15g0", err)
		}
	}
	return nil
}
