package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CC_Group_r17 struct {
	servCellIndexLower_r17  ServCellIndex                     `madatory`
	servCellIndexHigher_r17 *ServCellIndex                    `optional`
	defaultDC_Location_r17  DefaultDC_Location_r17            `madatory`
	offsetToDefault_r17     *CC_Group_r17_offsetToDefault_r17 `lb:1,ub:maxNrofReqComDC_Location_r17,optional`
}

func (ie *CC_Group_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.servCellIndexHigher_r17 != nil, ie.offsetToDefault_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.servCellIndexLower_r17.Encode(w); err != nil {
		return utils.WrapError("Encode servCellIndexLower_r17", err)
	}
	if ie.servCellIndexHigher_r17 != nil {
		if err = ie.servCellIndexHigher_r17.Encode(w); err != nil {
			return utils.WrapError("Encode servCellIndexHigher_r17", err)
		}
	}
	if err = ie.defaultDC_Location_r17.Encode(w); err != nil {
		return utils.WrapError("Encode defaultDC_Location_r17", err)
	}
	if ie.offsetToDefault_r17 != nil {
		if err = ie.offsetToDefault_r17.Encode(w); err != nil {
			return utils.WrapError("Encode offsetToDefault_r17", err)
		}
	}
	return nil
}

func (ie *CC_Group_r17) Decode(r *uper.UperReader) error {
	var err error
	var servCellIndexHigher_r17Present bool
	if servCellIndexHigher_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var offsetToDefault_r17Present bool
	if offsetToDefault_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.servCellIndexLower_r17.Decode(r); err != nil {
		return utils.WrapError("Decode servCellIndexLower_r17", err)
	}
	if servCellIndexHigher_r17Present {
		ie.servCellIndexHigher_r17 = new(ServCellIndex)
		if err = ie.servCellIndexHigher_r17.Decode(r); err != nil {
			return utils.WrapError("Decode servCellIndexHigher_r17", err)
		}
	}
	if err = ie.defaultDC_Location_r17.Decode(r); err != nil {
		return utils.WrapError("Decode defaultDC_Location_r17", err)
	}
	if offsetToDefault_r17Present {
		ie.offsetToDefault_r17 = new(CC_Group_r17_offsetToDefault_r17)
		if err = ie.offsetToDefault_r17.Decode(r); err != nil {
			return utils.WrapError("Decode offsetToDefault_r17", err)
		}
	}
	return nil
}
