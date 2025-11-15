package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TAG struct {
	tag_Id             TAG_Id             `madatory`
	timeAlignmentTimer TimeAlignmentTimer `madatory`
}

func (ie *TAG) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.tag_Id.Encode(w); err != nil {
		return utils.WrapError("Encode tag_Id", err)
	}
	if err = ie.timeAlignmentTimer.Encode(w); err != nil {
		return utils.WrapError("Encode timeAlignmentTimer", err)
	}
	return nil
}

func (ie *TAG) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.tag_Id.Decode(r); err != nil {
		return utils.WrapError("Decode tag_Id", err)
	}
	if err = ie.timeAlignmentTimer.Decode(r); err != nil {
		return utils.WrapError("Decode timeAlignmentTimer", err)
	}
	return nil
}
