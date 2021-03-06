package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"unicode"
	mathRand "math/rand"
	"path/filepath"
	"os"
	"io/ioutil"
	"mime/multipart"

	"github.com/gorilla/sessions"
	infrastructure "github.com/nugrohosam/goe2eds/services/infrastructure"
	viper "github.com/spf13/viper"
	redisStore "gopkg.in/boj/redistore.v1"
)

// MaxDepth ...
const MaxDepth = 32

// SetAuth ..
func SetAuth(auth *Auth) {
	AuthUser = auth
}

// GetAuth ..
func GetAuth() Auth {
	return *AuthUser
}

// ParseCert ..
func ParseCert(cert []byte) (*x509.Certificate, error) {
	parsedCert, err := x509.ParseCertificate(cert)
	if err != nil {
		return nil, err
	}

	return parsedCert, nil
}

// RandomString ..
func RandomString(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[mathRand.Intn(len(letterRunes))]
	}

	return string(b)
}

// DecodePublicKey ..
func DecodePublicKey(pemEncodedPub string) *rsa.PublicKey {
	blockPub, _ := pem.Decode([]byte(pemEncodedPub))
	x509EncodedPub := blockPub.Bytes
	publicKey, _ := x509.ParsePKCS1PublicKey(x509EncodedPub)

	return publicKey
}


// DecodePrivateKey ..
func DecodePrivateKey(pemEncoded string) *rsa.PrivateKey {
	block, _ := pem.Decode([]byte(pemEncoded))
	x509Encoded := block.Bytes
	privateKey, _ := x509.ParsePKCS1PrivateKey(x509Encoded)

	return privateKey
}

// EncodeKey ..
func EncodeKey(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) (string, string) {
	x509Encoded := x509.MarshalPKCS1PrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	x509EncodedPub := x509.MarshalPKCS1PublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	return string(pemEncoded), string(pemEncodedPub)
}

// StoreCache ..
func StoreCache(key string, data interface{}) error {
	return infrastructure.StoreCacheRedis(key, data)
}

// SetPath ..
func SetPath(paths ...string) string {
	return strings.Join(paths, "/")
}

// UnsetMap ..
func UnsetMap(data []interface{}, indexRemove int) []interface{} {
	firstSlice := data[:indexRemove]
	lastSlice := data[indexRemove+1:]

	return append(firstSlice, lastSlice)
}

// GetCache ..
func GetCache(key string) (interface{}, error) {
	return infrastructure.GetCacheRedis(key)
}

// MergeMap ...
func MergeMap(dst, src map[string]interface{}) map[string]interface{} {
	return merge(dst, src, 0)
}

func merge(dst, src map[string]interface{}, depth int) map[string]interface{} {
	if depth > MaxDepth {
		panic("too deep!")
	}
	for key, srcVal := range src {
		if dstVal, ok := dst[key]; ok {
			srcMap, srcMapOk := mapify(srcVal)
			dstMap, dstMapOk := mapify(dstVal)
			if srcMapOk && dstMapOk {
				srcVal = merge(dstMap, srcMap, depth+1)
			}
		}
		dst[key] = srcVal
	}
	return dst
}

// UcFirst ..
func UcFirst(s string) string {
	for index, value := range s {
		return string(unicode.ToUpper(value)) + s[index+1:]
	}
	return ""
}

// LcFirst ..
func LcFirst(s string) string {
	for index, value := range s {
		return string(unicode.ToLower(value)) + s[index+1:]
	}
	return ""
}

// GenerateLimitOffset ..
func GenerateLimitOffset(perPage, page string) (string, string) {
	limit := perPage
	pageInt, _ := strconv.Atoi(page)
	perPageInt, _ := strconv.Atoi(perPage)
	offset := perPageInt * (pageInt - 1)

	return limit, fmt.Sprint(offset)
}

// TypeName ..
func TypeName(t reflect.Type) string {
	return t.Name()
}

// SessionRedis ..
func SessionRedis(key string) {
	store, err := redisStore.NewRediStore(10, "tcp", ":6379", "", []byte(key))
	if err != nil {
		panic(err)
	}
	defer store.Close()
}

