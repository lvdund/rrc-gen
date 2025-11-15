package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TCI_UL_State_r17 struct {
	tci_UL_State_Id_r17        TCI_UL_State_Id_r17                  `madatory`
	servingCellId_r17          *ServCellIndex                       `optional`
	bwp_Id_r17                 *BWP_Id                              `optional`
	referenceSignal_r17        TCI_UL_State_r17_referenceSignal_r17 `madatory`
	additionalPCI_r17          *AdditionalPCIIndex_r17              `optional`
	ul_powerControl_r17        *Uplink_powerControlId_r17           `optional`
	pathlossReferenceRS_Id_r17 *PathlossReferenceRS_Id_r17          `optional`
}

func (ie *TCI_UL_State_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.servingCellId_r17 != nil, ie.bwp_Id_r17 != nil, ie.additionalPCI_r17 != nil, ie.ul_powerControl_r17 != nil, ie.pathlossReferenceRS_Id_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.tci_UL_State_Id_r17.Encode(w); err != nil {
		return utils.WrapError("Encode tci_UL_State_Id_r17", err)
	}
	if ie.servingCellId_r17 != nil {
		if err = ie.servingCellId_r17.Encode(w); err != nil {
			return utils.WrapError("Encode servingCellId_r17", err)
		}
	}
	if ie.bwp_Id_r17 != nil {
		if err = ie.bwp_Id_r17.Encode(w); err != nil {
			return utils.WrapError("Encode bwp_Id_r17", err)
		}
	}
	if err = ie.referenceSignal_r17.Encode(w); err != nil {
		return utils.WrapError("Encode referenceSignal_r17", err)
	}
	if ie.additionalPCI_r17 != nil {
		if err = ie.additionalPCI_r17.Encode(w); err != nil {
			return utils.WrapError("Encode additionalPCI_r17", err)
		}
	}
	if ie.ul_powerControl_r17 != nil {
		if err = ie.ul_powerControl_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ul_powerControl_r17", err)
		}
	}
	if ie.pathlossReferenceRS_Id_r17 != nil {
		if err = ie.pathlossReferenceRS_Id_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pathlossReferenceRS_Id_r17", err)
		}
	}
	return nil
}

func (ie *TCI_UL_State_r17) Decode(r *uper.UperReader) error {
	var err error
	var servingCellId_r17Present bool
	if servingCellId_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var bwp_Id_r17Present bool
	if bwp_Id_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var additionalPCI_r17Present bool
	if additionalPCI_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_powerControl_r17Present bool
	if ul_powerControl_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pathlossReferenceRS_Id_r17Present bool
	if pathlossReferenceRS_Id_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.tci_UL_State_Id_r17.Decode(r); err != nil {
		return utils.WrapError("Decode tci_UL_State_Id_r17", err)
	}
	if servingCellId_r17Present {
		ie.servingCellId_r17 = new(ServCellIndex)
		if err = ie.servingCellId_r17.Decode(r); err != nil {
			return utils.WrapError("Decode servingCellId_r17", err)
		}
	}
	if bwp_Id_r17Present {
		ie.bwp_Id_r17 = new(BWP_Id)
		if err = ie.bwp_Id_r17.Decode(r); err != nil {
			return utils.WrapError("Decode bwp_Id_r17", err)
		}
	}
	if err = ie.referenceSignal_r17.Decode(r); err != nil {
		return utils.WrapError("Decode referenceSignal_r17", err)
	}
	if additionalPCI_r17Present {
		ie.additionalPCI_r17 = new(AdditionalPCIIndex_r17)
		if err = ie.additionalPCI_r17.Decode(r); err != nil {
			return utils.WrapError("Decode additionalPCI_r17", err)
		}
	}
	if ul_powerControl_r17Present {
		ie.ul_powerControl_r17 = new(Uplink_powerControlId_r17)
		if err = ie.ul_powerControl_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ul_powerControl_r17", err)
		}
	}
	if pathlossReferenceRS_Id_r17Present {
		ie.pathlossReferenceRS_Id_r17 = new(PathlossReferenceRS_Id_r17)
		if err = ie.pathlossReferenceRS_Id_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pathlossReferenceRS_Id_r17", err)
		}
	}
	return nil
}
