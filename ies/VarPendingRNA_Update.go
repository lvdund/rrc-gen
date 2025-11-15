package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VarPendingRNA_Update struct {
	pendingRNA_Update *bool `optional`
}

func (ie *VarPendingRNA_Update) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.pendingRNA_Update != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.pendingRNA_Update != nil {
		if err = w.WriteBoolean(*ie.pendingRNA_Update); err != nil {
			return utils.WrapError("Encode pendingRNA_Update", err)
		}
	}
	return nil
}

func (ie *VarPendingRNA_Update) Decode(r *uper.UperReader) error {
	var err error
	var pendingRNA_UpdatePresent bool
	if pendingRNA_UpdatePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if pendingRNA_UpdatePresent {
		var tmp_bool_pendingRNA_Update bool
		if tmp_bool_pendingRNA_Update, err = r.ReadBoolean(); err != nil {
			return utils.WrapError("Decode pendingRNA_Update", err)
		}
		ie.pendingRNA_Update = &tmp_bool_pendingRNA_Update
	}
	return nil
}
