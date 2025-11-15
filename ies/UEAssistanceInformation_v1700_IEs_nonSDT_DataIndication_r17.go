package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UEAssistanceInformation_v1700_IEs_nonSDT_DataIndication_r17 struct {
	resumeCause_r17 *ResumeCause `optional`
}

func (ie *UEAssistanceInformation_v1700_IEs_nonSDT_DataIndication_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.resumeCause_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.resumeCause_r17 != nil {
		if err = ie.resumeCause_r17.Encode(w); err != nil {
			return utils.WrapError("Encode resumeCause_r17", err)
		}
	}
	return nil
}

func (ie *UEAssistanceInformation_v1700_IEs_nonSDT_DataIndication_r17) Decode(r *uper.UperReader) error {
	var err error
	var resumeCause_r17Present bool
	if resumeCause_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if resumeCause_r17Present {
		ie.resumeCause_r17 = new(ResumeCause)
		if err = ie.resumeCause_r17.Decode(r); err != nil {
			return utils.WrapError("Decode resumeCause_r17", err)
		}
	}
	return nil
}
