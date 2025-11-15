package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1560 struct {
	diffNumerologyWithinPUCCH_GroupLargerSCS *CA_ParametersNR_v1560_diffNumerologyWithinPUCCH_GroupLargerSCS `optional`
}

func (ie *CA_ParametersNR_v1560) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.diffNumerologyWithinPUCCH_GroupLargerSCS != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.diffNumerologyWithinPUCCH_GroupLargerSCS != nil {
		if err = ie.diffNumerologyWithinPUCCH_GroupLargerSCS.Encode(w); err != nil {
			return utils.WrapError("Encode diffNumerologyWithinPUCCH_GroupLargerSCS", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNR_v1560) Decode(r *uper.UperReader) error {
	var err error
	var diffNumerologyWithinPUCCH_GroupLargerSCSPresent bool
	if diffNumerologyWithinPUCCH_GroupLargerSCSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if diffNumerologyWithinPUCCH_GroupLargerSCSPresent {
		ie.diffNumerologyWithinPUCCH_GroupLargerSCS = new(CA_ParametersNR_v1560_diffNumerologyWithinPUCCH_GroupLargerSCS)
		if err = ie.diffNumerologyWithinPUCCH_GroupLargerSCS.Decode(r); err != nil {
			return utils.WrapError("Decode diffNumerologyWithinPUCCH_GroupLargerSCS", err)
		}
	}
	return nil
}
