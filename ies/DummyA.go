package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DummyA struct {
	maxNumberNZP_CSI_RS_PerCC                       int64                                                  `lb:1,ub:32,madatory`
	maxNumberPortsAcrossNZP_CSI_RS_PerCC            DummyA_maxNumberPortsAcrossNZP_CSI_RS_PerCC            `madatory`
	maxNumberCS_IM_PerCC                            DummyA_maxNumberCS_IM_PerCC                            `madatory`
	maxNumberSimultaneousCSI_RS_ActBWP_AllCC        DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC        `madatory`
	totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC `madatory`
}

func (ie *DummyA) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.maxNumberNZP_CSI_RS_PerCC, &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberNZP_CSI_RS_PerCC", err)
	}
	if err = ie.maxNumberPortsAcrossNZP_CSI_RS_PerCC.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberPortsAcrossNZP_CSI_RS_PerCC", err)
	}
	if err = ie.maxNumberCS_IM_PerCC.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberCS_IM_PerCC", err)
	}
	if err = ie.maxNumberSimultaneousCSI_RS_ActBWP_AllCC.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberSimultaneousCSI_RS_ActBWP_AllCC", err)
	}
	if err = ie.totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC.Encode(w); err != nil {
		return utils.WrapError("Encode totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC", err)
	}
	return nil
}

func (ie *DummyA) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_maxNumberNZP_CSI_RS_PerCC int64
	if tmp_int_maxNumberNZP_CSI_RS_PerCC, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberNZP_CSI_RS_PerCC", err)
	}
	ie.maxNumberNZP_CSI_RS_PerCC = tmp_int_maxNumberNZP_CSI_RS_PerCC
	if err = ie.maxNumberPortsAcrossNZP_CSI_RS_PerCC.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberPortsAcrossNZP_CSI_RS_PerCC", err)
	}
	if err = ie.maxNumberCS_IM_PerCC.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberCS_IM_PerCC", err)
	}
	if err = ie.maxNumberSimultaneousCSI_RS_ActBWP_AllCC.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberSimultaneousCSI_RS_ActBWP_AllCC", err)
	}
	if err = ie.totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC.Decode(r); err != nil {
		return utils.WrapError("Decode totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC", err)
	}
	return nil
}
