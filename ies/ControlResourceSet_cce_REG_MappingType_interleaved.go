package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ControlResourceSet_cce_REG_MappingType_interleaved struct {
	reg_BundleSize  ControlResourceSet_cce_REG_MappingType_interleaved_reg_BundleSize  `madatory`
	interleaverSize ControlResourceSet_cce_REG_MappingType_interleaved_interleaverSize `madatory`
	shiftIndex      *int64                                                             `lb:0,ub:maxNrofPhysicalResourceBlocks_1,optional`
}

func (ie *ControlResourceSet_cce_REG_MappingType_interleaved) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.shiftIndex != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.reg_BundleSize.Encode(w); err != nil {
		return utils.WrapError("Encode reg_BundleSize", err)
	}
	if err = ie.interleaverSize.Encode(w); err != nil {
		return utils.WrapError("Encode interleaverSize", err)
	}
	if ie.shiftIndex != nil {
		if err = w.WriteInteger(*ie.shiftIndex, &uper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
			return utils.WrapError("Encode shiftIndex", err)
		}
	}
	return nil
}

func (ie *ControlResourceSet_cce_REG_MappingType_interleaved) Decode(r *uper.UperReader) error {
	var err error
	var shiftIndexPresent bool
	if shiftIndexPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.reg_BundleSize.Decode(r); err != nil {
		return utils.WrapError("Decode reg_BundleSize", err)
	}
	if err = ie.interleaverSize.Decode(r); err != nil {
		return utils.WrapError("Decode interleaverSize", err)
	}
	if shiftIndexPresent {
		var tmp_int_shiftIndex int64
		if tmp_int_shiftIndex, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
			return utils.WrapError("Decode shiftIndex", err)
		}
		ie.shiftIndex = &tmp_int_shiftIndex
	}
	return nil
}
