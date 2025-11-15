package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCRelease_v1710_IEs struct {
	noLastCellUpdate_r17 *RRCRelease_v1710_IEs_noLastCellUpdate_r17 `optional`
	nonCriticalExtension interface{}                                `optional`
}

func (ie *RRCRelease_v1710_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.noLastCellUpdate_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.noLastCellUpdate_r17 != nil {
		if err = ie.noLastCellUpdate_r17.Encode(w); err != nil {
			return utils.WrapError("Encode noLastCellUpdate_r17", err)
		}
	}
	return nil
}

func (ie *RRCRelease_v1710_IEs) Decode(r *uper.UperReader) error {
	var err error
	var noLastCellUpdate_r17Present bool
	if noLastCellUpdate_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if noLastCellUpdate_r17Present {
		ie.noLastCellUpdate_r17 = new(RRCRelease_v1710_IEs_noLastCellUpdate_r17)
		if err = ie.noLastCellUpdate_r17.Decode(r); err != nil {
			return utils.WrapError("Decode noLastCellUpdate_r17", err)
		}
	}
	return nil
}
