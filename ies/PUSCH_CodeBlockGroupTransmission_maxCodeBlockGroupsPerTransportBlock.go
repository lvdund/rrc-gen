package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock_Enum_n2 uper.Enumerated = 0
	PUSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock_Enum_n4 uper.Enumerated = 1
	PUSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock_Enum_n6 uper.Enumerated = 2
	PUSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock_Enum_n8 uper.Enumerated = 3
)

type PUSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *PUSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode PUSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock", err)
	}
	return nil
}

func (ie *PUSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode PUSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
