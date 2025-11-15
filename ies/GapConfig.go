package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type GapConfig struct {
	gapOffset                 int64                           `lb:0,ub:159,madatory`
	mgl                       GapConfig_mgl                   `madatory`
	mgrp                      GapConfig_mgrp                  `madatory`
	mgta                      GapConfig_mgta                  `madatory`
	refServCellIndicator      *GapConfig_refServCellIndicator `optional,ext-1`
	refFR2ServCellAsyncCA_r16 *ServCellIndex                  `optional,ext-2`
	mgl_r16                   *GapConfig_mgl_r16              `optional,ext-2`
}

func (ie *GapConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.refServCellIndicator != nil || ie.refFR2ServCellAsyncCA_r16 != nil || ie.mgl_r16 != nil
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.gapOffset, &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
		return utils.WrapError("WriteInteger gapOffset", err)
	}
	if err = ie.mgl.Encode(w); err != nil {
		return utils.WrapError("Encode mgl", err)
	}
	if err = ie.mgrp.Encode(w); err != nil {
		return utils.WrapError("Encode mgrp", err)
	}
	if err = ie.mgta.Encode(w); err != nil {
		return utils.WrapError("Encode mgta", err)
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.refServCellIndicator != nil, ie.refFR2ServCellAsyncCA_r16 != nil || ie.mgl_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap GapConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.refServCellIndicator != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode refServCellIndicator optional
			if ie.refServCellIndicator != nil {
				if err = ie.refServCellIndicator.Encode(extWriter); err != nil {
					return utils.WrapError("Encode refServCellIndicator", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.refFR2ServCellAsyncCA_r16 != nil, ie.mgl_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode refFR2ServCellAsyncCA_r16 optional
			if ie.refFR2ServCellAsyncCA_r16 != nil {
				if err = ie.refFR2ServCellAsyncCA_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode refFR2ServCellAsyncCA_r16", err)
				}
			}
			// encode mgl_r16 optional
			if ie.mgl_r16 != nil {
				if err = ie.mgl_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mgl_r16", err)
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

func (ie *GapConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_gapOffset int64
	if tmp_int_gapOffset, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
		return utils.WrapError("ReadInteger gapOffset", err)
	}
	ie.gapOffset = tmp_int_gapOffset
	if err = ie.mgl.Decode(r); err != nil {
		return utils.WrapError("Decode mgl", err)
	}
	if err = ie.mgrp.Decode(r); err != nil {
		return utils.WrapError("Decode mgrp", err)
	}
	if err = ie.mgta.Decode(r); err != nil {
		return utils.WrapError("Decode mgta", err)
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
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

			refServCellIndicatorPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode refServCellIndicator optional
			if refServCellIndicatorPresent {
				ie.refServCellIndicator = new(GapConfig_refServCellIndicator)
				if err = ie.refServCellIndicator.Decode(extReader); err != nil {
					return utils.WrapError("Decode refServCellIndicator", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			refFR2ServCellAsyncCA_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mgl_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode refFR2ServCellAsyncCA_r16 optional
			if refFR2ServCellAsyncCA_r16Present {
				ie.refFR2ServCellAsyncCA_r16 = new(ServCellIndex)
				if err = ie.refFR2ServCellAsyncCA_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode refFR2ServCellAsyncCA_r16", err)
				}
			}
			// decode mgl_r16 optional
			if mgl_r16Present {
				ie.mgl_r16 = new(GapConfig_mgl_r16)
				if err = ie.mgl_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode mgl_r16", err)
				}
			}
		}
	}
	return nil
}
