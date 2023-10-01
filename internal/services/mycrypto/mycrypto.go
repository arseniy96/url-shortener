package mycrypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"

	"github.com/arseniy96/url-shortener/internal/logger"
)

const (
	CertFilePath    = "./cert/cert.pem"
	KeyFilePath     = "./cert/key.pem"
	FilePermissions = 0600
	CertNumber      = 1711
	KeySize         = 4096
)

func LoadCryptoFiles() (string, string, error) {
	_, err := os.Open(CertFilePath)
	if err == nil {
		_, err := os.Open(KeyFilePath)
		if err == nil {
			return CertFilePath, KeyFilePath, nil
		}
	}

	certFile, err := os.OpenFile(CertFilePath, os.O_CREATE|os.O_WRONLY, FilePermissions)
	if err != nil {
		return "", "", err
	}
	keyFile, err := os.OpenFile(KeyFilePath, os.O_CREATE|os.O_WRONLY, FilePermissions)
	if err != nil {
		return "", "", err
	}

	// создаём шаблон сертификата
	cert := &x509.Certificate{
		// указываем уникальный номер сертификата
		SerialNumber: big.NewInt(CertNumber),
		// заполняем базовую информацию о владельце сертификата
		Subject: pkix.Name{
			Organization: []string{"MyOrganization"},
			Country:      []string{"RU"},
		},
		// разрешаем использование сертификата для 127.0.0.1 и ::1
		//nolint:gomnd // it's localhost ip
		IPAddresses: []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback},
		// сертификат верен, начиная со времени создания
		NotBefore: time.Now(),
		// время жизни сертификата — 10 лет
		NotAfter:     time.Now().AddDate(0, 1, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		// устанавливаем использование ключа для цифровой подписи,
		// а также клиентской и серверной авторизации
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageDigitalSignature,
	}

	// создаём новый приватный RSA-ключ длиной 4096 бит
	// обратите внимание, что для генерации ключа и сертификата
	// используется rand.Reader в качестве источника случайных данных
	privateKey, err := rsa.GenerateKey(rand.Reader, KeySize)
	if err != nil {
		logger.Log.Error(err)
		return "", "", err
	}

	// создаём сертификат x.509
	certBytes, err := x509.CreateCertificate(rand.Reader, cert, cert, &privateKey.PublicKey, privateKey)
	if err != nil {
		logger.Log.Error(err)
		return "", "", err
	}

	// кодируем сертификат и ключ в формате PEM, который
	// используется для хранения и обмена криптографическими ключами
	err = pem.Encode(certFile, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})
	if err != nil {
		logger.Log.Error(err)
		return "", "", err
	}

	err = pem.Encode(keyFile, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	if err != nil {
		logger.Log.Error(err)
		return "", "", err
	}

	return CertFilePath, KeyFilePath, nil
}
