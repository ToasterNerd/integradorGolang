package avatar

import (
	"log"

	pkgEncoder "github.com/ToasterNerd/integradorGolang/avatar/encoder"
	pkgImages "github.com/ToasterNerd/integradorGolang/avatar/images"
)

//interfaz publica porque se va a exportar el pkg para que se use en otro lado
type CryptoEncoder interface {
	EncodeInfo(stringInfo string) uint64
}

//interfaz publica porque se va a exportar el pkg para que se use en otro lado
type ImageGenerator interface {
	BuildAndSaveImage(encodedInfo uint64) error
}

//struct publica porque se va a exportar el pkg para que se use en otro lado
type Service struct {
	EncoderCrypto  CryptoEncoder
	GeneratorImage ImageGenerator
}

//funcion publica porque se va a exportar el pkg para que se use en otro lado
func ServiceGenerator() *Service {
	return &Service{
		EncoderCrypto:  &pkgEncoder.EncoderCryptoStruct{},
		GeneratorImage: &pkgImages.GeneratorImageStruct{},
	}
}

//struct publica porque se va a exportar el pkg para que se use en otro lado
type Information struct {
	Email string
}

//funcion publica porque se va a exportar el pkg para que se use en otro lado
func (s *Service) GenerateAndSaveAvatar(information Information) error {

	defer func() {
		if err := recover(); err != nil {
			log.Println("panic-GenerateAndSaveAvatar:", err)
		}
	}()

	encodeado := s.EncoderCrypto.EncodeInfo(information.Email)
	s.GeneratorImage.BuildAndSaveImage(encodeado)

	return nil
}
