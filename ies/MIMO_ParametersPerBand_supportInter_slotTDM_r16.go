package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_supportInter_slotTDM_r16 struct {
	supportRepNumPDSCH_TDRA_r16 MIMO_ParametersPerBand_supportInter_slotTDM_r16_supportRepNumPDSCH_TDRA_r16 `madatory`
	maxTBS_Size_r16             MIMO_ParametersPerBand_supportInter_slotTDM_r16_maxTBS_Size_r16             `madatory`
	maxNumberTCI_states_r16     int64                                                                       `lb:1,ub:2,madatory`
}

func (ie *MIMO_ParametersPerBand_supportInter_slotTDM_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.supportRepNumPDSCH_TDRA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode supportRepNumPDSCH_TDRA_r16", err)
	}
	if err = ie.maxTBS_Size_r16.Encode(w); err != nil {
		return utils.WrapError("Encode maxTBS_Size_r16", err)
	}
	if err = w.WriteInteger(ie.maxNumberTCI_states_r16, &uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberTCI_states_r16", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_supportInter_slotTDM_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.supportRepNumPDSCH_TDRA_r16.Decode(r); err != nil {
		return utils.WrapError("Decode supportRepNumPDSCH_TDRA_r16", err)
	}
	if err = ie.maxTBS_Size_r16.Decode(r); err != nil {
		return utils.WrapError("Decode maxTBS_Size_r16", err)
	}
	var tmp_int_maxNumberTCI_states_r16 int64
	if tmp_int_maxNumberTCI_states_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberTCI_states_r16", err)
	}
	ie.maxNumberTCI_states_r16 = tmp_int_maxNumberTCI_states_r16
	return nil
}
