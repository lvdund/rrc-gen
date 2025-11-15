package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MobilityFromNRCommand_v1610_IEs struct {
	voiceFallbackIndication_r16 *MobilityFromNRCommand_v1610_IEs_voiceFallbackIndication_r16 `optional`
	nonCriticalExtension        interface{}                                                  `optional`
}

func (ie *MobilityFromNRCommand_v1610_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.voiceFallbackIndication_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.voiceFallbackIndication_r16 != nil {
		if err = ie.voiceFallbackIndication_r16.Encode(w); err != nil {
			return utils.WrapError("Encode voiceFallbackIndication_r16", err)
		}
	}
	return nil
}

func (ie *MobilityFromNRCommand_v1610_IEs) Decode(r *uper.UperReader) error {
	var err error
	var voiceFallbackIndication_r16Present bool
	if voiceFallbackIndication_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if voiceFallbackIndication_r16Present {
		ie.voiceFallbackIndication_r16 = new(MobilityFromNRCommand_v1610_IEs_voiceFallbackIndication_r16)
		if err = ie.voiceFallbackIndication_r16.Decode(r); err != nil {
			return utils.WrapError("Decode voiceFallbackIndication_r16", err)
		}
	}
	return nil
}
