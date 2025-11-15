package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1540 struct {
	simultaneousSRS_AssocCSI_RS_AllCC         *int64                                                           `lb:5,ub:32,optional`
	csi_RS_IM_ReceptionForFeedbackPerBandComb *CA_ParametersNR_v1540_csi_RS_IM_ReceptionForFeedbackPerBandComb `lb:1,ub:64,optional`
	simultaneousCSI_ReportsAllCC              *int64                                                           `lb:5,ub:32,optional`
	dualPA_Architecture                       *CA_ParametersNR_v1540_dualPA_Architecture                       `optional`
}

func (ie *CA_ParametersNR_v1540) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.simultaneousSRS_AssocCSI_RS_AllCC != nil, ie.csi_RS_IM_ReceptionForFeedbackPerBandComb != nil, ie.simultaneousCSI_ReportsAllCC != nil, ie.dualPA_Architecture != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.simultaneousSRS_AssocCSI_RS_AllCC != nil {
		if err = w.WriteInteger(*ie.simultaneousSRS_AssocCSI_RS_AllCC, &uper.Constraint{Lb: 5, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode simultaneousSRS_AssocCSI_RS_AllCC", err)
		}
	}
	if ie.csi_RS_IM_ReceptionForFeedbackPerBandComb != nil {
		if err = ie.csi_RS_IM_ReceptionForFeedbackPerBandComb.Encode(w); err != nil {
			return utils.WrapError("Encode csi_RS_IM_ReceptionForFeedbackPerBandComb", err)
		}
	}
	if ie.simultaneousCSI_ReportsAllCC != nil {
		if err = w.WriteInteger(*ie.simultaneousCSI_ReportsAllCC, &uper.Constraint{Lb: 5, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode simultaneousCSI_ReportsAllCC", err)
		}
	}
	if ie.dualPA_Architecture != nil {
		if err = ie.dualPA_Architecture.Encode(w); err != nil {
			return utils.WrapError("Encode dualPA_Architecture", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNR_v1540) Decode(r *uper.UperReader) error {
	var err error
	var simultaneousSRS_AssocCSI_RS_AllCCPresent bool
	if simultaneousSRS_AssocCSI_RS_AllCCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_RS_IM_ReceptionForFeedbackPerBandCombPresent bool
	if csi_RS_IM_ReceptionForFeedbackPerBandCombPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var simultaneousCSI_ReportsAllCCPresent bool
	if simultaneousCSI_ReportsAllCCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dualPA_ArchitecturePresent bool
	if dualPA_ArchitecturePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if simultaneousSRS_AssocCSI_RS_AllCCPresent {
		var tmp_int_simultaneousSRS_AssocCSI_RS_AllCC int64
		if tmp_int_simultaneousSRS_AssocCSI_RS_AllCC, err = r.ReadInteger(&uper.Constraint{Lb: 5, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode simultaneousSRS_AssocCSI_RS_AllCC", err)
		}
		ie.simultaneousSRS_AssocCSI_RS_AllCC = &tmp_int_simultaneousSRS_AssocCSI_RS_AllCC
	}
	if csi_RS_IM_ReceptionForFeedbackPerBandCombPresent {
		ie.csi_RS_IM_ReceptionForFeedbackPerBandComb = new(CA_ParametersNR_v1540_csi_RS_IM_ReceptionForFeedbackPerBandComb)
		if err = ie.csi_RS_IM_ReceptionForFeedbackPerBandComb.Decode(r); err != nil {
			return utils.WrapError("Decode csi_RS_IM_ReceptionForFeedbackPerBandComb", err)
		}
	}
	if simultaneousCSI_ReportsAllCCPresent {
		var tmp_int_simultaneousCSI_ReportsAllCC int64
		if tmp_int_simultaneousCSI_ReportsAllCC, err = r.ReadInteger(&uper.Constraint{Lb: 5, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode simultaneousCSI_ReportsAllCC", err)
		}
		ie.simultaneousCSI_ReportsAllCC = &tmp_int_simultaneousCSI_ReportsAllCC
	}
	if dualPA_ArchitecturePresent {
		ie.dualPA_Architecture = new(CA_ParametersNR_v1540_dualPA_Architecture)
		if err = ie.dualPA_Architecture.Decode(r); err != nil {
			return utils.WrapError("Decode dualPA_Architecture", err)
		}
	}
	return nil
}
