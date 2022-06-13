package encoder

import "hash/fnv"

//struct publica porque se va a llamar desde service.
type EncoderCryptoStruct struct {
}

//metodo publico porque se va a llamar desde service para "engancharlo" desde la interfaz"
func (m *EncoderCryptoStruct) EncodeInfo(stringInfo string) uint64 {

	hash := fnv.New64a()
	hash.Write([]byte(stringInfo))

	return hash.Sum64()
}
