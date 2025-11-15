package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17 struct {
	minBeamApplicationTime_r17 *MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_minBeamApplicationTime_r17 `optional`
	maxNumMAC_CE_PerCC         MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_maxNumMAC_CE_PerCC          `madatory`
}

func (ie *MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.minBeamApplicationTime_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.minBeamApplicationTime_r17 != nil {
		if err = ie.minBeamApplicationTime_r17.Encode(w); err != nil {
			return utils.WrapError("Encode minBeamApplicationTime_r17", err)
		}
	}
	if err = ie.maxNumMAC_CE_PerCC.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumMAC_CE_PerCC", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17) Decode(r *uper.UperReader) error {
	var err error
	var minBeamApplicationTime_r17Present bool
	if minBeamApplicationTime_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if minBeamApplicationTime_r17Present {
		ie.minBeamApplicationTime_r17 = new(MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_minBeamApplicationTime_r17)
		if err = ie.minBeamApplicationTime_r17.Decode(r); err != nil {
			return utils.WrapError("Decode minBeamApplicationTime_r17", err)
		}
	}
	if err = ie.maxNumMAC_CE_PerCC.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumMAC_CE_PerCC", err)
	}
	return nil
}
