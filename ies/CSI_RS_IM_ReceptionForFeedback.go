package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_RS_IM_ReceptionForFeedback struct {
	maxConfigNumberNZP_CSI_RS_PerCC              int64                                                      `lb:1,ub:64,madatory`
	maxConfigNumberPortsAcrossNZP_CSI_RS_PerCC   int64                                                      `lb:2,ub:256,madatory`
	maxConfigNumberCSI_IM_PerCC                  CSI_RS_IM_ReceptionForFeedback_maxConfigNumberCSI_IM_PerCC `madatory`
	maxNumberSimultaneousNZP_CSI_RS_PerCC        int64                                                      `lb:1,ub:64,madatory`
	totalNumberPortsSimultaneousNZP_CSI_RS_PerCC int64                                                      `lb:2,ub:256,madatory`
}

func (ie *CSI_RS_IM_ReceptionForFeedback) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.maxConfigNumberNZP_CSI_RS_PerCC, &uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("WriteInteger maxConfigNumberNZP_CSI_RS_PerCC", err)
	}
	if err = w.WriteInteger(ie.maxConfigNumberPortsAcrossNZP_CSI_RS_PerCC, &uper.Constraint{Lb: 2, Ub: 256}, false); err != nil {
		return utils.WrapError("WriteInteger maxConfigNumberPortsAcrossNZP_CSI_RS_PerCC", err)
	}
	if err = ie.maxConfigNumberCSI_IM_PerCC.Encode(w); err != nil {
		return utils.WrapError("Encode maxConfigNumberCSI_IM_PerCC", err)
	}
	if err = w.WriteInteger(ie.maxNumberSimultaneousNZP_CSI_RS_PerCC, &uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberSimultaneousNZP_CSI_RS_PerCC", err)
	}
	if err = w.WriteInteger(ie.totalNumberPortsSimultaneousNZP_CSI_RS_PerCC, &uper.Constraint{Lb: 2, Ub: 256}, false); err != nil {
		return utils.WrapError("WriteInteger totalNumberPortsSimultaneousNZP_CSI_RS_PerCC", err)
	}
	return nil
}

func (ie *CSI_RS_IM_ReceptionForFeedback) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_maxConfigNumberNZP_CSI_RS_PerCC int64
	if tmp_int_maxConfigNumberNZP_CSI_RS_PerCC, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("ReadInteger maxConfigNumberNZP_CSI_RS_PerCC", err)
	}
	ie.maxConfigNumberNZP_CSI_RS_PerCC = tmp_int_maxConfigNumberNZP_CSI_RS_PerCC
	var tmp_int_maxConfigNumberPortsAcrossNZP_CSI_RS_PerCC int64
	if tmp_int_maxConfigNumberPortsAcrossNZP_CSI_RS_PerCC, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 256}, false); err != nil {
		return utils.WrapError("ReadInteger maxConfigNumberPortsAcrossNZP_CSI_RS_PerCC", err)
	}
	ie.maxConfigNumberPortsAcrossNZP_CSI_RS_PerCC = tmp_int_maxConfigNumberPortsAcrossNZP_CSI_RS_PerCC
	if err = ie.maxConfigNumberCSI_IM_PerCC.Decode(r); err != nil {
		return utils.WrapError("Decode maxConfigNumberCSI_IM_PerCC", err)
	}
	var tmp_int_maxNumberSimultaneousNZP_CSI_RS_PerCC int64
	if tmp_int_maxNumberSimultaneousNZP_CSI_RS_PerCC, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberSimultaneousNZP_CSI_RS_PerCC", err)
	}
	ie.maxNumberSimultaneousNZP_CSI_RS_PerCC = tmp_int_maxNumberSimultaneousNZP_CSI_RS_PerCC
	var tmp_int_totalNumberPortsSimultaneousNZP_CSI_RS_PerCC int64
	if tmp_int_totalNumberPortsSimultaneousNZP_CSI_RS_PerCC, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 256}, false); err != nil {
		return utils.WrapError("ReadInteger totalNumberPortsSimultaneousNZP_CSI_RS_PerCC", err)
	}
	ie.totalNumberPortsSimultaneousNZP_CSI_RS_PerCC = tmp_int_totalNumberPortsSimultaneousNZP_CSI_RS_PerCC
	return nil
}
