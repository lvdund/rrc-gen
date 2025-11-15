package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_MultiTRP_SupportedCombinations_r17 struct {
	maxNumTx_Ports_r17                CSI_MultiTRP_SupportedCombinations_r17_maxNumTx_Ports_r17 `madatory`
	maxTotalNumCMR_r17                int64                                                     `lb:2,ub:64,madatory`
	maxTotalNumTx_PortsNZP_CSI_RS_r17 int64                                                     `lb:2,ub:256,madatory`
}

func (ie *CSI_MultiTRP_SupportedCombinations_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.maxNumTx_Ports_r17.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumTx_Ports_r17", err)
	}
	if err = w.WriteInteger(ie.maxTotalNumCMR_r17, &uper.Constraint{Lb: 2, Ub: 64}, false); err != nil {
		return utils.WrapError("WriteInteger maxTotalNumCMR_r17", err)
	}
	if err = w.WriteInteger(ie.maxTotalNumTx_PortsNZP_CSI_RS_r17, &uper.Constraint{Lb: 2, Ub: 256}, false); err != nil {
		return utils.WrapError("WriteInteger maxTotalNumTx_PortsNZP_CSI_RS_r17", err)
	}
	return nil
}

func (ie *CSI_MultiTRP_SupportedCombinations_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.maxNumTx_Ports_r17.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumTx_Ports_r17", err)
	}
	var tmp_int_maxTotalNumCMR_r17 int64
	if tmp_int_maxTotalNumCMR_r17, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 64}, false); err != nil {
		return utils.WrapError("ReadInteger maxTotalNumCMR_r17", err)
	}
	ie.maxTotalNumCMR_r17 = tmp_int_maxTotalNumCMR_r17
	var tmp_int_maxTotalNumTx_PortsNZP_CSI_RS_r17 int64
	if tmp_int_maxTotalNumTx_PortsNZP_CSI_RS_r17, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 256}, false); err != nil {
		return utils.WrapError("ReadInteger maxTotalNumTx_PortsNZP_CSI_RS_r17", err)
	}
	ie.maxTotalNumTx_PortsNZP_CSI_RS_r17 = tmp_int_maxTotalNumTx_PortsNZP_CSI_RS_r17
	return nil
}
