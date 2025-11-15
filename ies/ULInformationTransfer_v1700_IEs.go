package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ULInformationTransfer_v1700_IEs struct {
	dedicatedInfoF1c_r17 *DedicatedInfoF1c_r17 `optional`
	nonCriticalExtension interface{}           `optional`
}

func (ie *ULInformationTransfer_v1700_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dedicatedInfoF1c_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dedicatedInfoF1c_r17 != nil {
		if err = ie.dedicatedInfoF1c_r17.Encode(w); err != nil {
			return utils.WrapError("Encode dedicatedInfoF1c_r17", err)
		}
	}
	return nil
}

func (ie *ULInformationTransfer_v1700_IEs) Decode(r *uper.UperReader) error {
	var err error
	var dedicatedInfoF1c_r17Present bool
	if dedicatedInfoF1c_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if dedicatedInfoF1c_r17Present {
		ie.dedicatedInfoF1c_r17 = new(DedicatedInfoF1c_r17)
		if err = ie.dedicatedInfoF1c_r17.Decode(r); err != nil {
			return utils.WrapError("Decode dedicatedInfoF1c_r17", err)
		}
	}
	return nil
}
