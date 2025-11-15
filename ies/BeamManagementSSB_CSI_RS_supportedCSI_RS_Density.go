package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BeamManagementSSB_CSI_RS_supportedCSI_RS_Density_Enum_one         uper.Enumerated = 0
	BeamManagementSSB_CSI_RS_supportedCSI_RS_Density_Enum_three       uper.Enumerated = 1
	BeamManagementSSB_CSI_RS_supportedCSI_RS_Density_Enum_oneAndThree uper.Enumerated = 2
)

type BeamManagementSSB_CSI_RS_supportedCSI_RS_Density struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *BeamManagementSSB_CSI_RS_supportedCSI_RS_Density) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode BeamManagementSSB_CSI_RS_supportedCSI_RS_Density", err)
	}
	return nil
}

func (ie *BeamManagementSSB_CSI_RS_supportedCSI_RS_Density) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode BeamManagementSSB_CSI_RS_supportedCSI_RS_Density", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
