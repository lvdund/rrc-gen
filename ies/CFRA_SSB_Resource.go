package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CFRA_SSB_Resource struct {
	ssb                           SSB_Index `madatory`
	ra_PreambleIndex              int64     `lb:0,ub:63,madatory`
	msgA_PUSCH_Resource_Index_r16 *int64    `lb:0,ub:3071,optional,ext-1`
}

func (ie *CFRA_SSB_Resource) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.msgA_PUSCH_Resource_Index_r16 != nil
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.ssb.Encode(w); err != nil {
		return utils.WrapError("Encode ssb", err)
	}
	if err = w.WriteInteger(ie.ra_PreambleIndex, &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("WriteInteger ra_PreambleIndex", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.msgA_PUSCH_Resource_Index_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CFRA_SSB_Resource", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.msgA_PUSCH_Resource_Index_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode msgA_PUSCH_Resource_Index_r16 optional
			if ie.msgA_PUSCH_Resource_Index_r16 != nil {
				if err = extWriter.WriteInteger(*ie.msgA_PUSCH_Resource_Index_r16, &uper.Constraint{Lb: 0, Ub: 3071}, false); err != nil {
					return utils.WrapError("Encode msgA_PUSCH_Resource_Index_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *CFRA_SSB_Resource) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.ssb.Decode(r); err != nil {
		return utils.WrapError("Decode ssb", err)
	}
	var tmp_int_ra_PreambleIndex int64
	if tmp_int_ra_PreambleIndex, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("ReadInteger ra_PreambleIndex", err)
	}
	ie.ra_PreambleIndex = tmp_int_ra_PreambleIndex

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			msgA_PUSCH_Resource_Index_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode msgA_PUSCH_Resource_Index_r16 optional
			if msgA_PUSCH_Resource_Index_r16Present {
				var tmp_int_msgA_PUSCH_Resource_Index_r16 int64
				if tmp_int_msgA_PUSCH_Resource_Index_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3071}, false); err != nil {
					return utils.WrapError("Decode msgA_PUSCH_Resource_Index_r16", err)
				}
				ie.msgA_PUSCH_Resource_Index_r16 = &tmp_int_msgA_PUSCH_Resource_Index_r16
			}
		}
	}
	return nil
}
