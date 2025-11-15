package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1540_csi_RS_IM_ReceptionForFeedbackPerBandComb struct {
	maxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCC        *int64 `lb:1,ub:64,optional`
	totalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCC *int64 `lb:2,ub:256,optional`
}

func (ie *CA_ParametersNR_v1540_csi_RS_IM_ReceptionForFeedbackPerBandComb) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.maxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCC != nil, ie.totalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCC != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.maxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCC != nil {
		if err = w.WriteInteger(*ie.maxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCC, &uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
			return utils.WrapError("Encode maxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCC", err)
		}
	}
	if ie.totalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCC != nil {
		if err = w.WriteInteger(*ie.totalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCC, &uper.Constraint{Lb: 2, Ub: 256}, false); err != nil {
			return utils.WrapError("Encode totalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCC", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNR_v1540_csi_RS_IM_ReceptionForFeedbackPerBandComb) Decode(r *uper.UperReader) error {
	var err error
	var maxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCCPresent bool
	if maxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var totalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCCPresent bool
	if totalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if maxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCCPresent {
		var tmp_int_maxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCC int64
		if tmp_int_maxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCC, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode maxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCC", err)
		}
		ie.maxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCC = &tmp_int_maxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCC
	}
	if totalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCCPresent {
		var tmp_int_totalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCC int64
		if tmp_int_totalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCC, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 256}, false); err != nil {
			return utils.WrapError("Decode totalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCC", err)
		}
		ie.totalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCC = &tmp_int_totalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCC
	}
	return nil
}
