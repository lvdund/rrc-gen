package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_SemiPersistentOnPUSCH_TriggerState struct {
	associatedReportConfigInfo  CSI_ReportConfigId                                                  `madatory`
	sp_CSI_MultiplexingMode_r17 *CSI_SemiPersistentOnPUSCH_TriggerState_sp_CSI_MultiplexingMode_r17 `optional,ext-1`
}

func (ie *CSI_SemiPersistentOnPUSCH_TriggerState) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.sp_CSI_MultiplexingMode_r17 != nil
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.associatedReportConfigInfo.Encode(w); err != nil {
		return utils.WrapError("Encode associatedReportConfigInfo", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.sp_CSI_MultiplexingMode_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CSI_SemiPersistentOnPUSCH_TriggerState", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.sp_CSI_MultiplexingMode_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sp_CSI_MultiplexingMode_r17 optional
			if ie.sp_CSI_MultiplexingMode_r17 != nil {
				if err = ie.sp_CSI_MultiplexingMode_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sp_CSI_MultiplexingMode_r17", err)
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

func (ie *CSI_SemiPersistentOnPUSCH_TriggerState) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.associatedReportConfigInfo.Decode(r); err != nil {
		return utils.WrapError("Decode associatedReportConfigInfo", err)
	}

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

			sp_CSI_MultiplexingMode_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sp_CSI_MultiplexingMode_r17 optional
			if sp_CSI_MultiplexingMode_r17Present {
				ie.sp_CSI_MultiplexingMode_r17 = new(CSI_SemiPersistentOnPUSCH_TriggerState_sp_CSI_MultiplexingMode_r17)
				if err = ie.sp_CSI_MultiplexingMode_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sp_CSI_MultiplexingMode_r17", err)
				}
			}
		}
	}
	return nil
}
