package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_ConfigInfo_v1730_IEs struct {
	fr1_Carriers_MCG_r17 *int64      `lb:1,ub:32,optional`
	fr2_Carriers_MCG_r17 *int64      `lb:1,ub:32,optional`
	nonCriticalExtension interface{} `optional`
}

func (ie *CG_ConfigInfo_v1730_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.fr1_Carriers_MCG_r17 != nil, ie.fr2_Carriers_MCG_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.fr1_Carriers_MCG_r17 != nil {
		if err = w.WriteInteger(*ie.fr1_Carriers_MCG_r17, &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode fr1_Carriers_MCG_r17", err)
		}
	}
	if ie.fr2_Carriers_MCG_r17 != nil {
		if err = w.WriteInteger(*ie.fr2_Carriers_MCG_r17, &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode fr2_Carriers_MCG_r17", err)
		}
	}
	return nil
}

func (ie *CG_ConfigInfo_v1730_IEs) Decode(r *uper.UperReader) error {
	var err error
	var fr1_Carriers_MCG_r17Present bool
	if fr1_Carriers_MCG_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var fr2_Carriers_MCG_r17Present bool
	if fr2_Carriers_MCG_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if fr1_Carriers_MCG_r17Present {
		var tmp_int_fr1_Carriers_MCG_r17 int64
		if tmp_int_fr1_Carriers_MCG_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode fr1_Carriers_MCG_r17", err)
		}
		ie.fr1_Carriers_MCG_r17 = &tmp_int_fr1_Carriers_MCG_r17
	}
	if fr2_Carriers_MCG_r17Present {
		var tmp_int_fr2_Carriers_MCG_r17 int64
		if tmp_int_fr2_Carriers_MCG_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode fr2_Carriers_MCG_r17", err)
		}
		ie.fr2_Carriers_MCG_r17 = &tmp_int_fr2_Carriers_MCG_r17
	}
	return nil
}
