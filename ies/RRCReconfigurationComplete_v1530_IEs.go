package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfigurationComplete_v1530_IEs struct {
	uplinkTxDirectCurrentList *UplinkTxDirectCurrentList            `optional`
	nonCriticalExtension      *RRCReconfigurationComplete_v1560_IEs `optional`
}

func (ie *RRCReconfigurationComplete_v1530_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.uplinkTxDirectCurrentList != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.uplinkTxDirectCurrentList != nil {
		if err = ie.uplinkTxDirectCurrentList.Encode(w); err != nil {
			return utils.WrapError("Encode uplinkTxDirectCurrentList", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCReconfigurationComplete_v1530_IEs) Decode(r *uper.UperReader) error {
	var err error
	var uplinkTxDirectCurrentListPresent bool
	if uplinkTxDirectCurrentListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if uplinkTxDirectCurrentListPresent {
		ie.uplinkTxDirectCurrentList = new(UplinkTxDirectCurrentList)
		if err = ie.uplinkTxDirectCurrentList.Decode(r); err != nil {
			return utils.WrapError("Decode uplinkTxDirectCurrentList", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(RRCReconfigurationComplete_v1560_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
