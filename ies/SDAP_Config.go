package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SDAP_Config struct {
	pdu_Session              PDU_SessionID             `madatory`
	sdap_HeaderDL            SDAP_Config_sdap_HeaderDL `madatory`
	sdap_HeaderUL            SDAP_Config_sdap_HeaderUL `madatory`
	defaultDRB               bool                      `madatory`
	mappedQoS_FlowsToAdd     []QFI                     `lb:1,ub:maxNrofQFIs,optional`
	mappedQoS_FlowsToRelease []QFI                     `lb:1,ub:maxNrofQFIs,optional`
}

func (ie *SDAP_Config) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.mappedQoS_FlowsToAdd) > 0, len(ie.mappedQoS_FlowsToRelease) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.pdu_Session.Encode(w); err != nil {
		return utils.WrapError("Encode pdu_Session", err)
	}
	if err = ie.sdap_HeaderDL.Encode(w); err != nil {
		return utils.WrapError("Encode sdap_HeaderDL", err)
	}
	if err = ie.sdap_HeaderUL.Encode(w); err != nil {
		return utils.WrapError("Encode sdap_HeaderUL", err)
	}
	if err = w.WriteBoolean(ie.defaultDRB); err != nil {
		return utils.WrapError("WriteBoolean defaultDRB", err)
	}
	if len(ie.mappedQoS_FlowsToAdd) > 0 {
		tmp_mappedQoS_FlowsToAdd := utils.NewSequence[*QFI]([]*QFI{}, uper.Constraint{Lb: 1, Ub: maxNrofQFIs}, false)
		for _, i := range ie.mappedQoS_FlowsToAdd {
			tmp_mappedQoS_FlowsToAdd.Value = append(tmp_mappedQoS_FlowsToAdd.Value, &i)
		}
		if err = tmp_mappedQoS_FlowsToAdd.Encode(w); err != nil {
			return utils.WrapError("Encode mappedQoS_FlowsToAdd", err)
		}
	}
	if len(ie.mappedQoS_FlowsToRelease) > 0 {
		tmp_mappedQoS_FlowsToRelease := utils.NewSequence[*QFI]([]*QFI{}, uper.Constraint{Lb: 1, Ub: maxNrofQFIs}, false)
		for _, i := range ie.mappedQoS_FlowsToRelease {
			tmp_mappedQoS_FlowsToRelease.Value = append(tmp_mappedQoS_FlowsToRelease.Value, &i)
		}
		if err = tmp_mappedQoS_FlowsToRelease.Encode(w); err != nil {
			return utils.WrapError("Encode mappedQoS_FlowsToRelease", err)
		}
	}
	return nil
}

func (ie *SDAP_Config) Decode(r *uper.UperReader) error {
	var err error
	var mappedQoS_FlowsToAddPresent bool
	if mappedQoS_FlowsToAddPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var mappedQoS_FlowsToReleasePresent bool
	if mappedQoS_FlowsToReleasePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.pdu_Session.Decode(r); err != nil {
		return utils.WrapError("Decode pdu_Session", err)
	}
	if err = ie.sdap_HeaderDL.Decode(r); err != nil {
		return utils.WrapError("Decode sdap_HeaderDL", err)
	}
	if err = ie.sdap_HeaderUL.Decode(r); err != nil {
		return utils.WrapError("Decode sdap_HeaderUL", err)
	}
	var tmp_bool_defaultDRB bool
	if tmp_bool_defaultDRB, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean defaultDRB", err)
	}
	ie.defaultDRB = tmp_bool_defaultDRB
	if mappedQoS_FlowsToAddPresent {
		tmp_mappedQoS_FlowsToAdd := utils.NewSequence[*QFI]([]*QFI{}, uper.Constraint{Lb: 1, Ub: maxNrofQFIs}, false)
		fn_mappedQoS_FlowsToAdd := func() *QFI {
			return new(QFI)
		}
		if err = tmp_mappedQoS_FlowsToAdd.Decode(r, fn_mappedQoS_FlowsToAdd); err != nil {
			return utils.WrapError("Decode mappedQoS_FlowsToAdd", err)
		}
		ie.mappedQoS_FlowsToAdd = []QFI{}
		for _, i := range tmp_mappedQoS_FlowsToAdd.Value {
			ie.mappedQoS_FlowsToAdd = append(ie.mappedQoS_FlowsToAdd, *i)
		}
	}
	if mappedQoS_FlowsToReleasePresent {
		tmp_mappedQoS_FlowsToRelease := utils.NewSequence[*QFI]([]*QFI{}, uper.Constraint{Lb: 1, Ub: maxNrofQFIs}, false)
		fn_mappedQoS_FlowsToRelease := func() *QFI {
			return new(QFI)
		}
		if err = tmp_mappedQoS_FlowsToRelease.Decode(r, fn_mappedQoS_FlowsToRelease); err != nil {
			return utils.WrapError("Decode mappedQoS_FlowsToRelease", err)
		}
		ie.mappedQoS_FlowsToRelease = []QFI{}
		for _, i := range tmp_mappedQoS_FlowsToRelease.Value {
			ie.mappedQoS_FlowsToRelease = append(ie.mappedQoS_FlowsToRelease, *i)
		}
	}
	return nil
}
