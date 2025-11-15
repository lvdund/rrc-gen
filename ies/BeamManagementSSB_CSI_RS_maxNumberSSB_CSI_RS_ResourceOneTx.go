package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BeamManagementSSB_CSI_RS_maxNumberSSB_CSI_RS_ResourceOneTx_Enum_n0  uper.Enumerated = 0
	BeamManagementSSB_CSI_RS_maxNumberSSB_CSI_RS_ResourceOneTx_Enum_n8  uper.Enumerated = 1
	BeamManagementSSB_CSI_RS_maxNumberSSB_CSI_RS_ResourceOneTx_Enum_n16 uper.Enumerated = 2
	BeamManagementSSB_CSI_RS_maxNumberSSB_CSI_RS_ResourceOneTx_Enum_n32 uper.Enumerated = 3
	BeamManagementSSB_CSI_RS_maxNumberSSB_CSI_RS_ResourceOneTx_Enum_n64 uper.Enumerated = 4
)

type BeamManagementSSB_CSI_RS_maxNumberSSB_CSI_RS_ResourceOneTx struct {
	Value uper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *BeamManagementSSB_CSI_RS_maxNumberSSB_CSI_RS_ResourceOneTx) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode BeamManagementSSB_CSI_RS_maxNumberSSB_CSI_RS_ResourceOneTx", err)
	}
	return nil
}

func (ie *BeamManagementSSB_CSI_RS_maxNumberSSB_CSI_RS_ResourceOneTx) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode BeamManagementSSB_CSI_RS_maxNumberSSB_CSI_RS_ResourceOneTx", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
