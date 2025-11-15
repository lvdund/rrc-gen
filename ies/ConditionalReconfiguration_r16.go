package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ConditionalReconfiguration_r16 struct {
	attemptCondReconfig_r16      *ConditionalReconfiguration_r16_attemptCondReconfig_r16 `optional`
	condReconfigToRemoveList_r16 *CondReconfigToRemoveList_r16                           `optional`
	condReconfigToAddModList_r16 *CondReconfigToAddModList_r16                           `optional`
}

func (ie *ConditionalReconfiguration_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.attemptCondReconfig_r16 != nil, ie.condReconfigToRemoveList_r16 != nil, ie.condReconfigToAddModList_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.attemptCondReconfig_r16 != nil {
		if err = ie.attemptCondReconfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode attemptCondReconfig_r16", err)
		}
	}
	if ie.condReconfigToRemoveList_r16 != nil {
		if err = ie.condReconfigToRemoveList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode condReconfigToRemoveList_r16", err)
		}
	}
	if ie.condReconfigToAddModList_r16 != nil {
		if err = ie.condReconfigToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode condReconfigToAddModList_r16", err)
		}
	}
	return nil
}

func (ie *ConditionalReconfiguration_r16) Decode(r *uper.UperReader) error {
	var err error
	var attemptCondReconfig_r16Present bool
	if attemptCondReconfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var condReconfigToRemoveList_r16Present bool
	if condReconfigToRemoveList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var condReconfigToAddModList_r16Present bool
	if condReconfigToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if attemptCondReconfig_r16Present {
		ie.attemptCondReconfig_r16 = new(ConditionalReconfiguration_r16_attemptCondReconfig_r16)
		if err = ie.attemptCondReconfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode attemptCondReconfig_r16", err)
		}
	}
	if condReconfigToRemoveList_r16Present {
		ie.condReconfigToRemoveList_r16 = new(CondReconfigToRemoveList_r16)
		if err = ie.condReconfigToRemoveList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode condReconfigToRemoveList_r16", err)
		}
	}
	if condReconfigToAddModList_r16Present {
		ie.condReconfigToAddModList_r16 = new(CondReconfigToAddModList_r16)
		if err = ie.condReconfigToAddModList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode condReconfigToAddModList_r16", err)
		}
	}
	return nil
}
