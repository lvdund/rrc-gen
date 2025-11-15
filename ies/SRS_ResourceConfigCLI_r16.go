package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_ResourceConfigCLI_r16 struct {
	srs_Resource_r16     SRS_Resource      `madatory`
	srs_SCS_r16          SubcarrierSpacing `madatory`
	refServCellIndex_r16 *ServCellIndex    `optional`
	refBWP_r16           BWP_Id            `madatory`
}

func (ie *SRS_ResourceConfigCLI_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.refServCellIndex_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.srs_Resource_r16.Encode(w); err != nil {
		return utils.WrapError("Encode srs_Resource_r16", err)
	}
	if err = ie.srs_SCS_r16.Encode(w); err != nil {
		return utils.WrapError("Encode srs_SCS_r16", err)
	}
	if ie.refServCellIndex_r16 != nil {
		if err = ie.refServCellIndex_r16.Encode(w); err != nil {
			return utils.WrapError("Encode refServCellIndex_r16", err)
		}
	}
	if err = ie.refBWP_r16.Encode(w); err != nil {
		return utils.WrapError("Encode refBWP_r16", err)
	}
	return nil
}

func (ie *SRS_ResourceConfigCLI_r16) Decode(r *uper.UperReader) error {
	var err error
	var refServCellIndex_r16Present bool
	if refServCellIndex_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.srs_Resource_r16.Decode(r); err != nil {
		return utils.WrapError("Decode srs_Resource_r16", err)
	}
	if err = ie.srs_SCS_r16.Decode(r); err != nil {
		return utils.WrapError("Decode srs_SCS_r16", err)
	}
	if refServCellIndex_r16Present {
		ie.refServCellIndex_r16 = new(ServCellIndex)
		if err = ie.refServCellIndex_r16.Decode(r); err != nil {
			return utils.WrapError("Decode refServCellIndex_r16", err)
		}
	}
	if err = ie.refBWP_r16.Decode(r); err != nil {
		return utils.WrapError("Decode refBWP_r16", err)
	}
	return nil
}
