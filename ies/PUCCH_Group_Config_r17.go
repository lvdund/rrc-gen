package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_Group_Config_r17 struct {
	fr1_FR1_NonSharedTDD_r17 *PUCCH_Group_Config_r17_fr1_FR1_NonSharedTDD_r17 `optional`
	fr2_FR2_NonSharedTDD_r17 *PUCCH_Group_Config_r17_fr2_FR2_NonSharedTDD_r17 `optional`
	fr1_FR2_NonSharedTDD_r17 *PUCCH_Group_Config_r17_fr1_FR2_NonSharedTDD_r17 `optional`
}

func (ie *PUCCH_Group_Config_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.fr1_FR1_NonSharedTDD_r17 != nil, ie.fr2_FR2_NonSharedTDD_r17 != nil, ie.fr1_FR2_NonSharedTDD_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.fr1_FR1_NonSharedTDD_r17 != nil {
		if err = ie.fr1_FR1_NonSharedTDD_r17.Encode(w); err != nil {
			return utils.WrapError("Encode fr1_FR1_NonSharedTDD_r17", err)
		}
	}
	if ie.fr2_FR2_NonSharedTDD_r17 != nil {
		if err = ie.fr2_FR2_NonSharedTDD_r17.Encode(w); err != nil {
			return utils.WrapError("Encode fr2_FR2_NonSharedTDD_r17", err)
		}
	}
	if ie.fr1_FR2_NonSharedTDD_r17 != nil {
		if err = ie.fr1_FR2_NonSharedTDD_r17.Encode(w); err != nil {
			return utils.WrapError("Encode fr1_FR2_NonSharedTDD_r17", err)
		}
	}
	return nil
}

func (ie *PUCCH_Group_Config_r17) Decode(r *uper.UperReader) error {
	var err error
	var fr1_FR1_NonSharedTDD_r17Present bool
	if fr1_FR1_NonSharedTDD_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var fr2_FR2_NonSharedTDD_r17Present bool
	if fr2_FR2_NonSharedTDD_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var fr1_FR2_NonSharedTDD_r17Present bool
	if fr1_FR2_NonSharedTDD_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if fr1_FR1_NonSharedTDD_r17Present {
		ie.fr1_FR1_NonSharedTDD_r17 = new(PUCCH_Group_Config_r17_fr1_FR1_NonSharedTDD_r17)
		if err = ie.fr1_FR1_NonSharedTDD_r17.Decode(r); err != nil {
			return utils.WrapError("Decode fr1_FR1_NonSharedTDD_r17", err)
		}
	}
	if fr2_FR2_NonSharedTDD_r17Present {
		ie.fr2_FR2_NonSharedTDD_r17 = new(PUCCH_Group_Config_r17_fr2_FR2_NonSharedTDD_r17)
		if err = ie.fr2_FR2_NonSharedTDD_r17.Decode(r); err != nil {
			return utils.WrapError("Decode fr2_FR2_NonSharedTDD_r17", err)
		}
	}
	if fr1_FR2_NonSharedTDD_r17Present {
		ie.fr1_FR2_NonSharedTDD_r17 = new(PUCCH_Group_Config_r17_fr1_FR2_NonSharedTDD_r17)
		if err = ie.fr1_FR2_NonSharedTDD_r17.Decode(r); err != nil {
			return utils.WrapError("Decode fr1_FR2_NonSharedTDD_r17", err)
		}
	}
	return nil
}
