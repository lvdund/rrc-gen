package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MulticastRLC_BearerConfig_r17 struct {
	servedMBS_RadioBearer_r17 MRB_Identity_r17                                `madatory`
	isPTM_Entity_r17          *MulticastRLC_BearerConfig_r17_isPTM_Entity_r17 `optional`
}

func (ie *MulticastRLC_BearerConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.isPTM_Entity_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.servedMBS_RadioBearer_r17.Encode(w); err != nil {
		return utils.WrapError("Encode servedMBS_RadioBearer_r17", err)
	}
	if ie.isPTM_Entity_r17 != nil {
		if err = ie.isPTM_Entity_r17.Encode(w); err != nil {
			return utils.WrapError("Encode isPTM_Entity_r17", err)
		}
	}
	return nil
}

func (ie *MulticastRLC_BearerConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	var isPTM_Entity_r17Present bool
	if isPTM_Entity_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.servedMBS_RadioBearer_r17.Decode(r); err != nil {
		return utils.WrapError("Decode servedMBS_RadioBearer_r17", err)
	}
	if isPTM_Entity_r17Present {
		ie.isPTM_Entity_r17 = new(MulticastRLC_BearerConfig_r17_isPTM_Entity_r17)
		if err = ie.isPTM_Entity_r17.Decode(r); err != nil {
			return utils.WrapError("Decode isPTM_Entity_r17", err)
		}
	}
	return nil
}
