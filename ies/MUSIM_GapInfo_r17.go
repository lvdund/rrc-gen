package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MUSIM_GapInfo_r17 struct {
	musim_Starting_SFN_AndSubframe_r17 *MUSIM_Starting_SFN_AndSubframe_r17                 `optional`
	musim_GapLength_r17                *MUSIM_GapInfo_r17_musim_GapLength_r17              `optional`
	musim_GapRepetitionAndOffset_r17   *MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17 `lb:0,ub:19,optional`
}

func (ie *MUSIM_GapInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.musim_Starting_SFN_AndSubframe_r17 != nil, ie.musim_GapLength_r17 != nil, ie.musim_GapRepetitionAndOffset_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.musim_Starting_SFN_AndSubframe_r17 != nil {
		if err = ie.musim_Starting_SFN_AndSubframe_r17.Encode(w); err != nil {
			return utils.WrapError("Encode musim_Starting_SFN_AndSubframe_r17", err)
		}
	}
	if ie.musim_GapLength_r17 != nil {
		if err = ie.musim_GapLength_r17.Encode(w); err != nil {
			return utils.WrapError("Encode musim_GapLength_r17", err)
		}
	}
	if ie.musim_GapRepetitionAndOffset_r17 != nil {
		if err = ie.musim_GapRepetitionAndOffset_r17.Encode(w); err != nil {
			return utils.WrapError("Encode musim_GapRepetitionAndOffset_r17", err)
		}
	}
	return nil
}

func (ie *MUSIM_GapInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	var musim_Starting_SFN_AndSubframe_r17Present bool
	if musim_Starting_SFN_AndSubframe_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var musim_GapLength_r17Present bool
	if musim_GapLength_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var musim_GapRepetitionAndOffset_r17Present bool
	if musim_GapRepetitionAndOffset_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if musim_Starting_SFN_AndSubframe_r17Present {
		ie.musim_Starting_SFN_AndSubframe_r17 = new(MUSIM_Starting_SFN_AndSubframe_r17)
		if err = ie.musim_Starting_SFN_AndSubframe_r17.Decode(r); err != nil {
			return utils.WrapError("Decode musim_Starting_SFN_AndSubframe_r17", err)
		}
	}
	if musim_GapLength_r17Present {
		ie.musim_GapLength_r17 = new(MUSIM_GapInfo_r17_musim_GapLength_r17)
		if err = ie.musim_GapLength_r17.Decode(r); err != nil {
			return utils.WrapError("Decode musim_GapLength_r17", err)
		}
	}
	if musim_GapRepetitionAndOffset_r17Present {
		ie.musim_GapRepetitionAndOffset_r17 = new(MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17)
		if err = ie.musim_GapRepetitionAndOffset_r17.Decode(r); err != nil {
			return utils.WrapError("Decode musim_GapRepetitionAndOffset_r17", err)
		}
	}
	return nil
}
