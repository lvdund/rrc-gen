package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MUSIM_Assistance_r17 struct {
	musim_PreferredRRC_State_r17 *MUSIM_Assistance_r17_musim_PreferredRRC_State_r17 `optional`
	musim_GapPreferenceList_r17  *MUSIM_GapPreferenceList_r17                       `optional`
}

func (ie *MUSIM_Assistance_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.musim_PreferredRRC_State_r17 != nil, ie.musim_GapPreferenceList_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.musim_PreferredRRC_State_r17 != nil {
		if err = ie.musim_PreferredRRC_State_r17.Encode(w); err != nil {
			return utils.WrapError("Encode musim_PreferredRRC_State_r17", err)
		}
	}
	if ie.musim_GapPreferenceList_r17 != nil {
		if err = ie.musim_GapPreferenceList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode musim_GapPreferenceList_r17", err)
		}
	}
	return nil
}

func (ie *MUSIM_Assistance_r17) Decode(r *uper.UperReader) error {
	var err error
	var musim_PreferredRRC_State_r17Present bool
	if musim_PreferredRRC_State_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var musim_GapPreferenceList_r17Present bool
	if musim_GapPreferenceList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if musim_PreferredRRC_State_r17Present {
		ie.musim_PreferredRRC_State_r17 = new(MUSIM_Assistance_r17_musim_PreferredRRC_State_r17)
		if err = ie.musim_PreferredRRC_State_r17.Decode(r); err != nil {
			return utils.WrapError("Decode musim_PreferredRRC_State_r17", err)
		}
	}
	if musim_GapPreferenceList_r17Present {
		ie.musim_GapPreferenceList_r17 = new(MUSIM_GapPreferenceList_r17)
		if err = ie.musim_GapPreferenceList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode musim_GapPreferenceList_r17", err)
		}
	}
	return nil
}