// Find ..
func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// Encrypt ..
func Encrypt(stringToEncrypt string, keyString string) (encryptedString string) {
	//Since the key is in string, we need to convert decode it to bytes
	key, err := hex.DecodeString(keyString)
	if err != nil {
		panic(err)
	}

	plaintext := []byte(stringToEncrypt)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}

// StringInSlice ..
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}

	return false
}

// Decrypt ..
func Decrypt(encryptedString string, keyString string) (decryptedString string) {

	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(encryptedString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}

// RedisStoreSesssion ..
func RedisStoreSesssion() *redisStore.RediStore {
	redisKey := viper.GetString("reids.key")
	store, err := redisStore.NewRediStore(10, "tcp", ":6379", "", []byte(redisKey))
	if err != nil {
		panic(err)
	}

	return store
}

// StoreFile ..
func StoreFile(file []byte, filePath string) error {
	folderOfFile := filepath.Dir(filePath)
	FolderCheckAndCreate(folderOfFile)
	return ioutil.WriteFile(filePath, file, 0755)
}

// FolderCheckAndCreate ..
func FolderCheckAndCreate(folderPath string) {
	info, err := os.Stat(folderPath)
	if !(os.IsExist(err) && info.Mode().IsDir() && info.Mode().IsRegular()) {
		os.MkdirAll(folderPath, os.ModePerm)
	}
}

// StoreSessionString ..
func StoreSessionString(request *http.Request, writer http.ResponseWriter, nameSession string, data string) {
	if viper.GetString("session.driver") == "redis" {
		store := RedisStoreSesssion()
		defer store.Close()
	}

	sessionStore := sessions.NewCookieStore([]byte(viper.GetString("app.key")))
	sessionNow, err := sessionStore.Get(request, nameSession)
	if err != nil {
		panic(err)
	}

	sessionNow.Values["data"] = data
	sessionNow.Save(request, writer)
}

// ReadFileRequest ..
func ReadFileRequest(file *multipart.FileHeader) ([]byte, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	dst := viper.GetString("temp.root-path") + "/" + RandomString(5)
	os.Mkdir(dst, 0755)

	filePath := dst + "/" + file.Filename

	out, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer out.Close()

	_, err = io.Copy(out, src)

	return ioutil.ReadFile(filePath)
}

// SetPublicLink ..
func SetPublicLink(filePath string) string {
	host := viper.GetString("app.url")
	port := viper.GetString("app.port_exposed")
	urlLink := SetPath(host+":"+port, filePath)

	return urlLink
}

// GetSessionDataString ..
func GetSessionDataString(request *http.Request, writer http.ResponseWriter, nameSession string) string {
	if viper.GetString("session.driver") == "redis" {
		store := RedisStoreSesssion()
		defer store.Close()
	}

	sessionStore := sessions.NewCookieStore([]byte(viper.GetString("app.key")))
	sessionNow, err := sessionStore.Get(request, nameSession)
	if err != nil {
		return ""
	}

	return sessionNow.Values["data"].(string)
}

// DeleteSessionDataString ..
func DeleteSessionDataString(request *http.Request, writer http.ResponseWriter, nameSession string) error {
	if viper.GetString("session.driver") == "redis" {
		store := RedisStoreSesssion()
		defer store.Close()
	}

	sessionStore := sessions.NewCookieStore([]byte(viper.GetString("app.key")))
	sessionNow, err := sessionStore.Get(request, nameSession)
	if err != nil {
		panic(err)
	}

	sessionNow.Options.MaxAge = -1
	err = sessionNow.Save(request, writer)

	return err
}

// GetSessionData ..
func GetSessionData(request *http.Request, writer http.ResponseWriter, nameSession string) interface{} {
	sessionStore := sessions.NewCookieStore([]byte(viper.GetString("app.key")))
	sessionNow, err := sessionStore.Get(request, nameSession)
	if err != nil {
		return ""
	}

	return sessionNow.Values["data"]
}

func mapify(i interface{}) (map[string]interface{}, bool) {
	value := reflect.ValueOf(i)
	if value.Kind() == reflect.Map {
		m := map[string]interface{}{}
		for _, k := range value.MapKeys() {
			m[k.String()] = value.MapIndex(k).Interface()
		}
		return m, true
	}
	return map[string]interface{}{}, false
}
