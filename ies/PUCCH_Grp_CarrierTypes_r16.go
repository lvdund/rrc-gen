package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_Grp_CarrierTypes_r16 struct {
	fr1_NonSharedTDD_r16 *PUCCH_Grp_CarrierTypes_r16_fr1_NonSharedTDD_r16 `optional`
	fr1_SharedTDD_r16    *PUCCH_Grp_CarrierTypes_r16_fr1_SharedTDD_r16    `optional`
	fr1_NonSharedFDD_r16 *PUCCH_Grp_CarrierTypes_r16_fr1_NonSharedFDD_r16 `optional`
	fr2_r16              *PUCCH_Grp_CarrierTypes_r16_fr2_r16              `optional`
}

func (ie *PUCCH_Grp_CarrierTypes_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.fr1_NonSharedTDD_r16 != nil, ie.fr1_SharedTDD_r16 != nil, ie.fr1_NonSharedFDD_r16 != nil, ie.fr2_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.fr1_NonSharedTDD_r16 != nil {
		if err = ie.fr1_NonSharedTDD_r16.Encode(w); err != nil {
			return utils.WrapError("Encode fr1_NonSharedTDD_r16", err)
		}
	}
	if ie.fr1_SharedTDD_r16 != nil {
		if err = ie.fr1_SharedTDD_r16.Encode(w); err != nil {
			return utils.WrapError("Encode fr1_SharedTDD_r16", err)
		}
	}
	if ie.fr1_NonSharedFDD_r16 != nil {
		if err = ie.fr1_NonSharedFDD_r16.Encode(w); err != nil {
			return utils.WrapError("Encode fr1_NonSharedFDD_r16", err)
		}
	}
	if ie.fr2_r16 != nil {
		if err = ie.fr2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode fr2_r16", err)
		}
	}
	return nil
}

func (ie *PUCCH_Grp_CarrierTypes_r16) Decode(r *uper.UperReader) error {
	var err error
	var fr1_NonSharedTDD_r16Present bool
	if fr1_NonSharedTDD_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var fr1_SharedTDD_r16Present bool
	if fr1_SharedTDD_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var fr1_NonSharedFDD_r16Present bool
	if fr1_NonSharedFDD_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var fr2_r16Present bool
	if fr2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if fr1_NonSharedTDD_r16Present {
		ie.fr1_NonSharedTDD_r16 = new(PUCCH_Grp_CarrierTypes_r16_fr1_NonSharedTDD_r16)
		if err = ie.fr1_NonSharedTDD_r16.Decode(r); err != nil {
			return utils.WrapError("Decode fr1_NonSharedTDD_r16", err)
		}
	}
	if fr1_SharedTDD_r16Present {
		ie.fr1_SharedTDD_r16 = new(PUCCH_Grp_CarrierTypes_r16_fr1_SharedTDD_r16)
		if err = ie.fr1_SharedTDD_r16.Decode(r); err != nil {
			return utils.WrapError("Decode fr1_SharedTDD_r16", err)
		}
	}
	if fr1_NonSharedFDD_r16Present {
		ie.fr1_NonSharedFDD_r16 = new(PUCCH_Grp_CarrierTypes_r16_fr1_NonSharedFDD_r16)
		if err = ie.fr1_NonSharedFDD_r16.Decode(r); err != nil {
			return utils.WrapError("Decode fr1_NonSharedFDD_r16", err)
		}
	}
	if fr2_r16Present {
		ie.fr2_r16 = new(PUCCH_Grp_CarrierTypes_r16_fr2_r16)
		if err = ie.fr2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode fr2_r16", err)
		}
	}
	return nil
}
