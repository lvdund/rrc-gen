package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB19_r17 struct {
	ntn_Config_r17                   *NTN_Config_r17              `optional`
	t_Service_r17                    *int64                       `lb:0,ub:549755813887,optional`
	referenceLocation_r17            *ReferenceLocation_r17       `optional`
	distanceThresh_r17               *int64                       `lb:0,ub:65525,optional`
	ntn_NeighCellConfigList_r17      *NTN_NeighCellConfigList_r17 `optional`
	lateNonCriticalExtension         *[]byte                      `optional`
	ntn_NeighCellConfigListExt_v1720 *NTN_NeighCellConfigList_r17 `optional,ext-1`
}

func (ie *SIB19_r17) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.ntn_NeighCellConfigListExt_v1720 != nil
	preambleBits := []bool{hasExtensions, ie.ntn_Config_r17 != nil, ie.t_Service_r17 != nil, ie.referenceLocation_r17 != nil, ie.distanceThresh_r17 != nil, ie.ntn_NeighCellConfigList_r17 != nil, ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ntn_Config_r17 != nil {
		if err = ie.ntn_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ntn_Config_r17", err)
		}
	}
	if ie.t_Service_r17 != nil {
		if err = w.WriteInteger(*ie.t_Service_r17, &uper.Constraint{Lb: 0, Ub: 549755813887}, false); err != nil {
			return utils.WrapError("Encode t_Service_r17", err)
		}
	}
	if ie.referenceLocation_r17 != nil {
		if err = ie.referenceLocation_r17.Encode(w); err != nil {
			return utils.WrapError("Encode referenceLocation_r17", err)
		}
	}
	if ie.distanceThresh_r17 != nil {
		if err = w.WriteInteger(*ie.distanceThresh_r17, &uper.Constraint{Lb: 0, Ub: 65525}, false); err != nil {
			return utils.WrapError("Encode distanceThresh_r17", err)
		}
	}
	if ie.ntn_NeighCellConfigList_r17 != nil {
		if err = ie.ntn_NeighCellConfigList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ntn_NeighCellConfigList_r17", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.ntn_NeighCellConfigListExt_v1720 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SIB19_r17", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.ntn_NeighCellConfigListExt_v1720 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ntn_NeighCellConfigListExt_v1720 optional
			if ie.ntn_NeighCellConfigListExt_v1720 != nil {
				if err = ie.ntn_NeighCellConfigListExt_v1720.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ntn_NeighCellConfigListExt_v1720", err)
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

func (ie *SIB19_r17) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var ntn_Config_r17Present bool
	if ntn_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var t_Service_r17Present bool
	if t_Service_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var referenceLocation_r17Present bool
	if referenceLocation_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var distanceThresh_r17Present bool
	if distanceThresh_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ntn_NeighCellConfigList_r17Present bool
	if ntn_NeighCellConfigList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ntn_Config_r17Present {
		ie.ntn_Config_r17 = new(NTN_Config_r17)
		if err = ie.ntn_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ntn_Config_r17", err)
		}
	}
	if t_Service_r17Present {
		var tmp_int_t_Service_r17 int64
		if tmp_int_t_Service_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 549755813887}, false); err != nil {
			return utils.WrapError("Decode t_Service_r17", err)
		}
		ie.t_Service_r17 = &tmp_int_t_Service_r17
	}
	if referenceLocation_r17Present {
		ie.referenceLocation_r17 = new(ReferenceLocation_r17)
		if err = ie.referenceLocation_r17.Decode(r); err != nil {
			return utils.WrapError("Decode referenceLocation_r17", err)
		}
	}
	if distanceThresh_r17Present {
		var tmp_int_distanceThresh_r17 int64
		if tmp_int_distanceThresh_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 65525}, false); err != nil {
			return utils.WrapError("Decode distanceThresh_r17", err)
		}
		ie.distanceThresh_r17 = &tmp_int_distanceThresh_r17
	}
	if ntn_NeighCellConfigList_r17Present {
		ie.ntn_NeighCellConfigList_r17 = new(NTN_NeighCellConfigList_r17)
		if err = ie.ntn_NeighCellConfigList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ntn_NeighCellConfigList_r17", err)
		}
	}
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
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

			ntn_NeighCellConfigListExt_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ntn_NeighCellConfigListExt_v1720 optional
			if ntn_NeighCellConfigListExt_v1720Present {
				ie.ntn_NeighCellConfigListExt_v1720 = new(NTN_NeighCellConfigList_r17)
				if err = ie.ntn_NeighCellConfigListExt_v1720.Decode(extReader); err != nil {
					return utils.WrapError("Decode ntn_NeighCellConfigListExt_v1720", err)
				}
			}
		}
	}
	return nil
}
