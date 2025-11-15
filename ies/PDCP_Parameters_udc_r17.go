package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCP_Parameters_udc_r17 struct {
	standardDictionary_r17  *PDCP_Parameters_udc_r17_standardDictionary_r17  `optional`
	operatorDictionary_r17  *PDCP_Parameters_udc_r17_operatorDictionary_r17  `lb:0,ub:15,optional`
	continueUDC_r17         *PDCP_Parameters_udc_r17_continueUDC_r17         `optional`
	supportOfBufferSize_r17 *PDCP_Parameters_udc_r17_supportOfBufferSize_r17 `optional`
}

func (ie *PDCP_Parameters_udc_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.standardDictionary_r17 != nil, ie.operatorDictionary_r17 != nil, ie.continueUDC_r17 != nil, ie.supportOfBufferSize_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.standardDictionary_r17 != nil {
		if err = ie.standardDictionary_r17.Encode(w); err != nil {
			return utils.WrapError("Encode standardDictionary_r17", err)
		}
	}
	if ie.operatorDictionary_r17 != nil {
		if err = ie.operatorDictionary_r17.Encode(w); err != nil {
			return utils.WrapError("Encode operatorDictionary_r17", err)
		}
	}
	if ie.continueUDC_r17 != nil {
		if err = ie.continueUDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode continueUDC_r17", err)
		}
	}
	if ie.supportOfBufferSize_r17 != nil {
		if err = ie.supportOfBufferSize_r17.Encode(w); err != nil {
			return utils.WrapError("Encode supportOfBufferSize_r17", err)
		}
	}
	return nil
}

func (ie *PDCP_Parameters_udc_r17) Decode(r *uper.UperReader) error {
	var err error
	var standardDictionary_r17Present bool
	if standardDictionary_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var operatorDictionary_r17Present bool
	if operatorDictionary_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var continueUDC_r17Present bool
	if continueUDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportOfBufferSize_r17Present bool
	if supportOfBufferSize_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if standardDictionary_r17Present {
		ie.standardDictionary_r17 = new(PDCP_Parameters_udc_r17_standardDictionary_r17)
		if err = ie.standardDictionary_r17.Decode(r); err != nil {
			return utils.WrapError("Decode standardDictionary_r17", err)
		}
	}
	if operatorDictionary_r17Present {
		ie.operatorDictionary_r17 = new(PDCP_Parameters_udc_r17_operatorDictionary_r17)
		if err = ie.operatorDictionary_r17.Decode(r); err != nil {
			return utils.WrapError("Decode operatorDictionary_r17", err)
		}
	}
	if continueUDC_r17Present {
		ie.continueUDC_r17 = new(PDCP_Parameters_udc_r17_continueUDC_r17)
		if err = ie.continueUDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode continueUDC_r17", err)
		}
	}
	if supportOfBufferSize_r17Present {
		ie.supportOfBufferSize_r17 = new(PDCP_Parameters_udc_r17_supportOfBufferSize_r17)
		if err = ie.supportOfBufferSize_r17.Decode(r); err != nil {
			return utils.WrapError("Decode supportOfBufferSize_r17", err)
		}
	}
	return nil
}
