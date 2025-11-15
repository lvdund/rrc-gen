package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_codebookType_type2 struct {
	subType           *CodebookConfig_codebookType_type2_subType          `lb:16,ub:16,optional`
	phaseAlphabetSize CodebookConfig_codebookType_type2_phaseAlphabetSize `madatory`
	subbandAmplitude  bool                                                `madatory`
	numberOfBeams     CodebookConfig_codebookType_type2_numberOfBeams     `madatory`
}

func (ie *CodebookConfig_codebookType_type2) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.subType != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.subType != nil {
		if err = ie.subType.Encode(w); err != nil {
			return utils.WrapError("Encode subType", err)
		}
	}
	if err = ie.phaseAlphabetSize.Encode(w); err != nil {
		return utils.WrapError("Encode phaseAlphabetSize", err)
	}
	if err = w.WriteBoolean(ie.subbandAmplitude); err != nil {
		return utils.WrapError("WriteBoolean subbandAmplitude", err)
	}
	if err = ie.numberOfBeams.Encode(w); err != nil {
		return utils.WrapError("Encode numberOfBeams", err)
	}
	return nil
}

func (ie *CodebookConfig_codebookType_type2) Decode(r *uper.UperReader) error {
	var err error
	var subTypePresent bool
	if subTypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if subTypePresent {
		ie.subType = new(CodebookConfig_codebookType_type2_subType)
		if err = ie.subType.Decode(r); err != nil {
			return utils.WrapError("Decode subType", err)
		}
	}
	if err = ie.phaseAlphabetSize.Decode(r); err != nil {
		return utils.WrapError("Decode phaseAlphabetSize", err)
	}
	var tmp_bool_subbandAmplitude bool
	if tmp_bool_subbandAmplitude, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean subbandAmplitude", err)
	}
	ie.subbandAmplitude = tmp_bool_subbandAmplitude
	if err = ie.numberOfBeams.Decode(r); err != nil {
		return utils.WrapError("Decode numberOfBeams", err)
	}
	return nil
}
